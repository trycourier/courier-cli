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

var audiencesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the specified audience by id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "audience-id",
			Required: true,
		},
	},
	Action:          handleAudiencesRetrieve,
	HideHelpCommand: true,
}

var audiencesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Creates or updates audience.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "audience-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "A description of the audience",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter",
			Usage:    "Filter configuration for audience membership containing an array of filter rules",
			BodyPath: "filter",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "The name of the audience",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "operator",
			Usage:    "The logical operator (AND/OR) for the top-level filter",
			BodyPath: "operator",
		},
	},
	Action:          handleAudiencesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "filter.filters",
			Usage:      "Array of filter rules (single conditions or nested groups)",
			InnerField: "filters",
		},
	},
})

var audiencesList = cli.Command{
	Name:    "list",
	Usage:   "Get the audiences associated with the authorization token.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "A unique identifier that allows for fetching the next set of audiences",
			QueryPath: "cursor",
		},
	},
	Action:          handleAudiencesList,
	HideHelpCommand: true,
}

var audiencesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes the specified audience.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "audience-id",
			Required: true,
		},
	},
	Action:          handleAudiencesDelete,
	HideHelpCommand: true,
}

var audiencesListMembers = cli.Command{
	Name:    "list-members",
	Usage:   "Get list of members of an audience.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "audience-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "A unique identifier that allows for fetching the next set of members",
			QueryPath: "cursor",
		},
	},
	Action:          handleAudiencesListMembers,
	HideHelpCommand: true,
}

func handleAudiencesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("audience-id") && len(unusedArgs) > 0 {
		cmd.Set("audience-id", unusedArgs[0])
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
	_, err = client.Audiences.Get(ctx, cmd.Value("audience-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "audiences retrieve", obj, format, transform)
}

func handleAudiencesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("audience-id") && len(unusedArgs) > 0 {
		cmd.Set("audience-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.AudienceUpdateParams{}

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
	_, err = client.Audiences.Update(
		ctx,
		cmd.Value("audience-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "audiences update", obj, format, transform)
}

func handleAudiencesList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.AudienceListParams{}

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
	_, err = client.Audiences.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "audiences list", obj, format, transform)
}

func handleAudiencesDelete(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("audience-id") && len(unusedArgs) > 0 {
		cmd.Set("audience-id", unusedArgs[0])
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

	return client.Audiences.Delete(ctx, cmd.Value("audience-id").(string), options...)
}

func handleAudiencesListMembers(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("audience-id") && len(unusedArgs) > 0 {
		cmd.Set("audience-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.AudienceListMembersParams{}

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
	_, err = client.Audiences.ListMembers(
		ctx,
		cmd.Value("audience-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "audiences list-members", obj, format, transform)
}
