import React, {useEffect, useState} from 'react';
import {Box, Text} from 'ink';
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
};

interface IPayload {
  message: {
    to?: TRecipient | TRecipient[];
    content: {
      title?: string;
      body?: string;
    };
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
  title?: string;
  body?: string;
  message?: string;
  channel?: string;
  channels?: string[];
  route?: string;
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
    to.push({ phone_number: params.tel.toString()})
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

  let routing = undefined;
  if (params.channel) {
    routing = {
      method: 'single',
      channels: [params.channel]
    }
  }
  if (params.channels) {
    routing = {
      method: params.route || 'all',
      channels: [params.channels]
    }
  }
  if (!routing && params.email && params.email.length) {
    routing = {
      method: 'single',
      channels: ['email']
    }
  }
  if (!routing && params.tel && params.tel.length) {
    routing = {
      method: 'single',
      channels: ['sms']
    }
  }

  const {_, user, list, audience, email, tel, title, body,
    ...data} = params;  

	return {
    message: {
      to: to.length === 1 ? to[0] : to,
      content,
      data
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
	if (!payload.message.content.body || !payload.message.content.body?.length) {
		return <UhOh text="You must specify a body for the message." />;
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
