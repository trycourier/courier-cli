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

var workspacePreferencesTopicsCreate = requestflag.WithInnerFlags(cli.Command{
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
		&requestflag.Flag[map[string]any]{
			Name:     "digest",
			Usage:    "A topic's digest configuration: the template that renders it, the cadences it delivers on, and how collected events are retained.\n\nSend `null` for the whole object to turn a digest off, which unlinks the template and removes its schedules. There is no `enabled` flag, and `schedules: []` is rejected, because both states are un-deliverable rather than merely off.",
			BodyPath: "digest",
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
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			HeaderPath: "Idempotency-Key",
		},
		&requestflag.Flag[string]{
			Name:       "x-idempotency-expiration",
			HeaderPath: "x-idempotency-expiration",
		},
	},
	Action:          handleWorkspacePreferencesTopicsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"digest": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "digest.schedules",
			Usage:      "The cadences this digest delivers on. At least one is required: a digest with no schedule collects events into an instance that can never fire. Omitting the key on a replace leaves stored schedules untouched; sending `[]` is a `400`.",
			InnerField: "schedules",
		},
		&requestflag.InnerFlag[string]{
			Name:       "digest.template-id",
			Usage:      "The notification template that renders the digest. A digest with no template collects nothing, so this is required.",
			InnerField: "template_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "digest.audience-id",
			Usage:      "Optional audience the digest is scoped to.",
			InnerField: "audience_id",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "digest.categories",
			Usage:      "Retention rules per category key. Defaults to a single `digest` category retaining `FIRST`.",
			InnerField: "categories",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "digest.trigger-empty",
			Usage:      "Whether to deliver the digest even when nothing was collected.",
			InnerField: "trigger_empty",
		},
	},
})

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

var workspacePreferencesTopicsDeleteDigest = cli.Command{
	Name:    "delete-digest",
	Usage:   "Turn off a topic's digest, leaving the topic itself in place. The template is\nunlinked and the digest's schedules are removed along with their delivery rules.\nEquivalent to sending `digest: null` on a topic replace.",
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
	Action:          handleWorkspacePreferencesTopicsDeleteDigest,
	HideHelpCommand: true,
}

var workspacePreferencesTopicsReleaseDigest = cli.Command{
	Name:    "release-digest",
	Usage:   "Send one recipient's held digest now, instead of waiting for its schedule. Use\nit to preview what a digest will look like, or to let someone flush their own.",
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
			Name:     "user-id",
			Usage:    `The recipient whose digest to release. Required: there is no "release everyone on this topic" form, because a whole-schedule flush already has its own endpoint and a body-shaped difference between one recipient and all of them is too easy to get wrong.`,
			Required: true,
			BodyPath: "user_id",
		},
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Usage:    "The recipient's tenant, when they were sent to as part of one -- the same value returned as `tenant_id` on a digest instance and sent as `message.context.tenant_id`. It is part of the held digest's key, so a tenanted recipient cannot be found without it. Omit for an ordinary recipient.",
			BodyPath: "tenant_id",
		},
	},
	Action:          handleWorkspacePreferencesTopicsReleaseDigest,
	HideHelpCommand: true,
}

var workspacePreferencesTopicsReplace = requestflag.WithInnerFlags(cli.Command{
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
		&requestflag.Flag[map[string]any]{
			Name:     "digest",
			Usage:    "A topic's digest configuration: the template that renders it, the cadences it delivers on, and how collected events are retained.\n\nSend `null` for the whole object to turn a digest off, which unlinks the template and removes its schedules. There is no `enabled` flag, and `schedules: []` is rejected, because both states are un-deliverable rather than merely off.",
			BodyPath: "digest",
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
}, map[string][]requestflag.HasOuterFlag{
	"digest": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "digest.schedules",
			Usage:      "The cadences this digest delivers on. At least one is required: a digest with no schedule collects events into an instance that can never fire. Omitting the key on a replace leaves stored schedules untouched; sending `[]` is a `400`.",
			InnerField: "schedules",
		},
		&requestflag.InnerFlag[string]{
			Name:       "digest.template-id",
			Usage:      "The notification template that renders the digest. A digest with no template collects nothing, so this is required.",
			InnerField: "template_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "digest.audience-id",
			Usage:      "Optional audience the digest is scoped to.",
			InnerField: "audience_id",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "digest.categories",
			Usage:      "Retention rules per category key. Defaults to a single `digest` category retaining `FIRST`.",
			InnerField: "categories",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "digest.trigger-empty",
			Usage:      "Whether to deliver the digest even when nothing was collected.",
			InnerField: "trigger_empty",
		},
	},
})

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

func handleWorkspacePreferencesTopicsDeleteDigest(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.WorkspacePreferenceTopicDeleteDigestParams{
		SectionID: cmd.Value("section-id").(string),
	}

	return client.WorkspacePreferences.Topics.DeleteDigest(
		ctx,
		cmd.Value("topic-id").(string),
		params,
		options...,
	)
}

func handleWorkspacePreferencesTopicsReleaseDigest(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.WorkspacePreferenceTopicReleaseDigestParams{
		SectionID: cmd.Value("section-id").(string),
	}

	return client.WorkspacePreferences.Topics.ReleaseDigest(
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
