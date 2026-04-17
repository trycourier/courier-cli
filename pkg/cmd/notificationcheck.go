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

var notificationsChecksUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Replace the checks for a notification template submission.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "submission-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "check",
			Required: true,
			BodyPath: "checks",
		},
	},
	Action:          handleNotificationsChecksUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"check": {
		&requestflag.InnerFlag[string]{
			Name:       "check.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "check.status",
			Usage:      `Allowed values: "RESOLVED", "FAILED", "PENDING".`,
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "check.type",
			Usage:      `Allowed values: "custom".`,
			InnerField: "type",
		},
	},
})

var notificationsChecksList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve the checks for a notification template submission.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "submission-id",
			Required: true,
		},
	},
	Action:          handleNotificationsChecksList,
	HideHelpCommand: true,
}

var notificationsChecksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Cancel a notification template submission.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "submission-id",
			Required: true,
		},
	},
	Action:          handleNotificationsChecksDelete,
	HideHelpCommand: true,
}

func handleNotificationsChecksUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("submission-id") && len(unusedArgs) > 0 {
		cmd.Set("submission-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.NotificationCheckUpdateParams{
		ID: cmd.Value("id").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.Checks.Update(
		ctx,
		cmd.Value("submission-id").(string),
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
		Title:          "notifications:checks update",
		Transform:      transform,
	})
}

func handleNotificationsChecksList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("submission-id") && len(unusedArgs) > 0 {
		cmd.Set("submission-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.NotificationCheckListParams{
		ID: cmd.Value("id").(string),
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
	_, err = client.Notifications.Checks.List(
		ctx,
		cmd.Value("submission-id").(string),
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
		Title:          "notifications:checks list",
		Transform:      transform,
	})
}

func handleNotificationsChecksDelete(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("submission-id") && len(unusedArgs) > 0 {
		cmd.Set("submission-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.NotificationCheckDeleteParams{
		ID: cmd.Value("id").(string),
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

	return client.Notifications.Checks.Delete(
		ctx,
		cmd.Value("submission-id").(string),
		params,
		options...,
	)
}
