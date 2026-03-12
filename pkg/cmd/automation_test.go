// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestAutomationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "automations", "list",
			"--api-key", "string",
			"--cursor", "cursor",
			"--version", "published",
		)
	})
}
