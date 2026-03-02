// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
)

func TestNotificationsDraftRetrieveContent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:draft", "retrieve-content",
		"--api-key", "string",
		"--id", "id",
	)
}
