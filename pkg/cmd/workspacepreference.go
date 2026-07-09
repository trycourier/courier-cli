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

var workspacePreferencesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a workspace preference. The workspace preference id is generated and\nreturned. Topics are created inside a workspace preference via POST\n/preferences/sections/{section_id}/topics.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the workspace preference.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "Optional description shown under the section on the hosted preferences page.",
			BodyPath: "description",
		},
		&requestflag.Flag[*bool]{
			Name:     "has-custom-routing",
			Usage:    "Whether the workspace preference defines custom routing for its topics.",
			BodyPath: "has_custom_routing",
		},
		&requestflag.Flag[any]{
			Name:     "routing-option",
			Usage:    "Default channels for the workspace preference. Defaults to empty if omitted.",
			BodyPath: "routing_options",
		},
	},
	Action:          handleWorkspacePreferencesCreate,
	HideHelpCommand: true,
}

var workspacePreferencesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a workspace preference by id, including its topics.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
	},
	Action:          handleWorkspacePreferencesRetrieve,
	HideHelpCommand: true,
}

var workspacePreferencesList = cli.Command{
	Name:            "list",
	Usage:           "List the workspace's preferences. Each workspace preference embeds its topics.\nScoped to the workspace of the API key.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWorkspacePreferencesList,
	HideHelpCommand: true,
}

var workspacePreferencesArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a workspace preference. The workspace preference must be empty: delete\nits topics first, otherwise the request fails with 409.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
	},
	Action:          handleWorkspacePreferencesArchive,
	HideHelpCommand: true,
}

var workspacePreferencesPublish = cli.Command{
	Name:    "publish",
	Usage:   "Publish the workspace's preferences page. Takes a snapshot of every workspace\npreference with its topics under a new published version, making the current\nstate visible on the hosted preferences page (non-draft).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "brand-id",
			Usage:    `Brand for the hosted page - "default" (workspace default brand), "none" (no brand), or a specific brand id. Defaults to "default".`,
			BodyPath: "brand_id",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "Description shown under the heading on the hosted preferences page.",
			BodyPath: "description",
		},
		&requestflag.Flag[*string]{
			Name:     "heading",
			Usage:    "Heading shown at the top of the hosted preferences page.",
			BodyPath: "heading",
		},
	},
	Action:          handleWorkspacePreferencesPublish,
	HideHelpCommand: true,
}

var workspacePreferencesReplace = cli.Command{
	Name:    "replace",
	Usage:   "Replace a workspace preference. Full document replacement; missing optional\nfields are cleared. Topics attached to the workspace preference are unaffected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the workspace preference.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "Optional description shown under the section on the hosted preferences page. Omit to clear.",
			BodyPath: "description",
		},
		&requestflag.Flag[*bool]{
			Name:     "has-custom-routing",
			Usage:    "Whether the workspace preference defines custom routing for its topics.",
			BodyPath: "has_custom_routing",
		},
		&requestflag.Flag[any]{
			Name:     "routing-option",
			Usage:    "Default channels for the workspace preference. Omit to clear.",
			BodyPath: "routing_options",
		},
	},
	Action:          handleWorkspacePreferencesReplace,
	HideHelpCommand: true,
}

func handleWorkspacePreferencesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.WorkspacePreferenceNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WorkspacePreferences.New(ctx, params, options...)
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
		Title:          "workspace-preferences create",
		Transform:      transform,
	})
}

func handleWorkspacePreferencesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("section-id") && len(unusedArgs) > 0 {
		cmd.Set("section-id", unusedArgs[0])
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
	_, err = client.WorkspacePreferences.Get(ctx, cmd.Value("section-id").(string), options...)
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
		Title:          "workspace-preferences retrieve",
		Transform:      transform,
	})
}

func handleWorkspacePreferencesList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.WorkspacePreferences.List(ctx, options...)
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
		Title:          "workspace-preferences list",
		Transform:      transform,
	})
}

func handleWorkspacePreferencesArchive(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("section-id") && len(unusedArgs) > 0 {
		cmd.Set("section-id", unusedArgs[0])
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

	return client.WorkspacePreferences.Archive(ctx, cmd.Value("section-id").(string), options...)
}

func handleWorkspacePreferencesPublish(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.WorkspacePreferencePublishParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WorkspacePreferences.Publish(ctx, params, options...)
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
		Title:          "workspace-preferences publish",
		Transform:      transform,
	})
}

func handleWorkspacePreferencesReplace(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("section-id") && len(unusedArgs) > 0 {
		cmd.Set("section-id", unusedArgs[0])
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

	params := courier.WorkspacePreferenceReplaceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WorkspacePreferences.Replace(
		ctx,
		cmd.Value("section-id").(string),
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
		Title:          "workspace-preferences replace",
		Transform:      transform,
	})
}
