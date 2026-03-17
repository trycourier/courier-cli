// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestTenantsPreferencesItemsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants:preferences:items", "update",
			"--tenant-id", "tenant_id",
			"--topic-id", "topic_id",
			"--status", "OPTED_IN",
			"--custom-routing", "[inbox]",
			"--has-custom-routing=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"status: OPTED_IN\n" +
			"custom_routing:\n" +
			"  - inbox\n" +
			"has_custom_routing: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tenants:preferences:items", "update",
			"--tenant-id", "tenant_id",
			"--topic-id", "topic_id",
		)
	})
}

func TestTenantsPreferencesItemsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants:preferences:items", "delete",
			"--tenant-id", "tenant_id",
			"--topic-id", "topic_id",
		)
	})
}
