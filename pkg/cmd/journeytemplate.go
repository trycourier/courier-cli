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
	Usage:   "Create a notification template scoped to this journey. Defaults to `DRAFT`\nstate; pass `state: \"PUBLISHED\"` to publish on create.",
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
	Usage:   "List notification templates scoped to this journey. Journey-scoped notification\ntemplates can only be referenced from `send` nodes within the same journey.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor from a prior response.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Page size. Minimum 1, maximum 100.",
			QueryPath: "limit",
		},
	},
	Action:          handleJourneysTemplatesList,
	HideHelpCommand: true,
}

var journeysTemplatesArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive the journey-scoped notification template. Archived templates cannot be\nsent.",
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
	Usage:   "List published versions of the journey-scoped notification template, ordered\nmost recent first.",
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
	Usage:   "Publish the current draft of the journey-scoped notification template as a new\nversion. Optionally roll back to a prior version by passing\n`{ \"version\": \"vN\" }`.",
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

var journeysTemplatesPutContent = requestflag.WithInnerFlags(cli.Command{
	Name:    "put-content",
	Usage:   "Replace the elemental content of a journey-scoped notification template.\nOverwrites all elements in the template draft with the provided content.",
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
			Name:     "content",
			Usage:    "Elemental content payload. The server defaults `version` when omitted.",
			Required: true,
			BodyPath: "content",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "Template state. Defaults to `DRAFT`.",
			Default:  "DRAFT",
			BodyPath: "state",
		},
	},
	Action:          handleJourneysTemplatesPutContent,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"content": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "content.elements",
			InnerField: "elements",
		},
		&requestflag.InnerFlag[string]{
			Name:       "content.version",
			Usage:      "Content version identifier (e.g., `2022-01-01`). Optional; server defaults when omitted.",
			InnerField: "version",
		},
	},
})

var journeysTemplatesPutLocale = requestflag.WithInnerFlags(cli.Command{
	Name:    "put-locale",
	Usage:   "Set locale-specific content overrides for a journey-scoped notification\ntemplate. Each element override must reference an existing element by ID.",
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
			Name:      "locale-id",
			Required:  true,
			PathParam: "localeId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "element",
			Usage:    "Elements with locale-specific content overrides.",
			Required: true,
			BodyPath: "elements",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "Template state. Defaults to `DRAFT`.",
			Default:  "DRAFT",
			BodyPath: "state",
		},
	},
	Action:          handleJourneysTemplatesPutLocale,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"element": {
		&requestflag.InnerFlag[string]{
			Name:       "element.id",
			Usage:      "Target element ID.",
			InnerField: "id",
		},
	},
})

var journeysTemplatesReplace = requestflag.WithInnerFlags(cli.Command{
	Name:    "replace",
	Usage:   "Replace the journey-scoped notification template draft.",
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

var journeysTemplatesRetrieveContent = cli.Command{
	Name:    "retrieve-content",
	Usage:   "Retrieve the elemental content of a journey-scoped notification template. The\nresponse contains the versioned elements along with their content checksums,\nwhich can be used to detect changes between versions. Pass `?version=draft`\n(default `published`) to retrieve the working draft, or `?version=vN` for a\nhistorical version.",
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
			Name:      "version",
			Usage:     "Accepts `draft`, `published`, or a version string (e.g., `v001`). Defaults to `published`.",
			QueryPath: "version",
		},
	},
	Action:          handleJourneysTemplatesRetrieveContent,
	HideHelpCommand: true,
}

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

func handleJourneysTemplatesPutContent(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyTemplatePutContentParams{
		TemplateID: cmd.Value("template-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Templates.PutContent(
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
		Title:          "journeys:templates put-content",
		Transform:      transform,
	})
}

func handleJourneysTemplatesPutLocale(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("locale-id") && len(unusedArgs) > 0 {
		cmd.Set("locale-id", unusedArgs[0])
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

	params := courier.JourneyTemplatePutLocaleParams{
		TemplateID:     cmd.Value("template-id").(string),
		NotificationID: cmd.Value("notification-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Templates.PutLocale(
		ctx,
		cmd.Value("locale-id").(string),
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
		Title:          "journeys:templates put-locale",
		Transform:      transform,
	})
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

func handleJourneysTemplatesRetrieveContent(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyTemplateGetContentParams{
		TemplateID: cmd.Value("template-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Templates.GetContent(
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
		Title:          "journeys:templates retrieve-content",
		Transform:      transform,
	})
}
