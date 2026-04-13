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

var listsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a list based on the list ID provided.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
	},
	Action:          handleListsRetrieve,
	HideHelpCommand: true,
}

var listsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Create or replace an existing list with the supplied values.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "preferences",
			BodyPath: "preferences",
		},
	},
	Action:          handleListsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"preferences": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "preferences.categories",
			InnerField: "categories",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "preferences.notifications",
			InnerField: "notifications",
		},
	},
})

var listsList = cli.Command{
	Name:    "list",
	Usage:   "Returns all of the lists, with the ability to filter based on a pattern.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "A unique identifier that allows for fetching the next page of lists.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[any]{
			Name:      "pattern",
			Usage:     "\"A pattern used to filter the list items returned. Pattern types supported: exact match on `list_id` or a pattern of one or more pattern parts. you may replace a part with either: `*` to match all parts in that position, or `**` to signify a wildcard `endsWith` pattern match.\"",
			QueryPath: "pattern",
		},
	},
	Action:          handleListsList,
	HideHelpCommand: true,
}

var listsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a list by list ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
	},
	Action:          handleListsDelete,
	HideHelpCommand: true,
}

var listsRestore = cli.Command{
	Name:    "restore",
	Usage:   "Restore a previously deleted list.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
	},
	Action:          handleListsRestore,
	HideHelpCommand: true,
}

func handleListsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
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
	_, err = client.Lists.Get(ctx, cmd.Value("list-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lists retrieve", obj, format, transform)
}

func handleListsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ListUpdateParams{}

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

	return client.Lists.Update(
		ctx,
		cmd.Value("list-id").(string),
		params,
		options...,
	)
}

func handleListsList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ListListParams{}

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
	_, err = client.Lists.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "lists list", obj, format, transform)
}

func handleListsDelete(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
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

	return client.Lists.Delete(ctx, cmd.Value("list-id").(string), options...)
}

func handleListsRestore(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ListRestoreParams{}

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

	return client.Lists.Restore(
		ctx,
		cmd.Value("list-id").(string),
		params,
		options...,
	)
}
