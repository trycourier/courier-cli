import {Alert} from '@inkjs/ui';
import {InboundTrackEvent} from '@trycourier/courier/api/index.js';
import duckdb from 'duckdb';
import _ from 'lodash';
import React, {useEffect, useState} from 'react';
import {useBoolean, useCounter} from 'usehooks-ts';
import getDb from '../bulk.js';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';
import uuid from '../lib/uuid.js';

interface IParams {
	_: string[];
}

export default ({}: {}) => {
	const {parsedParams, courier} = useCliContext();
	const processing = useBoolean(true);
	const running = useBoolean(true);
	const [data, setData] = useState<duckdb.TableData | undefined>();
	const [data_errors, setDataError] = useState<string[]>([]);
	const counter = useCounter(0);

	const {
		_: [event_name, filename],
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
					messageId: uuid(),
					recipient: `${recipient ?? user_id}`,
					event: event_name,
					type: 'track',
					properties: {},
				} as InboundTrackEvent;
				_.keys(rest).forEach(key => {
					_.set(body, ['properties', key], rest[key]);
				});

				const res = await courier.inbound.track(body);
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
			<>
				<Alert
					variant={!counter.count ? 'info' : 'success'}
					title={`Track complete`}
				>
					Sent {counter.count} track events to {event_name}
				</Alert>
			</>
		);
	}
};
