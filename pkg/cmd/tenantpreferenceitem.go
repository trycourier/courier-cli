// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/trycourier/courier-cli/v3/internal/apiquery"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
	"github.com/trycourier/courier-go/v4"
	"github.com/urfave/cli/v3"
)

var tenantsPreferencesItemsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Create or Replace Default Preferences For Topic",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "topic-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    `Allowed values: "OPTED_OUT", "OPTED_IN", "REQUIRED".`,
			Required: true,
			BodyPath: "status",
		},
		&requestflag.Flag[any]{
			Name:     "custom-routing",
			Usage:    "The default channels to send to this tenant when has_custom_routing is enabled",
			BodyPath: "custom_routing",
		},
		&requestflag.Flag[*bool]{
			Name:     "has-custom-routing",
			Usage:    "Override channel routing with custom preferences. This will override any template preferences that are set, but a user can still customize their preferences",
			BodyPath: "has_custom_routing",
		},
	},
	Action:          handleTenantsPreferencesItemsUpdate,
	HideHelpCommand: true,
}

var tenantsPreferencesItemsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove Default Preferences For Topic",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "topic-id",
			Required: true,
		},
	},
	Action:          handleTenantsPreferencesItemsDelete,
	HideHelpCommand: true,
}

func handleTenantsPreferencesItemsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("topic-id") && len(unusedArgs) > 0 {
		cmd.Set("topic-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.TenantPreferenceItemUpdateParams{
		TenantID: cmd.Value("tenant-id").(string),
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

	return client.Tenants.Preferences.Items.Update(
		ctx,
		cmd.Value("topic-id").(string),
		params,
		options...,
	)
}

func handleTenantsPreferencesItemsDelete(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("topic-id") && len(unusedArgs) > 0 {
		cmd.Set("topic-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.TenantPreferenceItemDeleteParams{
		TenantID: cmd.Value("tenant-id").(string),
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

	return client.Tenants.Preferences.Items.Delete(
		ctx,
		cmd.Value("topic-id").(string),
		params,
		options...,
	)
}
