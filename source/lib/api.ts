interface IRequest {
	url: string;
	method: string;
	body?: object | string;
	options?: RequestInit;
}

interface IResponse {
	res: Response;
	json?: any;
	err?: Error;
}

const isString = (s: any): boolean => {
	return typeof s === 'string' || s instanceof String;
};

export default async (request: IRequest): Promise<IResponse> => {
	const baseUrl = process.env['COURIER_DOMAIN'] || 'https://api.courier.com';
	return fetch(`${baseUrl}${request.url}`, {
		method: request.method,
		headers: {
			Authorization: `Bearer ${process.env['COURIER_API_KEY']}`,
			'Content-Type': 'application/json',
			'User-Agent': `courier-cli/0.0.1`,
		},
		body:
			request.body && !isString(request.body)
				? JSON.stringify(request.body)
				: undefined,
		...request.options,
	}).then(res => {
		if (res.status > 400) {
			return {
				res,
				err: new Error(`${res.status}: ${res.statusText}`),
			};
		} else if (res.status === 204) {
			return {res};
		} else {
			return res.json().then(json => ({res, json}));
		}
	});
};
