// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestNotificationsChecksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications:checks", "update",
			"--id", "id",
			"--submission-id", "submissionId",
			"--check", "{id: id, status: RESOLVED, type: custom}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(notificationsChecksUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications:checks", "update",
			"--id", "id",
			"--submission-id", "submissionId",
			"--check.id", "id",
			"--check.status", "RESOLVED",
			"--check.type", "custom",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"checks:\n" +
			"  - id: id\n" +
			"    status: RESOLVED\n" +
			"    type: custom\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications:checks", "update",
			"--id", "id",
			"--submission-id", "submissionId",
		)
	})
}

func TestNotificationsChecksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications:checks", "list",
			"--id", "id",
			"--submission-id", "submissionId",
		)
	})
}

func TestNotificationsChecksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications:checks", "delete",
			"--id", "id",
			"--submission-id", "submissionId",
		)
	})
}
