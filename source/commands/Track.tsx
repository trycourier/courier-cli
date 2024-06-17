import React, {useEffect, useState} from 'react';
import UhOh from '../components/UhOh.js';
import Request from '../components/Request.js';
import Response from '../components/Response.js';
import api from '../lib/api.js';
import {useCliContext} from '../components/Context.js';
import uuid from '../lib/uuid.js';

interface IResponse {
	res: Response;
	json?: any;
	err?: Error;
}

export default ({params}: {params: any}) => {
	const {apikey, url} = useCliContext();
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
		type: 'track',
		event: eventId,
		messageId: uuid(),
		properties: {
			userId,
			...properties,
		},
	};
	const request = {
		method: 'POST',
		url: '/inbound/courier',
		body: payload,
	};

	useEffect(() => {
		api(request, url, apikey!).then(res => setResp(res));
	}, []);

	return (
		<>
			<Request request={request} response={resp} />
			<Response response={resp} />
		</>
	);
};
