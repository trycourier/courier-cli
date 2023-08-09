import React from 'react';
import UhOh from '../components/UhOh.js';
import {Text} from 'ink';
import fs from 'fs';

type Params = {
	apikey?: string;
	overwrite?: boolean;
};

export default ({params}: {params: Params}) => {
	const FILE_PATH = `${process.cwd()}/.courier`;

	if (!params?.apikey) {
		return (
			<UhOh text="You must specify your API key using --apikey <your-api-key>" />
		);
	}

	if (fs.existsSync(FILE_PATH) && !params.overwrite) {
		return (
			<UhOh
				text={`${FILE_PATH} already exists. Consider adding the --overwrite option`}
			/>
		);
	}

	fs.writeFileSync(FILE_PATH, `COURIER_API_KEY=${params.apikey}\n`);

	return (
		<Text bold={true} color="green">
			Your API key has been saved to {FILE_PATH}. Run "courier whoami" to verify
			API credentials.
		</Text>
	);
};
