// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
)

func TestTranslationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "translations", "retrieve",
			"--api-key", "string",
			"--domain", "domain",
			"--locale", "locale",
		)
	})
}

func TestTranslationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "translations", "update",
			"--api-key", "string",
			"--domain", "domain",
			"--locale", "locale",
			"--body", "body",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("body")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "translations", "update",
			"--api-key", "string",
			"--domain", "domain",
			"--locale", "locale",
		)
	})
}
