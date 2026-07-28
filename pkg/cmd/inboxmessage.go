// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/trycourier/courier-cli/v3/internal/apiquery"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
	"github.com/trycourier/courier-go/v4"
	"github.com/urfave/cli/v3"
)

var inboxMessagesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a user's inbox message. The message is removed from every inbox read (it\nstops appearing in the recipient's Inbox); it can be restored.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "message-id",
			Required:  true,
			PathParam: "message_id",
		},
	},
	Action:          handleInboxMessagesDelete,
	HideHelpCommand: true,
}

var inboxMessagesRestore = cli.Command{
	Name:    "restore",
	Usage:   "Restore a previously deleted inbox message.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "message-id",
			Required:  true,
			PathParam: "message_id",
		},
	},
	Action:          handleInboxMessagesRestore,
	HideHelpCommand: true,
}

func handleInboxMessagesDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Inbox.Messages.Delete(ctx, cmd.Value("message-id").(string), options...)
}

func handleInboxMessagesRestore(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.InboxMessageRestoreParams{}

	return client.Inbox.Messages.Restore(
		ctx,
		cmd.Value("message-id").(string),
		params,
		options...,
	)
}
