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

var usersTenantsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of user tenant associations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "Continue the pagination with the next cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[any]{
			Name:      "limit",
			Usage:     "The number of accounts to return \n(defaults to 20, maximum value of 100)",
			QueryPath: "limit",
		},
	},
	Action:          handleUsersTenantsList,
	HideHelpCommand: true,
}

var usersTenantsAddMultiple = requestflag.WithInnerFlags(cli.Command{
	Name:    "add-multiple",
	Usage:   "This endpoint is used to add a user to multiple tenants in one call. A custom\nprofile can also be supplied for each tenant. This profile will be merged with\nthe user's main profile when sending to the user with that tenant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "tenant",
			Required: true,
			BodyPath: "tenants",
		},
	},
	Action:          handleUsersTenantsAddMultiple,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"tenant": {
		&requestflag.InnerFlag[string]{
			Name:       "tenant.tenant-id",
			Usage:      "Tenant ID for the association between tenant and user",
			InnerField: "tenant_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "tenant.profile",
			Usage:      "Additional metadata to be applied to a user profile when used in a tenant context",
			InnerField: "profile",
		},
		&requestflag.InnerFlag[any]{
			Name:       "tenant.type",
			InnerField: "type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "tenant.user-id",
			Usage:      "User ID for the association between tenant and user",
			InnerField: "user_id",
		},
	},
})

var usersTenantsAddSingle = cli.Command{
	Name:    "add-single",
	Usage:   "This endpoint is used to add a single tenant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "profile",
			BodyPath: "profile",
		},
	},
	Action:          handleUsersTenantsAddSingle,
	HideHelpCommand: true,
}

var usersTenantsRemoveAll = cli.Command{
	Name:    "remove-all",
	Usage:   "Removes a user from any tenants they may have been associated with.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
	},
	Action:          handleUsersTenantsRemoveAll,
	HideHelpCommand: true,
}

var usersTenantsRemoveSingle = cli.Command{
	Name:    "remove-single",
	Usage:   "Removes a user from the supplied tenant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Required: true,
		},
	},
	Action:          handleUsersTenantsRemoveSingle,
	HideHelpCommand: true,
}

func handleUsersTenantsList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.UserTenantListParams{}

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
	_, err = client.Users.Tenants.List(
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
	return ShowJSON(os.Stdout, "users:tenants list", obj, format, transform)
}

func handleUsersTenantsAddMultiple(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.UserTenantAddMultipleParams{}

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

	return client.Users.Tenants.AddMultiple(
		ctx,
		cmd.Value("user-id").(string),
		params,
		options...,
	)
}

func handleUsersTenantsAddSingle(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tenant-id") && len(unusedArgs) > 0 {
		cmd.Set("tenant-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.UserTenantAddSingleParams{
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

	return client.Users.Tenants.AddSingle(
		ctx,
		cmd.Value("tenant-id").(string),
		params,
		options...,
	)
}

func handleUsersTenantsRemoveAll(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
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

	return client.Users.Tenants.RemoveAll(ctx, cmd.Value("user-id").(string), options...)
}

func handleUsersTenantsRemoveSingle(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tenant-id") && len(unusedArgs) > 0 {
		cmd.Set("tenant-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.UserTenantRemoveSingleParams{
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

	return client.Users.Tenants.RemoveSingle(
		ctx,
		cmd.Value("tenant-id").(string),
		params,
		options...,
	)
}
