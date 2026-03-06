// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
)

func TestAuditEventsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "audit-events", "retrieve",
			"--api-key", "string",
			"--audit-event-id", "audit-event-id",
		)
	})
}

func TestAuditEventsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "audit-events", "list",
			"--api-key", "string",
			"--cursor", "cursor",
		)
	})
}
