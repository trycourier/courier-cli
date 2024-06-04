import React from 'react';
import {Box, Text} from 'ink';
import {CourierError} from '@trycourier/courier';

type IProps<T> = {
	error?: CourierError;
	response?: T | string;
	content?: React.ReactNode;
};

const Spacing = () => <Text> </Text>;

const SdkResponse = <T extends object>({
	error,
	response,
	content,
}: IProps<T>) => {
	const status_code = error ? error.statusCode : '200';
	const response_text =
		typeof response === 'string' ? response : JSON.stringify(response, null, 2);

	return (
		<Box flexDirection="column">
			<Box
				flexDirection="column"
				borderStyle="bold"
				borderColor={error ? 'red' : 'green'}
			>
				{content ? (
					content
				) : (
					<>
						<Text>
							<Spacing />
							HTTP
							<Spacing />
							<Text bold={true}>{status_code}</Text>
						</Text>
						{}
						{error?.message?.length && (
							<>
								<Text color={'red'}>{error.message}</Text>
								<Text>{error.stack}</Text>
							</>
						)}
						{response && Object.keys(response).length > 0 ? (
							<>
								<Text>{response_text}</Text>
							</>
						) : (
							<>
								<Text color="grey">No response body</Text>
							</>
						)}
					</>
				)}
			</Box>
		</Box>
	);
};

export default SdkResponse;
