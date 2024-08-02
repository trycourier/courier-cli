import duckdb from 'duckdb'

export const installExtension = (db: duckdb.Database, type?: TFileType) => {
	if (['json', 'parquet'].includes(type || '')) {
		db.exec(`
            INSTALL ${type};
            LOAD ${type};`);
	}
};

export const getFrom = (filename: string, type: TFileType) => {
	switch (type) {
		case 'csv':
			return `read_csv(['${filename}'], union_by_name = true)`;
		case 'json':
			return `read_json_auto(['${filename}'])`;
		case 'parquet':
			return `read_parquet(['${filename}'])`;
	}
};

export const getFileType: (filename: string) => TFileType | undefined = (
	filename: string,
) => {
	if (filename.endsWith('.csv')) {
		return 'csv';
	} else if (filename.endsWith('.json') || filename.endsWith('.jsonl')) {
		return 'json';
	} else if (
		filename.endsWith('.parquet') ||
		filename.endsWith('.pq') ||
		filename.endsWith('.parq')
	) {
		return 'parquet';
	} else {
		return undefined;
	}
};

export const getSql = (filename: string, type: TFileType) => `SELECT * FROM ${getFrom(filename, type)} ;`;

const getDb = (filename: string) => {
    const filetype = getFileType(filename);
    const db = new duckdb.Database(':memory:'); // or a file name for a persistent DB
    installExtension(db, filetype);
    return {
        db,
        filetype,
        sql: getSql(filename, filetype || 'csv'),
    }
}

export const getChunk = (data: duckdb.TableData, chunk_size: number = 1) => {
	let rows = data.splice(0, chunk_size)
	return {
		rows,
		data
	}
}

export default getDb;