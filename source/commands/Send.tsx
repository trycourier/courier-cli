import React, {useEffect, useState} from 'react';
import {Box, Text} from 'ink';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';
import api from '../lib/api.js';

const constructPayload = (params: any) => {
  const to = [];
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
	if (!to.length) {
		return <UhOh text="You must specify a recipient." />;
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
  }
	if (!content.body || !content.body?.length) {
		return <UhOh text="You must specify a body for the message." />;
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

export default ({params}: {params: any}) => {
	const [resp, setResp] = useState<object | undefined>();
	const [error, setError] = useState<string | undefined>();
  const payload = constructPayload(params);

	useEffect(() => {
		api('/send', 'POST', {
			body: JSON.stringify(payload),
		}).then(
			({json}) => setResp(json),
			(err: Error) => setError(err.message)
		);
	}, []);

	if (error) {
		return (
			<Box flexDirection='column'>
				<UhOh text={error} />
				<Text>{JSON.stringify(payload, undefined, '  ')}</Text>
			</Box>
		);
	} else if (resp) {
		return (
			<Box flexDirection='column'>
        <Text color="green">Success 🎉</Text>
				<Text>{JSON.stringify(resp, undefined, '  ')}</Text>
			</Box>
    )
	} else {
		return (
			<Box flexDirection='column'>
				<Spinner text="air-mailing those bits & bytes..." />
				<Text>{JSON.stringify(payload, undefined, '  ')}</Text>
			</Box>
		);
	}
};
