import React from 'react';
import {Box, Text, Newline} from 'ink';
import constants from '../constants.js';
import Table from '../components/Table.js';
import {COMMON_OPTIONS} from '../mappings.js';

interface IMapping {
	params?: string;
	instructions?: string;
	example?: string | string[];
	options?: {
		option: string;
		value?: string;
		instructions?: string;
	}[];
	component: (params?: any) => React.ReactElement;
}

// const Space = () => <>{'  '}</>;

export default ({mappings}: {mappings: Map<string, IMapping>}) => {
	const keys = [...mappings.keys()]; // convert to array
	return (
		<Box flexDirection="column" paddingY={2}>
			<Text bold={true}>Usage</Text>
			<Text>
				$ <Text color={constants.colors.primary}>courier</Text>{' '}
				<Text color="gray">&lt;command&gt;</Text>
				<Newline />
			</Text>
			<Text bold={true}>Common Flags</Text>
			<Table
				disableRowSeparators={true}
				disableBorders={true}
				data={COMMON_OPTIONS}
				headerLabels={{option: 'Flags', value: 'Description'}}
				rowStyles={{wrap: 'wrap'}}
			/>
			<Newline />
			<Text bold={true}>Commands</Text>
			{keys.map(k => {
				const v = mappings.get(k);
				if (!v || !v.instructions || !v.instructions.length) {
					return null;
				}

				return (
					<React.Fragment key={k}>
						<Box paddingTop={2}>
							<Text>
								<Text color={constants.colors.primary}>{k}</Text>
								{v.params ? <Text color="gray"> {v.params}</Text> : null}
							</Text>
						</Box>
						<Box>
							<Text>{v.instructions}</Text>
						</Box>
						{v.options?.length ? (
							<Table
								disableRowSeparators={true}
								disableBorders={true}
								data={v.options}
								headerLabels={{option: 'Flags', value: 'Description'}}
								rowStyles={{wrap: 'wrap'}}
							/>
						) : (
							<></>
						)}
						{v.example && Array.isArray(v.example) ? (
							v.example.map((e, i) => (
								<Text color="cyan" key={i}>
									{e}
								</Text>
							))
						) : v.example ? (
							<Text color="cyan">{v.example}</Text>
						) : null}
					</React.Fragment>
				);
			})}
		</Box>
	);
};
