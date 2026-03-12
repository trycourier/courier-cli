// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestJourneysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "journeys", "list",
			"--api-key", "string",
			"--cursor", "cursor",
			"--version", "published",
		)
	})
}

func TestJourneysInvoke(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "journeys", "invoke",
			"--api-key", "string",
			"--template-id", "templateId",
			"--data", "{order_id: bar, amount: bar}",
			"--profile", "{foo: bar}",
			"--user-id", "user-123",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data:\n" +
			"  order_id: bar\n" +
			"  amount: bar\n" +
			"profile:\n" +
			"  foo: bar\n" +
			"user_id: user-123\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "journeys", "invoke",
			"--api-key", "string",
			"--template-id", "templateId",
		)
	})
}
