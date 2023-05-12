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
	res: Response;
	json?: any;
	err?: Error;
}

type Props = {
	request: IRequest;
	response?: IResponse;
	hideResponse?: boolean;
	spinnerText?: string;
};

export default (props: Props) => {
	return (
		<Box flexDirection="column">
			<Box borderStyle="bold" borderColor="blue">
				<Text>
					<Text bold={true}>{props.request.method}</Text>{' '}
					https://api.courier.com{props.request.url}
				</Text>
			</Box>
			{props.request.body && (
				<Text>{JSON.stringify(props.request.body, undefined, '  ')}</Text>
			)}
			{!props.request.body ? <Text color="gray">No request body</Text> : null}
			{!props.response ? (
				<Spinner text={props.spinnerText || defaultSpinnerText} />
			) : null}
			{props.response ? (
				<Box
					borderStyle="bold"
					borderColor={props.response?.err ? 'red' : 'green'}
				>
					<Text>
						HTTP{' '}
						<Text bold={true}>
							{props.response.err
								? props.response.err.message
								: props.response?.res?.status}
						</Text>
					</Text>
				</Box>
			) : null}
			{props.response && !props.hideResponse ? (
				props.response.json ? (
					<Text>{JSON.stringify(props.response.json, undefined, '  ')}</Text>
				) : (
					<Text color="gray">No response body</Text>
				)
			) : null}
		</Box>
	);
};
