import React from 'react';
import Help from './commands/Help.js';
import WhoAmI from './commands/WhoAmI.js';
import NotYetImplemented from './commands/NotYetImplemented.js';
import EventsTrack from './commands/EventsTrack.js';

interface IMapping {
	instructions?: string;
	component: (params?: any) => React.ReactElement;
}

const mappings: Map<string, IMapping> = new Map();

mappings.set('help', {
	component: () => {
		return <Help mappings={mappings} />;
	},
});
mappings.set('whoami', {
	instructions: 'Type "whoami" to see who you are.',
	component: () => {
		return <WhoAmI />;
	},
});
mappings.set('events:track', {
	instructions: 'Send an event to test your Courier Automations',
	component: params => {
		return <EventsTrack params={params} />;
	},
});
mappings.set('translations:push', {
	instructions: 'Push translation files to your Courier workspace',
	component: () => {
		return <NotYetImplemented />;
	},
});
mappings.set('translations:pull', {
	instructions: 'Pull translation files from your Courier workspace',
	component: () => {
		return <NotYetImplemented />;
	},
});

export default mappings;
