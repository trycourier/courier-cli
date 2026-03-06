// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
	"github.com/trycourier/courier-cli/internal/requestflag"
)

func TestProfilesListsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles:lists", "retrieve",
		"--api-key", "string",
		"--user-id", "user_id",
		"--cursor", "cursor",
	)
}

func TestProfilesListsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles:lists", "delete",
		"--api-key", "string",
		"--user-id", "user_id",
	)
}

func TestProfilesListsSubscribe(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles:lists", "subscribe",
		"--api-key", "string",
		"--user-id", "user_id",
		"--list", "{listId: listId, preferences: {categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(profilesListsSubscribe)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"profiles:lists", "subscribe",
		"--api-key", "string",
		"--user-id", "user_id",
		"--list.list-id", "listId",
		"--list.preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
	)
}
