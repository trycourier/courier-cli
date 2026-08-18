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

var broadcastsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a broadcast. Provisions a private notification template for the broadcast\nand returns the new broadcast in the draft state. Exactly one channel is\nrequired.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "channel",
			Usage:    "The single delivery channel for this broadcast.",
			Required: true,
			BodyPath: "channel",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleBroadcastsCreate,
	HideHelpCommand: true,
}

var broadcastsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a broadcast by ID. Archived broadcasts return 404.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
		},
	},
	Action:          handleBroadcastsRetrieve,
	HideHelpCommand: true,
}

var broadcastsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a broadcast's name. Content is edited via the broadcast's notification\ntemplate, not this endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "New human-readable name.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handleBroadcastsUpdate,
	HideHelpCommand: true,
}

var broadcastsList = cli.Command{
	Name:    "list",
	Usage:   "List broadcasts in your workspace. Cursor-paginated; returns broadcasts\nnewest-first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "cursor",
			Usage:     "Opaque pagination cursor from a previous response. Omit for the first page.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results per page.",
			QueryPath: "limit",
		},
	},
	Action:          handleBroadcastsList,
	HideHelpCommand: true,
}

var broadcastsArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a broadcast. This is a soft delete — the archived broadcast is returned\nand no longer appears in list results.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
		},
	},
	Action:          handleBroadcastsArchive,
	HideHelpCommand: true,
}

var broadcastsCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a broadcast's pending schedule, returning it to the draft state. Only\nvalid for a scheduled broadcast.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
		},
	},
	Action:          handleBroadcastsCancel,
	HideHelpCommand: true,
}

var broadcastsDuplicate = cli.Command{
	Name:    "duplicate",
	Usage:   "Duplicate a broadcast (and its template) into a new draft named \"{source name}\n(copy)\".",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
		},
	},
	Action:          handleBroadcastsDuplicate,
	HideHelpCommand: true,
}

var broadcastsPutContent = requestflag.WithInnerFlags(cli.Command{
	Name:    "put-content",
	Usage:   "Author the broadcast's content by replacing the draft elemental content of its\nprivate notification template. The draft is published automatically when the\nbroadcast is sent or scheduled.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
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
	Action:          handleBroadcastsPutContent,
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

var broadcastsRetrieveContent = cli.Command{
	Name:    "retrieve-content",
	Usage:   "Retrieve the broadcast's content — the elemental content of its private\nnotification template. Defaults to the working draft, since broadcast content is\nauthored as a draft until the broadcast is sent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
		},
		&requestflag.Flag[string]{
			Name:      "version",
			Usage:     "Accepts `draft`, `published`, or a version string (e.g. `v001`). Defaults to `draft`.",
			QueryPath: "version",
		},
	},
	Action:          handleBroadcastsRetrieveContent,
	HideHelpCommand: true,
}

var broadcastsSchedule = cli.Command{
	Name:    "schedule",
	Usage:   "Schedule a broadcast for a future send to a list or audience. Publishes the\nbroadcast template first. Not allowed once the broadcast is sending or sent. For\nan immediate send use POST /broadcasts/{broadcastId}/send.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-id",
			Usage:    "ID of the target list or audience.",
			Required: true,
			BodyPath: "recipient_id",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-type",
			Usage:    "Whether the broadcast targets a list or an audience.",
			Required: true,
			BodyPath: "recipient_type",
		},
		&requestflag.Flag[string]{
			Name:     "scheduled-to",
			Usage:    "Wall-clock timestamp of the future send, no timezone offset (e.g. \"2026-07-21T20:00:00\"). The zone is given by `timezone`.",
			Required: true,
			BodyPath: "scheduled_to",
		},
		&requestflag.Flag[string]{
			Name:     "timezone",
			Usage:    "IANA timezone for the scheduled send (e.g. America/New_York).",
			BodyPath: "timezone",
		},
	},
	Action:          handleBroadcastsSchedule,
	HideHelpCommand: true,
}

var broadcastsSend = cli.Command{
	Name:    "send",
	Usage:   "Send a broadcast immediately to a list or audience. Publishes the broadcast\ntemplate first. Not allowed once the broadcast is sending or sent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "broadcast-id",
			Required:  true,
			PathParam: "broadcastId",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-id",
			Usage:    "ID of the target list or audience.",
			Required: true,
			BodyPath: "recipient_id",
		},
		&requestflag.Flag[string]{
			Name:     "recipient-type",
			Usage:    "Whether the broadcast targets a list or an audience.",
			Required: true,
			BodyPath: "recipient_type",
		},
	},
	Action:          handleBroadcastsSend,
	HideHelpCommand: true,
}

func handleBroadcastsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.BroadcastNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.New(ctx, params, options...)
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
		Title:          "broadcasts create",
		Transform:      transform,
	})
}

func handleBroadcastsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
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
	_, err = client.Broadcasts.Get(ctx, cmd.Value("broadcast-id").(string), options...)
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
		Title:          "broadcasts retrieve",
		Transform:      transform,
	})
}

func handleBroadcastsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
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

	params := courier.BroadcastUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.Update(
		ctx,
		cmd.Value("broadcast-id").(string),
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
		Title:          "broadcasts update",
		Transform:      transform,
	})
}

func handleBroadcastsList(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.BroadcastListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.List(ctx, params, options...)
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
		Title:          "broadcasts list",
		Transform:      transform,
	})
}

func handleBroadcastsArchive(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
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
	_, err = client.Broadcasts.Archive(ctx, cmd.Value("broadcast-id").(string), options...)
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
		Title:          "broadcasts archive",
		Transform:      transform,
	})
}

func handleBroadcastsCancel(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
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
	_, err = client.Broadcasts.Cancel(ctx, cmd.Value("broadcast-id").(string), options...)
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
		Title:          "broadcasts cancel",
		Transform:      transform,
	})
}

func handleBroadcastsDuplicate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
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
	_, err = client.Broadcasts.Duplicate(ctx, cmd.Value("broadcast-id").(string), options...)
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
		Title:          "broadcasts duplicate",
		Transform:      transform,
	})
}

func handleBroadcastsPutContent(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
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

	params := courier.BroadcastPutContentParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.PutContent(
		ctx,
		cmd.Value("broadcast-id").(string),
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
		Title:          "broadcasts put-content",
		Transform:      transform,
	})
}

func handleBroadcastsRetrieveContent(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
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

	params := courier.BroadcastGetContentParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.GetContent(
		ctx,
		cmd.Value("broadcast-id").(string),
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
		Title:          "broadcasts retrieve-content",
		Transform:      transform,
	})
}

func handleBroadcastsSchedule(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
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

	params := courier.BroadcastScheduleParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.Schedule(
		ctx,
		cmd.Value("broadcast-id").(string),
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
		Title:          "broadcasts schedule",
		Transform:      transform,
	})
}

func handleBroadcastsSend(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("broadcast-id") && len(unusedArgs) > 0 {
		cmd.Set("broadcast-id", unusedArgs[0])
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

	params := courier.BroadcastSendParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Broadcasts.Send(
		ctx,
		cmd.Value("broadcast-id").(string),
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
		Title:          "broadcasts send",
		Transform:      transform,
	})
}
