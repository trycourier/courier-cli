import React, {useEffect, useState} from 'react';
import UhOh from '../components/UhOh.js';
import Request from '../components/Request.js';
import Response from '../components/Response.js';
import api from '../lib/api.js';
import {useCliContext} from '../components/Context.js';
import get from 'lodash/get.js';

interface IResponse {
	res: Response;
	json?: any;
	err?: Error;
}

export default ({params}: {params: any}) => {
	const {apikey, url, parsedParams} = useCliContext();
	const get_tenants = Boolean(get(parsedParams, ['tenants']));
	const [resp, setResp] = useState<IResponse | undefined>();
	const [tenants, setTenants] = useState<IResponse | undefined>();

	const userId = params?._?.[0];
	if (!userId) {
		return <UhOh text="You must specify a user ID." />;
	}

	const request = {
		method: 'GET',
		url: `/profiles/${userId}`,
	};

	const tenant_request = {
		method: 'GET',
		url: `/users/${userId}/tenants`,
	};

	useEffect(() => {
		api(request, url, apikey!).then(res => setResp(res));
		if (get_tenants) {
			api(tenant_request, url, apikey!).then(res => setTenants(res));
		}
	}, []);

	return (
		<>
			<Request request={request} response={resp} />
			<Response response={resp} />
			{get_tenants ? (
				<>
					<Request request={tenant_request} response={tenants} />
					<Response response={tenants} />
				</>
			) : (
				<></>
			)}
		</>
	);
};
