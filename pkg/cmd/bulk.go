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

var bulkAddUsers = requestflag.WithInnerFlags(cli.Command{
	Name:    "add-users",
	Usage:   "Ingest user data into a Bulk Job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "job-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "user",
			Required: true,
			BodyPath: "users",
		},
	},
	Action:          handleBulkAddUsers,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"user": {
		&requestflag.InnerFlag[any]{
			Name:       "user.data",
			Usage:      "User-specific data that will be merged with message.data",
			InnerField: "data",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "user.preferences",
			InnerField: "preferences",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "user.profile",
			Usage:      "User profile information. For email-based bulk jobs, `profile.email` is required \nfor provider routing to determine if the message can be delivered. The email \naddress should be provided here rather than in `to.email`.\n",
			InnerField: "profile",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "user.recipient",
			Usage:      "User ID (legacy field, use profile or to.user_id instead)",
			InnerField: "recipient",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "user.to",
			InnerField: "to",
		},
	},
})

var bulkCreateJob = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-job",
	Usage:   "Creates a new bulk job for sending messages to multiple recipients.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "message",
			Usage:    "Bulk message definition. Supports two formats:\n- V1 format: Requires `event` field (event ID or notification ID)\n- V2 format: Optionally use `template` (notification ID) or `content` (Elemental content) in addition to `event`\n",
			Required: true,
			BodyPath: "message",
		},
	},
	Action:          handleBulkCreateJob,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"message": {
		&requestflag.InnerFlag[string]{
			Name:       "message.event",
			Usage:      "Event ID or Notification ID (required). Can be either a \nNotification ID (e.g., \"FRH3QXM9E34W4RKP7MRC8NZ1T8V8\") or a custom Event ID \n(e.g., \"welcome-email\") mapped to a notification.\n",
			InnerField: "event",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "message.brand",
			InnerField: "brand",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.content",
			Usage:      "Elemental content (optional, for V2 format). When provided, this will be used \ninstead of the notification associated with the `event` field.\n",
			InnerField: "content",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.data",
			InnerField: "data",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.locale",
			InnerField: "locale",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.override",
			InnerField: "override",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "message.template",
			Usage:      "Notification ID or template ID (optional, for V2 format). When provided, \nthis will be used instead of the notification associated with the `event` field.\n",
			InnerField: "template",
		},
	},
})

var bulkListUsers = cli.Command{
	Name:    "list-users",
	Usage:   "Get Bulk Job Users",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "job-id",
			Required: true,
		},
		&requestflag.Flag[*string]{
			Name:      "cursor",
			Usage:     "A unique identifier that allows for fetching the next set of users added to the bulk job",
			QueryPath: "cursor",
		},
	},
	Action:          handleBulkListUsers,
	HideHelpCommand: true,
}

var bulkRetrieveJob = cli.Command{
	Name:    "retrieve-job",
	Usage:   "Get a bulk job",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "job-id",
			Required: true,
		},
	},
	Action:          handleBulkRetrieveJob,
	HideHelpCommand: true,
}

var bulkRunJob = cli.Command{
	Name:    "run-job",
	Usage:   "Run a bulk job",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "job-id",
			Required: true,
		},
	},
	Action:          handleBulkRunJob,
	HideHelpCommand: true,
}

func handleBulkAddUsers(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.BulkAddUsersParams{}

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

	return client.Bulk.AddUsers(
		ctx,
		cmd.Value("job-id").(string),
		params,
		options...,
	)
}

func handleBulkCreateJob(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.BulkNewJobParams{}

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
	_, err = client.Bulk.NewJob(ctx, params, options...)
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
		Title:          "bulk create-job",
		Transform:      transform,
	})
}

func handleBulkListUsers(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.BulkListUsersParams{}

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
	_, err = client.Bulk.ListUsers(
		ctx,
		cmd.Value("job-id").(string),
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
		Title:          "bulk list-users",
		Transform:      transform,
	})
}

func handleBulkRetrieveJob(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
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
	_, err = client.Bulk.GetJob(ctx, cmd.Value("job-id").(string), options...)
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
		Title:          "bulk retrieve-job",
		Transform:      transform,
	})
}

func handleBulkRunJob(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
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

	return client.Bulk.RunJob(ctx, cmd.Value("job-id").(string), options...)
}
