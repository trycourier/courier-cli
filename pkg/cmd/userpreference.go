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

var usersPreferencesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns a user's preference overrides with paging, one entry per subscription\ntopic they have set a choice for.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
		},
		&requestflag.Flag[*string]{
			Name:      "tenant-id",
			Usage:     "Query the preferences of a user for this specific tenant context.",
			QueryPath: "tenant_id",
		},
	},
	Action:          handleUsersPreferencesRetrieve,
	HideHelpCommand: true,
}

var usersPreferencesBulkReplace = requestflag.WithInnerFlags(cli.Command{
	Name:    "bulk-replace",
	Usage:   "Replaces a user's entire set of preference overrides. Any topic you leave out is\nreset to its default, so send the full set rather than a subset.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "topic",
			Usage:    "The complete set of topic overrides for the user. Up to 50 topics may be provided. Any existing override not listed here is reset to its topic default; an empty array resets every existing override.",
			Required: true,
			BodyPath: "topics",
		},
		&requestflag.Flag[*string]{
			Name:      "tenant-id",
			Usage:     "Replace the preferences of a user for this specific tenant context.",
			QueryPath: "tenant_id",
		},
	},
	Action:          handleUsersPreferencesBulkReplace,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"topic": {
		&requestflag.InnerFlag[string]{
			Name:       "topic.status",
			Usage:      "The subscription status to apply for this topic.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "topic.topic-id",
			Usage:      "A unique identifier associated with a subscription topic.",
			InnerField: "topic_id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "topic.custom-routing",
			Usage:      "The channels a user has chosen to receive notifications through for this topic.",
			InnerField: "custom_routing",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "topic.has-custom-routing",
			Usage:      "Whether the recipient has chosen specific delivery channels for this topic.",
			InnerField: "has_custom_routing",
		},
	},
})

var usersPreferencesBulkUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "bulk-update",
	Usage:   "Adds or updates a user's preferences for several subscription topics at once.\nTopics you leave out keep whatever they were set to before.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "topic",
			Usage:    "The topics to create or update. Between 1 and 50 topics may be provided in a single request.",
			Required: true,
			BodyPath: "topics",
		},
		&requestflag.Flag[*string]{
			Name:      "tenant-id",
			Usage:     "Update the preferences of a user for this specific tenant context.",
			QueryPath: "tenant_id",
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
	Action:          handleUsersPreferencesBulkUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"topic": {
		&requestflag.InnerFlag[string]{
			Name:       "topic.status",
			Usage:      "The subscription status to apply for this topic.",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "topic.topic-id",
			Usage:      "A unique identifier associated with a subscription topic.",
			InnerField: "topic_id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "topic.custom-routing",
			Usage:      "The channels a user has chosen to receive notifications through for this topic.",
			InnerField: "custom_routing",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "topic.has-custom-routing",
			Usage:      "Whether the recipient has chosen specific delivery channels for this topic.",
			InnerField: "has_custom_routing",
		},
	},
})

var usersPreferencesDeleteTopic = cli.Command{
	Name:    "delete-topic",
	Usage:   "Removes a user's override for one subscription topic, resetting it to the\neffective default from the tenant or workspace.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
		},
		&requestflag.Flag[string]{
			Name:      "topic-id",
			Required:  true,
			PathParam: "topic_id",
		},
		&requestflag.Flag[*string]{
			Name:      "tenant-id",
			Usage:     "Delete the preferences of a user for this specific tenant context.",
			QueryPath: "tenant_id",
		},
	},
	Action:          handleUsersPreferencesDeleteTopic,
	HideHelpCommand: true,
}

var usersPreferencesRetrieveTopic = cli.Command{
	Name:    "retrieve-topic",
	Usage:   "Returns a user's opt-in status and channel choices for one subscription topic,\nor the effective default if they have set no override.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
		},
		&requestflag.Flag[string]{
			Name:      "topic-id",
			Required:  true,
			PathParam: "topic_id",
		},
		&requestflag.Flag[*string]{
			Name:      "tenant-id",
			Usage:     "Query the preferences of a user for this specific tenant context.",
			QueryPath: "tenant_id",
		},
	},
	Action:          handleUsersPreferencesRetrieveTopic,
	HideHelpCommand: true,
}

var usersPreferencesUpdateOrCreateTopic = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-or-create-topic",
	Usage:   "Sets a user's opt-in status and channel choices for one subscription topic,\noverriding the tenant default for that topic only.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
		},
		&requestflag.Flag[string]{
			Name:      "topic-id",
			Required:  true,
			PathParam: "topic_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "topic",
			Required: true,
			BodyPath: "topic",
		},
		&requestflag.Flag[*string]{
			Name:      "tenant-id",
			Usage:     "Update the preferences of a user for this specific tenant context.",
			QueryPath: "tenant_id",
		},
	},
	Action:          handleUsersPreferencesUpdateOrCreateTopic,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"topic": {
		&requestflag.InnerFlag[string]{
			Name:       "topic.status",
			Usage:      `Allowed values: "OPTED_IN", "OPTED_OUT", "REQUIRED".`,
			InnerField: "status",
		},
		&requestflag.InnerFlag[any]{
			Name:       "topic.custom-routing",
			Usage:      "The channels to deliver this topic on when has_custom_routing is true. One or more of: direct_message, email, push, sms, webhook, inbox.",
			InnerField: "custom_routing",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "topic.has-custom-routing",
			Usage:      "Set to true to route this topic to the channels in custom_routing instead of the topic's default routing.",
			InnerField: "has_custom_routing",
		},
	},
})

func handleUsersPreferencesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
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

	params := courier.UserPreferenceGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Preferences.Get(
		ctx,
		cmd.Value("user-id").(string),
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
		Title:          "users:preferences retrieve",
		Transform:      transform,
	})
}

func handleUsersPreferencesBulkReplace(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
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

	params := courier.UserPreferenceBulkReplaceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Preferences.BulkReplace(
		ctx,
		cmd.Value("user-id").(string),
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
		Title:          "users:preferences bulk-replace",
		Transform:      transform,
	})
}

func handleUsersPreferencesBulkUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
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

	params := courier.UserPreferenceBulkUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Preferences.BulkUpdate(
		ctx,
		cmd.Value("user-id").(string),
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
		Title:          "users:preferences bulk-update",
		Transform:      transform,
	})
}

func handleUsersPreferencesDeleteTopic(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.UserPreferenceDeleteTopicParams{
		UserID: cmd.Value("user-id").(string),
	}

	return client.Users.Preferences.DeleteTopic(
		ctx,
		cmd.Value("topic-id").(string),
		params,
		options...,
	)
}

func handleUsersPreferencesRetrieveTopic(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.UserPreferenceGetTopicParams{
		UserID: cmd.Value("user-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Preferences.GetTopic(
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
		Title:          "users:preferences retrieve-topic",
		Transform:      transform,
	})
}

func handleUsersPreferencesUpdateOrCreateTopic(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.UserPreferenceUpdateOrNewTopicParams{
		UserID: cmd.Value("user-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Preferences.UpdateOrNewTopic(
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
		Title:          "users:preferences update-or-create-topic",
		Transform:      transform,
	})
}
