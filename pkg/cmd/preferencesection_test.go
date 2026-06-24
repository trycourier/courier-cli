// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestPreferenceSectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"preference-sections", "create",
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
			"preference-sections", "create",
		)
	})
}

func TestPreferenceSectionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"preference-sections", "retrieve",
			"--section-id", "section_id",
		)
	})
}

func TestPreferenceSectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"preference-sections", "list",
		)
	})
}

func TestPreferenceSectionsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"preference-sections", "archive",
			"--section-id", "section_id",
		)
	})
}

func TestPreferenceSectionsPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"preference-sections", "publish",
		)
	})
}

func TestPreferenceSectionsReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"preference-sections", "replace",
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
			"preference-sections", "replace",
			"--section-id", "section_id",
		)
	})
}
