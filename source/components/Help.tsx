import React from 'react';
import {Box, Text, Newline } from 'ink';

interface IMapping {
  instructions: string;
  component: any;
}

export default ({mappings}: {mappings: Map<string, IMapping>}) => {
	const keys = [...mappings.keys()]; // convert to array
  return <Box flexDirection="column">

 		<Text>Usage</Text>
 		<Text>  $ courier &lt;command&gt;<Newline /></Text>

 		<Text>Commands</Text>
		{keys.map((k) => (
			<Text key={k}>  {k}: {mappings.get(k)?.instructions}</Text>
		))}

  </Box>
}
