interface IResponse {
	res: Response;
	json?: any;
}

export default async (
	path: string,
	method: string,
	options?: RequestInit,
): Promise<IResponse> => {
	const baseUrl = process.env['COURIER_DOMAIN'] || 'https://api.courier.com';
	return fetch(`${baseUrl}${path}`, {
		method,
		headers: {
			Authorization: `Bearer ${process.env['COURIER_API_KEY']}`,
			'Content-Type': 'application/json',
			'User-Agent': `courier-cli/0.0.1`,
		},
		...options,
	}).then(res => {
		if (res.status > 400) {
			throw new Error(`${res.status}: ${res.statusText}`);
		} else if (res.status === 204) {
			return {res};
		} else {
			return res.json().then(json => ({res, json}));
		}
	});
};
