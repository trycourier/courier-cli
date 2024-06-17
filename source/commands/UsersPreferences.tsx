import {UserPreferencesListResponse} from '@trycourier/courier/api/resources/users/index.js';
import {Box, Text} from 'ink';
import lodash from 'lodash';
import React, {useEffect, useState} from 'react';
import {useCliContext} from '../components/Context.js';
import SdkResponse from '../components/SdkResponse.js';
import UhOh from '../components/UhOh.js';
import Spinner from '../components/Spinner.js';
import KVP from '../components/KVP.js';
import api from '../lib/api.js';
import constants from '../constants.js';
import Link from 'ink-link';
import InkSpinner from 'ink-spinner';

interface IParam {
	_: string[];
	tenant?: string | number;
	verbose?: boolean;
}

export default () => {
	const {courier, parsedParams, url: baseUrl, apikey} = useCliContext();
	const params = parsedParams as IParam;
	const [resp, setResp] = useState<UserPreferencesListResponse>();
	const [urlResp, setUrlResp] = useState<IResponseDebug | undefined>();
	const [err, setErr] = useState<Error>();

	const userId = lodash.get(params, ['_', 0]);
	const verbose = lodash.get(params, ['verbose'], false);
	const url = lodash.get(params, ['url'], false);
	const brand = lodash.get(params, ['brand']);

	useEffect(() => {
		if (url) {
			api(
				{
					url: '/debug',
					method: 'POST',
				},
				baseUrl,
				apikey!,
			).then(res => setUrlResp(res));
		}
	}, []);

	useEffect(() => {
		getPreferences();
	}, []);

	const getPreferences = async () => {
		try {
			const user = await courier.users.preferences.list(userId);
			setResp(user);
		} catch (e) {
			setErr(e as Error);
		}
	};

	if (!userId) {
		return <UhOh text="You must specify a user ID." />;
	} else {
		const short_preferences = lodash.map(lodash.get(resp, ['items']), item => {
			return (
				<KVP key={item.topic_id} label={item.topic_name} value={item.status} />
			);
		});
		const whoami = lodash.get(urlResp, ['json'], {} as any);
		const preference_url = urlResp
			? getUrl({...whoami, userId, brand})
			: undefined;

		return (
			<Box marginY={1} flexDirection="column">
				{!(resp || err) ? (
					<Spinner text={`Fetching user preferences for ${userId}...`} />
				) : (
					<>
						<Text color="green">User preferences for {userId}:</Text>
						<SdkResponse
							response={verbose ? resp : undefined}
							error={err}
							content={!(err || verbose) ? short_preferences : undefined}
						/>
					</>
				)}

				{url && (
					<Text>
						{!urlResp ? <InkSpinner type="dots" /> : <Text> </Text>}
						<Text> </Text>
						<Link url={preference_url || ''}>
							<Text color="cyan">Open Preference Page</Text>
						</Link>
					</Text>
				)}
			</Box>
		);
	}
};

const getUrl = ({
	userId,
	tenantId,
	scope,
	environment,
	brand,
}: {
	userId: string;
	tenantId: string;
	scope: string;
	environment: 'test' | 'production';
	brand?: string;
}) => {
	const is_draft = (scope || '').includes('draft');
	const tenant = environment === 'test' ? tenantId + '/test' : tenantId;
	const encoded = btoa([tenant, brand, userId, is_draft].join('#'));
	return `${constants.hosted_preference_page}/${encoded}`;
};
