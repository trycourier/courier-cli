// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestListsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"lists", "retrieve",
		"--api-key", "string",
		"--list-id", "list_id",
	)
}

func TestListsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"lists", "update",
		"--api-key", "string",
		"--list-id", "list_id",
		"--name", "name",
		"--preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(listsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"lists", "update",
		"--list-id", "list_id",
		"--name", "name",
		"--preferences.categories", "{foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}",
		"--preferences.notifications", "{foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}",
	)
}

func TestListsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"lists", "list",
		"--api-key", "string",
		"--cursor", "cursor",
		"--pattern", "pattern",
	)
}

func TestListsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"lists", "delete",
		"--api-key", "string",
		"--list-id", "list_id",
	)
}

func TestListsRestore(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"lists", "restore",
		"--api-key", "string",
		"--list-id", "list_id",
	)
}
