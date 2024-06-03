export default (argv: string[]): string[] => {
	if (!argv || !argv.length || !argv[0]) {
		return [];
	} else if (argv[0].endsWith('node')) {
		const [, , ...args] = argv;
		return args;
	} else {
		const [, ...args] = argv;
		return args;
	}
};