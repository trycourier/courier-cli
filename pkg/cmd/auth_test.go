// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestAuthIssueToken(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"auth", "issue-token",
			"--expires-in", "$YOUR_NUMBER days",
			"--scope", "user_id:$YOUR_USER_ID write:user-tokens inbox:read:messages inbox:write:events read:preferences write:preferences read:brands",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"expires_in: $YOUR_NUMBER days\n" +
			"scope: >-\n" +
			"  user_id:$YOUR_USER_ID write:user-tokens inbox:read:messages inbox:write:events\n" +
			"  read:preferences write:preferences read:brands\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"auth", "issue-token",
		)
	})
}
