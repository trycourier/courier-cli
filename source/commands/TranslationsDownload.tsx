import React, {useEffect, useState} from 'react';
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

	const request = {
		method: 'GET',
		url: `/translations/default/${locale}`,
	};

	useEffect(() => {
		api(request).then(res => {
			res.res.headers.get('content-type');
			setResp(res);
		});
	}, []);

	return <Request request={request} response={resp} />;
};
