// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestUsersPreferencesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:preferences", "retrieve",
			"--api-key", "string",
			"--user-id", "user_id",
			"--tenant-id", "tenant_id",
		)
	})
}

func TestUsersPreferencesRetrieveTopic(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:preferences", "retrieve-topic",
			"--api-key", "string",
			"--user-id", "user_id",
			"--topic-id", "topic_id",
			"--tenant-id", "tenant_id",
		)
	})
}

func TestUsersPreferencesUpdateOrCreateTopic(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:preferences", "update-or-create-topic",
			"--api-key", "string",
			"--user-id", "user_id",
			"--topic-id", "topic_id",
			"--topic", "{status: OPTED_IN, custom_routing: [inbox, email], has_custom_routing: true}",
			"--tenant-id", "tenant_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(usersPreferencesUpdateOrCreateTopic)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "users:preferences", "update-or-create-topic",
			"--api-key", "string",
			"--user-id", "user_id",
			"--topic-id", "topic_id",
			"--topic.status", "OPTED_IN",
			"--topic.custom-routing", "[inbox, email]",
			"--topic.has-custom-routing=true",
			"--tenant-id", "tenant_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"topic:\n" +
			"  status: OPTED_IN\n" +
			"  custom_routing:\n" +
			"    - inbox\n" +
			"    - email\n" +
			"  has_custom_routing: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "users:preferences", "update-or-create-topic",
			"--api-key", "string",
			"--user-id", "user_id",
			"--topic-id", "topic_id",
			"--tenant-id", "tenant_id",
		)
	})
}
