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

var notificationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a notification template. Requires all fields in the notification object.\nTemplates are created in draft state by default.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "notification",
			Usage:    "Core template fields used in POST and PUT request bodies (nested under a `notification` key) and returned at the top level in responses.",
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
		&requestflag.InnerFlag[map[string]any]{
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "notification.routing",
			Usage:      "Routing strategy reference, or null for none.",
			InnerField: "routing",
		},
		&requestflag.InnerFlag[map[string]any]{
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
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
		&requestflag.Flag[*string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response. Omit for the first page.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "event-id",
			Usage:     "Filter to templates linked to this event map ID.",
			QueryPath: "event_id",
		},
		&requestflag.Flag[*bool]{
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
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
			Name:      "id",
			Required:  true,
			PathParam: "id",
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

var notificationsPutContent = requestflag.WithInnerFlags(cli.Command{
	Name:    "put-content",
	Usage:   "Replace the elemental content of a notification template. Overwrites all\nelements in the template with the provided content. Only supported for V2\n(elemental) templates.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "content",
			Usage:    "Elemental content payload. The server defaults `version` when omitted.",
			Required: true,
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "Template state. Defaults to `DRAFT`.",
			Default:  "DRAFT",
			BodyPath: "state",
		},
	},
	Action:          handleNotificationsPutContent,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"content": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "content.elements",
			InnerField: "elements",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.version",
			Usage:      "Content version identifier (e.g., `2022-01-01`). Optional; server defaults when omitted.",
			InnerField: "version",
		},
	},
})

var notificationsPutElement = cli.Command{
	Name:    "put-element",
	Usage:   "Update a single element within a notification template. Only supported for V2\n(elemental) templates.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "element-id",
			Required:  true,
			PathParam: "elementId",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Element type (text, meta, action, image, etc.).",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[[]string]{
			Name:     "channel",
			BodyPath: "channels",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			BodyPath: "data",
		},
		&requestflag.Flag[string]{
			Name:     "if",
			BodyPath: "if",
		},
		&requestflag.Flag[string]{
			Name:     "loop",
			BodyPath: "loop",
		},
		&requestflag.Flag[string]{
			Name:     "ref",
			BodyPath: "ref",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "Template state. Defaults to `DRAFT`.",
			Default:  "DRAFT",
			BodyPath: "state",
		},
	},
	Action:          handleNotificationsPutElement,
	HideHelpCommand: true,
}

var notificationsPutLocale = requestflag.WithInnerFlags(cli.Command{
	Name:    "put-locale",
	Usage:   "Set locale-specific content overrides for a notification template. Each element\noverride must reference an existing element by ID. Only supported for V2\n(elemental) templates.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "locale-id",
			Required:  true,
			PathParam: "localeId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "element",
			Usage:    "Elements with locale-specific content overrides.",
			Required: true,
			BodyPath: "elements",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "Template state. Defaults to `DRAFT`.",
			Default:  "DRAFT",
			BodyPath: "state",
		},
	},
	Action:          handleNotificationsPutLocale,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"element": {
		&requestflag.InnerFlag[string]{
			Name:       "element.id",
			Usage:      "Target element ID.",
			InnerField: "id",
		},
	},
})

var notificationsReplace = requestflag.WithInnerFlags(cli.Command{
	Name:    "replace",
	Usage:   "Replace a notification template. All fields are required.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "notification",
			Usage:    "Core template fields used in POST and PUT request bodies (nested under a `notification` key) and returned at the top level in responses.",
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
		&requestflag.InnerFlag[map[string]any]{
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "notification.routing",
			Usage:      "Routing strategy reference, or null for none.",
			InnerField: "routing",
		},
		&requestflag.InnerFlag[map[string]any]{
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
	Usage:   "Retrieve the content of a notification template. The response shape depends on\nwhether the template uses V1 (blocks/channels) or V2 (elemental) content. Use\nthe `version` query parameter to select draft, published, or a specific\nhistorical version.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "version",
			Usage:     "Accepts `draft`, `published`, or a version string (e.g., `v001`). Defaults to `published`.",
			QueryPath: "version",
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

	params := courier.NotificationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.New(ctx, params, options...)
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
		Title:          "notifications create",
		Transform:      transform,
	})
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

	params := courier.NotificationGetParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "notifications retrieve",
		Transform:      transform,
	})
}

func handleNotificationsList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := courier.NotificationListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.List(ctx, params, options...)
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
		Title:          "notifications list",
		Transform:      transform,
	})
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

	params := courier.NotificationListVersionsParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "notifications list-versions",
		Transform:      transform,
	})
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

	params := courier.NotificationPublishParams{}

	return client.Notifications.Publish(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleNotificationsPutContent(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := courier.NotificationPutContentParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.PutContent(
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "notifications put-content",
		Transform:      transform,
	})
}

func handleNotificationsPutElement(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("element-id") && len(unusedArgs) > 0 {
		cmd.Set("element-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := courier.NotificationPutElementParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.PutElement(
		ctx,
		cmd.Value("element-id").(string),
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
		Title:          "notifications put-element",
		Transform:      transform,
	})
}

func handleNotificationsPutLocale(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("locale-id") && len(unusedArgs) > 0 {
		cmd.Set("locale-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := courier.NotificationPutLocaleParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.PutLocale(
		ctx,
		cmd.Value("locale-id").(string),
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
		Title:          "notifications put-locale",
		Transform:      transform,
	})
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

	params := courier.NotificationReplaceParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "notifications replace",
		Transform:      transform,
	})
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

	params := courier.NotificationGetContentParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.GetContent(
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "notifications retrieve-content",
		Transform:      transform,
	})
}
