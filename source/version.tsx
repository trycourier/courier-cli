import {execa} from 'execa';
import {Box, Text} from 'ink';
import _ from 'lodash';
import React, {useEffect} from 'react';
import {useCliContext} from './components/Context.js';
import constants from './constants.js';
import VERSION from './version.js';

const Version = async () => {
	const {version, setVersion} = useCliContext();
	useEffect(() => {
		getVersion();
	}, []);

	const getVersion = async () => {
		if (VERSION !== 'local') {
			try {
				const exc = await execa(
					'npm',
					['-g', 'outdated', constants.package_name, '--json'],
					{
						shell: true,
						reject: false,
					},
				);
				const stdout = JSON.parse(exc.stdout);
				setVersion({
					current: _.get(stdout, [constants.package_name, 'current']),
					latest: _.get(stdout, [constants.package_name, 'latest']),
				});
			} catch (e) {
				console.log(e);
			}
		}
	};

	const version_text =
		version.current && version.latest && version.current !== version.latest
			? `Upgrade available (${version.current} > ${version.latest}), run courier upgrade`
			: undefined;

	if (version_text?.length) {
		return (
			<Box
				flexDirection="column"
				marginY={1}
				flexShrink={1}
				width={version_text.length + 2}
				borderColor="gray"
				borderStyle={'single'}
			>
				<Text>
					Upgrade available ({version.current} {'>'} {version.latest}), run{' '}
					<Text color="green">courier upgrade</Text>
				</Text>
			</Box>
		);
	} else {
		return <></>;
	}
};
export default Version;
