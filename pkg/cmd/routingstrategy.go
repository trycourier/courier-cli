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

var routingStrategiesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a routing strategy. Requires a name and routing configuration at minimum.\nChannels and providers default to empty if omitted.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the routing strategy.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "routing",
			Required: true,
			BodyPath: "routing",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "channels",
			BodyPath: "channels",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Optional description of the routing strategy.",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "providers",
			BodyPath: "providers",
		},
		&requestflag.Flag[any]{
			Name:     "tag",
			Usage:    "Optional tags for categorization.",
			BodyPath: "tags",
		},
	},
	Action:          handleRoutingStrategiesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"routing": {
		&requestflag.InnerFlag[[]any]{
			Name:       "routing.channels",
			InnerField: "channels",
		},
		&requestflag.InnerFlag[string]{
			Name:       "routing.method",
			Usage:      `Allowed values: "all", "single".`,
			InnerField: "method",
		},
	},
})

var routingStrategiesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a routing strategy by ID. Returns the full entity including routing\ncontent and metadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRoutingStrategiesRetrieve,
	HideHelpCommand: true,
}

var routingStrategiesList = cli.Command{
	Name:    "list",
	Usage:   "List routing strategies in your workspace. Returns metadata only (no\nrouting/channels/providers content). Use GET /routing-strategies/{id} for full\ndetails.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response. Omit for the first page.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results per page. Default 20, max 100.",
			Default:   20,
			QueryPath: "limit",
		},
	},
	Action:          handleRoutingStrategiesList,
	HideHelpCommand: true,
}

var routingStrategiesArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a routing strategy. The strategy must not have associated notification\ntemplates. Unlink all templates before archiving.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleRoutingStrategiesArchive,
	HideHelpCommand: true,
}

var routingStrategiesListNotifications = cli.Command{
	Name:    "list-notifications",
	Usage:   "List notification templates associated with a routing strategy. Includes\ntemplate metadata only, not full content.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response. Omit for the first page.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results per page. Default 20, max 100.",
			Default:   20,
			QueryPath: "limit",
		},
	},
	Action:          handleRoutingStrategiesListNotifications,
	HideHelpCommand: true,
}

var routingStrategiesReplace = requestflag.WithInnerFlags(cli.Command{
	Name:    "replace",
	Usage:   "Replace a routing strategy. Full document replacement; the caller must send the\ncomplete desired state. Missing optional fields are cleared.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the routing strategy.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "routing",
			Required: true,
			BodyPath: "routing",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "channels",
			BodyPath: "channels",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Optional description. Omit or null to clear.",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "providers",
			BodyPath: "providers",
		},
		&requestflag.Flag[any]{
			Name:     "tag",
			Usage:    "Optional tags. Omit or null to clear.",
			BodyPath: "tags",
		},
	},
	Action:          handleRoutingStrategiesReplace,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"routing": {
		&requestflag.InnerFlag[[]any]{
			Name:       "routing.channels",
			InnerField: "channels",
		},
		&requestflag.InnerFlag[string]{
			Name:       "routing.method",
			Usage:      `Allowed values: "all", "single".`,
			InnerField: "method",
		},
	},
})

func handleRoutingStrategiesCreate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.RoutingStrategyNewParams{}

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
	_, err = client.RoutingStrategies.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "routing-strategies create", obj, format, explicitFormat, transform)
}

func handleRoutingStrategiesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.RoutingStrategies.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "routing-strategies retrieve", obj, format, explicitFormat, transform)
}

func handleRoutingStrategiesList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.RoutingStrategyListParams{}

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
	_, err = client.RoutingStrategies.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "routing-strategies list", obj, format, explicitFormat, transform)
}

func handleRoutingStrategiesArchive(ctx context.Context, cmd *cli.Command) error {
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

	return client.RoutingStrategies.Archive(ctx, cmd.Value("id").(string), options...)
}

func handleRoutingStrategiesListNotifications(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.RoutingStrategyListNotificationsParams{}

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
	_, err = client.RoutingStrategies.ListNotifications(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "routing-strategies list-notifications", obj, format, explicitFormat, transform)
}

func handleRoutingStrategiesReplace(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.RoutingStrategyReplaceParams{}

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
	_, err = client.RoutingStrategies.Replace(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "routing-strategies replace", obj, format, explicitFormat, transform)
}
