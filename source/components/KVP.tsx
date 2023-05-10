import React from 'react';
import { Box, Text } from 'ink';

type Props = {
  label?: string;
  value?: string;
  width?: number;
  indent?: string;
  labelColor?: string;
  labelBold?: boolean;
}

export default ({ label, labelColor, labelBold, value, width, indent }: Props) => {
	return <Box>
    <Box width={width}>
      <Text>{indent}</Text>
      <Text bold={labelBold === false ? false : true} color={labelColor}>{label}</Text>
    </Box>
    <Text> {value}</Text>
  </Box>
}
