// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
)

func TestTenantsTemplatesVersionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants:templates:versions", "retrieve",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
		"--template-id", "template_id",
		"--version", "version",
	)
}
