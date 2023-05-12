import React from 'react';
import {Box, Text} from 'ink';

interface IElement {
	type: string;
}

interface IElemental {
	version: string;
	elements: IElement[];
}

type Props = {
	elemental: IElemental;
};

type TProperty = {
	name: string;
	value: string;
};

type TBucket = {
	type: string;
	properties: TProperty[];
};

export default ({elemental}: Props) => {
	const buckets: TBucket[] = [];
	elemental.elements.map(element => {
		const {type, ...rest} = element;
		buckets.push({
			type: element.type,
			properties: Object.keys(rest).map(key => ({
				name: key,
				value: (rest as any)[key],
			})),
		});
	});

	return (
		<Box flexDirection="column">
			{buckets.map(bucket =>
				bucket.properties.map(property => (
					<Box>
						<Box width={16} paddingTop={1}>
							<Text>
								{bucket.type}.{property.name}
							</Text>
						</Box>
						<Box borderStyle="round" flexDirection="column" paddingX={1}>
							<Text>{property.value}</Text>
						</Box>
					</Box>
				)),
			)}
		</Box>
	);
};
