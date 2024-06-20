import React from 'react';
import Help from './commands/Help.js';
import Config from './commands/Config.js';
import WhoAmI from './commands/WhoAmI.js';
import Track from './commands/Track.js';
import UsersGet from './commands/UsersGet.js';
import UsersSet from './commands/UsersSet.js';
import Send from './commands/Send.js';
import TranslationsDownload from './commands/TranslationsDownload.js';
import TranslationsUpload from './commands/TranslationsUpload.js';
import UsersBulk from './commands/UsersBulk.js';
import UsersPreferences from './commands/UsersPreferences.js';
import Upgrade from './commands/Upgrade.js';
import TenantsBulk from './commands/TenantsBulk.js';
import MarkAllRead from './commands/Inbox/MarkAllRead.js';
import ArchiveAll from './commands/Inbox/ArchiveAll.js';
import ArchiveAllBulk from './commands/Inbox/ArchiveAllBulk.js';
import TemplatesList from './commands/Templates/List.js';
import TrackBulk from './commands/TrackBulk.js';
import AutomationInvokeBulk from './commands/AutomationInvokeBulk.js';

const mappings: Map<string, IMapping> = new Map();

export const COMMON_OPTIONS = [
	{
		option: '-M --mock',
		value:
			'Use the API key that simulates sending using the simulating routing',
	},
	{
		option: '-P --production',
		value: 'Use the production environment API key',
	},
	{
		option: '-D --draft',
		value:
			'Use the draft document scope API key. Use draft or submitted, will default to published key if neither are provided',
	},
	{
		option: '-S --submitted',
		value: 'Use the submitted document scope API key',
	},
	{
		option: '--apikey <Courier API Key>',
		value:
			'Use the provided Courier API key, otherwise use the approprate environment variable',
	},
	{
		option: '--apiurl <Courier API URL>',
		value:
			'Use the provided Courier API URL, otherwise use COURIER_API_URL environment variable. Default is https://api.courier.com',
	},
];

mappings.set('help', {
	component: () => {
		return <Help mappings={mappings} />;
	},
	noApiKeyRequired: true,
});

mappings.set('upgrade', {
	instructions: `Upgrade the Courier CLI to the latest versionw`,
	component: () => <Upgrade />,
});

mappings.set('config', {
	instructions:
		'Persist your Courier API key into a .courier file in your current working directory',
	component: () => {
		return <Config />;
	},
	options: [
		{
			option: '--overwrite',
			value: 'Overwrite this apikey in the existing .courier file',
		},
	],
	example: [
		`courier config --apikey MY_API_KEY -P --override`,
		`courier config --apikey MY_API_KEY --mock`,
		`courier config --apikey MY_API_KEY --draft`,
	],
	noApiKeyRequired: true,
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
			option: '--tenant <tenant ID>',
			value: 'ID of a Courier Tenant in your workspace. Will be used to send',
		},
		{
			option: '--include-children',
			value: 'When sending to a tenant, include all the children of the tenant',
		},
		{
			option: '--tenant-context <tenant ID>',
			value:
				'ID of a Courier Tenant in your workspace. Applies the tenant context to the message',
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
		`courier send -P --user=test123 --body "hello world" --title="hello" --channels=inbox`,
		`courier send --tenant=kewl --title=hello --body="hello world" --channel=inbox`,
		`courier send --user="1" --tenant-context=kewl --title=hello --body="hello world" --channel=inbox`,
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
mappings.set('track:bulk', {
	params: '<event> <filename>',
	instructions:
		'Bulk import track events from a file (csv, json, etc). user_id and recipient are mapped to recipient, all other additional fields are placed in properties',
	example: 'courier track "example event" examples/events.csv',
	component: () => {
		return <TrackBulk />;
	},
});
mappings.set('automation:invoke:bulk', {
	params: '<automation_template_id> <filename>',
	instructions:
		'Bulk invoke automations from a file (csv, json, etc). user_id and recipient are mapped to recipient, fields starting with profile. are routed to profile key. The rest is placed in data key',
	example:
		'courier automation:invoke:bulk "7ee13494-478e-4140-83bd-79143ebce02f" examples/events.csv',
	component: () => {
		return <AutomationInvokeBulk />;
	},
});

mappings.set('users:get', {
	params: '<user>',
	instructions: 'Fetch the data for a given user ID',
	example: `courier users:get user123`,
	component: params => {
		return <UsersGet params={params} />;
	},
});
mappings.set('users:set', {
	params: '<user>',
	instructions: "Overwrite a user's profile with the provided data",
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
			option: '--<key> <value>',
			value: 'arbitrary key/value properties for your user',
		},
	],
	example: `courier users:set user123 --email user@example.com`,
	component: params => {
		return <UsersSet params={params} />;
	},
});
mappings.set('users:bulk', {
	params: '<filename>',
	instructions:
		'Bulk import users from a file (csv, json, jsonl, xls, xlsx, .parquet)." For CSVs, we will unpack nested objects based on the header. E.g., "address.city" becomes {"address": {"city": "value"}}. Lodash path syntax is used for created the nested object. Supports wildcard syntax for multiple files, must surround with quotes (see examples)',
	options: [
		{
			option: '--replace',
			value:
				'Replace existing users with the same ID, if not set, will do a merge based on key',
		},
		{
			option: '--keep-flat',
			value:
				'When using a CSV, do not unpack nested objects based on the header. E.g., "address.city" stays as {"address.city": "value"}',
		},
		{
			option: '--remove-nulls',
			value: 'Remove null values from the object before updating the profile',
		},
		{
			option: '--list <List ID>',
			value:
				'Add all users to the specified list. Accepts comma-separated list',
		},
		{
			option: '--tenant <Tenant ID>',
			value:
				'Add all users to the specified tenant. Accepts comma-separated list. Note this will not automatically create the tenant, but the tenant memberships will exist and sending to this tenant_id will still succeed. ',
		},
	],
	example: [
		`courier users:bulk examples/users.csv --replace`,
		`courier users:bulk "examples/users/*.csv" --keep-flat`,
		`courier users:bulk "examples/*.json" --remove-nulls`,
		'courier users:bulk examples/users.parquet --list new-list-id',
		'courier users:bulk examples/users.xlsx --tenant new-tenant-id',
	],
	component: () => <UsersBulk />,
});
mappings.set('users:preferences', {
	params: '<user>',
	instructions: 'Fetch the preferences for a given user ID',
	example: `courier users:preferences user123`,
	options: [
		{
			option: '--verbose',
			value: 'Show the full preference object',
		},
		{
			option: '--url',
			value: 'Generate the Courier Hosted Preference Page URL',
		},
		{
			option: '--brand <brand ID>',
			value:
				'Only used with --url, optionally specify the brand_id. Generate the Courier Hosted Preference Page URL',
		},
		// TODO - add tenants when API is ready
	],
	component: () => {
		return <UsersPreferences />;
	},
});
// NOT IMPLEMENTED YET
// mappings.set('digests:flush', {
// 	params: '<user> <digest>',
// 	instructions: 'Flush any currently queued events for a given user + digest',
// 	example: `courier digests:flush user123 MY_DIGEST_TOPIC`,
// 	component: params => {
// 		return <DigestFlush params={params} />;
// 	},
// });
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
mappings.set('tenants:bulk', {
	params: '<filename>',
	instructions:
		'Bulk import tenants from a file (csv, json, jsonl, xls, xlsx, .parquet). Supports wildcard syntax for multiple files, must surround with quotes (see examples)',
	component: () => <TenantsBulk />,
	options: [
		{
			option: '--merge',
			value:
				'Create or merge existing tenants with the same ID. If the tenant exists, this will get the current values and merge the new values into the existing tenant',
		},
	],
});

