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

var inboundTrackEvent = cli.Command{
	Name:    "track-event",
	Usage:   "Courier Track Event",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "event",
			Usage:    "A descriptive name of the event. This name will appear as a trigger in the Courier Automation Trigger node.",
			Required: true,
			BodyPath: "event",
		},
		&requestflag.Flag[string]{
			Name:     "message-id",
			Usage:    "A required unique identifier that will be used to de-duplicate requests. If not unique, will respond with 409 Conflict status",
			Required: true,
			BodyPath: "messageId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "properties",
			Required: true,
			BodyPath: "properties",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "track".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[any]{
			Name:     "user-id",
			Usage:    "The user id associated with the track",
			BodyPath: "userId",
		},
	},
	Action:          handleInboundTrackEvent,
	HideHelpCommand: true,
}

func handleInboundTrackEvent(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.InboundTrackEventParams{}

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
	_, err = client.Inbound.TrackEvent(ctx, params, options...)
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
		Title:          "inbound track-event",
		Transform:      transform,
	})
}
