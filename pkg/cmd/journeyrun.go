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

var journeysRunsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch one Journey run by id. Returns `404` for an unknown run, a run belonging\nto another workspace, a run past the 95-day retention window, or an Automation\nrun id — the same body in every case, so the response never reveals whether a\nrun exists elsewhere.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "run-id",
			Required:  true,
			PathParam: "run_id",
		},
	},
	Action:          handleJourneysRunsRetrieve,
	HideHelpCommand: true,
}

var journeysRunsList = cli.Command{
	Name:    "list",
	Usage:   "List runs of the workspace's Journeys, newest first, filtered by status,\nJourney, or date range and paged by cursor. Runs of v2 Automations are listed by\n`GET /automations/runs` instead — the two surfaces never return each other's\nruns. Runs are retained for 95 days.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "A cursor token for pagination. Use the `next_cursor` from the previous response to fetch the next page of results. Treat it as opaque.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "end-date",
			Usage:     "An inclusive upper bound on `created_at`, in the same format as `start_date`.",
			QueryPath: "end_date",
		},
		&requestflag.Flag[string]{
			Name:      "limit",
			Usage:     "The number of runs to return per page, between `1` and `50`. Defaults to `20`. Values outside the range are clamped, and a non-numeric value falls back to `20`.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "start-date",
			Usage:     "An inclusive lower bound on `created_at`, as an ISO 8601 date or timestamp (e.g. `2026-08-18` or `2026-08-18T20:06:36.259Z`). Any other format returns `400`.",
			QueryPath: "start_date",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "A comma-separated list of run statuses to filter on, e.g. `PROCESSED,ERROR`.",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "template-id",
			Usage:     "A comma-separated list of Journey ids to filter on.",
			QueryPath: "template_id",
		},
	},
	Action:          handleJourneysRunsList,
	HideHelpCommand: true,
}

var journeysRunsListSteps = cli.Command{
	Name:    "list-steps",
	Usage:   "List the per-node state of one Journey run, in full — this endpoint is not\npaginated. Each step's `node_id` is the id of the node in the published Journey,\nso a step maps directly onto the Journey graph. `message_id` is present on send\nsteps that produced a message; follow it to `GET /messages/{message_id}` for\ndelivery status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "run-id",
			Required:  true,
			PathParam: "run_id",
		},
	},
	Action:          handleJourneysRunsListSteps,
	HideHelpCommand: true,
}

func handleJourneysRunsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
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
	_, err = client.Journeys.Runs.Get(ctx, cmd.Value("run-id").(string), options...)
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
		Title:          "journeys:runs retrieve",
		Transform:      transform,
	})
}

func handleJourneysRunsList(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.JourneyRunListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Journeys.Runs.List(ctx, params, options...)
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
		Title:          "journeys:runs list",
		Transform:      transform,
	})
}

func handleJourneysRunsListSteps(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
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
	_, err = client.Journeys.Runs.ListSteps(ctx, cmd.Value("run-id").(string), options...)
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
		Title:          "journeys:runs list-steps",
		Transform:      transform,
	})
}