mappings.set('inbox:mark-all-read', {
	params: '<user_id>',
	instructions: 'Mark all messages in the inbox as read for a given user',
	example: [`courier inbox:mark-all-read user123`],
	options: [
		// TODO - coming soon when mark all read with tags/tenants is ready
		// {
		// 	option: '--tenant <tenant_id>',
		// 	value: 'The tenant_id to mark all messages as read',
		// },
		// {
		// 	option: '--tag <tag>',
		// 	value: 'The tag to mark all messages as read. Can provide multiple',
		// },
	],
	component: () => {
		return <MarkAllRead />;
	},
});

mappings.set('inbox:archive-all', {
	params: '<user_id>',
	instructions: 'Archive all messages in the inbox for a given user',
	example: [
		`courier inbox:archive-all user123 --before="7 day"`,
		`courier inbox:archive-all user123 --tenant=workspace123`,
		`courier inbox:archive-all user123 --tag product --tag marketing"`,
	],
	options: [
		{
			option: '--tenant <tenant_id>',
			value: 'The tenant_id to mark all messages as read',
		},
		{
			option: '--tag <tag>',
			value: 'The tag to mark all messages as read. Can provide multiple tags',
		},
		{
			option: '--before',
			value:
				'Archive all messages before a given duration using ms: https://www.npmjs.com/package/ms',
		},
		{
			option: '--batch-size',
			value: 'Control the batch size of the archive operation',
		},
		{
			option: '--include-pinned',
			value: 'Include pinned messages in the archive operation',
		},
	],
	component: () => {
		return <ArchiveAll />;
	},
});

mappings.set('inbox:archive-all:bulk', {
	params: '<filename>',
	instructions: 'Archive all messages in the inbox for each user in the file',
	example: [
		`courier inbox:archive-all:bulk archive-inbox.csv --before="7 day"`,
		`courier inbox:archive-all user123 --tenant=workspace123`,
		`courier inbox:archive-all user123 --tag product --tag marketing"`,
	],
	options: mappings.get('inbox:archive-all')!.options,
	component: () => {
		return <ArchiveAllBulk />;
	},
});

mappings.set('templates:list', {
	instructions: 'List all templates in your workspace and export it',
	options: [
		{
			option: '--csv',
			value: 'Output the templates in CSV format',
		},
		{
			option: '--json',
			value: 'Output the templates in JSON format',
		},
		{
			option: '--webhook <webhook_url>',
			value: 'Where to send the JSON output',
		},
		{
			option: '--filename <filename>',
			value: 'Name of the file to save the output',
		},
	],
	component: () => {
		return <TemplatesList />;
	},
});

export default mappings;
