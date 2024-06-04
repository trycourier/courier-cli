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
import {
	SubscriptionTopic,
	SubscriptionTopicNew,
	Tenant,
} from '@trycourier/courier/api/index.js';
import delay from '../lib/delay.js';

type SubscriptionTopicPref = {[key: string]: SubscriptionTopicNew};

interface TTenantSubscriptionModified extends Tenant {
	default_preferences: SubscriptionTopicPref;
}

const tenantToModifed = (t: Tenant) => {
	const modified: TTenantSubscriptionModified = {
		...t,
		default_preferences: _.get(t, ['default_preferences', 'items'], []).reduce(
			(p: TTenantSubscriptionModified, {id, ...r}: SubscriptionTopic) => {
				_.set(p, [id], r);
				return p;
			},
			{} as TTenantSubscriptionModified,
		),
	};
	return modified;
};
const modifiedToTenant = (t: TTenantSubscriptionModified) => {
	const tenant: Tenant = {
		...t,
		default_preferences: {
			items: Object.entries(t.default_preferences).map(([id, r]) => {
				return {id, ...r};
			}),
		},
	};
	return tenant;
};

export default () => {
	const {parsedParams, courier} = useCliContext();
	const [error, setError] = useState<string | undefined>();
	const processing = useBoolean(false);
	const [data, setData] = useState<duckdb.TableData | undefined>();
	const [data_errors, setDataErrors] = useState<string[]>([]);
	const counter = useCounter(0);

	const filename = String(_.get(parsedParams, ['_', 0], ''));
	const {db, filetype, sql} = getDb(filename);

	const merge = Boolean(parsedParams['merge']);

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
				let {
					tenant_id,
					parent_tenant_id,
					name,
					brand_id,
					user_profile = {},
					default_preferences = {},
					properties = {},
					...rest
				} = data[i] || {};
				tenant_id = tenant_id ? String(tenant_id) : undefined;
				parent_tenant_id = parent_tenant_id
					? String(parent_tenant_id)
					: undefined;
				brand_id = brand_id ? String(brand_id) : undefined;
				name = name ? String(name) : tenant_id;
				if (!tenant_id) {
					setDataErrors(p => [...p, `tenant_id not found in index ${i}`]);
				} else {
					try {
						// handle rest, all should go into properties unlness they are user_profile or default_preferences
						Object.entries(rest).forEach(([key, value]) => {
							if (key.startsWith('user_profile.')) {
								_.set(
									user_profile,
									key.slice('user_profile.'.length - 1),
									value,
								);
							} else if (key.startsWith('default_preferences.')) {
								_.set(
									default_preferences,
									key.slice('default_preferences.'.length - 1),
									value,
								);
							} else {
								_.set(properties, key, value);
							}
						});
						let curr: Tenant | undefined;
						let next: Partial<Tenant> = {
							parent_tenant_id,
							name,
							brand_id,
							user_profile,
							default_preferences,
							properties,
						};
						if (merge) {
							try {
								curr = await courier.tenants.get(tenant_id);
							} catch (e) {
								curr = undefined;
							} finally {
								if (curr) {
									const modified = tenantToModifed(curr);
									next = {
										...curr,
										...next,
										user_profile: {
											...next.user_profile,
											...curr.user_profile,
										},
									};

									next = {
										...next,
										default_preferences: {
											...modified?.default_preferences,
											...next.default_preferences,
										},
										user_profile: {
											...modified?.user_profile,
											...next.user_profile,
										},
										properties: {
											...modified?.properties,
											...next.properties,
										},
									};
								}
							}
						}
						const add = modifiedToTenant(next as TTenantSubscriptionModified);
						console.log(JSON.stringify(add, null, 2));
						await delay(20000);
						await courier.tenants.createOrReplace(
							tenant_id,
							modifiedToTenant(next as TTenantSubscriptionModified),
						);

						counter.increment();
					} catch (err) {
						setDataErrors(p => [
							...p,
							`tenant_id (${tenant_id}) failed to update in index ${i}: ${String(
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
