import React, {useEffect, useState} from 'react';
import fs from 'fs';
import UhOh from '../components/UhOh.js';
import Request from '../components/Request.js';
import Response from '../components/Response.js';
import Elemental from '../components/Elemental.js';
import api from '../lib/api.js';
import {Box, Text} from 'ink';
import {useCliContext} from '../components/Context.js';
import lodash from 'lodash';

type TRecipient =
	| {
			user_id: string;
	  }
	| {
			list_id: string;
	  }
	| {
			audience_id: string;
	  }
	| {
			tenant_id: string;
	  }
	| {
			include_children: boolean;
	  }
	| {
			context: {
				tenant_id: string;
			};
	  }
	| {
			email: string;
	  }
	| {
			phone_number: string;
	  }
	| {
			apn: {
				token: string;
			};
	  }
	| {
			firebaseToken: string;
	  };

type TElemental = any;

interface IPayloadElemental {
	type: 'elemental';
	message: {
		to?: TRecipient | TRecipient[];
		content:
			| {
					title?: string;
					body?: string;
			  }
			| TElemental;
		routing?: {
			method: string;
			channels: string[];
		};
		data?: any;
	};
}

interface IPayloadTemplate {
	type: 'template';
	message: {
		to?: TRecipient | TRecipient[];
		template: string;
		routing?: {
			method: string;
			channels: string[];
		};
		data?: any;
	};
}

type Params = {
	_?: string[];
	user?: string;
	list?: string;
	audience?: string;
	tenant?: string;
	email?: string;
	tel?: string;
	apn?: string;
	fcm?: string;
	title?: string;
	body?: string;
	message?: string;
	template?: string;
	channel?: string;
	channels?: string;
	all?: boolean;
	elemental?: string;
	mock?: boolean;
	'include-children'?: boolean;
	includeChildren?: boolean;
	'tenant-context'?: string;
	tenantContext?: string;
};

const constructPayload = (
	params: Params,
): IPayloadElemental | IPayloadTemplate => {
	const to: TRecipient[] = [];
	const {
		_,
		user,
		list,
		audience,
		email,
		tel,
		apn,
		fcm,
		title,
		body,
		message,
		template,
		channel,
		channels,
		all,
		elemental,
		mock,
		tenant,
		'include-children': includeChildren,
		'tenant-context': tenantContext,
		includeChildren: _includeChildren,
		tenantContext: _tenantContext,
		...data
	} = params;
	if (user) {
		to.push({user_id: String(user)});
	}
	if (list) {
		to.push({list_id: list});
	}
	if (audience) {
		to.push({audience_id: audience});
	}
	if (tenant) {
		to.push({tenant_id: tenant});
		if (includeChildren) {
			to.push({include_children: true});
		}
	}
	if (email) {
		to.push({email: email});
	}
	if (tel) {
		to.push({phone_number: tel});
	}
	if (apn) {
		to.push({apn: {token: apn}});
	}
	if (fcm) {
		to.push({firebaseToken: fcm});
	}
	if (to.length && tenantContext) {
		to.forEach((_t, i) => {
			lodash.set(to, [i, 'context', 'tenant_id'], tenantContext);
		});
	}

	let contentElemental: {title?: string; body?: string} = {
		title: undefined,
		body: undefined,
	};
	if (title) {
		contentElemental.title = title;
	}
	if (body) {
		contentElemental.body = body;
	} else if (message) {
		contentElemental.body = message;
	}

	let routing: {channels: string[]; method: string} = {
		channels: [],
		method: params.all ? 'all' : 'single',
	};
	if (channel) {
		routing.channels = channel.split(',');
	} else if (channels) {
		routing.channels = channels.split(',');
	}
	if (email && email.length) {
		routing.channels.push('email');
	}
	if (tel && tel.length) {
		routing.channels.push('sms');
	}
	if ((apn && apn.length) || (fcm && fcm.length)) {
		routing.channels.push('push');
	}

	if (template) {
		return {
			type: 'template',
			message: {
				to: to.length === 1 ? to[0] : to,
				template,
				routing: routing.channels.length ? routing : undefined,
				data: data ? data : undefined,
			},
		};
	} else {
		return {
			type: 'elemental',
			message: {
				to: to.length === 1 ? to[0] : to,
				content: contentElemental,
				routing: routing.channels.length ? routing : undefined,
				data: data ? data : undefined,
			},
		};
	}
};

interface IResponse {
	res: IHttpResponse;
	json?: any;
	err?: Error;
}

interface IHttpResponse {
	status: number;
	statusText: string;
}

export default ({params}: {params: any}) => {
	const {apikey, url} = useCliContext();
	const [resp, setResp] = useState<IResponse | undefined>();

	if (!params.body && !params.template && !params.elemental) {
		return (
			<UhOh text="You must specify a message body, template, or Elemental file path." />
		);
	}

	let payload: IPayloadElemental | IPayloadTemplate = constructPayload(params);
	if (Array.isArray(payload.message.to) && !payload.message.to.length) {
		return <UhOh text="You must specify a recipient." />;
	}
	if (
		payload.type === 'elemental' &&
		!params.elemental &&
		(!payload.message.content.body || !payload.message.content.body?.length)
	) {
		throw new Error('You must specify a body for the message.');
	}

	if (params.elemental && payload.type === 'elemental') {
		if (!fs.existsSync(params.elemental)) {
			throw new Error('Invalid file path to Elemental document.');
		}
		payload.message.content = JSON.parse(
			fs.readFileSync(params.elemental, 'utf8'),
		);
	}

	const request = {
		method: 'POST',
		url: '/send',
		body: {
			message: payload.message,
		},
	};

	useEffect(() => {
		if (params.mock) {
			setResp({res: {status: 999, statusText: 'MOCKED'}});
			return;
		}
		api(request, url, apikey!).then(res => setResp(res));
	}, []);

	return (
		<Box flexDirection="column">
			{params.elemental ? (
				<>
					<Box borderStyle="bold" borderColor="white">
						<Text> Elemental</Text>
					</Box>
					<Box>
						{payload.type === 'elemental' ? (
							<Elemental elemental={payload.message.content} />
						) : null}
					</Box>
				</>
			) : null}
			<>
				<Request request={request} response={resp} />
				<Response response={resp} />
			</>
		</Box>
	);
};
