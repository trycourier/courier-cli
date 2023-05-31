import React, {useEffect, useState} from 'react';
import fs from 'fs';
import UhOh from '../components/UhOh.js';
import Request from '../components/Request.js';
import api from '../lib/api.js';

interface IResponse {
	res: Response;
	json?: any;
	err?: Error;
}

export default ({params}: {params: any}) => {
	const [resp, setResp] = useState<IResponse | undefined>();

	const locale = params?._?.[0];
	if (!locale) {
		return <UhOh text="You must specify a locale, e.g. en-US." />;
	}

	const filepath = params?._?.[1];
	if (!filepath) {
		return <UhOh text="You must specify a path to the .PO file." />;
	} else if (!fs.existsSync(filepath)) {
		return <UhOh text="No .PO file found at given path." />;
	}

	const po = fs.readFileSync(filepath, 'utf8');
	const request = {
		method: 'PUT',
		url: `/translations/default/${locale}`,
		headers: {
			'Content-Type': 'text/plain',
		},
		body: po,
	};

	useEffect(() => {
		api(request).then(res => setResp(res));
	}, []);

	return <Request request={request} response={resp} />;
};
