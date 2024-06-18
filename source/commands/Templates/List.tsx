import React, {useEffect, useState} from 'react';
import {useCliContext} from '../../components/Context.js';
import {Notification} from '@trycourier/courier/api/index.js';
import Table from '../../components/Table.js';
import {useBoolean} from 'usehooks-ts';
import Spinner from '../../components/Spinner.js';
import {Text} from 'ink';
import _ from 'lodash';
import {DateTime} from 'luxon';
import fs from 'fs/promises';
import {stringify} from 'csv-stringify/sync';
import UhOh from '../../components/UhOh.js';

const MAX_PAGES = 100;
const FILENAME = 'templates';

interface IParams {
	csv?: boolean;
	json?: boolean;
	webhook?: string;
	filename?: string;
}

const TemplatesList = () => {
	const [templates, setTemplates] = useState<Notification[]>([]);
	const {courier, parsedParams} = useCliContext();
	const running = useBoolean(true);
	const processing = useBoolean(true);
	const [error, setError] = useState<string | undefined>();

	const {csv, filename, json, webhook} = parsedParams as IParams;

	const out_file = filename || FILENAME + (csv ? '.csv' : '.json');

	useEffect(() => {
		getTemplates();
	}, []);

	const getTemplates = async () => {
		let cursor: string | undefined;
		let page = -1;
		while (page < MAX_PAGES) {
			page++;
			const templates = await courier.notifications.list({cursor});
			setTemplates(p => [...p, ...templates.results]);
			if (templates.paging.more) {
				cursor = templates.paging.cursor;
			} else {
				break;
			}
		}
		processing.setFalse();
	};

	useEffect(() => {
		if (!processing.value) {
			runExport();
		}
	}, [processing.value, templates]);

	const runExport = async () => {
		if (csv) {
			await fs.writeFile(
				out_file,
				stringify(flattenData(templates), {header: true}),
			);
		} else if (json) {
			await fs.writeFile(
				out_file,
				JSON.stringify(flattenData(templates), null, 2),
				{
					encoding: 'utf-8',
				},
			);
		}
		if (webhook?.length) {
			try {
				await fetch(webhook, {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
					},
					body: JSON.stringify(flattenData(templates)),
				});
			} catch (e) {
				setError(e instanceof Error ? e.message : String(e));
			}
		}
		running.setFalse();
	};

	if (running.value) {
		return <Spinner text="Retrieving notification templates" />;
	} else if (error?.length) {
		return <UhOh text={error} />;
	} else if (json || csv) {
		return <Text>Saved to {out_file}</Text>;
	} else if (webhook?.length) {
		return <Text>Webhook sent to {webhook}</Text>;
	} else {
		return (
			<Table
				data={templates.map(
					({tags, routing, created_at, updated_at, ...rest}) => {
						return {
							...rest,
							tags: _.map(tags?.data, t => t.name).join(', '),
							routing: JSON.stringify(routing, null, 2),
							created_at: DateTime.fromMillis(created_at, {
								zone: 'utc',
							}).toRelative(),
							updated_at: DateTime.fromMillis(updated_at, {
								zone: 'utc',
							}).toRelative(),
						};
					},
				)}
			/>
		);
	}
};

const flattenData = (data: Notification[]) => {
	return data.map(({tags, routing, created_at, updated_at, ...rest}) => {
		return {
			...rest,
			tags: _.map(tags?.data, t => t.name).join(', '),
			routing: JSON.stringify(routing, null, 2),
			created_at: DateTime.fromMillis(created_at, {zone: 'utc'}).toISO(),
			updated_at: DateTime.fromMillis(updated_at, {zone: 'utc'}).toISO(),
		};
	});
};

export default TemplatesList;
