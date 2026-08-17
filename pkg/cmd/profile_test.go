// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v4/internal/mocktest"
	"github.com/trycourier/courier-cli/v4/internal/requestflag"
)

func TestProfilesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"profiles", "create",
			"--user-id", "user_id",
			"--profile", "{email: bar, phone_number: bar}",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"profile:\n" +
			"  email: bar\n" +
			"  phone_number: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"profiles", "create",
			"--user-id", "user_id",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestProfilesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"profiles", "retrieve",
			"--user-id", "user_id",
		)
	})
}

func TestProfilesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"profiles", "update",
			"--user-id", "user_id",
			"--patch", "{op: replace, path: /email, value: jdoe@example.com}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(profilesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"profiles", "update",
			"--user-id", "user_id",
			"--patch.op", "replace",
			"--patch.path", "/email",
			"--patch.value", "jdoe@example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"patch:\n" +
			"  - op: replace\n" +
			"    path: /email\n" +
			"    value: jdoe@example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"profiles", "update",
			"--user-id", "user_id",
		)
	})
}

func TestProfilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"profiles", "delete",
			"--user-id", "user_id",
		)
	})
}

func TestProfilesReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"profiles", "replace",
			"--user-id", "user_id",
			"--profile", "{email: bar, phone_number: bar, locale: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"profile:\n" +
			"  email: bar\n" +
			"  phone_number: bar\n" +
			"  locale: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"profiles", "replace",
			"--user-id", "user_id",
		)
	})
}
