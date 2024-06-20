import {Alert} from '@inkjs/ui';
import {AutomationInvokeParams} from '@trycourier/courier/api/index.js';
import duckdb from 'duckdb';
import _ from 'lodash';
import React, {useEffect, useState} from 'react';
import {useBoolean, useCounter} from 'usehooks-ts';
import getDb from '../bulk.js';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';
import {Text} from 'ink';

interface IParams {
	_: string[];
}

const AutomationInvokeBulk = () => {
	const {parsedParams, courier} = useCliContext();
	const processing = useBoolean(true);
	const running = useBoolean(true);
	const [data, setData] = useState<duckdb.TableData | undefined>();
	const [data_errors, setDataError] = useState<string[]>([]);
	const counter = useCounter(0);

	const {
		_: [template_id, filename],
	} = parsedParams as IParams;
	const {db, filetype, sql} = getDb(filename || '');

	useEffect(() => {
		if (filetype) {
			getData();
		} else {
			setDataError(p => [...p, 'File type not supported.']);
		}
	}, []);

	useEffect(() => {
		if (data) {
			processData();
		}
	}, [data]);

	useEffect(() => {
		if (data_errors?.length) {
			running.setFalse();
			processing.setFalse();
		}
	}, [data_errors]);

	const getData = () => {
		db.all(sql, (err, result) => {
			if (err) {
				setDataError(p => [...p, err.message]);
			} else {
				setData(result);
			}
		});
		processing.setFalse();
	};

	const processData = async () => {
		if (data?.length) {
			for (let i = 0; i < data.length; i++) {
				let {user_id, recipient, ...rest} = data[i] || {};
				let body = {
					recipient: `${recipient ?? user_id}`,
					data: {},
					profile: {},
				} as AutomationInvokeParams;
				_.keys(rest).forEach(key => {
					if (key.startsWith('profile.')) {
						_.set(body, ['profile', key.replace('profile.', '')], rest[key]);
					} else {
						_.set(body, ['data', key], rest[key]);
					}
				});

				const res = await courier.automations.invokeAutomationTemplate(
					template_id || '',
					body,
				);
				if (res instanceof Error) {
					setDataError([res.message]);
				}
				counter.increment();
			}
		}
		running.setFalse();
	};

	if (data_errors?.length) {
		return <UhOh text={data_errors.join('\n')} />;
	} else if (processing.value) {
		return <Spinner text="Processing data..." />;
	} else if (running.value) {
		return (
			<Spinner text={`Processing ${counter.count} / ${data?.length} records`} />
		);
	} else {
		return (
			<Alert
				variant={!counter.count ? 'info' : 'success'}
				title={`Invoke complete`}
			>
				<Text>
					Sent {counter.count} automation invokes to {template_id}
				</Text>
			</Alert>
		);
	}
};

export default AutomationInvokeBulk;
