import React from 'react';
import {Box, Text, Newline} from 'ink';
import KVP from '../components/KVP.js';
import constants from '../constants.js';

interface IMapping {
	instructions?: string;
	component: (params?: any) => React.ReactElement;
}

export default ({mappings}: {mappings: Map<string, IMapping>}) => {
	const keys = [...mappings.keys()]; // convert to array
	return (
		<Box flexDirection="column">
			<Text bold={true}>Usage</Text>
			<Text>
				$ <Text color={constants.colors.primary}>courier</Text>{' '}
				<Text color="gray">&lt;command&gt;</Text>
				<Newline />
			</Text>

			<Text bold={true}>Commands</Text>
			{keys.map(k => {
				const v = mappings.get(k);
				if (!v || !v.instructions || !v.instructions.length) {
					return null;
				}

				return (
					<KVP
						key={k}
						width={22}
						indent="  • "
						label={k}
						labelColor={constants.colors.primary}
						labelBold={false}
						value={v.instructions}
					/>
				);
			})}
		</Box>
	);
};
