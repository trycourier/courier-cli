// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestWorkspacePreferencesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "create",
			"--name", "Account Notifications",
			"--has-custom-routing=true",
			"--routing-option", "[direct_message]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Account Notifications\n" +
			"has_custom_routing: true\n" +
			"routing_options:\n" +
			"  - direct_message\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-preferences", "create",
		)
	})
}

func TestWorkspacePreferencesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "retrieve",
			"--section-id", "section_id",
		)
	})
}

func TestWorkspacePreferencesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "list",
		)
	})
}

func TestWorkspacePreferencesArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "archive",
			"--section-id", "section_id",
		)
	})
}

func TestWorkspacePreferencesPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "publish",
		)
	})
}

func TestWorkspacePreferencesReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "replace",
			"--section-id", "section_id",
			"--name", "name",
			"--has-custom-routing=true",
			"--routing-option", "[direct_message]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"has_custom_routing: true\n" +
			"routing_options:\n" +
			"  - direct_message\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-preferences", "replace",
			"--section-id", "section_id",
		)
	})
}
