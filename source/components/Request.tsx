import React from 'react';
import {Box, Text} from 'ink';
import Spinner from './Spinner.js';

const defaultSpinnerText = 'blipping those bits and blooping those bytes';

interface IRequest {
	url: string;
	method: string;
	body?: object | string;
}

interface IResponse {
	res: IHttpResponse;
	json?: any;
	err?: Error;
}

interface IHttpResponse {
	status: number;
	statusText: string;
	headers?: Headers;
}

type Props = {
	request: IRequest;
	response?: IResponse;
	spinnerText?: string;
};

export default (props: Props) => {
  const url = `${process.env['COURIER_DOMAIN'] || 'https://api.courier.com'}${props.request.url}`;
	return (
		<Box flexDirection="column">
			<Box borderStyle="bold" borderColor="blue">
				<Text>
					{' '}
					<Text bold={true}>{props.request.method}</Text>{' '}
          {url}
				</Text>
			</Box>
			{props.request.body ? (
				typeof props.request.body === 'string' ? (
					<Text>{props.request.body}</Text>
				) : (
					<Text>{JSON.stringify(props.request.body, undefined, '  ')}</Text>
				)
			) : null}
			{!props.request.body ? <Text color="gray">No request body</Text> : null}
			{!props.response ? (
				<Spinner text={props.spinnerText || defaultSpinnerText} />
			) : null}
		</Box>
	);
};
