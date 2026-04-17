// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/trycourier/courier-cli/v3/internal/autocomplete"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "courier",
		Usage:     "CLI for the Courier API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Sources: cli.EnvVars("COURIER_API_KEY"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "send",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sendMessage,
				},
			},
			{
				Name:     "audiences",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&audiencesRetrieve,
					&audiencesUpdate,
					&audiencesList,
					&audiencesDelete,
					&audiencesListMembers,
				},
			},
			{
				Name:     "providers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&providersCreate,
					&providersRetrieve,
					&providersUpdate,
					&providersList,
					&providersDelete,
				},
			},
			{
				Name:     "providers:catalog",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&providersCatalogList,
				},
			},
			{
				Name:     "audit-events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&auditEventsRetrieve,
					&auditEventsList,
				},
			},
			{
				Name:     "auth",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&authIssueToken,
				},
			},
			{
				Name:     "automations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&automationsList,
				},
			},
			{
				Name:     "automations:invoke",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&automationsInvokeInvokeAdHoc,
					&automationsInvokeInvokeByTemplate,
				},
			},
			{
				Name:     "journeys",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&journeysList,
					&journeysInvoke,
				},
			},
			{
				Name:     "brands",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&brandsCreate,
					&brandsRetrieve,
					&brandsUpdate,
					&brandsList,
					&brandsDelete,
				},
			},
			{
				Name:     "bulk",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&bulkAddUsers,
					&bulkCreateJob,
					&bulkListUsers,
					&bulkRetrieveJob,
					&bulkRunJob,
				},
			},
			{
				Name:     "inbound",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&inboundTrackEvent,
				},
			},
			{
				Name:     "lists",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&listsRetrieve,
					&listsUpdate,
					&listsList,
					&listsDelete,
					&listsRestore,
				},
			},
			{
				Name:     "lists:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&listsSubscriptionsList,
					&listsSubscriptionsAdd,
					&listsSubscriptionsSubscribe,
					&listsSubscriptionsSubscribeUser,
					&listsSubscriptionsUnsubscribeUser,
				},
			},
			{
				Name:     "messages",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&messagesRetrieve,
					&messagesList,
					&messagesCancel,
					&messagesContent,
					&messagesHistory,
				},
			},
			{
				Name:     "requests",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&requestsArchive,
				},
			},
			{
				Name:     "notifications",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&notificationsCreate,
					&notificationsRetrieve,
					&notificationsList,
					&notificationsArchive,
					&notificationsListVersions,
					&notificationsPublish,
					&notificationsPutContent,
					&notificationsPutElement,
					&notificationsPutLocale,
					&notificationsReplace,
					&notificationsRetrieveContent,
				},
			},
			{
				Name:     "notifications:checks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&notificationsChecksUpdate,
					&notificationsChecksList,
					&notificationsChecksDelete,
				},
			},
			{
				Name:     "routing-strategies",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&routingStrategiesCreate,
					&routingStrategiesRetrieve,
					&routingStrategiesList,
					&routingStrategiesArchive,
					&routingStrategiesListNotifications,
					&routingStrategiesReplace,
				},
			},
			{
				Name:     "profiles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&profilesCreate,
					&profilesRetrieve,
					&profilesUpdate,
					&profilesDelete,
					&profilesReplace,
				},
			},
			{
				Name:     "profiles:lists",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&profilesListsRetrieve,
					&profilesListsDelete,
					&profilesListsSubscribe,
				},
			},
			{
				Name:     "tenants",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&tenantsRetrieve,
					&tenantsUpdate,
					&tenantsList,
					&tenantsDelete,
					&tenantsListUsers,
				},
			},
			{
				Name:     "tenants:preferences:items",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&tenantsPreferencesItemsUpdate,
					&tenantsPreferencesItemsDelete,
				},
			},
			{
				Name:     "tenants:templates",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&tenantsTemplatesRetrieve,
					&tenantsTemplatesList,
					&tenantsTemplatesPublish,
					&tenantsTemplatesReplace,
				},
			},
			{
				Name:     "tenants:templates:versions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&tenantsTemplatesVersionsRetrieve,
				},
			},
			{
				Name:     "translations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&translationsRetrieve,
					&translationsUpdate,
				},
			},
			{
				Name:     "users:preferences",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersPreferencesRetrieve,
					&usersPreferencesRetrieveTopic,
					&usersPreferencesUpdateOrCreateTopic,
				},
			},
			{
				Name:     "users:tenants",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersTenantsList,
					&usersTenantsAddMultiple,
					&usersTenantsAddSingle,
					&usersTenantsRemoveAll,
					&usersTenantsRemoveSingle,
				},
			},
			{
				Name:     "users:tokens",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersTokensRetrieve,
					&usersTokensUpdate,
					&usersTokensList,
					&usersTokensDelete,
					&usersTokensAddMultiple,
					&usersTokensAddSingle,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "courier @manpages [-o courier.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "courier.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "courier.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
