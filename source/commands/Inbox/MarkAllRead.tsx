import {Inbox} from '@trycourier/client-graphql';
import {Box, Text} from 'ink';
import React, {useEffect, useState} from 'react';
import {useBoolean} from 'usehooks-ts';
import {useCliContext} from '../../components/Context.js';
import Spinner from '../../components/Spinner.js';
import UhOh from '../../components/UhOh.js';
import uuid from '../../lib/uuid.js';

interface IMarkAllRead {}

interface IParam {
	_: string[];
	tenant?: string | number;
	tag?: string | string[];
}

const MarkAllRead = ({}: IMarkAllRead) => {
	const [error, setError] = useState<string | undefined>();
	const [jwt, setJwt] = useState<string>('');
	const [unread, setUnread] = useState<number>();
	const running = useBoolean(true);
	const {getJWT, parsedParams} = useCliContext();
	const [response, setResponse] = useState<any>();
	const {
		tenant,
		tag,
		_: [user_id, ...args],
	} = parsedParams as IParam;

	const handleError = (text: string) => {
		setError(text + '\n' + JSON.stringify({tenant, tag, user_id, args}));
	};

	useEffect(() => {
		runJwt();
	}, [user_id]);

	useEffect(() => {
		if (jwt?.length) {
			runMarkRead();
		}
	}, [jwt]);

	useEffect(() => {
		if (error?.length) {
			running.setFalse();
		} else if (response) {
			running.setFalse();
		}
	}, [error, response]);

	const runJwt = async () => {
		try {
			if (user_id) {
				const jwt = await getJWT(user_id, [
					'inbox:read:messages',
					'inbox:write:events',
				]);
				if (jwt.token) {
					setJwt(jwt.token);
				} else {
					handleError('No JWT token returned.');
				}
			} else {
				handleError('No user ID provided.');
			}
		} catch (e) {
			handleError(
				e instanceof Error ? e.message : String(e) || 'An error occurred.',
			);
		}
	};

	const runMarkRead = async () => {
		if (jwt) {
			const gql_client = Inbox({
				authorization: jwt,
				userId: user_id,
				clientSourceId: uuid(),
			});
			try {
				const res = await gql_client.getInboxCount({
					status: 'unread',
				});
				if (res) {
					if (typeof res.count === 'number') {
						setUnread(res?.count || 0);
						const res2 = await gql_client.markAllRead();
						if (res2) {
							setResponse(res2);
						} else {
							handleError('No response from markAllRead');
						}
					} else {
						handleError('No count returned from getInboxCount');
					}
				} else {
					handleError('No response from getInboxCount');
				}
			} catch (e) {
				handleError(
					e instanceof Error ? e.message : String(e) || 'An error occurred.',
				);
			}
		} else {
			handleError('No JWT generated, but GQL client called.');
		}
	};

	if (error?.length) {
		return <UhOh text={error} />;
	} else if (running.value) {
		return <Spinner text="Marking all read" />;
	} else {
		return (
			<Box>
				<Text>Marked {unread} messages as read</Text>
			</Box>
		);
	}
};

export default MarkAllRead;
