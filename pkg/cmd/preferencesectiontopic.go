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

var preferenceSectionsTopicsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a subscription preference topic inside a section. Fails with 404 if the\nsection does not exist. The topic id is generated and returned.",
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
	Action:          handlePreferenceSectionsTopicsCreate,
	HideHelpCommand: true,
}

var preferenceSectionsTopicsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a topic within a section. Returns 404 if the section does not exist,\nthe topic does not exist, or the topic belongs to a different section.",
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
	Action:          handlePreferenceSectionsTopicsRetrieve,
	HideHelpCommand: true,
}

var preferenceSectionsTopicsList = cli.Command{
	Name:    "list",
	Usage:   "List the topics in a preference section.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "section-id",
			Required:  true,
			PathParam: "section_id",
		},
	},
	Action:          handlePreferenceSectionsTopicsList,
	HideHelpCommand: true,
}

var preferenceSectionsTopicsArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a topic and remove it from its section. Same 404 rules as GET.",
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
	Action:          handlePreferenceSectionsTopicsArchive,
	HideHelpCommand: true,
}

var preferenceSectionsTopicsReplace = cli.Command{
	Name:    "replace",
	Usage:   "Replace a topic within a section. Full document replacement; missing optional\nfields are cleared. Same 404 rules as GET.",
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
	Action:          handlePreferenceSectionsTopicsReplace,
	HideHelpCommand: true,
}

func handlePreferenceSectionsTopicsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.PreferenceSectionTopicNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PreferenceSections.Topics.New(
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
		Title:          "preference-sections:topics create",
		Transform:      transform,
	})
}

func handlePreferenceSectionsTopicsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.PreferenceSectionTopicGetParams{
		SectionID: cmd.Value("section-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PreferenceSections.Topics.Get(
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
		Title:          "preference-sections:topics retrieve",
		Transform:      transform,
	})
}

func handlePreferenceSectionsTopicsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.PreferenceSections.Topics.List(ctx, cmd.Value("section-id").(string), options...)
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
		Title:          "preference-sections:topics list",
		Transform:      transform,
	})
}

func handlePreferenceSectionsTopicsArchive(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.PreferenceSectionTopicArchiveParams{
		SectionID: cmd.Value("section-id").(string),
	}

	return client.PreferenceSections.Topics.Archive(
		ctx,
		cmd.Value("topic-id").(string),
		params,
		options...,
	)
}

func handlePreferenceSectionsTopicsReplace(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.PreferenceSectionTopicReplaceParams{
		SectionID: cmd.Value("section-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PreferenceSections.Topics.Replace(
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
		Title:          "preference-sections:topics replace",
		Transform:      transform,
	})
}
