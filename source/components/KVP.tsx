import React from 'react';
import { Box, Text } from 'ink';

type Props = {
  label: string;
  value: string;
}

export default ({ label, value }: Props) => {
	return <Box>
    <Box width={20}><Text bold={true}>{label}:</Text></Box>
    <Text> {value}</Text>
  </Box>
}
