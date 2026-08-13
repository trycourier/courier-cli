// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestWorkspacePreferencesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "create",
			"--name", "Account Notifications",
			"--description", "description",
			"--has-custom-routing=true",
			"--routing-option", "[direct_message]",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Account Notifications\n" +
			"description: description\n" +
			"has_custom_routing: true\n" +
			"routing_options:\n" +
			"  - direct_message\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-preferences", "create",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestWorkspacePreferencesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "retrieve",
			"--section-id", "section_id",
		)
	})
}

func TestWorkspacePreferencesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "list",
		)
	})
}

func TestWorkspacePreferencesArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "archive",
			"--section-id", "section_id",
		)
	})
}

func TestWorkspacePreferencesPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "publish",
			"--brand-id", "bnd_01kx4mrd0pfzw8wt7pn7p2fzag",
			"--description", "Choose what you hear from us about.",
			"--heading", "Notification Preferences",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"brand_id: bnd_01kx4mrd0pfzw8wt7pn7p2fzag\n" +
			"description: Choose what you hear from us about.\n" +
			"heading: Notification Preferences\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-preferences", "publish",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestWorkspacePreferencesReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences", "replace",
			"--section-id", "section_id",
			"--name", "Account Notifications",
			"--description", "description",
			"--has-custom-routing=true",
			"--routing-option", "[email, push]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Account Notifications\n" +
			"description: description\n" +
			"has_custom_routing: true\n" +
			"routing_options:\n" +
			"  - email\n" +
			"  - push\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-preferences", "replace",
			"--section-id", "section_id",
		)
	})
}
