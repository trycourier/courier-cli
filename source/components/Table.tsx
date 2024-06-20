// Table.tsx
import React from 'react';
import {Box, Text, TextProps} from 'ink';

type Scalar = string | number | boolean | null | undefined;

type ScalarDict = {
	[key: string]: Scalar;
};

type Column = {
	key: string;
	label?: string;
	width: number;
};

type TableProps = {
	data: ScalarDict[];
	showHeaders?: boolean;
	headerLabels?: {[key: string]: string};
	disableRowSeparators?: boolean;
	disableBorders?: boolean;
	headerStyles?: TextProps;
	rowStyles?: TextProps;
};

const Table = ({
	data,
	disableBorders,
	disableRowSeparators,
	headerStyles,
	headerLabels,
	rowStyles,
	showHeaders = true,
}: TableProps) => {
	// Determine columns and their widths
	const columns: Column[] = getColumns(data, headerLabels);

	return (
		<Box flexDirection="column" width="100%">
			{!disableBorders && renderHeaderSeparators(columns)}

			{showHeaders && (
				<>
					{renderRow(
						columns.reduce((p, v) => {
							p[v.key] = v.label || v.key;
							return p;
						}, {} as ScalarDict),
						columns,
						disableBorders,
						{
							color: 'blue',
							bold: true,
							...headerStyles,
						},
					)}
					{!disableRowSeparators ? renderRowSeparators(columns) : null}
				</>
			)}

			{data.map((row, index) => (
				<React.Fragment key={`row-${index}`}>
					{index !== 0 && !disableRowSeparators && renderRowSeparators(columns)}
					{renderRow(row, columns, disableBorders, rowStyles)}
				</React.Fragment>
			))}
			{!disableBorders && renderFooterSeparators(columns)}
		</Box>
	);
};

// Helper function to determine columns and their widths
function getColumns(
	data: ScalarDict[],
	headerLabels?: {[key: string]: string},
): Column[] {
	let columnWidths: {[key: string]: number} = {};

	data.forEach(row => {
		Object.keys(row).forEach(key => {
			const valueLength = row[key]?.toString().length || 0;
			columnWidths[key] = Math.max(
				columnWidths[key] || key.length,
				valueLength,
			);
		});
	});

	return Object.keys(columnWidths).map(key => ({
		key: key,
		label: headerLabels?.[key] || key,
		width: (columnWidths[key] ?? 0) + 2, // adding padding
	}));
}

// Helper function to render a row with separators
function renderRow(
	row: ScalarDict,
	columns: Column[],
	disableBorders?: boolean,
	textStyles?: TextProps & {disableBorders?: boolean},
) {
	return (
		<Box flexDirection="row">
			{!disableBorders && <Text>│</Text>}
			{columns.map((column, index) => (
				<React.Fragment key={column.key}>
					{index !== 0 && !disableBorders && <Text>│</Text>}
					{/* Add separator before each cell except the first one */}
					<Box width={column.width} justifyContent="flex-start">
						<Text {...textStyles}>{row[column.key]?.toString() || ''}</Text>
					</Box>
				</React.Fragment>
			))}
			{!disableBorders && <Text>│</Text>}
		</Box>
	);
}

function renderHeaderSeparators(columns: Column[]) {
	return renderRowSeparators(columns, '┌', '┬', '┐');
}

function renderFooterSeparators(columns: Column[]) {
	return renderRowSeparators(columns, '└', '┴', '┘');
}

function renderRowSeparators(
	columns: Column[],
	leftChar = '├',
	midChar = '┼',
	rightChar = '┤',
) {
	return (
		<Box flexDirection="row">
			<Text>{leftChar}</Text>
			{columns.map((column, index) => (
				<React.Fragment key={column.key}>
					<Text>{'─'.repeat(column.width)}</Text>
					{index < columns.length - 1 ? (
						<Text>{midChar}</Text>
					) : (
						<Text>{rightChar}</Text>
					)}
				</React.Fragment>
			))}
		</Box>
	);
}

export default Table;
