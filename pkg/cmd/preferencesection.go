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

var preferenceSectionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a preference section in your workspace. The section id is generated and\nreturned. Topics are created inside a section via POST\n/preferences/sections/{section_id}/topics.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the section.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*bool]{
			Name:     "has-custom-routing",
			Usage:    "Whether the section defines custom routing for its topics.",
			BodyPath: "has_custom_routing",
		},
		&requestflag.Flag[any]{
			Name:     "routing-option",
			Usage:    "Default channels for the section. Defaults to empty if omitted.",
			BodyPath: "routing_options",
		},
	},
	Action:          handlePreferenceSectionsCreate,
	HideHelpCommand: true,
}

var preferenceSectionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a preference section by id, including its topics.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
	},
	Action:          handlePreferenceSectionsRetrieve,
	HideHelpCommand: true,
}

var preferenceSectionsList = cli.Command{
	Name:            "list",
	Usage:           "List the workspace's preference sections. Each section embeds its topics. Scoped\nto the workspace of the API key.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePreferenceSectionsList,
	HideHelpCommand: true,
}

var preferenceSectionsArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a preference section. The section must be empty: delete its topics\nfirst, otherwise the request fails with 409.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
	},
	Action:          handlePreferenceSectionsArchive,
	HideHelpCommand: true,
}

var preferenceSectionsPublish = cli.Command{
	Name:            "publish",
	Usage:           "Publish the workspace's preferences page. Takes a snapshot of every section with\nits topics under a new published version, making the current state visible on\nthe hosted preferences page (non-draft).",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePreferenceSectionsPublish,
	HideHelpCommand: true,
}

var preferenceSectionsReplace = cli.Command{
	Name:    "replace",
	Usage:   "Replace a preference section. Full document replacement; missing optional fields\nare cleared. Topics attached to the section are unaffected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the section.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*bool]{
			Name:     "has-custom-routing",
			Usage:    "Whether the section defines custom routing for its topics.",
			BodyPath: "has_custom_routing",
		},
		&requestflag.Flag[any]{
			Name:     "routing-option",
			Usage:    "Default channels for the section. Omit to clear.",
			BodyPath: "routing_options",
		},
	},
	Action:          handlePreferenceSectionsReplace,
	HideHelpCommand: true,
}

func handlePreferenceSectionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.PreferenceSectionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PreferenceSections.New(ctx, params, options...)
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
		Title:          "preference-sections create",
		Transform:      transform,
	})
}

func handlePreferenceSectionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PreferenceSections.Get(ctx, cmd.Value("section-id").(string), options...)
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
		Title:          "preference-sections retrieve",
		Transform:      transform,
	})
}

func handlePreferenceSectionsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PreferenceSections.List(ctx, options...)
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
		Title:          "preference-sections list",
		Transform:      transform,
	})
}

func handlePreferenceSectionsArchive(ctx context.Context, cmd *cli.Command) error {
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

	return client.PreferenceSections.Archive(ctx, cmd.Value("section-id").(string), options...)
}

func handlePreferenceSectionsPublish(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PreferenceSections.Publish(ctx, options...)
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
		Title:          "preference-sections publish",
		Transform:      transform,
	})
}

func handlePreferenceSectionsReplace(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.PreferenceSectionReplaceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PreferenceSections.Replace(
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
		Title:          "preference-sections replace",
		Transform:      transform,
	})
}
