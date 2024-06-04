#!/usr/bin/env node
import React from 'react';
import {Box, Text, render} from 'ink';
import args from './lib/args.js';
import loadEnv from './lib/load-env.js';
import {CliContextProvider} from './components/Context.js';
import {execa} from 'execa';
import _ from 'lodash';
import VERSION from './version.js';
import constants from './constants.js';
2;

const CLI = async () => {
	process.removeAllListeners('warning');
	await loadEnv();
	const params = args(process.argv);
	const mappings = (await import('./mappings.js')).default;
	const Router = (await import('./components/Router.js')).default;
	let current: string | undefined;
	let latest: string | undefined;

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
			current = _.get(stdout, ['npm', 'current']);
			latest = _.get(stdout, ['npm', 'latest']);
		} catch (e) {
			console.log(e);
		}
	}

	const version_text = `Upgrade available (${current} > ${latest}), run courier upgrade`;

	render(
		<CliContextProvider args={params} {...{mappings, current, latest}}>
			<Router />
			{current && latest && current !== latest ? (
				<Box
					flexDirection="column"
					marginY={1}
					flexShrink={1}
					width={version_text.length + 2}
					borderColor="gray"
					borderStyle={'single'}
				>
					<Text>
						Upgrade available ({current} {'>'} {latest}), run{' '}
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
