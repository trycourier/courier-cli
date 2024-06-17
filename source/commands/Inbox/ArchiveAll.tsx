import {IGetInboxMessagesParams, Inbox} from '@trycourier/client-graphql';
import {Box, Text} from 'ink';
import React, {useEffect, useState} from 'react';
import {useBoolean, useCounter} from 'usehooks-ts';
import {useCliContext} from '../../components/Context.js';
import Spinner from '../../components/Spinner.js';
import UhOh from '../../components/UhOh.js';
import uuid from '../../lib/uuid.js';
import {IInboxMessagePreview} from '@trycourier/core';
// @ts-ignore
import ms, {StringValue} from 'ms';
import _ from 'lodash';

interface IArchiveAll {
	user_id_override?: string;
}

interface IParam {
	_: string[];
	tenant?: string | number;
	tag?: string | string[];
	before?: StringValue;
	batchSize?: string | number;
	includePinned?: boolean;
}

const LIMIT = 10;

const ArchiveAll = ({user_id_override}: IArchiveAll) => {
	const [error, setError] = useState<string | undefined>();
	const [InboxClient, setInboxClient] = useState<ReturnType<typeof Inbox>>();
	const [jwt, setJwt] = useState<string>('');
	const total_messages = useCounter(0);
	const running = useBoolean(true);
	const gathering_messages = useBoolean(true);
	const pages = useCounter(0);
	const {getJWT, parsedParams} = useCliContext();
	const archived_messages = useCounter(0);
	const [messages, setMessages] = useState<IInboxMessagePreview[]>([]);
	const {
		tenant,
		tag,
		before,
		batchSize,
		includePinned,
		_: [userId, ...args],
	} = parsedParams as IParam;
	const user_id = user_id_override || userId;
	const tenant_id = tenant ? String(tenant) : undefined;

	const until = before
		? new Date(new Date().getTime() - Math.abs(ms(before)))
		: undefined;

	const limit = Number(batchSize) || LIMIT;

	const handleError = (text: string) => {
		setError(text + '\n' + JSON.stringify({tenant, tag, user_id, args}));
	};

	useEffect(() => {
		getInboxClient();
	}, [user_id]);

	useEffect(() => {
		if (jwt?.length && InboxClient) {
			getAllMessages();
		}
	}, [jwt, InboxClient]);

	useEffect(() => {
		if (!gathering_messages.value) {
			handleMessageArchive();
		}
	}, [gathering_messages.value]);

	useEffect(() => {
		if (error?.length) {
			running.setFalse();
		}
	}, [error]);

	const getInboxClient = async () => {
		try {
			if (user_id) {
				const jwt = await getJWT(user_id, [
					'inbox:read:messages',
					'inbox:write:events',
				]);
				if (jwt.token) {
					setJwt(jwt.token);
					setInboxClient(
						Inbox({
							authorization: jwt.token,
							userId: user_id,
							clientSourceId: uuid(),
						}),
					);
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

	const handleMessageArchive = async () => {
		if (messages.length) {
			for (var i = 0; i < messages.length; i++) {
				try {
					const message = messages[i];
					if (message?.messageId) {
						await InboxClient?.markArchive(message.messageId);
					}
					archived_messages.increment();
				} catch (e) {
					handleError(
						`Failed to archive message ${
							messages[i]?.messageId || 'unknown'
						}: ${
							e instanceof Error ? e.message : String(e) || 'An error occurred.'
						}`,
					);
				}
			}
		}
		running.setFalse();
	};

	const getAllMessages = async () => {
		let cursor: string | undefined;
		let counter = -1;
		if (jwt?.length && InboxClient) {
			let params: IGetInboxMessagesParams = {
				archived: false,
				limit,
			};
			if (tag) {
				params.tags = Array.isArray(tag) ? tag : [tag];
			}
			if (tenant_id) {
				params.tenantId = tenant_id;
			}
			while (counter < 0) {
				counter++;
				pages.increment();
				let filtered_messages: IInboxMessagePreview[] = [];
				const r = await InboxClient.getMessages(params, cursor);
				_.get(r, 'messages', []).forEach(m => {
					if (until) {
						if (new Date(m.created) < until) {
							filtered_messages.push(m);
						}
					} else {
						filtered_messages.push(m);
					}
				});
				// handle pins on first iteration, ignore after
				if (counter === 0 && Boolean(includePinned)) {
					(_.get(r, 'pinned', []) as IInboxMessagePreview[]).forEach(m => {
						if (until) {
							if (new Date(m.created) < until) {
								filtered_messages.push(m);
							}
						} else {
							filtered_messages.push(m);
						}
					});
				}
				if (filtered_messages.length) {
					total_messages.setCount(c => c + filtered_messages.length);
					setMessages(p => [...p, ...filtered_messages]);
				}
				if (!r?.startCursor) {
					break;
				} else {
					cursor = r.startCursor;
				}
			}
		} else {
			handleError('No JWT generated, but GQL client called.');
		}
		gathering_messages.setFalse();
	};

	if (error?.length) {
		return <UhOh text={error} />;
	} else if (gathering_messages.value) {
		return (
			<Spinner
				text={`Finding messages: Found ${messages.length} in ${pages.count} pages `}
			/>
		);
	} else if (running.value) {
		return (
			<Spinner
				text={`Archiving ${archived_messages.count} / ${messages.length} for ${user_id}`}
			/>
		);
	} else {
		return (
			<Box>
				<Text>
					Marked {archived_messages.count} messages as archived for {user_id}
				</Text>
			</Box>
		);
	}
};

export default ArchiveAll;
