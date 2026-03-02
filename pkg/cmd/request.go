// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/trycourier/courier-cli/internal/apiquery"
	"github.com/trycourier/courier-cli/internal/requestflag"
	"github.com/trycourier/courier-go/v4"
	"github.com/urfave/cli/v3"
)

var requestsArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archive message",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "request-id",
			Required: true,
		},
	},
	Action:          handleRequestsArchive,
	HideHelpCommand: true,
}

func handleRequestsArchive(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("request-id") && len(unusedArgs) > 0 {
		cmd.Set("request-id", unusedArgs[0])
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

	return client.Requests.Archive(ctx, cmd.Value("request-id").(string), options...)
}
