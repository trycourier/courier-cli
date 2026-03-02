// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestTranslationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"translations", "retrieve",
		"--api-key", "string",
		"--domain", "domain",
		"--locale", "locale",
	)
}

func TestTranslationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"translations", "update",
		"--api-key", "string",
		"--domain", "domain",
		"--locale", "locale",
		"--body", "body",
	)
}
