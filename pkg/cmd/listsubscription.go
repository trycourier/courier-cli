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

var listsSubscriptionsList = cli.Command{
	Name:    "list",
	Usage:   "Get the list's subscriptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "A unique identifier that allows for fetching the next set of list subscriptions",
			QueryPath: "cursor",
		},
	},
	Action:          handleListsSubscriptionsList,
	HideHelpCommand: true,
}

var listsSubscriptionsAdd = requestflag.WithInnerFlags(cli.Command{
	Name:    "add",
	Usage:   "Subscribes additional users to the list, without modifying existing\nsubscriptions. If the list does not exist, it will be automatically created.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "recipient",
			Required: true,
			BodyPath: "recipients",
		},
	},
	Action:          handleListsSubscriptionsAdd,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"recipient": {
		&requestflag.InnerFlag[string]{
			Name:       "recipient.recipient-id",
			InnerField: "recipientId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "recipient.preferences",
			InnerField: "preferences",
		},
	},
})

var listsSubscriptionsSubscribe = requestflag.WithInnerFlags(cli.Command{
	Name:    "subscribe",
	Usage:   "Subscribes the users to the list, overwriting existing subscriptions. If the\nlist does not exist, it will be automatically created.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "recipient",
			Required: true,
			BodyPath: "recipients",
		},
	},
	Action:          handleListsSubscriptionsSubscribe,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"recipient": {
		&requestflag.InnerFlag[string]{
			Name:       "recipient.recipient-id",
			InnerField: "recipientId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "recipient.preferences",
			InnerField: "preferences",
		},
	},
})

var listsSubscriptionsSubscribeUser = requestflag.WithInnerFlags(cli.Command{
	Name:    "subscribe-user",
	Usage:   "Subscribe a user to an existing list (note: if the List does not exist, it will\nbe automatically created).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "preferences",
			BodyPath: "preferences",
		},
	},
	Action:          handleListsSubscriptionsSubscribeUser,
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

var listsSubscriptionsUnsubscribeUser = cli.Command{
	Name:    "unsubscribe-user",
	Usage:   "Delete a subscription to a list by list ID and user ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "list-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
	},
	Action:          handleListsSubscriptionsUnsubscribeUser,
	HideHelpCommand: true,
}

func handleListsSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ListSubscriptionListParams{}

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
	_, err = client.Lists.Subscriptions.List(
		ctx,
		cmd.Value("list-id").(string),
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "lists:subscriptions list",
		Transform:      transform,
	})
}

func handleListsSubscriptionsAdd(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ListSubscriptionAddParams{}

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

	return client.Lists.Subscriptions.Add(
		ctx,
		cmd.Value("list-id").(string),
		params,
		options...,
	)
}

func handleListsSubscriptionsSubscribe(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("list-id") && len(unusedArgs) > 0 {
		cmd.Set("list-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ListSubscriptionSubscribeParams{}

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

	return client.Lists.Subscriptions.Subscribe(
		ctx,
		cmd.Value("list-id").(string),
		params,
		options...,
	)
}

func handleListsSubscriptionsSubscribeUser(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ListSubscriptionSubscribeUserParams{
		ListID: cmd.Value("list-id").(string),
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

	return client.Lists.Subscriptions.SubscribeUser(
		ctx,
		cmd.Value("user-id").(string),
		params,
		options...,
	)
}

func handleListsSubscriptionsUnsubscribeUser(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.ListSubscriptionUnsubscribeUserParams{
		ListID: cmd.Value("list-id").(string),
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

	return client.Lists.Subscriptions.UnsubscribeUser(
		ctx,
		cmd.Value("user-id").(string),
		params,
		options...,
	)
}
