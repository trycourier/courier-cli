// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v4/internal/mocktest"
)

func TestProvidersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"providers", "create",
			"--provider", "sendgrid",
			"--alias", "alias",
			"--settings", "{api_key: bar}",
			"--title", "Production SendGrid",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"provider: sendgrid\n" +
			"alias: alias\n" +
			"settings:\n" +
			"  api_key: bar\n" +
			"title: Production SendGrid\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"providers", "create",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestProvidersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"providers", "retrieve",
			"--id", "id",
		)
	})
}

func TestProvidersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"providers", "update",
			"--id", "id",
			"--provider", "sendgrid",
			"--alias", "alias",
			"--settings", "{api_key: bar}",
			"--title", "Production SendGrid",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"provider: sendgrid\n" +
			"alias: alias\n" +
			"settings:\n" +
			"  api_key: bar\n" +
			"title: Production SendGrid\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"providers", "update",
			"--id", "id",
		)
	})
}

func TestProvidersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"providers", "list",
			"--cursor", "cursor",
		)
	})
}

func TestProvidersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"providers", "delete",
			"--id", "id",
		)
	})
}
