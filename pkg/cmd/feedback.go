// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/trycourier/courier-cli/v6/internal/apiquery"
	"github.com/trycourier/courier-cli/v6/internal/requestflag"
	"github.com/trycourier/courier-go/v4"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/urfave/cli/v3"
)

var feedbackCreate = cli.Command{
	Name:    "create",
	Usage:   "Reports a suspected bug, confusing behavior or a missing feature in Courier. It\nis built for AI agents and developer tools to call when Courier does something\nunexpected, and people can call it too.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "category",
			Usage:    "What kind of report this is: `possible_bug` for behavior that looks wrong, `confusing_behavior` for behavior that works but was hard to understand, `feature_request` for something Courier does not do yet, and `other`.",
			Required: true,
			BodyPath: "category",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "What happened, in enough detail for someone at Courier to follow it.",
			Required: true,
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "surface",
			Usage:    "Where you met the problem: the REST `api`, an `sdk`, the Courier `mcp` server, the Courier app `ui`, the `docs`, or `other`.",
			Required: true,
			BodyPath: "surface",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Usage:    "A one-line summary of the report.",
			Required: true,
			BodyPath: "title",
		},
		&requestflag.Flag[string]{
			Name:     "expected-behavior",
			Usage:    "What you expected to happen instead.",
			BodyPath: "expected_behavior",
		},
		&requestflag.Flag[string]{
			Name:     "harness",
			Usage:    "The agent or tool that submitted the report, such as `Claude Code`.",
			BodyPath: "harness",
		},
		&requestflag.Flag[string]{
			Name:     "location",
			Usage:    "Where it happened, such as an endpoint (`POST /send`), an SDK method, an MCP tool name or a docs URL.",
			BodyPath: "location",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "The AI model that submitted the report, such as `claude-opus-5-5`.",
			BodyPath: "model",
		},
		&requestflag.Flag[string]{
			Name:     "request-id",
			Usage:    "The request ID of an API call that shows the problem, if there is one.",
			BodyPath: "request_id",
		},
	},
	Action:          handleFeedbackCreate,
	HideHelpCommand: true,
}

func handleFeedbackCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := courier.FeedbackNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Feedback.New(ctx, params, options...)
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
		Title:          "feedback create",
		Transform:      transform,
	})
}
