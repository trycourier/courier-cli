// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v5/internal/mocktest"
	"github.com/trycourier/courier-cli/v5/internal/requestflag"
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

func TestProvidersPatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"providers", "patch",
			"--id", "id",
			"--patch", "{op: replace, path: /title, from: from, value: Production SendGrid}",
			"--patch", "{op: replace, path: /settings/api_key, from: from, value: SG.xxxxxxxx}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(providersPatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"providers", "patch",
			"--id", "id",
			"--patch.op", "replace",
			"--patch.path", "/title",
			"--patch.from", "from",
			"--patch.value", "Production SendGrid",
			"--patch.op", "replace",
			"--patch.path", "/settings/api_key",
			"--patch.from", "from",
			"--patch.value", "SG.xxxxxxxx",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"patch:\n" +
			"  - op: replace\n" +
			"    path: /title\n" +
			"    from: from\n" +
			"    value: Production SendGrid\n" +
			"  - op: replace\n" +
			"    path: /settings/api_key\n" +
			"    from: from\n" +
			"    value: SG.xxxxxxxx\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"providers", "patch",
			"--id", "id",
		)
	})
}
