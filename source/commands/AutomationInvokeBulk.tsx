import {Alert} from '@inkjs/ui';
import {AutomationInvokeParams} from '@trycourier/courier/api/index.js';
import {stringify} from 'csv-stringify/sync';
import duckdb from 'duckdb';
import fs from 'fs/promises';
import {Newline, Text} from 'ink';
import _ from 'lodash';
import React, {useEffect, useRef, useState} from 'react';
import {useBoolean, useCounter} from 'usehooks-ts';
import getDb from '../bulk.js';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';

const FILENAME = 'invokes';

interface IParams {
	_: string[];
	csv?: boolean;
	json?: boolean;
	webhook?: string;
	filename?: string;
}

interface AutomationInvokeParamsWithRunId extends AutomationInvokeParams {
	runId: string;
}

const AutomationInvokeBulk = () => {
	const {parsedParams, courier} = useCliContext();
	const processing = useBoolean(true);
	const running = useBoolean(true);
	const [data, setData] = useState<duckdb.TableData | undefined>();
	const [data_errors, setDataError] = useState<string[]>([]);
	const counter = useCounter(0);
	const invokes = useRef<AutomationInvokeParamsWithRunId[]>([]);
	const addInvoke = (invoke: AutomationInvokeParamsWithRunId) =>
		invokes.current.push(invoke);

	const {
		_: [template_id, infile],
		csv,
		filename: outfile,
		json,
		webhook,
	} = parsedParams as IParams;
	const {db, filetype, sql} = getDb(infile || '');

	const out_file = outfile || FILENAME + (csv ? '.csv' : '.json');

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

	const runExport = async () => {
		if (csv) {
			await fs.writeFile(out_file, stringify(invokes.current, {header: true}));
		} else if (json) {
			await fs.writeFile(out_file, JSON.stringify(invokes.current, null, 2), {
				encoding: 'utf-8',
			});
		}
		if (webhook?.length) {
			try {
				await fetch(webhook, {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
					},
					body: JSON.stringify(invokes),
				});
			} catch (e) {
				setDataError(p => [...p, e instanceof Error ? e.message : String(e)]);
			}
		}
		running.setFalse();
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
				addInvoke({...body, runId: res.runId});

				if (res instanceof Error) {
					setDataError([res.message]);
				}
				counter.increment();
			}
		}
		if (csv || json) {
			runExport();
		} else {
			running.setFalse();
		}
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
				{csv ||
					(json && (
						<Text>
							<Newline />
							Saved to {out_file}
						</Text>
					))}
				{webhook?.length && (
					<Text>
						<Newline />
						Sent to {webhook}
					</Text>
				)}
			</Alert>
		);
	}
};

export default AutomationInvokeBulk;
