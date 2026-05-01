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

var sendMessage = requestflag.WithInnerFlags(cli.Command{
	Name:    "message",
	Usage:   "Send a message to one or more recipients.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "message",
			Usage:    "The message property has the following primary top-level properties. They define the destination and content of the message.",
			Required: true,
			BodyPath: "message",
		},
	},
	Action:          handleSendMessage,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"message": {
		&requestflag.InnerFlag[*string]{
			Name:       "message.brand-id",
			InnerField: "brand_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.channels",
			InnerField: "channels",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.content",
			Usage:      "Describes content that will work for email, inbox, push, chat, or any channel id.",
			InnerField: "content",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.context",
			InnerField: "context",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.data",
			InnerField: "data",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.delay",
			InnerField: "delay",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.expiry",
			InnerField: "expiry",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.metadata",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.preferences",
			InnerField: "preferences",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.providers",
			InnerField: "providers",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.routing",
			Usage:      "Customize which channels/providers Courier may deliver the message through.",
			InnerField: "routing",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "message.template",
			InnerField: "template",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "message.timeout",
			InnerField: "timeout",
		},
		&requestflag.InnerFlag[any]{
			Name:       "message.to",
			Usage:      "The recipient or a list of recipients of the message",
			InnerField: "to",
		},
	},
})

func handleSendMessage(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.SendMessageParams{}

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
	_, err = client.Send.Message(ctx, params, options...)
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
		Title:          "send message",
		Transform:      transform,
	})
}
