#!/usr/bin/env node
import {render} from 'ink';
import React from 'react';
import {CliContextProvider} from './components/Context.js';
import args from './lib/args.js';
import loadEnv from './lib/load-env.js';
import Version from './version.js';

const CLI = async () => {
	process.removeAllListeners('warning');
	await loadEnv();
	const params = args(process.argv);
	const mappings = (await import('./mappings.js')).default;
	const Router = (await import('./components/Router.js')).default;

	render(
		<CliContextProvider args={params} {...{mappings}}>
			<Router />
			<Version />
		</CliContextProvider>,
	);
};

(async () => {
	await CLI();
})();
