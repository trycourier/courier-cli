import React from 'react';
import {Text} from 'ink';

export default ({text}: {text: string}) => {
	if (!text || !text.length) {
		return null;
	} else {
		return (
			<Text bold={true} color="red">
				{text}
			</Text>
		);
	}
};
