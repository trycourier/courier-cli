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

var messagesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch the status of a message you've previously sent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
	},
	Action:          handleMessagesRetrieve,
	HideHelpCommand: true,
}

var messagesList = cli.Command{
	Name:    "list",
	Usage:   "Fetch the statuses of messages you've previously sent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "archived",
			Usage:     "A boolean value that indicates whether archived messages should be included in the response.",
			QueryPath: "archived",
		},
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "A unique identifier that allows for fetching the next set of messages.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[any]{
			Name:      "enqueued-after",
			Usage:     "The enqueued datetime of a message to filter out messages received before.",
			QueryPath: "enqueued_after",
		},
		&requestflag.Flag[any]{
			Name:      "event",
			Usage:     "A unique identifier representing the event that was used to send the event.",
			QueryPath: "event",
		},
		&requestflag.Flag[any]{
			Name:      "list",
			Usage:     "A unique identifier representing the list the message was sent to.",
			QueryPath: "list",
		},
		&requestflag.Flag[any]{
			Name:      "message-id",
			Usage:     "A unique identifier representing the message_id returned from either /send or /send/list.",
			QueryPath: "messageId",
		},
		&requestflag.Flag[any]{
			Name:      "notification",
			Usage:     "A unique identifier representing the notification that was used to send the event.",
			QueryPath: "notification",
		},
		&requestflag.Flag[[]any]{
			Name:      "provider",
			Usage:     "The key assocated to the provider you want to filter on. E.g., sendgrid, inbox, twilio, slack, msteams, etc. Allows multiple values to be set in query parameters.",
			QueryPath: "provider",
		},
		&requestflag.Flag[any]{
			Name:      "recipient",
			Usage:     "A unique identifier representing the recipient associated with the requested profile.",
			QueryPath: "recipient",
		},
		&requestflag.Flag[[]any]{
			Name:      "status",
			Usage:     "An indicator of the current status of the message. Allows multiple values to be set in query parameters.",
			QueryPath: "status",
		},
		&requestflag.Flag[[]any]{
			Name:      "tag",
			Usage:     "A tag placed in the metadata.tags during a notification send. Allows multiple values to be set in query parameters.",
			QueryPath: "tag",
		},
		&requestflag.Flag[any]{
			Name:      "tags",
			Usage:     "A comma delimited list of 'tags'. Messages will be returned if they match any of the tags passed in.",
			QueryPath: "tags",
		},
		&requestflag.Flag[any]{
			Name:      "tenant-id",
			Usage:     "Messages sent with the context of a Tenant",
			QueryPath: "tenant_id",
		},
		&requestflag.Flag[any]{
			Name:      "trace-id",
			Usage:     "The unique identifier used to trace the requests",
			QueryPath: "traceId",
		},
	},
	Action:          handleMessagesList,
	HideHelpCommand: true,
}

var messagesCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a message that is currently in the process of being delivered. A\nwell-formatted API call to the cancel message API will return either `200`\nstatus code for a successful cancellation or `409` status code for an\nunsuccessful cancellation. Both cases will include the actual message record in\nthe response body (see details below).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
	},
	Action:          handleMessagesCancel,
	HideHelpCommand: true,
}

var messagesContent = cli.Command{
	Name:    "content",
	Usage:   "Get message content",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
	},
	Action:          handleMessagesContent,
	HideHelpCommand: true,
}

var messagesHistory = cli.Command{
	Name:    "history",
	Usage:   "Fetch the array of events of a message you've previously sent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "type",
			Usage:     "A supported Message History type that will filter the events returned.",
			QueryPath: "type",
		},
	},
	Action:          handleMessagesHistory,
	HideHelpCommand: true,
}

func handleMessagesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
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
	_, err = client.Messages.Get(ctx, cmd.Value("message-id").(string), options...)
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
		Title:          "messages retrieve",
		Transform:      transform,
	})
}

func handleMessagesList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.MessageListParams{}

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
	_, err = client.Messages.List(ctx, params, options...)
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
		Title:          "messages list",
		Transform:      transform,
	})
}

func handleMessagesCancel(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
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
	_, err = client.Messages.Cancel(ctx, cmd.Value("message-id").(string), options...)
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
		Title:          "messages cancel",
		Transform:      transform,
	})
}

func handleMessagesContent(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
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
	_, err = client.Messages.Content(ctx, cmd.Value("message-id").(string), options...)
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
		Title:          "messages content",
		Transform:      transform,
	})
}

func handleMessagesHistory(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-id") && len(unusedArgs) > 0 {
		cmd.Set("message-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.MessageHistoryParams{}

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
	_, err = client.Messages.History(
		ctx,
		cmd.Value("message-id").(string),
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
		Title:          "messages history",
		Transform:      transform,
	})
}
