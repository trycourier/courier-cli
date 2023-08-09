#!/usr/bin/env node
import React from 'react';
import {render} from 'ink';
import args from './lib/args.js';
import loadEnv from './lib/load-env.js';

const CLI = async () => {
	await loadEnv();
	const params = args(process.argv);
	const mappings = (await import('./mappings.js')).default;
	const Router = (await import('./components/Router.js')).default;

	render(<Router args={params} mappings={mappings} />);
};

(async () => {
	await CLI();
})();
