// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestDigestsSchedulesListInstances(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"digests:schedules", "list-instances",
			"--schedule-id", "schedule_id",
			"--cursor", "cursor",
			"--limit", "100",
		)
	})
}

func TestDigestsSchedulesRelease(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"digests:schedules", "release",
			"--schedule-id", "schedule_id",
		)
	})
}
