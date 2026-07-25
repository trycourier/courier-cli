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

var workspacePreferencesTopicsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a subscription topic inside a workspace preference. The default status\nsets whether users start opted in, opted out, or required.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
		&requestflag.Flag[string]{
			Name:     "default-status",
			Usage:    "The default subscription status applied when a recipient has not set their own.",
			Required: true,
			BodyPath: "default_status",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the preference topic.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "allowed-preference",
			Usage:    "Preference controls a recipient may customize for this topic. Defaults to empty if omitted.",
			BodyPath: "allowed_preferences",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "Optional description shown under the topic on the hosted preferences page.",
			BodyPath: "description",
		},
		&requestflag.Flag[*bool]{
			Name:     "include-unsubscribe-header",
			Usage:    "Whether to include a list-unsubscribe header on emails for this topic.",
			BodyPath: "include_unsubscribe_header",
		},
		&requestflag.Flag[any]{
			Name:     "routing-option",
			Usage:    "Default channels delivered for this topic. Defaults to empty if omitted.",
			BodyPath: "routing_options",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "topic-data",
			Usage:    "Arbitrary metadata associated with the topic.",
			BodyPath: "topic_data",
		},
	},
	Action:          handleWorkspacePreferencesTopicsCreate,
	HideHelpCommand: true,
}

var workspacePreferencesTopicsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns one subscription topic with its default status, routing options, allowed\npreferences, and unsubscribe header setting.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
		&requestflag.Flag[string]{
			Name:      "topic-id",
			Required:  true,
			PathParam: "topic_id",
		},
	},
	Action:          handleWorkspacePreferencesTopicsRetrieve,
	HideHelpCommand: true,
}

var workspacePreferencesTopicsList = cli.Command{
	Name:    "list",
	Usage:   "Returns the subscription topics inside a workspace preference, each with its\ndefault status and routing options.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
	},
	Action:          handleWorkspacePreferencesTopicsList,
	HideHelpCommand: true,
}

var workspacePreferencesTopicsArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archives a subscription topic and removes it from its workspace preference,\naddressed by section id and topic id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
		&requestflag.Flag[string]{
			Name:      "topic-id",
			Required:  true,
			PathParam: "topic_id",
		},
	},
	Action:          handleWorkspacePreferencesTopicsArchive,
	HideHelpCommand: true,
}

var workspacePreferencesTopicsReplace = cli.Command{
	Name:    "replace",
	Usage:   "Replace a topic within a workspace preference. Full document replacement;\nmissing optional fields are cleared. Same 404 rules as GET.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
		&requestflag.Flag[string]{
			Name:      "topic-id",
			Required:  true,
			PathParam: "topic_id",
		},
		&requestflag.Flag[string]{
			Name:     "default-status",
			Usage:    "The default subscription status applied when a recipient has not set their own.",
			Required: true,
			BodyPath: "default_status",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the preference topic.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "allowed-preference",
			Usage:    "Preference controls a recipient may customize. Omit to clear.",
			BodyPath: "allowed_preferences",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "Optional description shown under the topic on the hosted preferences page. Omit to clear.",
			BodyPath: "description",
		},
		&requestflag.Flag[*bool]{
			Name:     "include-unsubscribe-header",
			Usage:    "Whether to include a list-unsubscribe header on emails for this topic.",
			BodyPath: "include_unsubscribe_header",
		},
		&requestflag.Flag[any]{
			Name:     "routing-option",
			Usage:    "Default channels delivered for this topic. Omit to clear.",
			BodyPath: "routing_options",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "topic-data",
			Usage:    "Arbitrary metadata associated with the topic. Omit to clear.",
			BodyPath: "topic_data",
		},
	},
	Action:          handleWorkspacePreferencesTopicsReplace,
	HideHelpCommand: true,
}

func handleWorkspacePreferencesTopicsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.WorkspacePreferenceTopicNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WorkspacePreferences.Topics.New(
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
		Title:          "workspace-preferences:topics create",
		Transform:      transform,
	})
}

func handleWorkspacePreferencesTopicsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("topic-id") && len(unusedArgs) > 0 {
		cmd.Set("topic-id", unusedArgs[0])
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

	params := courier.WorkspacePreferenceTopicGetParams{
		SectionID: cmd.Value("section-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WorkspacePreferences.Topics.Get(
		ctx,
		cmd.Value("topic-id").(string),
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
		Title:          "workspace-preferences:topics retrieve",
		Transform:      transform,
	})
}

func handleWorkspacePreferencesTopicsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.WorkspacePreferences.Topics.List(ctx, cmd.Value("section-id").(string), options...)
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
		Title:          "workspace-preferences:topics list",
		Transform:      transform,
	})
}

func handleWorkspacePreferencesTopicsArchive(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("topic-id") && len(unusedArgs) > 0 {
		cmd.Set("topic-id", unusedArgs[0])
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

	params := courier.WorkspacePreferenceTopicArchiveParams{
		SectionID: cmd.Value("section-id").(string),
	}

	return client.WorkspacePreferences.Topics.Archive(
		ctx,
		cmd.Value("topic-id").(string),
		params,
		options...,
	)
}

func handleWorkspacePreferencesTopicsReplace(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("topic-id") && len(unusedArgs) > 0 {
		cmd.Set("topic-id", unusedArgs[0])
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

	params := courier.WorkspacePreferenceTopicReplaceParams{
		SectionID: cmd.Value("section-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WorkspacePreferences.Topics.Replace(
		ctx,
		cmd.Value("topic-id").(string),
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
		Title:          "workspace-preferences:topics replace",
		Transform:      transform,
	})
}
