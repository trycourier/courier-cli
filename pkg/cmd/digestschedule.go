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

var digestsSchedulesListInstances = cli.Command{
	Name:    "list-instances",
	Usage:   "List the digest instances for a schedule. Each instance represents the events\naccumulated for a single user against the schedule, and can be used to monitor\ndigest accumulation before the digest is released.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "schedule-id",
			Required:  true,
			PathParam: "schedule_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "A cursor token from a previous response, used to fetch the next page of results.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "The maximum number of digest instances to return. Defaults to 20, with a maximum of 100.",
			Default:   20,
			QueryPath: "limit",
		},
	},
	Action:          handleDigestsSchedulesListInstances,
	HideHelpCommand: true,
}

var digestsSchedulesRelease = cli.Command{
	Name:    "release",
	Usage:   "Send a digest now instead of waiting for its scheduled time, so your users get\nwhat they have collected so far right away.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "schedule-id",
			Required:  true,
			PathParam: "schedule_id",
		},
	},
	Action:          handleDigestsSchedulesRelease,
	HideHelpCommand: true,
}

func handleDigestsSchedulesListInstances(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("schedule-id") && len(unusedArgs) > 0 {
		cmd.Set("schedule-id", unusedArgs[0])
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

	params := courier.DigestScheduleListInstancesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Digests.Schedules.ListInstances(
		ctx,
		cmd.Value("schedule-id").(string),
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
		Title:          "digests:schedules list-instances",
		Transform:      transform,
	})
}

func handleDigestsSchedulesRelease(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("schedule-id") && len(unusedArgs) > 0 {
		cmd.Set("schedule-id", unusedArgs[0])
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

	return client.Digests.Schedules.Release(ctx, cmd.Value("schedule-id").(string), options...)
}
