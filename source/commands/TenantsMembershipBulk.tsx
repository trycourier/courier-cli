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

const DEFAULT_DELAY = 10000;
const DEFAULT_CHUNK_SIZE = 500;
const DEFAULT_TIMEOUT = 10;
const DEFAULT_REQ_OPTIONS = {maxRetries: 5, timeoutInSeconds: DEFAULT_TIMEOUT};

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

	const log_errors = Boolean(parsedParams['errors']);
	const remove_membership = Boolean(parsedParams['remove-membership']);
	const keep_flat = Boolean(parsedParams['keep-flat']);
	const remove_nulls = Boolean(parsedParams['remove-nulls']);

	const tenants = String(_.get(parsedParams, ['tenant'], ''))
		.split(',')
		.map(l => l.trim())
		.filter(t => t.length > 0);

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
			let {user_id, tenant_id, ...profile} = row || {};
			if (!user_id) {
				return Promise.resolve({
					success: false,
					userId: '__unknown__',
					error: `user_id not found in index ${curr_index}`,
					index: curr_index,
				} as RowResponse);
			} else {
				Object.entries(profile).forEach(([key, value]) => {
					if (filetype === 'csv' && !keep_flat) {
						_.unset(profile, key);
						_.set(profile, key, value);
					}
					if (value === null && remove_nulls) {
						_.unset(profile, key);
					}
				});
				return processRow(
					String(user_id),
					profile,
					curr_index,
					tenant_id?.length ? String(tenant_id) : undefined,
				);
			}
		});
	};

	const processRow: (
		userId: string,
		profile: any,
		index: number,
		tenant_id?: string,
	) => Promise<RowResponse> = async (userId, profile, index, tenant_id) => {
		return new Promise(async resolve => {
			let promises: Promise<
				| SubscribeToListsResponse
				| MergeProfileResponse
				| ReplaceProfileResponse
				| void
			>[] = [];

			try {
				if (!remove_membership) {
					if (tenants.length) {
						promises.push(
							courier.users.tenants.addMultple(
								userId,
								{
									tenants: tenants.map(t => ({tenant_id: t, profile})),
								},
								DEFAULT_REQ_OPTIONS,
							),
						);
					} else if (tenant_id) {
						promises.push(
							courier.users.tenants.add(
								userId,
								tenant_id,
								{profile},
								DEFAULT_REQ_OPTIONS,
							),
						);
					}
				} else {
					if (tenants.length) {
						tenants.forEach(t => {
							promises.push(
								courier.users.tenants.remove(userId, t, DEFAULT_REQ_OPTIONS),
							);
						});
					} else if (tenant_id?.length) {
						promises.push(
							courier.users.tenants.remove(
								userId,
								tenant_id,
								DEFAULT_REQ_OPTIONS,
							),
						);
					}
				}
				await Promise.all(promises);
				counter.increment();
				return resolve({userId, success: true, index});
			} catch (error: any) {
				counter.increment();
				return resolve({
					userId,
					success: false,
					index,
					error:
						(String(error) ??
							error?.message ??
							error.message ??
							'Unknown Error') + `+ ${userId}`,
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
