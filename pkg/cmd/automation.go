// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
	"github.com/trycourier/courier-cli/internal/apiquery"
	"github.com/trycourier/courier-cli/internal/requestflag"
	"github.com/trycourier/courier-go/v4"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/urfave/cli/v3"
)

var automationsList = cli.Command{
	Name:    "list",
	Usage:   "Get the list of automations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "A cursor token for pagination. Use the cursor from the previous response to fetch the next page of results.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "version",
			Usage:     "The version of templates to retrieve. Accepted values are published (for published templates) or draft (for draft templates). Defaults to published.",
			Default:   "published",
			QueryPath: "version",
		},
	},
	Action:          handleAutomationsList,
	HideHelpCommand: true,
}

func handleAutomationsList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.AutomationListParams{}

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
	_, err = client.Automations.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "automations list", obj, format, transform)
}
