import React from 'react';
import {Box, Spacer, Text} from 'ink';

type Props = {
	label?: string;
	value?: string;
	width?: number;
	indent?: string;
	labelColor?: string;
	labelBold?: boolean;
};

export default ({
	label,
	labelColor,
	labelBold,
	value,
	width,
	indent,
}: Props) => {
	return (
		<Box flexDirection="column">
			<Box width={width} flexDirection="row">
				<Text>{indent}</Text>
				<Text bold={labelBold === false ? false : true} color={labelColor}>
					{label}
				</Text>
				<Spacer />
				<Text> {value}</Text>
			</Box>
		</Box>
	);
};
