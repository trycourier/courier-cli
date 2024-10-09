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

const FILENAME = 'tenent_members';

interface IParams extends IParamsOutputOptions {
	_: string[];
	maxPages?: string | number;
}

const TenantsGetMembership = () => {
	const {parsedParams, courier} = useCliContext();
	const processing = useBoolean(true);
	const running = useBoolean(true);
	const counter = useCounter(0);
	const [members, setMembers] = useState<
		Omit<Courier.UserTenantAssociation, 'tenant_id' | 'type'>[]
	>([]);
	const [error, setError] = useState<string | undefined>();

	const {
		maxPages,
		json,
		csv,
		webhook,
		filename,
		_: [tenant_id],
	} = parsedParams as IParams;

	const out_file =
		(filename?.length
			? filename.substring(
					0,
					filename.includes('.') ? filename.lastIndexOf('.') : filename.length,
			  )
			: FILENAME) + (csv ? '.csv' : '.json');

	const MAX_PAGES = Number(maxPages) || 100;

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
		getTenantMemberships();
	}, []);

	const getTenantMemberships = async (cursor?: string, count: number = 0) => {
		if (!tenant_id) {
			setError('No Tenant Specified');

			processing.setFalse();
		} else {
			counter.increment();
			const r = await courier.tenants.getUsersByTenant(tenant_id, {
				cursor,
				limit: 100,
			});
			let items = r.items;
			if (items?.length) {
				setMembers(previous => {
					return [
						...previous,
						...items.map(({tenant_id, type, ...rest}) => rest),
					];
				});
				if (r.has_more && count < MAX_PAGES) {
					await getTenantMemberships(r.cursor, count + 1);
				} else {
					processing.setFalse();
				}
			}
		}
	};

	const runExport = async () => {
		const flat = csv ? flattenData(members) : members;
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
		return <Spinner text={`Fetching tenent_members - page ${counter.count}`} />;
	} else {
		return (
			<>
				<Alert variant="success" title={`Finished ${counter.count} pages`}>
					{csv || json
						? `Output ${members.length} tenent_members to ${out_file}`
						: JSON.stringify(members, null, 2)}
				</Alert>
			</>
		);
	}
};

const flattenData = (
	data: Omit<Courier.UserTenantAssociation, 'tenant_id' | 'type'>[],
) => {
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

export default TenantsGetMembership;
