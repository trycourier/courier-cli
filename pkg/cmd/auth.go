// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/trycourier/courier-cli/v3/internal/apiquery"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
	"github.com/trycourier/courier-go/v4"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/urfave/cli/v3"
)

var authIssueToken = cli.Command{
	Name:    "issue-token",
	Usage:   "Returns a new access token.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "expires-in",
			Usage:    "Duration for token expiration. Accepts various time formats:\n- \"2 hours\" - 2 hours from now\n- \"1d\" - 1 day\n- \"3 days\" - 3 days\n- \"10h\" - 10 hours\n- \"2.5 hrs\" - 2.5 hours\n- \"1m\" - 1 minute\n- \"5s\" - 5 seconds\n- \"1y\" - 1 year",
			Required: true,
			BodyPath: "expires_in",
		},
		&requestflag.Flag[string]{
			Name:     "scope",
			Usage:    "Available scopes:\n- `user_id:<user-id>` - Defines which user the token will be scoped to. Multiple can be listed if needed. Ex `user_id:pigeon user_id:bluebird`.\n- `read:messages` - Read messages.\n- `read:user-tokens` - Read user push tokens.\n- `write:user-tokens` - Write user push tokens.\n- `read:brands[:<brand_id>]` - Read brands, optionally restricted to a specific brand_id. Examples `read:brands`, `read:brands:my_brand`.\n- `write:brands[:<brand_id>]` - Write brands, optionally restricted to a specific brand_id. Examples `write:brands`, `write:brands:my_brand`.\n- `inbox:read:messages` - Read inbox messages.\n- `inbox:write:events` - Write inbox events, such as mark message as read.\n- `read:preferences` - Read user preferences.\n- `write:preferences` - Write user preferences.\nExample: `user_id:user123 write:user-tokens inbox:read:messages inbox:write:events read:preferences write:preferences read:brands`",
			Required: true,
			BodyPath: "scope",
		},
	},
	Action:          handleAuthIssueToken,
	HideHelpCommand: true,
}

func handleAuthIssueToken(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := courier.AuthIssueTokenParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Auth.IssueToken(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "auth issue-token",
		Transform:      transform,
	})
}
