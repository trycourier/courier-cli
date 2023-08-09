import React, {useEffect, useState} from 'react';
import UhOh from '../components/UhOh.js';
import Request from '../components/Request.js';
import Response from '../components/Response.js';
import api from '../lib/api.js';

interface IResponse {
	res: Response;
	json?: any;
	err?: Error;
}

export default ({params}: {params: any}) => {
	const [resp, setResp] = useState<IResponse | undefined>();

	const userId = params?._?.[0];
	if (!userId) {
		return <UhOh text="You must specify a user ID." />;
	}

	const request = {
		method: 'GET',
		url: `/profiles/${userId}`,
	};

	useEffect(() => {
		api(request).then(res => setResp(res));
	}, []);

	return (
		<>
			<Request request={request} response={resp} />
			<Response response={resp} />
		</>
	);
};
