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

var usersTokensRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get single token available for a `:token`",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "token",
			Required: true,
		},
	},
	Action:          handleUsersTokensRetrieve,
	HideHelpCommand: true,
}

var usersTokensUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Apply a JSON Patch (RFC 6902) to the specified token.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "token",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "patch",
			Required: true,
			BodyPath: "patch",
		},
	},
	Action:          handleUsersTokensUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"patch": {
		&requestflag.InnerFlag[string]{
			Name:       "patch.op",
			Usage:      "The operation to perform.",
			InnerField: "op",
		},
		&requestflag.InnerFlag[string]{
			Name:       "patch.path",
			Usage:      "The JSON path specifying the part of the profile to operate on.",
			InnerField: "path",
		},
		&requestflag.InnerFlag[any]{
			Name:       "patch.value",
			Usage:      "The value for the operation.",
			InnerField: "value",
		},
	},
})

var usersTokensList = cli.Command{
	Name:    "list",
	Usage:   "Gets all tokens available for a :user_id",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
	},
	Action:          handleUsersTokensList,
	HideHelpCommand: true,
}

var usersTokensDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete User Token",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "token",
			Required: true,
		},
	},
	Action:          handleUsersTokensDelete,
	HideHelpCommand: true,
}

var usersTokensAddMultiple = cli.Command{
	Name:    "add-multiple",
	Usage:   "Adds multiple tokens to a user and overwrites matching existing tokens.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
	},
	Action:          handleUsersTokensAddMultiple,
	HideHelpCommand: true,
}

var usersTokensAddSingle = requestflag.WithInnerFlags(cli.Command{
	Name:    "add-single",
	Usage:   "Adds a single token to a user and overwrites a matching existing token.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "token",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "provider-key",
			Usage:    `Allowed values: "firebase-fcm", "apn", "expo", "onesignal".`,
			Required: true,
			BodyPath: "provider_key",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "device",
			Usage:    "Information about the device the token came from.",
			BodyPath: "device",
		},
		&requestflag.Flag[any]{
			Name:     "expiry-date",
			Usage:    "ISO 8601 formatted date the token expires. Defaults to 2 months. Set to false to disable expiration.",
			BodyPath: "expiry_date",
		},
		&requestflag.Flag[any]{
			Name:     "properties",
			Usage:    "Properties about the token.",
			BodyPath: "properties",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "tracking",
			Usage:    "Tracking information about the device the token came from.",
			BodyPath: "tracking",
		},
	},
	Action:          handleUsersTokensAddSingle,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"device": {
		&requestflag.InnerFlag[any]{
			Name:       "device.ad-id",
			Usage:      "Id of the advertising identifier",
			InnerField: "ad_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "device.app-id",
			Usage:      "Id of the application the token is used for",
			InnerField: "app_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "device.device-id",
			Usage:      "Id of the device the token is associated with",
			InnerField: "device_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "device.manufacturer",
			Usage:      "The device manufacturer",
			InnerField: "manufacturer",
		},
		&requestflag.InnerFlag[any]{
			Name:       "device.model",
			Usage:      "The device model",
			InnerField: "model",
		},
		&requestflag.InnerFlag[any]{
			Name:       "device.platform",
			Usage:      "The device platform i.e. android, ios, web",
			InnerField: "platform",
		},
	},
	"tracking": {
		&requestflag.InnerFlag[any]{
			Name:       "tracking.ip",
			Usage:      "The IP address of the device",
			InnerField: "ip",
		},
		&requestflag.InnerFlag[any]{
			Name:       "tracking.lat",
			Usage:      "The latitude of the device",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[any]{
			Name:       "tracking.long",
			Usage:      "The longitude of the device",
			InnerField: "long",
		},
		&requestflag.InnerFlag[any]{
			Name:       "tracking.os-version",
			Usage:      "The operating system version",
			InnerField: "os_version",
		},
	},
})

func handleUsersTokensRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("token") && len(unusedArgs) > 0 {
		cmd.Set("token", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.UserTokenGetParams{
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Tokens.Get(
		ctx,
		cmd.Value("token").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:tokens retrieve", obj, format, transform)
}

func handleUsersTokensUpdate(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("token") && len(unusedArgs) > 0 {
		cmd.Set("token", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.UserTokenUpdateParams{
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

	return client.Users.Tokens.Update(
		ctx,
		cmd.Value("token").(string),
		params,
		options...,
	)
}

func handleUsersTokensList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Tokens.List(ctx, cmd.Value("user-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:tokens list", obj, format, transform)
}

func handleUsersTokensDelete(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("token") && len(unusedArgs) > 0 {
		cmd.Set("token", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.UserTokenDeleteParams{
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

	return client.Users.Tokens.Delete(
		ctx,
		cmd.Value("token").(string),
		params,
		options...,
	)
}

func handleUsersTokensAddMultiple(ctx context.Context, cmd *cli.Command) error {
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

	return client.Users.Tokens.AddMultiple(ctx, cmd.Value("user-id").(string), options...)
}

func handleUsersTokensAddSingle(ctx context.Context, cmd *cli.Command) error {
	client := courier.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("token") && len(unusedArgs) > 0 {
		cmd.Set("token", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courier.UserTokenAddSingleParams{
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

	return client.Users.Tokens.AddSingle(
		ctx,
		cmd.Value("token").(string),
		params,
		options...,
	)
}
