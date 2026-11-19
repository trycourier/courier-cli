// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v6/internal/mocktest"
)

func TestFeedbackCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"feedback", "create",
			"--category", "possible_bug",
			"--description", "GET /lists/{list_id}/subscriptions returned items: [] with paging.more set to true, and following the cursor returned the remaining subscribers.",
			"--surface", "api",
			"--title", "List subscriptions returns an empty page with a next cursor",
			"--expected-behavior", "Every page before the last one contains at least one item.",
			"--harness", "Claude Code",
			"--location", "GET /lists/{list_id}/subscriptions",
			"--model", "claude-opus-5-5",
			"--request-id", "1-67f340b1-58b7b231d9485ef0cda0a38b",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"category: possible_bug\n" +
			"description: >-\n" +
			"  GET /lists/{list_id}/subscriptions returned items: [] with paging.more set to\n" +
			"  true, and following the cursor returned the remaining subscribers.\n" +
			"surface: api\n" +
			"title: List subscriptions returns an empty page with a next cursor\n" +
			"expected_behavior: Every page before the last one contains at least one item.\n" +
			"harness: Claude Code\n" +
			"location: GET /lists/{list_id}/subscriptions\n" +
			"model: claude-opus-5-5\n" +
			"request_id: 1-67f340b1-58b7b231d9485ef0cda0a38b\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"feedback", "create",
		)
	})
}
