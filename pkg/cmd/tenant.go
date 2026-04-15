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

var tenantsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a Tenant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Required: true,
		},
	},
	Action:          handleTenantsRetrieve,
	HideHelpCommand: true,
}

var tenantsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Create or Replace a Tenant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the tenant.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "brand-id",
			Usage:    "Brand to be used for the account when one is not specified by the send call.",
			BodyPath: "brand_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default-preferences",
			BodyPath: "default_preferences",
		},
		&requestflag.Flag[any]{
			Name:     "parent-tenant-id",
			Usage:    "Tenant's parent id (if any).",
			BodyPath: "parent_tenant_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "properties",
			Usage:    "Arbitrary properties accessible to a template.",
			BodyPath: "properties",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "user-profile",
			Usage:    "A user profile object merged with user profile on send.",
			BodyPath: "user_profile",
		},
	},
	Action:          handleTenantsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"default-preferences": {
		&requestflag.InnerFlag[any]{
			Name:       "default-preferences.items",
			InnerField: "items",
		},
	},
})

var tenantsList = cli.Command{
	Name:    "list",
	Usage:   "Get a List of Tenants",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "Continue the pagination with the next cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[any]{
			Name:      "limit",
			Usage:     "The number of tenants to return \n(defaults to 20, maximum value of 100)",
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "parent-tenant-id",
			Usage:     "Filter the list of tenants by parent_id",
			QueryPath: "parent_tenant_id",
		},
	},
	Action:          handleTenantsList,
	HideHelpCommand: true,
}

var tenantsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Tenant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Required: true,
		},
	},
	Action:          handleTenantsDelete,
	HideHelpCommand: true,
}

var tenantsListUsers = cli.Command{
	Name:    "list-users",
	Usage:   "Get Users in Tenant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tenant-id",
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
	Action:          handleTenantsListUsers,
	HideHelpCommand: true,
}

func handleTenantsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tenants.Get(ctx, cmd.Value("tenant-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "tenants retrieve", obj, format, explicitFormat, transform)
}

func handleTenantsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tenant-id") && len(unusedArgs) > 0 {
		cmd.Set("tenant-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.TenantUpdateParams{}

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
	_, err = client.Tenants.Update(
		ctx,
		cmd.Value("tenant-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "tenants update", obj, format, explicitFormat, transform)
}

func handleTenantsList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.TenantListParams{}

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
	_, err = client.Tenants.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "tenants list", obj, format, explicitFormat, transform)
}

func handleTenantsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Tenants.Delete(ctx, cmd.Value("tenant-id").(string), options...)
}

func handleTenantsListUsers(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tenant-id") && len(unusedArgs) > 0 {
		cmd.Set("tenant-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.TenantListUsersParams{}

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
	_, err = client.Tenants.ListUsers(
		ctx,
		cmd.Value("tenant-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "tenants list-users", obj, format, explicitFormat, transform)
}
