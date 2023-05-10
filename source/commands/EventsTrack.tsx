import React, {useEffect, useState} from 'react';
import {Box, Text} from 'ink';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';
import api from '../lib/api.js';

export default ({params}: {params: any}) => {
	const [resp, setResp] = useState<object | undefined>();
	const [error, setError] = useState<string | undefined>();

	const eventId = params?._?.[0];
	if (!eventId) {
		return <UhOh text="You must specify an event ID to trigger." />;
	}

	const userId = params?._?.[1];
	if (!userId) {
		return <UhOh text="You must specify a user ID for the event." />;
	}

	const {_, ...properties} = params;
	const body = {
		event_id: eventId,
		user_id: userId,
		properties,
	};

	useEffect(() => {

		api('/events/track', 'POST', {
			body: JSON.stringify(body),
		}).then(
			({json}) => setResp(json),
			(err: Error) => setError(err.message),
		);
	}, []);

	if (error) {
		return (
			<Box flexDirection='column'>
				<UhOh text={error} />
				<Text>{JSON.stringify(body, undefined, '  ')}</Text>
			</Box>
		);
	} else if (resp) {
		return <Text color="green">Success 🎉</Text>;
	} else {
		return (
			<Box flexDirection='column'>
				<Spinner text="air-mailing those bits & bytes..." />
				<Text>{JSON.stringify(body, undefined, '  ')}</Text>
			</Box>
		);
	}
};
