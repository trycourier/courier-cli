import React from 'react';
import {Box, Text, Newline} from 'ink';
import constants from '../constants.js';

interface IMapping {
	params?: string;
	instructions?: string;
	example?: string;
	options?: {
		option: string;
		value?: string;
		instructions?: string;
	}[];
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
					<React.Fragment key={k}>
						<Box>
							<Text>
								{'   • '}
								<Text color={constants.colors.primary}>{k}</Text>
								{v.params ? <Text color="gray"> {v.params}</Text> : null}
							</Text>
						</Box>
						<Box>
							<Text>
								{'     '}
								{v.instructions}
							</Text>
						</Box>
						{v.options
							? v.options.map(o => (
									<React.Fragment key={o.option}>
										<Box>
											<Text color="gray">
												{'     '}
												{o.option}
											</Text>
											<Text> {o.value}</Text>
										</Box>
									</React.Fragment>
							  ))
							: null}
						{v.example ? (
							<Text color="cyan">
								{'     '}
								{v.example}
							</Text>
						) : null}
					</React.Fragment>
				);
			})}
		</Box>
	);
};
