export default async (path: string, method: string, options?: RequestInit) => {
  const baseUrl = process.env['COURIER_DOMAIN'] || 'https://api.courier.com';
  return fetch(`${baseUrl}${path}`, {
    method,
    headers: {
      'Authorization': `Bearer ${process.env['API_KEY']}`,
      'Content-Type': 'application/json',
      "User-Agent": `courier-cli/0.0.1`,
    },
    ...options
  });
}