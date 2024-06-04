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
				<>
					<KVP width={20} label="Workspace Name" value={resp.json.tenantName} />
					<KVP width={20} label="Workspace ID" value={resp.json.tenantId} />
					<KVP
						width={20}
						label="API Key Environment"
						value={resp.json.environment}
					/>
					<KVP width={20} label="API Key Scope" value={resp.json.scope} />
					{resp.json.mock && (
						<KVP
							width={20}
							label="API Key Simulated (Mock)"
							value={resp.json.scope}
						/>
					)}
				</>
			) : null}
		</Box>
	);
};
