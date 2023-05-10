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
	const UhOh = (await import('./components/UhOh.js')).default;

	const apiKey = process.env['COURIER_API_KEY'];
	if (apiKey && apiKey.length) {
		render(<Router args={params} mappings={mappings} />);
	} else {
		render(
			<UhOh text="No COURIER_API_KEY specified; add that to your environment or ~/.courier file" />,
		);
	}
};

(async () => {
	await CLI();
})();
