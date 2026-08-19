// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v5/internal/mocktest"
)

func TestJourneysRunsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:runs", "retrieve",
			"--run-id", "x",
		)
	})
}

func TestJourneysRunsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:runs", "list",
			"--cursor", "cursor",
			"--end-date", "end_date",
			"--limit", "321669910225",
			"--start-date", "start_date",
			"--status", "status",
			"--template-id", "template_id",
		)
	})
}

func TestJourneysRunsListSteps(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:runs", "list-steps",
			"--run-id", "x",
		)
	})
}
