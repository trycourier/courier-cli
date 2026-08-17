// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/trycourier/courier-cli/v4/internal/apiquery"
	"github.com/trycourier/courier-cli/v4/internal/requestflag"
	"github.com/trycourier/courier-go/v4"
	"github.com/trycourier/courier-go/v4/option"
	"github.com/urfave/cli/v3"
)

var usersTenantsList = cli.Command{
	Name:    "list",
	Usage:   "Returns the tenants a user belongs to, with cursor paging. A user can belong to\nmany tenants, each with its own profile and preferences.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
		},
		&requestflag.Flag[*string]{
			Name:      "cursor",
			Usage:     "Continue the pagination with the next cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[*int64]{
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
	Usage:   "Adds a user to several tenants in one call, each optionally with a per-tenant\nprofile that overrides their workspace profile.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "tenant.profile",
			Usage:      "Additional metadata to be applied to a user profile when used in a tenant context",
			InnerField: "profile",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "tenant.type",
			Usage:      `Allowed values: "user".`,
			InnerField: "type",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "tenant.user-id",
			Usage:      "User ID for the association between tenant and user",
			InnerField: "user_id",
		},
	},
})

var usersTenantsAddSingle = cli.Command{
	Name:    "add-single",
	Usage:   "Adds a user to one tenant, optionally with a tenant-specific profile that\noverrides their workspace profile for sends in that tenant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
		},
		&requestflag.Flag[string]{
			Name:      "tenant-id",
			Required:  true,
			PathParam: "tenant_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "profile",
			BodyPath: "profile",
		},
	},
	Action:          handleUsersTenantsAddSingle,
	HideHelpCommand: true,
}

var usersTenantsRemoveAll = cli.Command{
	Name:    "remove-all",
	Usage:   "Removes a user from every tenant they belong to in one call. Their\nworkspace-level profile is a separate resource.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
		},
	},
	Action:          handleUsersTenantsRemoveAll,
	HideHelpCommand: true,
}

var usersTenantsRemoveSingle = cli.Command{
	Name:    "remove-single",
	Usage:   "Removes a user from one tenant. Their other tenant memberships and workspace\nprofile are managed through separate endpoints.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "user_id",
		},
		&requestflag.Flag[string]{
			Name:      "tenant-id",
			Required:  true,
			PathParam: "tenant_id",
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

	params := courier.UserTenantListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "users:tenants list",
		Transform:      transform,
	})
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

	params := courier.UserTenantAddMultipleParams{}

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

	params := courier.UserTenantAddSingleParams{
		UserID: cmd.Value("user-id").(string),
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

	params := courier.UserTenantRemoveSingleParams{
		UserID: cmd.Value("user-id").(string),
	}

	return client.Users.Tenants.RemoveSingle(
		ctx,
		cmd.Value("tenant-id").(string),
		params,
		options...,
	)
}
