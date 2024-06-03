import React from 'react';
import UhOh from '../components/UhOh.js';
import {Text} from 'ink';
import fs from 'fs';
import {useCliContext} from '../components/Context.js';
import dotenv from 'dotenv';

type Params = {};

export default ({}: Params) => {
	const {apikey, parsedParams, env_var} = useCliContext();
	const FILE_PATH = `${process.cwd()}/.courier`;
	const file_exists = fs.existsSync(FILE_PATH);

	const overwrite: boolean = Boolean(parsedParams['overwrite']);

	if (!apikey?.length) {
		return (
			<UhOh text="You must specify your API key using --apikey <your-api-key>" />
		);
	} else if (!file_exists) {
		fs.writeFileSync(FILE_PATH, `${env_var}=${apikey}\n`);
		return (
			<Text bold={true} color="green">
				Your API key ({env_var}) has been saved to {FILE_PATH}. Run "courier
				whoami" to verify API credentials.
			</Text>
		);
	} else {
		// open file as buffer
		const buffer = fs.readFileSync(FILE_PATH);
		let parsed = dotenv.parse(buffer);
		if (typeof parsed[env_var] !== 'undefined') {
			if (!overwrite) {
				return (
					<UhOh
						text={`${env_var} already exists. Consider adding the --overwrite option`}
					/>
				);
			}
		}
		parsed[env_var] = apikey;

		let new_file = Object.entries(parsed)
			.map(([key, value]) => `${key}=${value}`)
			.join('\n');
		fs.writeFileSync(FILE_PATH, new_file);
		return (
			<Text bold={true} color="green">
				Your API key ({env_var}) has been saved to {FILE_PATH}. Run "courier
				whoami" to verify API credentials.
			</Text>
		);
	}
};
