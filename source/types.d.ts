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