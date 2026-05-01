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

var automationsInvokeInvokeAdHoc = requestflag.WithInnerFlags(cli.Command{
	Name:    "invoke-ad-hoc",
	Usage:   "Invoke an ad hoc automation run. This endpoint accepts a JSON payload with a\nseries of automation steps. For information about what steps are available,\ncheckout the ad hoc automation guide\n[here](https://www.courier.com/docs/automations/steps/).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "automation",
			Required: true,
			BodyPath: "automation",
		},
		&requestflag.Flag[*string]{
			Name:     "brand",
			BodyPath: "brand",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			BodyPath: "data",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "profile",
			BodyPath: "profile",
		},
		&requestflag.Flag[*string]{
			Name:     "recipient",
			BodyPath: "recipient",
		},
		&requestflag.Flag[*string]{
			Name:     "template",
			BodyPath: "template",
		},
	},
	Action:          handleAutomationsInvokeInvokeAdHoc,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"automation": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "automation.steps",
			InnerField: "steps",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "automation.cancelation-token",
			InnerField: "cancelation_token",
		},
	},
})

var automationsInvokeInvokeByTemplate = cli.Command{
	Name:    "invoke-by-template",
	Usage:   "Invoke an automation run from an automation template.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[*string]{
			Name:     "recipient",
			Required: true,
			BodyPath: "recipient",
		},
		&requestflag.Flag[*string]{
			Name:     "brand",
			BodyPath: "brand",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			BodyPath: "data",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "profile",
			BodyPath: "profile",
		},
		&requestflag.Flag[*string]{
			Name:     "template",
			BodyPath: "template",
		},
	},
	Action:          handleAutomationsInvokeInvokeByTemplate,
	HideHelpCommand: true,
}

func handleAutomationsInvokeInvokeAdHoc(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.AutomationInvokeInvokeAdHocParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Automations.Invoke.InvokeAdHoc(ctx, params, options...)
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
		Title:          "automations:invoke invoke-ad-hoc",
		Transform:      transform,
	})
}

func handleAutomationsInvokeInvokeByTemplate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("template-id") && len(unusedArgs) > 0 {
		cmd.Set("template-id", unusedArgs[0])
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

	params := courier.AutomationInvokeInvokeByTemplateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Automations.Invoke.InvokeByTemplate(
		ctx,
		cmd.Value("template-id").(string),
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
		Title:          "automations:invoke invoke-by-template",
		Transform:      transform,
	})
}
