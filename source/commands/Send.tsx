import React, {useEffect, useState} from 'react';
import {Box, Text} from 'ink';
import fs from 'fs';
import UhOh from '../components/UhOh.js';
import Request from '../components/Request.js';
import api from '../lib/api.js';

type TRecipient = {
  user_id: string;
} | {
  list_id: string;
} | {
  audience_id: string;
} | {
  email: string;
} | {
  phone_number: string;
} | {
  apn: {
    token: string;
  };
} | {
  firebaseToken: string;
};

type TElemental = any;

interface IPayload {
  message: {
    to?: TRecipient | TRecipient[];
    content: {
      title?: string;
      body?: string;
    } | TElemental;
    routing?: {
      method: string;
      channels: string[];
    };
    data?: any;
  }
}

type Params = {
  _?: string[];
  user?: string;
  list?: string;
  audience?: string;
  email?: string;
  tel?: string;
  apn?: string;
  fcm?: string;
  title?: string;
  body?: string;
  message?: string;
  channel?: string;
  channels?: string;
  all?: boolean;
  elemental?: string;
}

const constructPayload = (params: Params): IPayload => {
  const to: TRecipient[] = [];
  if (params.user) {
    to.push({ user_id: params.user })
  }
  if (params.list) {
    to.push({ list_id: params.list })
  }
  if (params.audience) {
    to.push({ audience_id: params.audience})
  }
  if (params.email) {
    to.push({ email: params.email})
  }
  if (params.tel) {
    to.push({ phone_number: params.tel})
  }
  if (params.apn) {
    to.push({ apn: { token: params.apn } })
  }
  if (params.fcm) {
    to.push({ firebaseToken: params.fcm })
  }

  let content : { title?: string, body?: string } = {
    title: undefined,
    body: undefined
  };
  if (params.title) {
    content.title = params.title;
  }
  if (params.body) {
    content.body = params.body;
  } else if (params.message) {
    content.body = params.message
  }

  let routing : { channels: string[], method: string } = {
    channels: [],
    method: params.all ? 'all' : 'single'
  };
  if (params.channel) {
    routing.channels = params.channel.split(',')
  } else if (params.channels) {
    routing.channels = params.channels.split(',')
  }
  if (params.email && params.email.length) {
    routing.channels.push('email')
  }
  if (params.tel && params.tel.length) {
    routing.channels.push('sms')
  }
  if ((params.apn && params.apn.length) || (params.fcm && params.fcm.length)) {
    routing.channels.push('push')
  }

  const {_, user, list, audience, email, tel, apn, fcm, title, body, elemental,
    ...data} = params;  

	return {
    message: {
      to: to.length === 1 ? to[0] : to,
      content,
      routing: routing.channels.length ? routing : undefined,
      data: data ? data : undefined
    }
	};
}

interface IResponse {
	res: Response,
	json?: any,
	err?: Error
}

export default ({params}: {params: any}) => {
	const [resp, setResp] = useState<IResponse | undefined>();

  const payload = constructPayload(params);
	if (!payload.message.to) {
		return <UhOh text="You must specify a recipient." />;
	}
	if (!params.elemental && (!payload.message.content.body || !payload.message.content.body?.length)) {
		return <UhOh text="You must specify a body for the message." />;
	}
  if (params.elemental) {
    if (!fs.existsSync(params.elemental)) {
      return <UhOh text="Invalid file path to Elemental document." />;
    }
    payload.message.content = JSON.parse(fs.readFileSync(params.elemental, 'utf8'));
  }

  const request = {
    method: 'POST',
    url: '/send',
    body: payload
  };

	useEffect(() => {
		api(request).then((res) => setResp(res));
	}, []);

	return <Box flexDirection='column'>
		<Request request={request} response={resp} />
		{resp && resp.json ? <>
      <Text>{JSON.stringify(resp.json, undefined, '  ')}</Text>
		</> : null}
	</Box>;
};
