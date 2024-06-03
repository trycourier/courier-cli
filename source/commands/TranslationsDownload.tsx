import React, {useEffect, useState} from 'react';
import {Box, Text} from 'ink';
import UhOh from '../components/UhOh.js';
import Request from '../components/Request.js';
import Response from '../components/Response.js';
import api from '../lib/api.js';
import {useCliContext} from '../components/Context.js';

interface IResponse {
	res: Response;
	json?: any;
	text?: string;
	err?: Error;
}

type Params = {
	_?: string[];
	text?: boolean;
	domain?: string;
};

export default ({params}: {params: Params}) => {
	const {apikey, url} = useCliContext();
	const [resp, setResp] = useState<IResponse | undefined>();

	const locale = params?._?.[0];
	if (!locale) {
		return <UhOh text="You must specify a locale, e.g. en-US." />;
	}

	const isText = params.text || false;

	const request = {
		method: 'GET',
		url: `/translations/${params.domain || 'default'}/${locale}`,
	};

	useEffect(() => {
		api(request, url, apikey!).then(res => {
			res.res.headers.get('content-type');
			setResp(res);
		});
	}, []);

	return isText ? (
		<Box>
			<Text>{resp?.text}</Text>
		</Box>
	) : (
		<>
			<Request request={request} response={resp} />
			<Response response={resp} />
		</>
	);
};
