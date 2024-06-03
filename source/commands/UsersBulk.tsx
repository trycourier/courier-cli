import duckdb from 'duckdb';
import {Box, Text} from 'ink';
import React, {useEffect, useState} from 'react';
import {useBoolean, useCounter} from 'usehooks-ts';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';
import _ from 'lodash';
import delay from '../lib/delay.js';
import {ProgressBar} from '@inkjs/ui';

type TFileType = 'csv' | 'json' | 'parquet';

const installExtension = (db: duckdb.Database, type?: TFileType) => {
	if (['json', 'parquet'].includes(type || '')) {
		db.exec(`
            INSTALL ${type};
            LOAD ${type};`);
	}
};

const getFrom = (filename: string, type: TFileType) => {
	switch (type) {
		case 'csv':
			return `read_csv(['${filename}'], union_by_name = true)`;
		case 'json':
			return `read_json_auto(['${filename}'])`;
		case 'parquet':
			return `read_parquet(['${filename}'])`;
	}
};

const getFileType: (filename: string) => TFileType | undefined = (
	filename: string,
) => {
	if (filename.endsWith('.csv')) {
		return 'csv';
	} else if (filename.endsWith('.json') || filename.endsWith('.jsonl')) {
		return 'json';
	} else if (
		filename.endsWith('.parquet') ||
		filename.endsWith('.pq') ||
		filename.endsWith('.parq')
	) {
		return 'parquet';
	} else {
		return undefined;
	}
};

export default () => {
	const {parsedParams, courier} = useCliContext();
	const [error, setError] = useState<string | undefined>();
	const processing = useBoolean(false);
	const [data, setData] = useState<duckdb.TableData | undefined>();
	const [data_errors, setDataErrors] = useState<string[]>([]);
	const counter = useCounter(0);
	// const [resp, setResp] = useState<IResponse | undefined>();
	const filename = String(_.get(parsedParams, ['_', 0], ''));
	const filetype = getFileType(filename);
	const db = new duckdb.Database(':memory:'); // or a file name for a persistent DB
	installExtension(db, filetype);
	const keep_flat = Boolean(parsedParams['keep-flat']);
	const remove_nulls = Boolean(parsedParams['remove-nulls']);
	const replace = Boolean(parsedParams['replace']);
	const lists = String(_.get(parsedParams, ['list'], ''))
		.split(',')
		.map(l => l.trim())
		.filter(l => l.length > 0);
	const tenants = String(_.get(parsedParams, ['tenant'], ''))
		.split(',')
		.map(l => l.trim())
		.filter(t => t.length > 0);

	useEffect(() => {
		if (filetype) {
			getData(filetype);
		} else {
			setError('File type not supported.');
		}
	}, []);

	useEffect(() => {
		if (data) {
			processData();
		}
	}, [data]);

	const getData = (type: TFileType) => {
		processing.setTrue();
		const sql = `SELECT * FROM ${getFrom(filename, type)} ;`;

		db.all(sql, (err, result) => {
			if (err) {
				setError(err.message);
			} else {
				setData(result);
			}
		});
	};

	const processData = async () => {
		if (data?.length) {
			for (let i = 0; i < data.length; i++) {
				let {user_id, ...profile} = data[i] || {};
				let userId = user_id ? String(user_id) : undefined;
				if (!userId) {
					setDataErrors(p => [...p, `user_id not found in index ${i}`]);
				} else {
					try {
						Object.entries(profile).forEach(([key, value]) => {
							if (filetype === 'csv' && !keep_flat) {
								_.unset(profile, key);
								_.set(profile, key, value);
							}
							if (value === null && remove_nulls) {
								_.unset(profile, key);
							}
						});
						if (replace) {
							await courier.profiles.replace(userId, {profile});
						} else {
							await courier.profiles.create(userId, {profile});
						}

						if (lists.length) {
							await courier.profiles.subscribeToLists(userId, {
								lists: lists.map(l => ({listId: l})),
							});
						}
						if (tenants.length) {
							await courier.users.tenants.addMultple(userId, {
								tenants: tenants.map(t => ({tenant_id: t})),
							});
						}
						counter.increment();
						delay(10000);
					} catch (err) {
						setDataErrors(p => [
							...p,
							`user_id (${user_id}) failed to update in index ${i}: ${String(
								err,
							)}`,
						]);
					}
				}
			}
		}
		processing.setFalse();
	};

	if (!filename?.length) {
		return <UhOh text="You must specify a filename." />;
	} else if (error?.length) {
		return <UhOh text={error} />;
	} else if (data && processing.value) {
		return (
			<>
				<ProgressBar
					value={Math.floor((counter.count + 1 / data.length) * 100)}
				/>
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
			</Box>
		);
	}
};
