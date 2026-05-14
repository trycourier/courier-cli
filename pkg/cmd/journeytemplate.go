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

var journeysTemplatesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a notification template scoped to this journey. The template is created\nin DRAFT state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:     "channel",
			Required: true,
			BodyPath: "channel",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "notification",
			Required: true,
			BodyPath: "notification",
		},
		&requestflag.Flag[string]{
			Name:     "provider-key",
			BodyPath: "providerKey",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			BodyPath: "state",
		},
	},
	Action:          handleJourneysTemplatesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"notification": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "notification.brand",
			InnerField: "brand",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "notification.content",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "notification.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "notification.subscription",
			InnerField: "subscription",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "notification.tags",
			InnerField: "tags",
		},
	},
})

var journeysTemplatesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a journey-scoped notification template by id. Pass `?version=draft`\n(default `published`) to retrieve the working draft, or `?version=vN` for a\nhistorical version.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:      "notification-id",
			Required:  true,
			PathParam: "notificationId",
		},
	},
	Action:          handleJourneysTemplatesRetrieve,
	HideHelpCommand: true,
}

var journeysTemplatesList = cli.Command{
	Name:    "list",
	Usage:   "List notification templates scoped to this journey. Templates scoped to a\njourney can only be referenced from `send` nodes of the same journey.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			QueryPath: "limit",
		},
	},
	Action:          handleJourneysTemplatesList,
	HideHelpCommand: true,
}

var journeysTemplatesArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a journey-scoped notification template. Archived templates cannot be\nsent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:      "notification-id",
			Required:  true,
			PathParam: "notificationId",
		},
	},
	Action:          handleJourneysTemplatesArchive,
	HideHelpCommand: true,
}

var journeysTemplatesListVersions = cli.Command{
	Name:    "list-versions",
	Usage:   "List published versions of a journey-scoped notification template, ordered most\nrecent first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:      "notification-id",
			Required:  true,
			PathParam: "notificationId",
		},
	},
	Action:          handleJourneysTemplatesListVersions,
	HideHelpCommand: true,
}

var journeysTemplatesPublish = cli.Command{
	Name:    "publish",
	Usage:   "Publish the current draft of a journey-scoped notification template.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:      "notification-id",
			Required:  true,
			PathParam: "notificationId",
		},
		&requestflag.Flag[string]{
			Name:     "version",
			BodyPath: "version",
		},
	},
	Action:          handleJourneysTemplatesPublish,
	HideHelpCommand: true,
}

var journeysTemplatesReplace = requestflag.WithInnerFlags(cli.Command{
	Name:    "replace",
	Usage:   "Replace a journey-scoped notification template draft.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:      "notification-id",
			Required:  true,
			PathParam: "notificationId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "notification",
			Required: true,
			BodyPath: "notification",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			BodyPath: "state",
		},
	},
	Action:          handleJourneysTemplatesReplace,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"notification": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "notification.brand",
			InnerField: "brand",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "notification.content",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "notification.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "notification.subscription",
			InnerField: "subscription",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "notification.tags",
			InnerField: "tags",
		},
	},
})

func handleJourneysTemplatesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyTemplateNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Templates.New(
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
		Title:          "journeys:templates create",
		Transform:      transform,
	})
}

func handleJourneysTemplatesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("notification-id") && len(unusedArgs) > 0 {
		cmd.Set("notification-id", unusedArgs[0])
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

	params := courier.JourneyTemplateGetParams{
		TemplateID: cmd.Value("template-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Templates.Get(
		ctx,
		cmd.Value("notification-id").(string),
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
		Title:          "journeys:templates retrieve",
		Transform:      transform,
	})
}

func handleJourneysTemplatesList(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := courier.JourneyTemplateListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Templates.List(
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
		Title:          "journeys:templates list",
		Transform:      transform,
	})
}

func handleJourneysTemplatesArchive(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("notification-id") && len(unusedArgs) > 0 {
		cmd.Set("notification-id", unusedArgs[0])
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

	params := courier.JourneyTemplateArchiveParams{
		TemplateID: cmd.Value("template-id").(string),
	}

	return client.Journeys.Templates.Archive(
		ctx,
		cmd.Value("notification-id").(string),
		params,
		options...,
	)
}

func handleJourneysTemplatesListVersions(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("notification-id") && len(unusedArgs) > 0 {
		cmd.Set("notification-id", unusedArgs[0])
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

	params := courier.JourneyTemplateListVersionsParams{
		TemplateID: cmd.Value("template-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Templates.ListVersions(
		ctx,
		cmd.Value("notification-id").(string),
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
		Title:          "journeys:templates list-versions",
		Transform:      transform,
	})
}

func handleJourneysTemplatesPublish(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("notification-id") && len(unusedArgs) > 0 {
		cmd.Set("notification-id", unusedArgs[0])
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

	params := courier.JourneyTemplatePublishParams{
		TemplateID: cmd.Value("template-id").(string),
	}

	return client.Journeys.Templates.Publish(
		ctx,
		cmd.Value("notification-id").(string),
		params,
		options...,
	)
}

func handleJourneysTemplatesReplace(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("notification-id") && len(unusedArgs) > 0 {
		cmd.Set("notification-id", unusedArgs[0])
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

	params := courier.JourneyTemplateReplaceParams{
		TemplateID: cmd.Value("template-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Templates.Replace(
		ctx,
		cmd.Value("notification-id").(string),
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
		Title:          "journeys:templates replace",
		Transform:      transform,
	})
}
