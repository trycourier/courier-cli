#!/usr/bin/env node
import React, {useEffect, useState} from 'react';
import {Box, Text, render} from 'ink';
import args from './lib/args.js';
import loadEnv from './lib/load-env.js';
import {CliContextProvider} from './components/Context.js';
import {execa} from 'execa';
import _ from 'lodash';
import VERSION from './version.js';
import constants from './constants.js';

const CLI = async () => {
	process.removeAllListeners('warning');
	await loadEnv();
	const params = args(process.argv);
	const mappings = (await import('./mappings.js')).default;
	const Router = (await import('./components/Router.js')).default;
	const [version, setVersion] = useState<{current?: string; latest?: string}>(
		{},
	);

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

	render(
		<CliContextProvider args={params} {...{mappings}} {...version}>
			<Router />
			{version_text ? (
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
			) : (
				<></>
			)}
		</CliContextProvider>,
	);
};

(async () => {
	await CLI();
})();
