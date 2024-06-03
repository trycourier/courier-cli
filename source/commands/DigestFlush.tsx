import React, {useEffect, useState} from 'react';
import UhOh from '../components/UhOh.js';
import Request from '../components/Request.js';
import Response from '../components/Response.js';
import api from '../lib/api.js';
import {useCliContext} from '../components/Context.js';

interface IResponse {
	res: Response;
	json?: any;
	err?: Error;
}

export default ({params}: {params: any}) => {
	const {apikey, url} = useCliContext();
	const [resp, setResp] = useState<IResponse | undefined>();

	const userId = params?._?.[0];
	if (!userId) {
		return (
			<UhOh text="You must specify the ID of the user whose digest you want to flush." />
		);
	}

	const topicId = params?._?.[1];
	if (!topicId) {
		return (
			<UhOh text="You must specify the ID of the Digest-enabled Subscription Topic you want to flush for the given user." />
		);
	}

	const request = {
		method: 'POST',
		url: `/users/${userId}/preferences/${topicId}/flush`,
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
