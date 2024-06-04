interface IResponse {
	res: Response;
	json?: IDebug;
	err?: Error;
}

interface IMapping {
	noApiKeyRequired?: boolean;
	params?: string;
	instructions?: string;
	example?: string | string[];
	options?: {
		option: string;
		value?: string;
		instructions?: string;
	}[];
	component: (params?: any) => React.ReactElement;
}

type TFileType = 'csv' | 'json' | 'parquet';

interface IDebug {
	environment: string;
	scope: string;
	tenantId: string;
	tenantName: string;
	mock: boolean;
}

interface IResponseDebug extends IResponse {
	json?: IDebug;
}