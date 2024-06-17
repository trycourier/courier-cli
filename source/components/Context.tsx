import {CourierClient} from '@trycourier/courier';
import {IssueTokenResponse} from '@trycourier/courier/api/index.js';
import React, {createContext, useContext, useState} from 'react';
import yargsParser from 'yargs-parser';

const CliContext = createContext({});

export type TEnvironment = 'test' | 'production';
export type TAPIRouting = 'normal' | 'simulated';
export type TDocumentScope = 'Published' | 'Draft' | 'Submitted';

const PREFIX = 'COURIER_AUTH_TOKEN';

type TGetJWT = (
	user_id: string,
	additional_scopes: TJWTScope[],
	other?: {
		write_brands: string[];
		expires_in?: string;
	},
) => Promise<IssueTokenResponse>;

interface ICliContextProvider {
	args: string[];
	mappings: Map<string, IMapping>;
	current?: string;
	latest?: string;
	children: JSX.Element[] | JSX.Element;
}

interface ICliContext {
	map?: string;
	env_var: string;
	apikey?: string;
	environment: TEnvironment;
	routing: TAPIRouting;
	document_scope: TDocumentScope;
	mappings: Map<string, IMapping>;
	mapping?: IMapping;
	parsedParams: yargsParser.Arguments;
	args: string[];
	url: string;
	courier: CourierClient;
	version: {current?: string; latest?: string};
	setVersion: React.Dispatch<
		React.SetStateAction<{current?: string; latest?: string}>
	>;
	getJWT: TGetJWT;
}
type IUseCliContext = () => ICliContext;

export const useCliContext: IUseCliContext = () =>
	useContext(CliContext as any);

export const CliContextProvider = ({
	args,
	mappings,
	children,
}: ICliContextProvider) => {
	const [version, setVersion] = useState<{current?: string; latest?: string}>(
		{},
	);
	const [map, ...params] = args;
	const mapping = mappings.get(map || '');
	const parsedParams = yargsParser(params);
	const environment =
		parsedParams['production'] || parsedParams['P'] ? 'production' : 'test';
	const routing =
		parsedParams['mock'] || parsedParams['M'] ? 'simulated' : 'normal';

	let document_scope: TDocumentScope = 'Published';

	if (parsedParams['draft'] || parsedParams['D']) {
		document_scope = 'Draft';
	} else if (parsedParams['submitted'] || parsedParams['S']) {
		document_scope = 'Submitted';
	}

	const env_var = getApiKeyVariable({environment, routing, document_scope});
	const apikey = parsedParams['apikey'] ?? process.env[env_var];

	const url: string =
		parsedParams['apiurl'] ??
		process.env['COURIER_API_URL'] ??
		'https://api.courier.com';

	const courier = new CourierClient({
		authorizationToken: apikey,
		environment: url,
	});

	const getJWT: TGetJWT = async (user_id, additional_scopes, options) => {
		const {write_brands = [], expires_in = '5 min'} = options || {};
		let scopes = [`user_id:${user_id}`, ...additional_scopes];
		if (write_brands.length) {
			write_brands.forEach(brand => {
				scopes.push(`write:brands:${brand}`);
			});
		}
		const token = await courier.authTokens.issueToken({
			scope: scopes.join(' '),
			expires_in,
		});
		return token;
	};

	const context: ICliContext = {
		apikey,
		mappings,
		mapping,
		parsedParams,
		args,
		environment,
		routing,
		document_scope,
		env_var,
		courier,
		url,
		map,
		getJWT,
		version,
		setVersion,
	};

	return <CliContext.Provider value={context}>{children}</CliContext.Provider>;
};

export const getApiKeyFlags = ({
	environment,
	routing,
	document_scope,
}: {
	environment: TEnvironment;
	routing: TAPIRouting;
	document_scope: TDocumentScope;
}) => {
	let extra_flags = [];
	if (environment === 'production') {
		extra_flags.push('-P');
	}
	if (routing === 'simulated') {
		extra_flags.push('-M');
	}
	if (document_scope === 'Draft') {
		extra_flags.push('-D');
	} else if (document_scope === 'Submitted') {
		extra_flags.push('-S');
	}
	return extra_flags.join(' ');
};

export const getApiKeyVariable = ({
	environment,
	routing,
	document_scope,
}: {
	environment: TEnvironment;
	routing: TAPIRouting;
	document_scope: TDocumentScope;
}) => {
	let var_name = [PREFIX];
	if (environment === 'test') {
		var_name.push('TEST');
	}
	if (routing === 'simulated') {
		var_name.push('MOCK');
	}
	if (document_scope !== 'Published') {
		var_name.push(document_scope.toUpperCase());
	}
	return var_name.join('_').toUpperCase();
};
