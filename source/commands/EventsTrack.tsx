import React, {useEffect, useState} from 'react';
import {Box} from 'ink';
import UhOh from '../components/UhOh.js';
import Request from '../components/Request.js';
import api from '../lib/api.js';

interface IResponse {
	res: Response,
	json?: any,
	err?: Error
}

export default ({params}: {params: any}) => {
	const [resp, setResp] = useState<IResponse | undefined>();

	const eventId = params?._?.[0];
	if (!eventId) {
		return <UhOh text="You must specify an event ID to trigger." />;
	}

	const userId = params?._?.[1];
	if (!userId) {
		return <UhOh text="You must specify a user ID for the event." />;
	}

	const {_, ...properties} = params;
	const payload = {
		event_id: eventId,
		user_id: userId,
		properties,
	};
	const request = {
		method: 'POST',
		url: '/events/track',
		body: payload
	};

	useEffect(() => {
		api(request).then((res) => setResp(res));
	}, []);

	return <Box flexDirection='column'>
		<Request request={request} response={resp} />
		{resp && resp.json ? <>
		</> : null}
	</Box>;
};
