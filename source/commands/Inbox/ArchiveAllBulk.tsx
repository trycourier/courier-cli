import duckdb from 'duckdb';
import _ from 'lodash';
import React, {useEffect, useState} from 'react';
import {useBoolean, useCounter} from 'usehooks-ts';
import {useCliContext} from '../../components/Context.js';
import Spinner from '../../components/Spinner.js';
import UhOh from '../../components/UhOh.js';
// @ts-ignore
import {Alert} from '@inkjs/ui';
import getDb from '../../bulk.js';
import ArchiveAll from './ArchiveAll.js';

interface IArchiveAllBulk {}

interface IParam {
	_: string[];
}

const ArchiveAllBulk = ({}: IArchiveAllBulk) => {
	const [error, setError] = useState<string | undefined>();
	const running = useBoolean(true);
	const processing = useBoolean(false);
	const {parsedParams} = useCliContext();
	const [data, setData] = useState<duckdb.TableData | undefined>();
	const current = useCounter(0);
	const [finished, setFinished] = useState<
		{
			type: 'error' | 'success';
			message: string;
		}[]
	>([]);
	const {
		_: [filename],
	} = parsedParams as IParam;

	const {db, filetype, sql} = getDb(filename || '');

	const getData = () => {
		db.all(sql, (err, result) => {
			if (err) {
				setError(err.message);
			} else {
				setData(result);
			}
		});
	};

	useEffect(() => {
		if (filetype) {
			getData();
		} else {
			setError('File type not supported.');
		}
	}, []);

	useEffect(() => {
		if (data) {
			processData();
		}
	}, [data]);

	useEffect(() => {
		if (error?.length) {
			running.setFalse();
		}
	}, [error]);

	const processData = () => {
		if (data?.length) {
			processing.setTrue();
		}
	};

	const getNext = () => {
		if (data?.length && current.count < data.length - 1) {
			current.increment();
		} else {
			running.setFalse();
			processing.setFalse();
		}
	};

	if (error?.length) {
		return <UhOh text={error} />;
	} else if (processing.value) {
		const current_user = _.get(data, [current.count, 'user_id']);
		if (!current_user) {
			getNext();
			return <UhOh text="No user_id found" />;
		} else {
			return (
				<>
					{finished.map((t, i) => (
						<Alert key={i} variant={t.type}>
							{t.message}
						</Alert>
					))}
					<Spinner text={`Processing ${current_user}`} />
					<ArchiveAll
						key={current_user}
						user_id_override={current_user}
						onError={(error: string) => {
							const text = `${current_user} - ${error}`;
							setFinished(p => [...p, {type: 'error', message: text}]);
							getNext();
						}}
						user_finished={(messages: number) => {
							const text = `${current_user} - Archived ${messages} messages`;
							setFinished(p => [...p, {type: 'success', message: text}]);
							getNext();
						}}
					/>
				</>
			);
		}
	} else if (running.value) {
		return <Spinner text="Gathering data" />;
	} else if (finished.length) {
		return (
			<>
				{finished.map((t, i) => (
					<Alert key={i} variant={t.type}>
						{t.message}
					</Alert>
				))}
			</>
		);
	} else {
		return <UhOh text="No data found" />;
	}
};

export default ArchiveAllBulk;
