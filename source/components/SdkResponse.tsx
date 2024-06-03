import React from 'react';
import {Box, Text} from 'ink';
import {CourierError} from '@trycourier/courier';

type IProps<T> = {
	error?: CourierError;
	response?: T;
};

const Spacing = () => <Text> </Text>;

const SdkResponse = <T extends object>({error, response}: IProps<T>) => {
	const status_code = error ? error.statusCode : '200';
	return (
		<Box flexDirection="column">
			<Box
				flexDirection="column"
				borderStyle="bold"
				borderColor={error ? 'red' : 'green'}
			>
				<Text>
					<Spacing />
					HTTP
					<Spacing />
					<Text bold={true}>{status_code}</Text>
				</Text>
				{error?.message?.length && (
					<>
						<Text color={'red'}>{error.message}</Text>
						<Text>{error.stack}</Text>
					</>
				)}
				{response && Object.keys(response).length > 0 ? (
					<>
						<Text>{JSON.stringify(response, null, 2)}</Text>
					</>
				) : (
					<>
						<Text color="grey">No response body</Text>
					</>
				)}
			</Box>
		</Box>
	);
};

export default SdkResponse;
