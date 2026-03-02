// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
	"github.com/trycourier/courier-cli/internal/requestflag"
)

func TestNotificationsChecksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:checks", "update",
		"--api-key", "string",
		"--id", "id",
		"--submission-id", "submissionId",
		"--check", "{id: id, status: RESOLVED, type: custom}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(notificationsChecksUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:checks", "update",
		"--id", "id",
		"--submission-id", "submissionId",
		"--check.id", "id",
		"--check.status", "RESOLVED",
		"--check.type", "custom",
	)
}

func TestNotificationsChecksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:checks", "list",
		"--api-key", "string",
		"--id", "id",
		"--submission-id", "submissionId",
	)
}

func TestNotificationsChecksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"notifications:checks", "delete",
		"--api-key", "string",
		"--id", "id",
		"--submission-id", "submissionId",
	)
}
