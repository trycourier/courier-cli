// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/trycourier/courier-cli/v5/internal/apiquery"
	"github.com/trycourier/courier-cli/v5/internal/requestflag"
	"github.com/trycourier/courier-go/v4"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/urfave/cli/v3"
)

var notificationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a notification template. Requires all fields in the notification object.\nTemplates are created in draft state by default.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "notification",
			Usage:    "Template fields accepted in POST and PUT request bodies, nested under a `notification` key.",
			Required: true,
			BodyPath: "notification",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    `Template state after creation. Case-insensitive input, normalized to uppercase in the response. Defaults to "DRAFT".`,
			Default:  "DRAFT",
			BodyPath: "state",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
		&requestflag.Flag[string]{
			Name:       "x-idempotency-expiration",
			HeaderPath: "x-idempotency-expiration",
		},
	},
	Action:          handleNotificationsCreate,
	HideHelpCommand: true,
}

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
	Usage:   "Lists the workspace's notification templates. Each carries a name, tags, brand,\nrouting, and its draft or published state.",
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
	Usage:   "Archives a notification template, preventing new sends from referencing it. The\ntemplate stays retrievable for its version history.",
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

var notificationsGetMetrics = cli.Command{
	Name:    "get-metrics",
	Usage:   "Fetch the delivery funnel for one Notification Template as a time series — sent,\ndelivered, opened, clicked, errors, and undeliverable — broken out per provider\nand channel inside each bucket. Sum the entries in a bucket for its totals;\nthere is no bucket-level total.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:      "end",
			Usage:     "The end of the window, as an ISO 8601 timestamp with an offset. Must be supplied together with `start`. An `end` in the future is accepted and not clamped — the trailing buckets come back empty.",
			QueryPath: "end",
		},
		&requestflag.Flag[string]{
			Name:      "granularity",
			Usage:     "The size of each bucket in the series. Defaults to `DAY`. `WEEK` buckets start on Sunday. A fine granularity caps the window it can cover: `HOUR` spans at most 7 days and `DAY` at most 90 days, and a wider window returns `400` — request a coarser granularity instead. `WEEK` and `MONTH` are uncapped, subject to the 1000-bucket limit on a single response.",
			Default:   "DAY",
			QueryPath: "granularity",
		},
		&requestflag.Flag[string]{
			Name:      "lookback",
			Usage:     "The length of the window, counted back from now, as an ISO 8601 duration (`P30D`, `P12W`, `PT12H`). Defaults to `P30D`, and is ignored when `start` and `end` are supplied. A malformed or non-positive duration returns `400`.",
			Default:   "P30D",
			QueryPath: "lookback",
		},
		&requestflag.Flag[any]{
			Name:      "start",
			Usage:     "The inclusive start of the window, as an ISO 8601 timestamp with an offset (`2026-04-01T00:00:00Z`). Must be supplied together with `end` and be earlier than it; either one alone returns `400`.",
			QueryPath: "start",
		},
	},
	Action:          handleNotificationsGetMetrics,
	HideHelpCommand: true,
}

var notificationsListVersions = cli.Command{
	Name:    "list-versions",
	Usage:   "Returns a notification template's published versions, most recent first, for\ncomparison or rollback. Paged.",
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
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
		&requestflag.Flag[string]{
			Name:       "x-idempotency-expiration",
			HeaderPath: "x-idempotency-expiration",
		},
	},
	Action:          handleNotificationsPublish,
	HideHelpCommand: true,
}

var notificationsPutContent = requestflag.WithInnerFlags(cli.Command{
	Name:    "put-content",
	Usage:   "Replaces all Elemental content in a template, overwriting every existing\nelement. Supported for V2 templates only, not V1 blocks and channels.",
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
	Usage:   "Replaces one Elemental element in a template, addressed by its element id.\nSupported for V2 templates only, not V1 blocks and channels.",
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
	Usage:   "Sets locale-specific content overrides for a template. Each override must\nreference an element that already exists in the default content.",
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

var notificationsReplace = cli.Command{
	Name:    "replace",
	Usage:   "Replaces a notification template in full, so send every field rather than only\nthe ones you want changed. Publish separately to make it live.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "notification",
			Usage:    "Template fields accepted in POST and PUT request bodies, nested under a `notification` key.",
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
}

var notificationsRetrieveContent = cli.Command{
	Name:    "retrieve-content",
	Usage:   "Returns a template's content and checksum. V2 templates return Elemental\nelements, while V1 templates return blocks and channels instead.",
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

func handleNotificationsGetMetrics(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.NotificationGetMetricsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.GetMetrics(
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
		Title:          "notifications get-metrics",
		Transform:      transform,
	})
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
