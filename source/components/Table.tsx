// Table.tsx
import React from 'react';
import {Box, Text} from 'ink';

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
	headerStyles?: {
		color?: string;
		backgroundColor?: string;
		bold?: boolean;
		italic?: boolean;
		underline?: boolean;
		inverse?: boolean;
		strikethrough?: boolean;
		dimColor?: boolean;
	};
};

const Table = ({
	data,
	showHeaders = true,
	headerStyles,
	headerLabels,
}: TableProps) => {
	// Determine columns and their widths
	const columns: Column[] = getColumns(data, headerLabels);

	return (
		<Box flexDirection="column" width="100%">
			{renderHeaderSeparators(columns)}

			{showHeaders && (
				<>
					{renderRow(
						columns.reduce((p, v) => {
							p[v.key] = v.label || v.key;
							return p;
						}, {} as ScalarDict),
						columns,
						{
							color: 'blue',
							bold: true,
							...headerStyles,
						},
					)}
					{renderRowSeparators(columns)}
				</>
			)}

			{data.map((row, index) => (
				<React.Fragment key={`row-${index}`}>
					{index !== 0 && renderRowSeparators(columns)}
					{renderRow(row, columns, {wrap: 'wrap'})}
				</React.Fragment>
			))}
			{renderFooterSeparators(columns)}
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
function renderRow(row: ScalarDict, columns: Column[], textStyles?: any) {
	return (
		<Box flexDirection="row">
			<Text>│</Text>
			{columns.map((column, index) => (
				<React.Fragment key={column.key}>
					{index !== 0 && <Text>│</Text>}
					{/* Add separator before each cell except the first one */}
					<Box width={column.width} justifyContent="flex-start">
						<Text {...textStyles}>{row[column.key]?.toString() || ''}</Text>
					</Box>
				</React.Fragment>
			))}
			<Text>│</Text>
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
