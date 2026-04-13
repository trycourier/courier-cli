// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
	"github.com/trycourier/courier-cli/v3/internal/apiquery"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
	"github.com/trycourier/courier-go/v4"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/urfave/cli/v3"
)

var providersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new provider configuration. The `provider` field must be a known\nCourier provider key (see catalog).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    `The provider key identifying the type (e.g. "sendgrid", "twilio"). Must be a known Courier provider — see the catalog endpoint for valid keys.`,
			Required: true,
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "alias",
			Usage:    "Optional alias for this configuration.",
			BodyPath: "alias",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "settings",
			Usage:    "Provider-specific settings (snake_case keys). Defaults to an empty object when omitted. Use the catalog endpoint to discover required fields for a given provider — omitting a required field returns a 400 validation error.",
			BodyPath: "settings",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Usage:    `Optional display title. Omit to use "Default Configuration".`,
			BodyPath: "title",
		},
	},
	Action:          handleProvidersCreate,
	HideHelpCommand: true,
}

var providersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a single provider configuration by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleProvidersRetrieve,
	HideHelpCommand: true,
}

var providersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Replace an existing provider configuration. The `provider` key is required and\ndetermines which provider-specific settings schema is applied. All other fields\nare optional — omitted fields are cleared from the stored configuration (this is\na full replacement, not a partial merge). Changing the provider type for an\nexisting configuration is not supported.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "provider",
			Usage:    "The provider key identifying the type. Required on every request because it selects the provider-specific settings schema for validation.",
			Required: true,
			BodyPath: "provider",
		},
		&requestflag.Flag[string]{
			Name:     "alias",
			Usage:    "Updated alias. Omit to clear.",
			BodyPath: "alias",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "settings",
			Usage:    "Provider-specific settings (snake_case keys). Replaces the full settings object — omitted settings fields are removed. Use the catalog endpoint to check required fields.",
			BodyPath: "settings",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Usage:    "Updated display title.",
			BodyPath: "title",
		},
	},
	Action:          handleProvidersUpdate,
	HideHelpCommand: true,
}

var providersList = cli.Command{
	Name:    "list",
	Usage:   "List configured provider integrations for the current workspace. Supports\ncursor-based pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque cursor for fetching the next page.",
			QueryPath: "cursor",
		},
	},
	Action:          handleProvidersList,
	HideHelpCommand: true,
}

var providersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a provider configuration. Returns 409 if the provider is still referenced\nby routing or notifications.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleProvidersDelete,
	HideHelpCommand: true,
}

func handleProvidersCreate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ProviderNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Providers.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "providers create", obj, format, transform)
}

func handleProvidersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Providers.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "providers retrieve", obj, format, transform)
}

func handleProvidersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ProviderUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Providers.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "providers update", obj, format, transform)
}

func handleProvidersList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ProviderListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Providers.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "providers list", obj, format, transform)
}

func handleProvidersDelete(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Providers.Delete(ctx, cmd.Value("id").(string), options...)
}
