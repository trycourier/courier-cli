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

var brandsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new brand",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "id",
			BodyPath: "id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "settings",
			BodyPath: "settings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "snippets",
			BodyPath: "snippets",
		},
	},
	Action:          handleBrandsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"settings": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "settings.colors",
			InnerField: "colors",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "settings.email",
			InnerField: "email",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "settings.inapp",
			InnerField: "inapp",
		},
	},
	"snippets": {
		&requestflag.InnerFlag[any]{
			Name:       "snippets.items",
			InnerField: "items",
		},
	},
})

var brandsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch a specific brand by brand ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleBrandsRetrieve,
	HideHelpCommand: true,
}

var brandsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Replace an existing brand with the supplied values.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the brand.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "settings",
			BodyPath: "settings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "snippets",
			BodyPath: "snippets",
		},
	},
	Action:          handleBrandsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"settings": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "settings.colors",
			InnerField: "colors",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "settings.email",
			InnerField: "email",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "settings.inapp",
			InnerField: "inapp",
		},
	},
	"snippets": {
		&requestflag.InnerFlag[any]{
			Name:       "snippets.items",
			InnerField: "items",
		},
	},
})

var brandsList = cli.Command{
	Name:    "list",
	Usage:   "Get the list of brands.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "cursor",
			Usage:     "A unique identifier that allows for fetching the next set of brands.",
			QueryPath: "cursor",
		},
	},
	Action:          handleBrandsList,
	HideHelpCommand: true,
}

var brandsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a brand by brand ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Required: true,
		},
	},
	Action:          handleBrandsDelete,
	HideHelpCommand: true,
}

func handleBrandsCreate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.BrandNewParams{}

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
	_, err = client.Brands.New(ctx, params, options...)
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
		Title:          "brands create",
		Transform:      transform,
	})
}

func handleBrandsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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
	_, err = client.Brands.Get(ctx, cmd.Value("brand-id").(string), options...)
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
		Title:          "brands retrieve",
		Transform:      transform,
	})
}

func handleBrandsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.BrandUpdateParams{}

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
	_, err = client.Brands.Update(
		ctx,
		cmd.Value("brand-id").(string),
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
		Title:          "brands update",
		Transform:      transform,
	})
}

func handleBrandsList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.BrandListParams{}

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
	_, err = client.Brands.List(ctx, params, options...)
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
		Title:          "brands list",
		Transform:      transform,
	})
}

func handleBrandsDelete(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("brand-id") && len(unusedArgs) > 0 {
		cmd.Set("brand-id", unusedArgs[0])
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

	return client.Brands.Delete(ctx, cmd.Value("brand-id").(string), options...)
}
