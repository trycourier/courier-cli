import {Box, Text} from 'ink';
import _ from 'lodash';
import React, {useEffect, useState} from 'react';
import {useBoolean} from 'usehooks-ts';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';

interface IParams {
	_: string[];
	scopes?: string;
	expiration?: string | number;
	all?: boolean;
}

const VALID_SCOPE_PREFIXES = [
	'read:messages',
	'read:user-tokens',
	'write:user-tokens',
	'read:brands',
	'write:brands',
	'inbox:read:messages',
	'inbox:write:events',
	'read:preferences',
	'write:preferences',
];

const ALL = VALID_SCOPE_PREFIXES.filter(s => !s.endsWith('brand'));

const UserToken = () => {
	const {parsedParams, getJWT} = useCliContext();
	const running = useBoolean(true);
	const [jwt, setJWT] = useState<string>();
	const [final_scopes, setFinalScopes] = useState<string[]>([]);
	const [error, setError] = useState<string | undefined>();

	const {
		expiration,
		scopes,
		all,
		_: [user_id],
	} = parsedParams as IParams;

	useEffect(() => {
		getUserJWT();
	}, []);

	const getUserJWT = async () => {
		const exp = Number(expiration);

		const scope_input = scopes?.split(',') || [];
		const sc = [...new Set([...scope_input, ...(all ? ALL : [])])];
		const invalid_scopes = sc.filter(
			scope => !_.some(VALID_SCOPE_PREFIXES, val => scope.startsWith(val)),
		);
		setFinalScopes([`user_id:${user_id}`, ...sc]);
		if (!user_id) {
			setError('No User Specified');
		} else if (!(sc && sc.length)) {
			setError('No scopes provided. They are required');
			return;
		} else if (invalid_scopes.length) {
			setError(`Found invalid scopes (${invalid_scopes.join(', ')})`);
		} else if (exp && Number.isNaN(exp)) {
			setError('Not a valid number');
		} else {
			try {
				const r = await getJWT(user_id, sc as TJWTScope[], {
					expires_in: `${exp || 5} mins`,
					write_brands: [],
				});
				setJWT(r.token);
			} catch (e) {
				setError(String(e));
			}
		}
		running.setFalse();
	};

	if (error?.length) {
		return <UhOh text={error} />;
	} else if (running.value) {
		return <Spinner text={`Fetching JWT`} />;
	} else {
		return (
			<>
				<Text>Token has the following scopes:</Text>
				<Text>{final_scopes.join(' ')}</Text>
				<Box
					flexDirection="column"
					marginY={1}
					borderColor="gray"
					borderStyle={'single'}
					borderTop={false}
					borderLeft={false}
					borderRight={false}
				></Box>
				<Text>{jwt}</Text>
				<Box
					flexDirection="column"
					marginY={1}
					borderColor="gray"
					borderStyle={'single'}
					borderBottom={false}
					borderLeft={false}
					borderRight={false}
				></Box>
			</>
		);
	}
};

export default UserToken;
