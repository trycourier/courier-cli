import {Text} from 'ink';
import _ from 'lodash';
import React, {useEffect, useState} from 'react';
import {useBoolean} from 'usehooks-ts';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import UhOh from '../components/UhOh.js';
import {stdout} from 'process';

interface IParams {
	_: string[];
	scopes?: string;
	expiration?: string | number;
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

const UserToken = () => {
	const {parsedParams, getJWT} = useCliContext();
	const running = useBoolean(true);
	const [jwt, setJWT] = useState<string>();
	const [error, setError] = useState<string | undefined>();
	const {
		expiration,
		scopes,
		_: [user_id],
	} = parsedParams as IParams;

	useEffect(() => {
		getUserJWT();
	}, []);

	const getUserJWT = async () => {
		const exp = Number(expiration);
		const sc = (scopes?.split(',') || []).map(s => s.trim());
		const invalid_scopes = sc.filter(
			scope => !_.some(VALID_SCOPE_PREFIXES, val => scope.startsWith(val)),
		);

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
		stdout.write(jwt ?? '');
		return <Text>{jwt}</Text>;
	}
};

export default UserToken;
