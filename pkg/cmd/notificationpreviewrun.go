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

var notificationsPreviewsRunsCreate = cli.Command{
	Name:    "create",
	Usage:   "Render this template's email content on each of the requested devices.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			Usage:    "Template variables to render with, the same shape as the `data` object on a send.",
			BodyPath: "data",
		},
		&requestflag.Flag[[]string]{
			Name:     "device-id",
			Usage:    "The devices to render on, by `PreviewDevice.id`, for a one-off run. Mutually exclusive with `device_set_id`.",
			BodyPath: "device_ids",
		},
		&requestflag.Flag[string]{
			Name:     "device-set-id",
			Usage:    "A saved device set naming the devices to render on. Mutually exclusive with `device_ids`.",
			BodyPath: "device_set_id",
		},
		&requestflag.Flag[string]{
			Name:     "locale",
			Usage:    `Render the template's content for this locale, e.g. "fr-FR".`,
			BodyPath: "locale",
		},
		&requestflag.Flag[string]{
			Name:     "template-version",
			Usage:    "Which version of the template to render. Omit for the latest saved draft, which always exists and is what the editor shows. `published` renders the live version; a zero-padded `v002` renders that specific publish. Versions are 1-based, so `v000` is not a version, and the unpadded `v2` is rejected — that spelling belongs to journeys' AutomationVersionId, a different scheme in which `v0` means published.",
			BodyPath: "template_version",
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
	Action:          handleNotificationsPreviewsRunsCreate,
	HideHelpCommand: true,
}

var notificationsPreviewsRunsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve one of this template's preview runs together with its per-device\nresults.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "preview-run-id",
			Required:  true,
			PathParam: "previewRunId",
		},
	},
	Action:          handleNotificationsPreviewsRunsRetrieve,
	HideHelpCommand: true,
}

var notificationsPreviewsRunsList = cli.Command{
	Name:    "list",
	Usage:   "List this template's preview runs, newest first. Cursor-paginated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
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
	Action:          handleNotificationsPreviewsRunsList,
	HideHelpCommand: true,
}

func handleNotificationsPreviewsRunsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.NotificationPreviewRunNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.Previews.Runs.New(
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
		Title:          "notifications:previews:runs create",
		Transform:      transform,
	})
}

func handleNotificationsPreviewsRunsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("preview-run-id") && len(unusedArgs) > 0 {
		cmd.Set("preview-run-id", unusedArgs[0])
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

	params := courier.NotificationPreviewRunGetParams{
		ID: cmd.Value("id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.Previews.Runs.Get(
		ctx,
		cmd.Value("preview-run-id").(string),
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
		Title:          "notifications:previews:runs retrieve",
		Transform:      transform,
	})
}

func handleNotificationsPreviewsRunsList(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.NotificationPreviewRunListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.Previews.Runs.List(
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
		Title:          "notifications:previews:runs list",
		Transform:      transform,
	})
}
