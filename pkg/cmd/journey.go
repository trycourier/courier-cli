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

var journeysCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a journey. Defaults to `DRAFT` state; pass `state: \"PUBLISHED\"` to\npublish on create. Send nodes are not allowed on `POST`. The standard flow is:\ncreate the journey shell here, add notification templates with\n`POST /journeys/{templateId}/templates`, then wire them into the journey with\n`PUT /journeys/{templateId}`. Call `POST /journeys/{templateId}/publish` to\npublish a draft after the fact.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "node",
			Required: true,
			BodyPath: "nodes",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			BodyPath: "enabled",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "Lifecycle state of a journey.",
			BodyPath: "state",
		},
	},
	Action:          handleJourneysCreate,
	HideHelpCommand: true,
}

var journeysRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a journey by id. Pass `?version=draft` (default `published`) to retrieve\nthe working draft, or `?version=vN` to retrieve a historical version.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:      "version",
			Usage:     "Version selector: `draft`, `published` (default), or `vN`.",
			QueryPath: "version",
		},
	},
	Action:          handleJourneysRetrieve,
	HideHelpCommand: true,
}

var journeysList = cli.Command{
	Name:    "list",
	Usage:   "Get the list of journeys.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "A cursor token for pagination. Use the cursor from the previous response to fetch the next page of results.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "version",
			Usage:     "The version of journeys to retrieve. Accepted values are published (for published journeys) or draft (for draft journeys). Defaults to published.",
			Default:   "published",
			QueryPath: "version",
		},
	},
	Action:          handleJourneysList,
	HideHelpCommand: true,
}

var journeysArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive a journey. Archived journeys cannot be invoked. Existing journey runs\ncontinue to completion.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
	},
	Action:          handleJourneysArchive,
	HideHelpCommand: true,
}

var journeysCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel journey runs. The request body must contain EXACTLY ONE of\n`cancelation_token` (cancels every run associated with the token) or `run_id`\n(cancels a single tenant-scoped run). Supplying both or neither is a `400`. A\n`run_id` that does not exist for the caller's tenant returns `404`. Cancelation\nis idempotent and non-clobbering: a run that has already finished\n(`PROCESSED`/`ERROR`) or was already `CANCELED` is left untouched and its\ncurrent status is echoed back.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cancelation-token",
			BodyPath: "cancelation_token",
		},
		&requestflag.Flag[string]{
			Name:     "run-id",
			BodyPath: "run_id",
		},
	},
	Action:          handleJourneysCancel,
	HideHelpCommand: true,
}

var journeysInvoke = cli.Command{
	Name:    "invoke",
	Usage:   "Invoke a journey by id or alias to start a new run. The response includes a\n`runId` identifying the run.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			Usage:    "Data payload passed to the journey. The expected shape can be predefined using the schema builder in the journey editor. This data is available in journey steps for condition evaluation and template variable interpolation. Can also contain user identifiers (user_id, userId, anonymousId) if not provided elsewhere.",
			BodyPath: "data",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "profile",
			Usage:    "Profile data for the user. Can contain contact information (email, phone_number), user identifiers (user_id, userId, anonymousId), or any custom profile fields. Profile fields are merged with any existing stored profile for the user. Include context.tenant_id to load a tenant-scoped profile for multi-tenant scenarios.",
			BodyPath: "profile",
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    "A unique identifier for the user. If not provided, the system will attempt to resolve the user identifier from profile or data objects.",
			BodyPath: "user_id",
		},
	},
	Action:          handleJourneysInvoke,
	HideHelpCommand: true,
}

var journeysListVersions = cli.Command{
	Name:    "list-versions",
	Usage:   "List published versions of a journey, ordered most recent first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
	},
	Action:          handleJourneysListVersions,
	HideHelpCommand: true,
}

var journeysPublish = cli.Command{
	Name:    "publish",
	Usage:   "Publish the current draft as a new version. Body is optional; pass\n`{ \"version\": \"vN\" }` to roll back to a prior version instead. Returns 404 if\nthe journey has no draft to publish.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:     "version",
			BodyPath: "version",
		},
	},
	Action:          handleJourneysPublish,
	HideHelpCommand: true,
}

var journeysReplace = cli.Command{
	Name:    "replace",
	Usage:   "Replace the journey draft. Updates the working draft only; call\n`POST /journeys/{templateId}/publish` to make it live, or pass\n`state: \"PUBLISHED\"` in this request to publish immediately. Send-node\n`template` ids must already exist and be scoped to this journey, and node ids\nmust not be claimed by another journey.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-id",
			Required:  true,
			PathParam: "templateId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "node",
			Required: true,
			BodyPath: "nodes",
		},
		&requestflag.Flag[bool]{
			Name:     "enabled",
			BodyPath: "enabled",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "Lifecycle state of a journey.",
			BodyPath: "state",
		},
	},
	Action:          handleJourneysReplace,
	HideHelpCommand: true,
}

func handleJourneysCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.New(ctx, params, options...)
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
		Title:          "journeys create",
		Transform:      transform,
	})
}

func handleJourneysRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Get(
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
		Title:          "journeys retrieve",
		Transform:      transform,
	})
}

func handleJourneysList(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.List(ctx, params, options...)
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
		Title:          "journeys list",
		Transform:      transform,
	})
}

func handleJourneysArchive(ctx context.Context, cmd *cli.Command) error {
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

	return client.Journeys.Archive(ctx, cmd.Value("template-id").(string), options...)
}

func handleJourneysCancel(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyCancelParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Cancel(ctx, params, options...)
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
		Title:          "journeys cancel",
		Transform:      transform,
	})
}

func handleJourneysInvoke(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyInvokeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Invoke(
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
		Title:          "journeys invoke",
		Transform:      transform,
	})
}

func handleJourneysListVersions(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.ListVersions(ctx, cmd.Value("template-id").(string), options...)
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
		Title:          "journeys list-versions",
		Transform:      transform,
	})
}

func handleJourneysPublish(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyPublishParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Publish(
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
		Title:          "journeys publish",
		Transform:      transform,
	})
}

func handleJourneysReplace(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyReplaceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Replace(
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
		Title:          "journeys replace",
		Transform:      transform,
	})
}
