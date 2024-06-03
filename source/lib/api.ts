import VERSION from "../version.js";

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

export default async (request: IRequest, baseUrl: string, apikey: string): Promise<IResponse> => {
	const req = {
		method: request.method,
		headers: {
			Authorization: `Bearer ${apikey}`,
			'Content-Type': 'application/json',
			'User-Agent': `courier-cli/${VERSION}`,
			'X-Courier-CLI-Version': VERSION,
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
