import {Text} from 'ink';
import React from 'react';
import Help from '../commands/Help.js';
import {getApiKeyFlags, useCliContext} from './Context.js';
import VERSION from '../version.js';

type Props = {};

export default ({}: Props) => {
	const {
		args,
		mapping,
		parsedParams,
		mappings,
		apikey,
		env_var,
		environment,
		routing,
		document_scope,
		map,
	} = useCliContext();

	if (args.filter(a => ['--version', '-V'].includes(a)).length) {
		return <Text>courier v{VERSION}</Text>;
	} else if (!args.length || !args[0]) {
		return <Help mappings={mappings} />;
	} else {
		if (mapping) {
			if (parsedParams['help']) {
				return <Help mappings={new Map([[map!, mapping]])} />;
			} else if ((apikey && apikey.length) || mapping.noApiKeyRequired) {
				return mapping.component(parsedParams);
			} else {
				const extra_flags = getApiKeyFlags({
					environment,
					routing,
					document_scope,
				});
				return (
					<>
						<Text bold={true} color="red">
							No {env_var} specified, please set via one of these options:
						</Text>
						<Text bold={true} color="red">
							• running "courier config --apikey &lt;your-api-key&gt;{' '}
							{extra_flags}"
						</Text>
						<Text bold={true} color="red">
							• setting {env_var} in your shell via "export {env_var}
							=&lt;your-api-key&gt;"
						</Text>
						<Text bold={true} color="red">
							• setting {env_var} in a ".courier" file in your current working
							directory
						</Text>
						<Text bold={true} color="red">
							• setting {env_var} in a ".courier" file in your user's home
							directory
						</Text>
					</>
				);
			}
		} else {
			return <Help mappings={mappings} />;
		}
	}
};
