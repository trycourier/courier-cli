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

var providersCatalogList = cli.Command{
	Name:    "list",
	Usage:   "Returns the catalog of available provider types with their display names,\ndescriptions, and configuration schema fields (snake_case, with `type` and\n`required`). Providers with no configurable schema return only `provider`,\n`name`, and `description`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "channel",
			Usage:     "Exact match (case-insensitive) against the provider channel taxonomy (e.g. `email`, `sms`, `push`).",
			QueryPath: "channel",
		},
		&requestflag.Flag[string]{
			Name:      "keys",
			Usage:     "Comma-separated provider keys to filter by (e.g. `sendgrid,twilio`).",
			QueryPath: "keys",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Case-insensitive substring match against the provider display name.",
			QueryPath: "name",
		},
	},
	Action:          handleProvidersCatalogList,
	HideHelpCommand: true,
}

func handleProvidersCatalogList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ProviderCatalogListParams{}

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
	_, err = client.Providers.Catalog.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "providers:catalog list", obj, format, explicitFormat, transform)
}
