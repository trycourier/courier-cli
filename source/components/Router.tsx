import React from 'react';
import yargs from 'yargs-parser';
import {Text} from 'ink';
import Help from '../commands/Help.js';

interface IMapping {
	component: (params?: any) => React.ReactElement;
	noApiKeyRequired?: boolean;
}

type Props = {
	args: string[];
	mappings: Map<string, IMapping>;
};

export default ({args, mappings}: Props) => {
	if (!args.length || !args[0]) {
		return <Help mappings={mappings} />;
	}

	const mapping = mappings.get(args[0]);
	const [, ...params] = args;
	const parsedParams = params.length ? yargs(params) : undefined;

	if (mapping) {
		const apiKey = process.env['COURIER_API_KEY'];
		if ((apiKey && apiKey.length) || mapping.noApiKeyRequired) {
			return mapping.component(parsedParams);
		} else {
			return <>
				<Text bold={true} color="red">
					No COURIER_API_KEY specified, please set via one of these options:
				</Text>
				<Text bold={true} color="red">
					• running "courier config --apikey &lt;your-api-key&gt;"
				</Text>
				<Text bold={true} color="red">
					• setting COURIER_API_KEY in your shell via "export COURIER_API_KEY=&lt;your-api-key&gt;"
				</Text>
				<Text bold={true} color="red">
					• setting COURIER_API_KEY in a ".courier" file in your current working directory
				</Text>
				<Text bold={true} color="red">
					• setting COURIER_API_KEY in a ".courier" file in your user's home directory
				</Text>
			</>
		}
	} else {
		return <Help mappings={mappings} />;
	}
};
