// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
	"github.com/trycourier/courier-cli/v3/internal/apiquery"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
	"github.com/trycourier/courier-go/v4"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/urfave/cli/v3"
)

var usersPreferencesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch all user preferences.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "tenant-id",
			Usage:     "Query the preferences of a user for this specific tenant context.",
			QueryPath: "tenant_id",
		},
	},
	Action:          handleUsersPreferencesRetrieve,
	HideHelpCommand: true,
}

var usersPreferencesRetrieveTopic = cli.Command{
	Name:    "retrieve-topic",
	Usage:   "Fetch user preferences for a specific subscription topic.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "topic-id",
			Required: true,
		},
		&requestflag.Flag[any]{
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
	Usage:   "Update or Create user preferences for a specific subscription topic.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "topic-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "topic",
			Required: true,
			BodyPath: "topic",
		},
		&requestflag.Flag[any]{
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
			InnerField: "status",
		},
		&requestflag.InnerFlag[any]{
			Name:       "topic.custom-routing",
			Usage:      "The Channels a user has chosen to receive notifications through for this topic",
			InnerField: "custom_routing",
		},
		&requestflag.InnerFlag[any]{
			Name:       "topic.has-custom-routing",
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

	params := courier.UserPreferenceGetParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:preferences retrieve", obj, format, transform)
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

	params := courier.UserPreferenceGetTopicParams{
		UserID: cmd.Value("user-id").(string),
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:preferences retrieve-topic", obj, format, transform)
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

	params := courier.UserPreferenceUpdateOrNewTopicParams{
		UserID: cmd.Value("user-id").(string),
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:preferences update-or-create-topic", obj, format, transform)
}
