import React, {useEffect, useState} from 'react';
import {Box, Text} from 'ink';
import spinners from 'cli-spinners';
import constants from '../constants.js';

/*
some of my favorites:
- spinners.dots
- spinners.triangle
- spinners.binary
- spinners.arc
- spinners.arrow
- spinners.bouncingBar
- spinners.pong
- spinners.aesthetic
*/
const spinner = spinners.aesthetic;
const spinnerColor = constants.colors.primary;
const textColor = 'white';

type Props = {
	text?: string;
};

export default ({text}: Props) => {
	const [index, setIndex] = useState(0);

	useEffect(() => {
		const timer = setInterval(() => {
			setIndex(i => {
				if (i + 1 >= spinner.frames.length) {
					return 0;
				} else {
					return i + 1;
				}
			});
		}, spinner.interval);

		return () => {
			clearInterval(timer);
		};
	}, []);

	return (
		<Box borderStyle="classic" borderColor="gray">
			<Text bold={true} color={spinnerColor}>
				{spinner.frames[index]}
			</Text>
			{text?.length ? <Text color={textColor}> {text}</Text> : <></>}
		</Box>
	);
};
