function delay(ms: number) {
	if (ms === 0) {
		return Promise.resolve();
	}
	return new Promise(resolve => setTimeout(resolve, ms));
}

export default delay;
