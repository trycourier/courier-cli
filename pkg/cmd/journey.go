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

var journeysList = cli.Command{
	Name:    "list",
	Usage:   "Get the list of journeys.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "A cursor token for pagination. Use the cursor from the previous response to fetch the next page of results.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "version",
			Usage:     "The version of journeys to retrieve. Accepted values are published (for published journeys) or draft (for draft journeys). Defaults to published.",
			Default:   "published",
			QueryPath: "version",
		},
	},
	Action:          handleJourneysList,
	HideHelpCommand: true,
}

var journeysInvoke = cli.Command{
	Name:    "invoke",
	Usage:   "Invoke a journey run from a journey template.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "template-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			Usage:    "Data payload passed to the journey. The expected shape can be predefined using the schema builder in the journey editor. This data is available in journey steps for condition evaluation and template variable interpolation. Can also contain user identifiers (user_id, userId, anonymousId) if not provided elsewhere.",
			BodyPath: "data",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "profile",
			Usage:    "Profile data for the user. Can contain contact information (email, phone_number), user identifiers (user_id, userId, anonymousId), or any custom profile fields. Profile fields are merged with any existing stored profile for the user. Include context.tenant_id to load a tenant-scoped profile for multi-tenant scenarios.",
			BodyPath: "profile",
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    "A unique identifier for the user. If not provided, the system will attempt to resolve the user identifier from profile or data objects.",
			BodyPath: "user_id",
		},
	},
	Action:          handleJourneysInvoke,
	HideHelpCommand: true,
}

func handleJourneysList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.JourneyListParams{}

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
	_, err = client.Journeys.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "journeys list", obj, format, transform)
}

func handleJourneysInvoke(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("template-id") && len(unusedArgs) > 0 {
		cmd.Set("template-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.JourneyInvokeParams{}

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
	_, err = client.Journeys.Invoke(
		ctx,
		cmd.Value("template-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "journeys invoke", obj, format, transform)
}
