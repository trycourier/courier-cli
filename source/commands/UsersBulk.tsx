import {ProgressBar} from '@inkjs/ui';
import duckdb from 'duckdb';
import {Box, Text} from 'ink';
import _ from 'lodash';
import React, {useEffect, useState} from 'react';
import {useBoolean, useCounter} from 'usehooks-ts';
import getDb from '../bulk.js';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';
import delay from '../lib/delay.js';

export default () => {
	const {parsedParams, courier} = useCliContext();
	const [error, setError] = useState<string | undefined>();
	const processing = useBoolean(false);
	const [data, setData] = useState<duckdb.TableData | undefined>();
	const [data_errors, setDataErrors] = useState<string[]>([]);
	const counter = useCounter(0);

	const filename = String(_.get(parsedParams, ['_', 0], ''));
	const {db, filetype, sql} = getDb(filename);

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

	const getData = () => {
		processing.setTrue();

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
