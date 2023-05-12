import React from 'react';
import Help from './commands/Help.js';
import WhoAmI from './commands/WhoAmI.js';
import EventsTrack from './commands/EventsTrack.js';
import Send from './commands/Send.js';

interface IMapping {
	params?: string;
	instructions?: string;
	example?: string;
	options?: {
		option: string;
		value?: string;
		instructions?: string;
	}[];
	component: (params?: any) => React.ReactElement;
}

const mappings: Map<string, IMapping> = new Map();

mappings.set('help', {
	component: () => {
		return <Help mappings={mappings} />;
	},
});
mappings.set('whoami', {
	instructions: 'Display the currently authenticated workspace',
	component: () => {
		return <WhoAmI />;
	},
});
mappings.set('send', {
	instructions: 'Send a notification to a user, list, or audience',
	options: [
		{
			option: '--email <email address>',
			value: ''
		},
		{
			option: '--tel <phone number>',
			value: ''
		},
		{
			option: '--user <user ID>',
			value: ''
		},
		{
			option: '--list <list ID>',
			value: ''
		},
		{
			option: '--audience <audience ID>',
			value: ''
		},
		{
			option: '--body <message body>',
			value: ''
		},
		{
			option: '--title <message title/subject>',
			value: ''
		},
		{
			option: '--channels <channel>',
			value: 'comma-delimted list of channels to send to'
		},
		{
			option: '--all',
			value: 'send to all channels for each recipient (default is "single")'
		}
	],
	example: `courier send --tel 555-867-5309 --body "Hey Jenny\\!"`,
	component: params => {
		return <Send params={params} />;
	},
});
mappings.set('track', {
	params: '<event> <user>',
	instructions: 'Send a track event to trigger a Courier Automations',
	options: [
		{
			option: '--<key> <value>',
			value: ''
		}
	],
	example: `courier track EXAMPLE_EVENT user123 --name "Pip the Pigeon"`,
	component: params => {
		return <EventsTrack params={params} />;
	},
});

export default mappings;
