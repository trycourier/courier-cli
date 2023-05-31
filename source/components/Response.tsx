import React from 'react';
import {Box, Text} from 'ink';

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
	response?: IResponse;
};

function getContentType(res: IResponse): string {
	return res.res.headers?.get('content-type') || 'none';
}

export default (props: Props) => {
	if (!props.response) return null;
	return (
		<Box flexDirection="column">
			<Box
				flexDirection="column"
				borderStyle="bold"
				borderColor={props.response?.err ? 'red' : 'green'}
			>
				<Text>
					{' '}
					HTTP{' '}
					<Text bold={true}>
						{props.response.err
							? props.response.err.message
							: props.response?.res?.status}
					</Text>{' '}
					{!props.response.err ? (
						<>
							<Text>• {props.response.res.statusText}</Text>
						</>
					) : null}
				</Text>
				{!props.response.err &&
				getContentType(props.response) !== 'application/json' ? (
					<Box>
						<Text> Content-Type: {getContentType(props.response)}</Text>
					</Box>
				) : null}
			</Box>
			{props.response.json ? (
				<Text>{JSON.stringify(props.response.json, undefined, '  ')}</Text>
			) : (
				<Text color="gray">No response body</Text>
			)}
		</Box>
	);
};
