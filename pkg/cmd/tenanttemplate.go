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

var tenantsTemplatesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a Template in Tenant",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "template-id",
			Required: true,
		},
	},
	Action:          handleTenantsTemplatesRetrieve,
	HideHelpCommand: true,
}

var tenantsTemplatesList = cli.Command{
	Name:    "list",
	Usage:   "List Templates in Tenant",
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
			Usage:     "The number of templates to return (defaults to 20, maximum value of 100)",
			QueryPath: "limit",
		},
	},
	Action:          handleTenantsTemplatesList,
	HideHelpCommand: true,
}

var tenantsTemplatesPublish = cli.Command{
	Name:    "publish",
	Usage:   "Publishes a specific version of a notification template for a tenant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "template-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "version",
			Usage:    `The version of the template to publish (e.g., "v1", "v2", "latest"). If not provided, defaults to "latest".`,
			Default:  "latest",
			BodyPath: "version",
		},
	},
	Action:          handleTenantsTemplatesPublish,
	HideHelpCommand: true,
}

var tenantsTemplatesReplace = requestflag.WithInnerFlags(cli.Command{
	Name:    "replace",
	Usage:   "Creates or updates a notification template for a tenant.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tenant-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "template-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "template",
			Usage:    "Template configuration for creating or updating a tenant notification template",
			Required: true,
			BodyPath: "template",
		},
		&requestflag.Flag[bool]{
			Name:     "published",
			Usage:    "Whether to publish the template immediately after saving. When true, the template becomes the active/published version. When false (default), the template is saved as a draft.",
			Default:  false,
			BodyPath: "published",
		},
	},
	Action:          handleTenantsTemplatesReplace,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"template": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "template.content",
			InnerField: "content",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "template.channels",
			InnerField: "channels",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "template.providers",
			InnerField: "providers",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "template.routing",
			InnerField: "routing",
		},
	},
})

func handleTenantsTemplatesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("template-id") && len(unusedArgs) > 0 {
		cmd.Set("template-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.TenantTemplateGetParams{
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tenants.Templates.Get(
		ctx,
		cmd.Value("template-id").(string),
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
		Title:          "tenants:templates retrieve",
		Transform:      transform,
	})
}

func handleTenantsTemplatesList(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tenant-id") && len(unusedArgs) > 0 {
		cmd.Set("tenant-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.TenantTemplateListParams{}

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
	_, err = client.Tenants.Templates.List(
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		Title:          "tenants:templates list",
		Transform:      transform,
	})
}

func handleTenantsTemplatesPublish(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("template-id") && len(unusedArgs) > 0 {
		cmd.Set("template-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.TenantTemplatePublishParams{
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tenants.Templates.Publish(
		ctx,
		cmd.Value("template-id").(string),
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
		Title:          "tenants:templates publish",
		Transform:      transform,
	})
}

func handleTenantsTemplatesReplace(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("template-id") && len(unusedArgs) > 0 {
		cmd.Set("template-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.TenantTemplateReplaceParams{
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tenants.Templates.Replace(
		ctx,
		cmd.Value("template-id").(string),
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
		Title:          "tenants:templates replace",
		Transform:      transform,
	})
}
