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

var translationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get translations by locale",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain",
			Required:  true,
			PathParam: "domain",
		},
		&requestflag.Flag[string]{
			Name:      "locale",
			Required:  true,
			PathParam: "locale",
		},
	},
	Action:          handleTranslationsRetrieve,
	HideHelpCommand: true,
}

var translationsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a translation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "domain",
			Required:  true,
			PathParam: "domain",
		},
		&requestflag.Flag[string]{
			Name:      "locale",
			Required:  true,
			PathParam: "locale",
		},
		&requestflag.Flag[string]{
			Name:     "body",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleTranslationsUpdate,
	HideHelpCommand: true,
}

func handleTranslationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("locale") && len(unusedArgs) > 0 {
		cmd.Set("locale", unusedArgs[0])
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

	params := courier.TranslationGetParams{
		Domain: cmd.Value("domain").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Translations.Get(
		ctx,
		cmd.Value("locale").(string),
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
		Title:          "translations retrieve",
		Transform:      transform,
	})
}

func handleTranslationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("locale") && len(unusedArgs) > 0 {
		cmd.Set("locale", unusedArgs[0])
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

	params := courier.TranslationUpdateParams{
		Domain: cmd.Value("domain").(string),
	}

	return client.Translations.Update(
		ctx,
		cmd.Value("locale").(string),
		params,
		options...,
	)
}
