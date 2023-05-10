import React from 'react';
import {Text} from 'ink';

export default ({text}: {text: string}) => {
	return (
		<Text bold={true} color="red">
			{text}
		</Text>
	);
};
