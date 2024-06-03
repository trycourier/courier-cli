import {UserPreferencesListResponse} from '@trycourier/courier/api/resources/users/index.js';
import {Box, Text} from 'ink';
import lodash from 'lodash';
import React, {useEffect, useState} from 'react';
import {useCliContext} from '../components/Context.js';
import SdkResponse from '../components/SdkResponse.js';
import UhOh from '../components/UhOh.js';
import Spinner from '../components/Spinner.js';

interface IParam {
	_: string[];
	tenant?: string | number;
	verbose?: boolean;
}

export default () => {
	const {courier, parsedParams} = useCliContext();
	const params = parsedParams as IParam;
	const [resp, setResp] = useState<UserPreferencesListResponse>();
	const [err, setErr] = useState<Error>();

	const userId = lodash.get(params, ['_', 0]);
	const verbose = lodash.get(params, ['verbose'], false);

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
			return `${item.topic_name}: ${item.status}`;
		});
		return (
			<Box marginY={1} flexDirection="column">
				{!(resp || err) ? (
					<Spinner text={`Fetching user preferences for ${userId}...`} />
				) : (
					<>
						<Text color="green">User preferences for {userId}:</Text>
						<SdkResponse
							response={verbose ? resp : short_preferences}
							error={err}
						/>
					</>
				)}
			</Box>
		);
	}
};
