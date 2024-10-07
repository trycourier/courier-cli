import {ProgressBar} from '@inkjs/ui';
import {
	MergeProfileResponse,
	ReplaceProfileResponse,
	SubscribeToListsResponse,
} from '@trycourier/courier/api/index.js';
import duckdb from 'duckdb';
import {Box, Text} from 'ink';
import _ from 'lodash';
import fs from 'fs/promises';
import React, {useEffect, useState} from 'react';
import {useBoolean, useCounter} from 'usehooks-ts';
import getDb, {getChunk} from '../bulk.js';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';
import delay from '../lib/delay.js';
import {UserToken} from '@trycourier/courier/api/resources/users/index.js';

const DEFAULT_DELAY = 5000;
const DEFAULT_CHUNK_SIZE = 500;
const DEFAULT_TIMEOUT = 10;
const DEFAULT_ERROR_FILENAME = 'errors.json';

interface RowResponse {
	userId: string;
	success: Boolean;
	error?: string;
	index: number;
}

export default () => {
	const {parsedParams, courier} = useCliContext();
	const [error, setError] = useState<string | undefined>();
	const processing = useBoolean(true);
	const running = useBoolean(true);
	const [data, setData] = useState<duckdb.TableData | undefined>();
	const [data_errors, setDataErrors] = useState<string[]>([]);
	const counter = useCounter(0);
	const [row_errors, setRowErrors] = useState<duckdb.RowData[]>([]);

	const filename = String(_.get(parsedParams, ['_', 0], ''));
	const {db, filetype, sql} = getDb(filename);

	const delay_between_chunks = Number(parsedParams['delay']) ?? DEFAULT_DELAY;
	const chunk_size = parsedParams['chunk_size']
		? Number(parsedParams['chunk_size'])
		: DEFAULT_CHUNK_SIZE;

	const log_errors = true;

	useEffect(() => {
		if (filetype) {
			getData();
		} else {
			setError('File type not supported.');
		}
	}, []);

	useEffect(() => {
		if (data) {
			processData();
		}
	}, [data]);

	useEffect(() => {
		if (!processing.value) {
			handleErrors();
		}
	}, [processing.value]);

	const getData = () => {
		db.all(sql, (err, result) => {
			if (err) {
				setError(err.message);
			} else {
				setData(result);
			}
		});
	};

	const processChunkRows = (data: duckdb.RowData[], start_index: number) => {
		return data.map((row, i) => {
			const curr_index = start_index + i;
			let {
				user_id,
				token,
				provider_key,
				device,
				tracking,
				expiry_date,
				...properties
			} = row || {};
			if (!device) device = {};
			if (!tracking) tracking = {};
			if (!properties) properties = {};
			if (!user_id) {
				return Promise.resolve({
					success: false,
					userId: '__unknown__',
					error: `user_id not found in index ${curr_index}`,
					index: curr_index,
				} as RowResponse);
			} else if (!provider_key) {
				return Promise.resolve({
					success: false,
					userId: user_id,
					error: `provider_key not found in index ${curr_index}`,
					index: curr_index,
				} as RowResponse);
			} else if (!token) {
				return Promise.resolve({
					success: false,
					userId: user_id,
					error: `token not found in index ${curr_index}`,
					index: curr_index,
				} as RowResponse);
			} else {
				Object.entries(properties).forEach(([key, value]) => {
					if (key.startsWith('device.')) {
						_.unset(properties, key);
						_.set(device, key.replace('device.', ''), value);
					} else if (key.startsWith('tracking.')) {
						_.unset(properties, key);
						_.set(tracking, key.replace('tracking.', ''), value);
					} else {
						_.unset(properties, key);
						_.set(properties, key.replace('properties.', ''), value);
					}
				});
				return processRow({
					user_id: String(user_id),
					token: token,
					provider_key,
					device,
					tracking,
					expiry_date,
					properties,
					index: curr_index,
				});
			}
		});
	};

	const processRow: (
		props: {
			user_id: string;
			token: string;
			index: number;
		} & UserToken,
	) => Promise<RowResponse> = async ({user_id, token, index, ...body}) => {
		return new Promise(async resolve => {
			let promises: Promise<
				| SubscribeToListsResponse
				| MergeProfileResponse
				| ReplaceProfileResponse
				| void
			>[] = [];

			try {
				promises.push(
					courier.users.tokens.add(user_id, token, body, {
						maxRetries: 5,
						timeoutInSeconds: DEFAULT_TIMEOUT,
					}),
				);

				await Promise.all(promises);
				counter.increment();
				return resolve({userId: user_id, success: true, index});
			} catch (error: any) {
				counter.increment();
				return resolve({
					userId: user_id,
					success: false,
					index,
					error:
						(String(error) ??
							error?.message ??
							error.message ??
							'Unknown Error') + `+ ${user_id}`,
				});
			}
		});
	};

	const processData = async () => {
		if (data?.length) {
			let data_copy = [...data];
			let counter = 0;
			let {rows, data: rest} = getChunk(data_copy, chunk_size);
			while (rows?.length) {
				const chunk = processChunkRows(rows, counter);
				const processed_chunks = await Promise.all(chunk);
				const errors = processed_chunks.filter(r => !r.success);
				if (errors.length) {
					setDataErrors(p => [
						...p,
						...errors.map(r => {
							return `user_id (${r.userId}) failed to update in index ${
								r.index
							}: ${String(r.error)}`;
						}),
					]);
					setRowErrors(r => [
						...r,
						...errors.map(e => data[e.index]! as duckdb.RowData),
					]);
				}
				if (rest.length > 0) {
					await delay(delay_between_chunks);
					counter += rows.length;
					const next = getChunk(rest, chunk_size);
					rows = next.rows;
					rest = next.data;
				} else {
					processing.setFalse();
					break;
				}
			}
		}
	};

	const handleErrors = async () => {
		if (row_errors.length && log_errors) {
			await fs.writeFile(
				DEFAULT_ERROR_FILENAME,
				JSON.stringify(row_errors, null, 2),
				{
					encoding: 'utf-8',
				},
			);
			running.setFalse();
		} else {
			running.setFalse();
		}
	};

	if (!filename?.length) {
		return <UhOh text="You must specify a filename." />;
	} else if (error?.length) {
		return <UhOh text={error} />;
	} else if (data && running.value) {
		return (
			<>
				<ProgressBar value={Math.floor((counter.count / data.length) * 100)} />
				<Spinner text={`Completed Rows: ${counter.count} / ${data.length}`} />
			</>
		);
	} else {
		return (
			<Box flexDirection="column" marginY={1}>
				<Text color={'green'}>{`Completed Rows: ${counter.count} / ${
					data?.length || 0
				}`}</Text>
				{data_errors.map((err, i) => {
					return <UhOh key={i} text={err} />;
				})}
				{log_errors && data_errors.length ? (
					<Text>Errors output to {DEFAULT_ERROR_FILENAME}</Text>
				) : (
					<></>
				)}
			</Box>
		);
	}
};
