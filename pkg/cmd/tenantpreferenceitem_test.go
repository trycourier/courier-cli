// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
)

func TestTenantsPreferencesItemsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants:preferences:items", "update",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
		"--topic-id", "topic_id",
		"--status", "OPTED_IN",
		"--custom-routing", "[inbox]",
		"--has-custom-routing=true",
	)
}

func TestTenantsPreferencesItemsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants:preferences:items", "delete",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
		"--topic-id", "topic_id",
	)
}
