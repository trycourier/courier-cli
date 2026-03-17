// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestListsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists", "retrieve",
			"--list-id", "list_id",
		)
	})
}

func TestListsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists", "update",
			"--list-id", "list_id",
			"--name", "name",
			"--preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(listsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists", "update",
			"--list-id", "list_id",
			"--name", "name",
			"--preferences.categories", "{foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}",
			"--preferences.notifications", "{foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"preferences:\n" +
			"  categories:\n" +
			"    foo:\n" +
			"      status: OPTED_IN\n" +
			"      channel_preferences:\n" +
			"        - channel: direct_message\n" +
			"      rules:\n" +
			"        - until: until\n" +
			"          start: start\n" +
			"  notifications:\n" +
			"    foo:\n" +
			"      status: OPTED_IN\n" +
			"      channel_preferences:\n" +
			"        - channel: direct_message\n" +
			"      rules:\n" +
			"        - until: until\n" +
			"          start: start\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"lists", "update",
			"--list-id", "list_id",
		)
	})
}

func TestListsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists", "list",
			"--cursor", "cursor",
			"--pattern", "pattern",
		)
	})
}

func TestListsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists", "delete",
			"--list-id", "list_id",
		)
	})
}

func TestListsRestore(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists", "restore",
			"--list-id", "list_id",
		)
	})
}
