const VERSION = '1.0.0';

interface IRequest {
	url: string;
	method: string;
	body?: object | string;
	headers?: {
		[key: string]: string;
	};
	options?: RequestInit;
}

interface IResponse {
	res: Response;
	json?: any;
	text?: string;
	err?: Error;
}

const isString = (s: any): boolean => {
	return typeof s === 'string' || s instanceof String;
};

export default async (request: IRequest): Promise<IResponse> => {
	const baseUrl = process.env['COURIER_DOMAIN'] || 'https://api.courier.com';
	const req = {
		method: request.method,
		headers: {
			Authorization: `Bearer ${process.env['COURIER_API_KEY']}`,
			'Content-Type': 'application/json',
			'User-Agent': `courier-cli/${VERSION}`,
			...request.headers,
		},
		body:
			request.body && !isString(request.body)
				? JSON.stringify(request.body)
				: request.body?.toString(),
		...request.options,
	};
	return fetch(`${baseUrl}${request.url}`, req).then(res => {
		if (res.status > 400) {
			return {
				res,
				err: new Error(`${res.status}: ${res.statusText}`),
			};
		} else if (res.status === 204) {
			return {res};
		} else {
			if (res.headers.get('Content-Type')?.includes('application/json')) {
				return res.json().then(json => ({res, json}));
			} else {
				return res.text().then(text => ({res, text}));
			}
		}
	});
};
