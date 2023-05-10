import React from 'react';
import Help from '../commands/Help.js';
import yargs from 'yargs-parser';

interface IMapping {
	instructions?: string;
	component: (params?: any) => React.ReactElement;
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
		return mapping.component(parsedParams);
	} else {
		return <Help mappings={mappings} />;
	}
};
