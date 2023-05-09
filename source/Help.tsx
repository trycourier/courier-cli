import React from 'react';
import {Box, Text, Newline } from 'ink';

export default () => (
  <Box flexDirection="column">
 		<Text>Usage $ %NAME%<Newline /></Text>
 		<Text>Options</Text>
 		<Text>  --name Your name<Newline /></Text>
 		<Text>Examples</Text>
 		<Text>  $ %NAME% --name=Jane</Text>
 		<Text>  Hello, Jane</Text>
  </Box>
)
