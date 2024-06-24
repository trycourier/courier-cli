import {Alert} from '@inkjs/ui';
import {
	ListMessagesRequest,
	MessageDetails,
} from '@trycourier/courier/api/index.js';
import {stringify} from 'csv-stringify/sync';
import fs from 'fs/promises';
import _ from 'lodash';
import {DateTime} from 'luxon';
import React, {useEffect, useState} from 'react';
import {useBoolean, useCounter} from 'usehooks-ts';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';

const FILENAME = 'messages';

interface IParams extends IParamsOutputOptions {
	user?: string | number;
	from?: string;
	enqueued_after?: string;
	tag?: string | string[];
	status?: string | string[];
	maxPages?: string | number;
}

const MessagesSearch = () => {
	const {parsedParams, courier} = useCliContext();
	const processing = useBoolean(true);
	const running = useBoolean(true);
	const counter = useCounter(0);
	const [messages, setMessages] = useState<MessageDetails[]>([]);
	const [error, setError] = useState<string | undefined>();

	const {
		user,
		from,
		tag,
		status,
		enqueued_after,
		maxPages,
		json,
		csv,
		webhook,
		filename,
	} = parsedParams as IParams;

	let searchParams: ListMessagesRequest = {};
	if (user) searchParams['recipient'] = String(user);
	if (enqueued_after || from)
		searchParams['enqueued_after'] = enqueued_after ?? from;
	if (tag) searchParams['tag'] = tag;
	if (status) searchParams['status'] = status;

	const out_file =
		(filename?.length
			? filename.substring(
					0,
					filename.includes('.') ? filename.lastIndexOf('.') : filename.length,
			  )
			: FILENAME) + (csv ? '.csv' : '.json');

	const MAX_PAGES = Number(maxPages) || 10;

	useEffect(() => {
		if (!processing.value) {
			if (json || csv || webhook?.length) {
				runExport();
			} else {
				running.setFalse();
			}
		}
	}, [processing.value]);

	useEffect(() => {
		getMessages();
	}, []);

	const getMessages = async (cursor?: string, count: number = 0) => {
		counter.increment();
		const r = await courier.messages.list({
			...searchParams,
			cursor,
		});
		setMessages(p => [...p, ...r.results]);
		if (r.paging.more && count < MAX_PAGES) {
			await getMessages(r.paging.cursor, count + 1);
		} else {
			processing.setFalse();
		}
	};

	const runExport = async () => {
		const flat = flattenData(messages);
		if (csv) {
			await fs.writeFile(out_file, stringify(flat, {header: true}));
		} else if (json) {
			await fs.writeFile(out_file, JSON.stringify(flat, null, 2), {
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
					body: JSON.stringify(flat),
				});
			} catch (e) {
				setError(e instanceof Error ? e.message : String(e));
			}
		}
		running.setFalse();
	};

	if (error?.length) {
		return <UhOh text={error} />;
	} else if (running.value) {
		return <Spinner text={`Fetching messages - page ${counter.count}`} />;
	} else {
		return (
			<>
				<Alert variant="success" title={`Finished ${counter.count} pages`}>
					{csv || json
						? `Output ${messages.length} messages to ${out_file}`
						: JSON.stringify(messages, null, 2)}
				</Alert>
			</>
		);
	}
};

const flattenData = (data: MessageDetails[]) => {
	return data.map(row => {
		return Object.keys(row).reduce((p, key) => {
			const v = _.get(row, [key]);
			if (typeof v === 'number') {
				p[key] = DateTime.fromMillis(v, {zone: 'utc'}).toISO();
			} else if (typeof v === 'object') {
				p[key] = JSON.stringify(v);
			} else if (v) {
				p[key] = v;
			}
			return p;
		}, {} as any);
	});
};

export default MessagesSearch;
