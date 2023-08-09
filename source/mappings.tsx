import React from 'react';
import Help from './commands/Help.js';
import Config from './commands/Config.js';
import WhoAmI from './commands/WhoAmI.js';
import Track from './commands/Track.js';
import Send from './commands/Send.js';
import DigestFlush from './commands/DigestFlush.js';
import TranslationsDownload from './commands/TranslationsDownload.js';
import TranslationsUpload from './commands/TranslationsUpload.js';

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

const mappings: Map<string, IMapping> = new Map();

mappings.set('help', {
	component: () => {
		return <Help mappings={mappings} />;
	},
	noApiKeyRequired: true
});
mappings.set('config', {
	instructions: 'Persist your Courier API key into a .courier file in your current working directory',
	component: params => {
		return <Config params={params} />;
	},
	options: [
		{
			option: '--apikey <Courier API Key>',
			value: '',
		},
		{
			option: '--overwrite',
			value: 'overwrite an existing .courier file',
		},
	],
	example: `courier config --apikey MY_API_KEY`,
	noApiKeyRequired: true
});
mappings.set('whoami', {
	instructions: 'Display the currently authenticated workspace',
	component: () => {
		return <WhoAmI />;
	},
});
mappings.set('send', {
	instructions:
		'Send a notification to a user, list, or audience. Unrecognized parameters will be sent as message data.',
	options: [
		{
			option: '--email <email address>',
			value: '',
		},
		{
			option: '--tel <phone number>',
			value: '',
		},
		{
			option: '--apn <Apple push token>',
			value: '',
		},
		{
			option: '--fcm <Firebase push token>',
			value: '',
		},
		{
			option: '--user <user ID>',
			value: 'ID of a Courier User in your workspace',
		},
		{
			option: '--list <list ID>',
			value: 'ID of a Courier List in your workspace',
		},
		{
			option: '--audience <audience ID>',
			value: 'ID of a Courier Audience in your workspace',
		},
		{
			option: '--body <message body>',
			value: '',
		},
		{
			option: '--title <message title/subject>',
			value: '',
		},
		{
			option: '--elemental <filepath>',
			value: 'path to Courier Elemental JSON file',
		},
		{
			option: '--template <template ID or alias>',
			value: 'ID or alias of a template stored in Courier',
		},
		{
			option: '--channels <channel>',
			value: 'comma-delimted list of channels to send to',
		},
		{
			option: '--all',
			value: 'send to all channels for each recipient (default is "single")',
		},
	],
	example: [
		`courier send --tel 555-867-5309 --body "Hey Jenny\\!"`,
		`courier send --user user123 --template my-template-id --foo bar`,
	],
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
			value: 'arbitrary key/value properties for your event',
		},
	],
	example: `courier track EXAMPLE_EVENT user123 --name "Pip the Pigeon"`,
	component: params => {
		return <Track params={params} />;
	},
});
mappings.set('digests:flush', {
	params: '<user> <digest>',
	instructions: 'Flush any currently queued events for a given user + digest',
	example: `courier digests:flush user123 MY_DIGEST_TOPIC`,
	component: params => {
		return <DigestFlush params={params} />;
	},
});
mappings.set('translations:upload', {
	params: '<locale> <filepath>',
	instructions: 'Upload a .PO file to Courier for a given locale',
	options: [
		{
			option: '--domain <domain>',
			value: 'a custom language domain (default is "default")',
		},
	],
	example: `courier translations:upload en-US ./translations/en-US.po`,
	component: params => {
		return <TranslationsUpload params={params} />;
	},
});
mappings.set('translations:download', {
	params: '<locale>',
	instructions: 'Download a .PO file to Courier for a given locale',
	options: [
		{
			option: '--domain <domain>',
			value: 'a custom language domain (default is "default")',
		},
		{
			option: '--text',
			value: 'only return plain text (e.g. for piping into a file)',
		},
	],
	example: `courier translations:download en-US --text > example.en-US.po`,
	component: params => {
		return <TranslationsDownload params={params} />;
	},
});

export default mappings;
