import {Alert} from '@inkjs/ui';
import React, {useEffect, useState} from 'react';
import {useBoolean} from 'usehooks-ts';
import {useCliContext} from '../components/Context.js';
import Spinner from '../components/Spinner.js';
import {Box} from 'ink';
import {ExecaError, execa} from 'execa';
import constants from '../constants.js';

export default () => {
	const {version} = useCliContext();
	const running = useBoolean(true);
	const [error, setError] = useState<string | undefined>();

	useEffect(() => {
		if (version.latest) {
			upgrade();
		}
	}, [version.latest]);

	const upgrade = async () => {
		try {
			await execa(
				'npm',
				['install', '-g', `${constants.package_name}`, '--', 'latest'],
				{
					shell: true,
				},
			);
		} catch (e) {
			setError((e as ExecaError).message);
		} finally {
			running.setFalse();
		}
	};

	if (running.value) {
		return <Spinner text={`Upgrading Courier CLI ${version.latest}`} />;
	} else {
		const text = `Courier CLI upgraded to ${version.latest}`;
		return (
			<Box width={error ? undefined : text.length + 8}>
				{error?.length ? (
					<Alert variant="error" title={'Error'}>
						{error}
					</Alert>
				) : (
					<Alert variant="success" title={'Upgraded'}>
						{text}
					</Alert>
				)}
			</Box>
		);
	}
};
