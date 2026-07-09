// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestWorkspacePreferencesTopicsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "create",
			"--section-id", "section_id",
			"--default-status", "OPTED_OUT",
			"--name", "Marketing",
			"--allowed-preference", "[snooze]",
			"--description", "description",
			"--include-unsubscribe-header=true",
			"--routing-option", "[direct_message]",
			"--topic-data", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"default_status: OPTED_OUT\n" +
			"name: Marketing\n" +
			"allowed_preferences:\n" +
			"  - snooze\n" +
			"description: description\n" +
			"include_unsubscribe_header: true\n" +
			"routing_options:\n" +
			"  - direct_message\n" +
			"topic_data:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-preferences:topics", "create",
			"--section-id", "section_id",
		)
	})
}

func TestWorkspacePreferencesTopicsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "retrieve",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
		)
	})
}

func TestWorkspacePreferencesTopicsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "list",
			"--section-id", "section_id",
		)
	})
}

func TestWorkspacePreferencesTopicsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "archive",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
		)
	})
}

func TestWorkspacePreferencesTopicsReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "replace",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
			"--default-status", "OPTED_OUT",
			"--name", "name",
			"--allowed-preference", "[snooze]",
			"--description", "description",
			"--include-unsubscribe-header=true",
			"--routing-option", "[direct_message]",
			"--topic-data", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"default_status: OPTED_OUT\n" +
			"name: name\n" +
			"allowed_preferences:\n" +
			"  - snooze\n" +
			"description: description\n" +
			"include_unsubscribe_header: true\n" +
			"routing_options:\n" +
			"  - direct_message\n" +
			"topic_data:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-preferences:topics", "replace",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
		)
	})
}
