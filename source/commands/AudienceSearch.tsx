import {Alert} from '@inkjs/ui';
import {Courier} from '@trycourier/courier';
import {stringify} from 'csv-stringify/sync';
import fs from 'fs/promises';
import _ from 'lodash';
import {DateTime} from 'luxon';
import React, {useEffect, useState} from 'react';
import {useBoolean, useCounter} from 'usehooks-ts';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';

const FILENAME = 'audiences';

interface IParams extends IParamsOutputOptions {
	id?: string;
	name?: string;
	maxPages?: string | number;
}

const AudienceSearch = () => {
	const {parsedParams, courier} = useCliContext();
	const processing = useBoolean(true);
	const running = useBoolean(true);
	const counter = useCounter(0);
	const [audiences, setAudiences] = useState<Courier.Audience[]>([]);
	const [error, setError] = useState<string | undefined>();

	const {maxPages, json, csv, webhook, filename, name, id} =
		parsedParams as IParams;

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
		getAudiences();
	}, []);

	const getAudiences = async (cursor?: string, count: number = 0) => {
		counter.increment();
		const r = await courier.audiences.listAudiences({
			cursor,
		});
		let items = r.items;
		if (name) {
			items = items.filter(a => a.name.includes(name));
		}
		if (id) {
			items = items.filter(a => a.id.includes(id));
		}

		setAudiences(p => [...p, ...items]);
		if (r.paging.more && count < MAX_PAGES) {
			await getAudiences(r.paging.cursor, count + 1);
		} else {
			processing.setFalse();
		}
	};

	const runExport = async () => {
		const flat = csv ? flattenData(audiences) : audiences;
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
		return <Spinner text={`Fetching audiences - page ${counter.count}`} />;
	} else {
		return (
			<>
				<Alert variant="success" title={`Finished ${counter.count} pages`}>
					{csv || json
						? `Output ${audiences.length} audiences to ${out_file}`
						: JSON.stringify(audiences, null, 2)}
				</Alert>
			</>
		);
	}
};

const flattenData = (data: Courier.Audience[]) => {
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

export default AudienceSearch;
