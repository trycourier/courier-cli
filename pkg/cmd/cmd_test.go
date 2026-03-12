package cmd

import (
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
)

func TestCommandStructure(t *testing.T) {
	t.Parallel()

	require.NotNil(t, Command)
	assert.Equal(t, "courier", Command.Name)
	assert.Equal(t, Version, Command.Version)
}

func TestCommandHasExpectedSubcommands(t *testing.T) {
	t.Parallel()

	expectedNames := []string{
		"send",
		"audiences",
		"audit-events",
		"auth",
		"automations",
		"automations:invoke",
		"brands",
		"bulk",
		"inbound",
		"lists",
		"lists:subscriptions",
		"messages",
		"requests",
		"notifications",
		"notifications:draft",
		"notifications:checks",
		"profiles",
		"profiles:lists",
		"tenants",
		"tenants:preferences:items",
		"tenants:templates",
		"tenants:templates:versions",
		"translations",
		"users:preferences",
		"users:tenants",
		"users:tokens",
	}

	actualNames := make([]string, 0, len(Command.Commands))
	for _, cmd := range Command.Commands {
		if !cmd.Hidden {
			actualNames = append(actualNames, cmd.Name)
		}
	}

	for _, expected := range expectedNames {
		assert.True(t, slices.Contains(actualNames, expected), "missing subcommand: %s", expected)
	}
}

func TestCommandHiddenSubcommands(t *testing.T) {
	t.Parallel()

	hiddenNames := []string{"@manpages", "__complete", "@completion"}

	for _, name := range hiddenNames {
		t.Run(name, func(t *testing.T) {
			var found bool
			for _, cmd := range Command.Commands {
				if cmd.Name == name {
					found = true
					assert.True(t, cmd.Hidden, "subcommand %s should be hidden", name)
				}
			}
			assert.True(t, found, "hidden subcommand %s not found", name)
		})
	}
}

func TestCommandGlobalFlags(t *testing.T) {
	t.Parallel()

	flagNames := make([]string, 0)
	for _, flag := range Command.Flags {
		flagNames = append(flagNames, flag.Names()...)
	}

	expectedFlags := []string{"debug", "base-url", "format", "format-error", "transform", "transform-error", "api-key"}
	for _, expected := range expectedFlags {
		assert.True(t, slices.Contains(flagNames, expected), "missing global flag: %s", expected)
	}
}

func TestCommandFormatValidator(t *testing.T) {
	t.Parallel()

	// Find the format flag
	var formatFlag *cli.StringFlag
	for _, f := range Command.Flags {
		if slices.Contains(f.Names(), "format") {
			if sf, ok := f.(*cli.StringFlag); ok {
				formatFlag = sf
			}
		}
	}

	require.NotNil(t, formatFlag, "format flag should exist")

	for _, valid := range OutputFormats {
		assert.NoError(t, formatFlag.Validator(valid), "format %q should be valid", valid)
	}

	assert.Error(t, formatFlag.Validator("invalid"), "invalid format should fail validation")
	assert.Error(t, formatFlag.Validator(""), "empty format should fail validation")
}

func TestCommandSuggestEnabled(t *testing.T) {
	t.Parallel()

	assert.True(t, Command.Suggest, "root command should have Suggest enabled")

	for _, cmd := range Command.Commands {
		if !cmd.Hidden {
			assert.True(t, cmd.Suggest, "subcommand %s should have Suggest enabled", cmd.Name)
		}
	}
}

func TestSubcommandCategories(t *testing.T) {
	t.Parallel()

	for _, cmd := range Command.Commands {
		if !cmd.Hidden && cmd.Category != "" {
			assert.Equal(t, "API RESOURCE", cmd.Category, "subcommand %s has unexpected category", cmd.Name)
		}
	}
}

func TestOutputFormats(t *testing.T) {
	t.Parallel()

	expected := []string{"auto", "explore", "json", "jsonl", "pretty", "raw", "yaml"}
	assert.Equal(t, expected, OutputFormats)
}

func TestSmokeHelpOutput(t *testing.T) {
	t.Parallel()

	// Run the command with --help; it should not error
	err := Command.Run(context.Background(), []string{"courier", "--help"})
	assert.NoError(t, err)
}

func TestSmokeVersion(t *testing.T) {
	t.Parallel()

	assert.True(t, strings.HasPrefix(Version, "3."), "version should start with 3.")
	assert.NotEmpty(t, Version)
}

func TestEachResourceHasSubcommands(t *testing.T) {
	t.Parallel()

	for _, cmd := range Command.Commands {
		if cmd.Hidden {
			continue
		}

		t.Run(cmd.Name, func(t *testing.T) {
			assert.Greater(t, len(cmd.Commands), 0, "resource %s should have subcommands", cmd.Name)

			for _, sub := range cmd.Commands {
				assert.NotEmpty(t, sub.Name, "subcommand under %s should have a name", cmd.Name)
				assert.NotNil(t, sub.Action, "subcommand %s/%s should have an action", cmd.Name, sub.Name)
			}
		})
	}
}
