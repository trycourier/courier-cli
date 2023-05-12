import React, {useEffect, useState} from 'react';
import {Box} from 'ink';
import KVP from '../components/KVP.js';
import Request from '../components/Request.js';
import api from '../lib/api.js';
// import constants from '../constants.js';

interface IDebug {
	environment: string;
	scope: string;
	tenantId: string;
	tenantName: string;
}

interface IResponse {
	res: Response;
	json?: IDebug;
	err?: Error;
}

export default () => {
	const [resp, setResp] = useState<IResponse | undefined>();

	const request = {
		url: '/debug',
		method: 'POST',
	};

	useEffect(() => {
		api(request).then(res => setResp(res));
	}, []);

	return (
		<Box flexDirection="column">
			<Request request={request} response={resp} hideResponse={true} />
			{resp && resp.json ? (
				<>
					<KVP width={20} label="Workspace Name" value={resp.json.tenantName} />
					<KVP width={20} label="Workspace ID" value={resp.json.tenantId} />
					<KVP
						width={20}
						label="API Key Environment"
						value={resp.json.environment}
					/>
					<KVP width={20} label="API Key Scope" value={resp.json.scope} />
				</>
			) : null}
		</Box>
	);
};
