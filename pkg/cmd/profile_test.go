// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
	"github.com/trycourier/courier-cli/internal/requestflag"
)

func TestProfilesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "create",
		"--api-key", "string",
		"--user-id", "user_id",
		"--profile", "{foo: bar}",
	)
}

func TestProfilesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "retrieve",
		"--api-key", "string",
		"--user-id", "user_id",
	)
}

func TestProfilesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "update",
		"--api-key", "string",
		"--user-id", "user_id",
		"--patch", "{op: op, path: path, value: value}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(profilesUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "update",
		"--user-id", "user_id",
		"--patch.op", "op",
		"--patch.path", "path",
		"--patch.value", "value",
	)
}

func TestProfilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "delete",
		"--api-key", "string",
		"--user-id", "user_id",
	)
}

func TestProfilesReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles", "replace",
		"--api-key", "string",
		"--user-id", "user_id",
		"--profile", "{foo: bar}",
	)
}
