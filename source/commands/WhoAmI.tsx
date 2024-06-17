import React, {useEffect, useState} from 'react';
import {Box} from 'ink';
import KVP from '../components/KVP.js';
import Request from '../components/Request.js';
import Response from '../components/Response.js';
import api from '../lib/api.js';
import {useCliContext} from '../components/Context.js';
// import constants from '../constants.js';

export default () => {
	const {apikey, url} = useCliContext();
	const [resp, setResp] = useState<IResponseDebug | undefined>();

	const request = {
		url: '/debug',
		method: 'POST',
	};

	useEffect(() => {
		api(request, url, apikey!).then(res => setResp(res));
	}, []);

	return (
		<Box flexDirection="column" width="100%">
			<Request request={request} response={resp} />
			{resp && resp.err ? (
				<Response response={resp} />
			) : resp && resp.json ? (
				<Box flexDirection="column" width="100%">
					<KVP label="Workspace Name" value={resp.json.tenantName} />
					<KVP label="Workspace ID" value={resp.json.tenantId} />
					<KVP label="API Key Environment" value={resp.json.environment} />
					<KVP label="API Key Scope" value={resp.json.scope} />
					{resp.json.mock && (
						<KVP label="API Key Simulated (Mock)" value={resp.json.scope} />
					)}
				</Box>
			) : null}
		</Box>
	);
};
