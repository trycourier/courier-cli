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

var notificationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a notification template. Requires all fields in the notification object.\nTemplates are created in draft state by default.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "notification",
			Usage:    "Full document shape used in POST and PUT request bodies, and returned inside the GET response envelope.",
			Required: true,
			BodyPath: "notification",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    `Template state after creation. Case-insensitive input, normalized to uppercase in the response. Defaults to "DRAFT".`,
			Default:  "DRAFT",
			BodyPath: "state",
		},
	},
	Action:          handleNotificationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"notification": {
		&requestflag.InnerFlag[any]{
			Name:       "notification.brand",
			Usage:      "Brand reference, or null for no brand.",
			InnerField: "brand",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "notification.content",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "notification.name",
			Usage:      "Display name for the template.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[any]{
			Name:       "notification.routing",
			Usage:      "Routing strategy reference, or null for none.",
			InnerField: "routing",
		},
		&requestflag.InnerFlag[any]{
			Name:       "notification.subscription",
			Usage:      "Subscription topic reference, or null for none.",
			InnerField: "subscription",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "notification.tags",
			Usage:      "Tags for categorization. Send empty array for none.",
			InnerField: "tags",
		},
	},
})

var notificationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a notification template by ID. Returns the published version by\ndefault. Pass version=draft to retrieve an unpublished template.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "version",
			Usage:     `Version to retrieve. One of "draft", "published", or a version string like "v001". Defaults to "published".`,
			QueryPath: "version",
		},
	},
	Action:          handleNotificationsRetrieve,
	HideHelpCommand: true,
}

var notificationsList = cli.Command{
	Name:    "list",
	Usage:   "List notification templates in your workspace.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response. Omit for the first page.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "event-id",
			Usage:     "Filter to templates linked to this event map ID.",
			QueryPath: "event_id",
		},
		&requestflag.Flag[any]{
			Name:      "notes",
			Usage:     "Include template notes in the response. Only applies to legacy templates.",
			QueryPath: "notes",
		},
	},
	Action:          handleNotificationsList,
	HideHelpCommand: true,
}

var notificationsArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a notification template.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleNotificationsArchive,
	HideHelpCommand: true,
}

var notificationsListVersions = cli.Command{
	Name:    "list-versions",
	Usage:   "List versions of a notification template.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response. Omit for the first page.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of versions to return per page. Default 10, max 10.",
			Default:   10,
			QueryPath: "limit",
		},
	},
	Action:          handleNotificationsListVersions,
	HideHelpCommand: true,
}

var notificationsPublish = cli.Command{
	Name:    "publish",
	Usage:   "Publish a notification template. Publishes the current draft by default. Pass a\nversion in the request body to publish a specific historical version.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "version",
			Usage:    `Historical version to publish (e.g. "v001"). Omit to publish the current draft.`,
			BodyPath: "version",
		},
	},
	Action:          handleNotificationsPublish,
	HideHelpCommand: true,
}

var notificationsReplace = requestflag.WithInnerFlags(cli.Command{
	Name:    "replace",
	Usage:   "Replace a notification template. All fields are required.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "notification",
			Usage:    "Full document shape used in POST and PUT request bodies, and returned inside the GET response envelope.",
			Required: true,
			BodyPath: "notification",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    `Template state after update. Case-insensitive input, normalized to uppercase in the response. Defaults to "DRAFT".`,
			Default:  "DRAFT",
			BodyPath: "state",
		},
	},
	Action:          handleNotificationsReplace,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"notification": {
		&requestflag.InnerFlag[any]{
			Name:       "notification.brand",
			Usage:      "Brand reference, or null for no brand.",
			InnerField: "brand",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "notification.content",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "notification.name",
			Usage:      "Display name for the template.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[any]{
			Name:       "notification.routing",
			Usage:      "Routing strategy reference, or null for none.",
			InnerField: "routing",
		},
		&requestflag.InnerFlag[any]{
			Name:       "notification.subscription",
			Usage:      "Subscription topic reference, or null for none.",
			InnerField: "subscription",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "notification.tags",
			Usage:      "Tags for categorization. Send empty array for none.",
			InnerField: "tags",
		},
	},
})

var notificationsRetrieveContent = cli.Command{
	Name:    "retrieve-content",
	Usage:   "Perform retrieve-content operation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleNotificationsRetrieveContent,
	HideHelpCommand: true,
}

func handleNotificationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.NotificationNewParams{}

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
	_, err = client.Notifications.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notifications create", obj, format, transform)
}

func handleNotificationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.NotificationGetParams{}

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
	_, err = client.Notifications.Get(
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
	return ShowJSON(os.Stdout, "notifications retrieve", obj, format, transform)
}

func handleNotificationsList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.NotificationListParams{}

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
	_, err = client.Notifications.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notifications list", obj, format, transform)
}

func handleNotificationsArchive(ctx context.Context, cmd *cli.Command) error {
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

	return client.Notifications.Archive(ctx, cmd.Value("id").(string), options...)
}

func handleNotificationsListVersions(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.NotificationListVersionsParams{}

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
	_, err = client.Notifications.ListVersions(
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
	return ShowJSON(os.Stdout, "notifications list-versions", obj, format, transform)
}

func handleNotificationsPublish(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.NotificationPublishParams{}

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

	return client.Notifications.Publish(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleNotificationsReplace(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.NotificationReplaceParams{}

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
	_, err = client.Notifications.Replace(
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
	return ShowJSON(os.Stdout, "notifications replace", obj, format, transform)
}

func handleNotificationsRetrieveContent(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Notifications.GetContent(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notifications retrieve-content", obj, format, transform)
}
