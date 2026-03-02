// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
)

func TestNotificationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications", "list",
		"--api-key", "string",
		"--cursor", "cursor",
		"--notes=true",
	)
}

func TestNotificationsRetrieveContent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications", "retrieve-content",
		"--api-key", "string",
		"--id", "id",
	)
}
