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

var previewsArchiveDeviceSet = cli.Command{
	Name:    "archive-device-set",
	Usage:   "Archive a device set. This is a soft delete — the archived set is returned and\nno longer appears in list results. Runs already created against it keep their\nown copy of the device list and are unaffected. The Courier-provided default set\ncannot be archived and returns 409.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "device-set-id",
			Required:  true,
			PathParam: "deviceSetId",
		},
	},
	Action:          handlePreviewsArchiveDeviceSet,
	HideHelpCommand: true,
}

var previewsCreateDeviceSet = cli.Command{
	Name:    "create-device-set",
	Usage:   "Create a named, reusable set of preview devices. Every id must be one listed by\n`GET /previews/devices`; any other is a 422.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "device-id",
			Usage:    "The devices the set contains, by `PreviewDevice.id`. At least one is required.",
			Required: true,
			BodyPath: "device_ids",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handlePreviewsCreateDeviceSet,
	HideHelpCommand: true,
}

var previewsListDeviceSets = cli.Command{
	Name:            "list-device-sets",
	Usage:           "List the workspace's preview sets. Archived sets are not returned.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePreviewsListDeviceSets,
	HideHelpCommand: true,
}

var previewsListDevices = cli.Command{
	Name:            "list-devices",
	Usage:           "List the devices a preview can be rendered on. Reference data, identical for\nevery workspace — these ids are what a device set is built from and what a run\nreports results for.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePreviewsListDevices,
	HideHelpCommand: true,
}

var previewsRetrieveDeviceSet = cli.Command{
	Name:    "retrieve-device-set",
	Usage:   "Retrieve a preview set by ID. Archived sets return 404.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "device-set-id",
			Required:  true,
			PathParam: "deviceSetId",
		},
	},
	Action:          handlePreviewsRetrieveDeviceSet,
	HideHelpCommand: true,
}

var previewsUpdateDeviceSet = cli.Command{
	Name:    "update-device-set",
	Usage:   "Replace a device set. This is a full replace, not a patch — both the name and\nthe device list are always written. The Courier-provided default set cannot be\nchanged and returns 409.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "device-set-id",
			Required:  true,
			PathParam: "deviceSetId",
		},
		&requestflag.Flag[[]string]{
			Name:     "device-id",
			Usage:    "The devices the set contains, by `PreviewDevice.id`. At least one is required.",
			Required: true,
			BodyPath: "device_ids",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name.",
			Required: true,
			BodyPath: "name",
		},
	},
	Action:          handlePreviewsUpdateDeviceSet,
	HideHelpCommand: true,
}

func handlePreviewsArchiveDeviceSet(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("device-set-id") && len(unusedArgs) > 0 {
		cmd.Set("device-set-id", unusedArgs[0])
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
	_, err = client.Previews.ArchiveDeviceSet(ctx, cmd.Value("device-set-id").(string), options...)
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
		Title:          "previews archive-device-set",
		Transform:      transform,
	})
}

func handlePreviewsCreateDeviceSet(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.PreviewNewDeviceSetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Previews.NewDeviceSet(ctx, params, options...)
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
		Title:          "previews create-device-set",
		Transform:      transform,
	})
}

func handlePreviewsListDeviceSets(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Previews.ListDeviceSets(ctx, options...)
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
		Title:          "previews list-device-sets",
		Transform:      transform,
	})
}

func handlePreviewsListDevices(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Previews.ListDevices(ctx, options...)
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
		Title:          "previews list-devices",
		Transform:      transform,
	})
}

func handlePreviewsRetrieveDeviceSet(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("device-set-id") && len(unusedArgs) > 0 {
		cmd.Set("device-set-id", unusedArgs[0])
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
	_, err = client.Previews.GetDeviceSet(ctx, cmd.Value("device-set-id").(string), options...)
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
		Title:          "previews retrieve-device-set",
		Transform:      transform,
	})
}

func handlePreviewsUpdateDeviceSet(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("device-set-id") && len(unusedArgs) > 0 {
		cmd.Set("device-set-id", unusedArgs[0])
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

	params := courier.PreviewUpdateDeviceSetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Previews.UpdateDeviceSet(
		ctx,
		cmd.Value("device-set-id").(string),
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
		Title:          "previews update-device-set",
		Transform:      transform,
	})
}
