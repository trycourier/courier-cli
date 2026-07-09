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
			t,
			"--api-key", "string",
			"users:preferences", "retrieve",
			"--user-id", "user_id",
			"--tenant-id", "tenant_id",
		)
	})
}

func TestUsersPreferencesBulkReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:preferences", "bulk-replace",
			"--user-id", "user_id",
			"--topic", "{status: OPTED_IN, topic_id: 74Q4QGFBEX481DP6JRPMV751H4XT, custom_routing: [inbox, email], has_custom_routing: true}",
			"--tenant-id", "tenant_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(usersPreferencesBulkReplace)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:preferences", "bulk-replace",
			"--user-id", "user_id",
			"--topic.status", "OPTED_IN",
			"--topic.topic-id", "74Q4QGFBEX481DP6JRPMV751H4XT",
			"--topic.custom-routing", "[inbox, email]",
			"--topic.has-custom-routing=true",
			"--tenant-id", "tenant_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"topics:\n" +
			"  - status: OPTED_IN\n" +
			"    topic_id: 74Q4QGFBEX481DP6JRPMV751H4XT\n" +
			"    custom_routing:\n" +
			"      - inbox\n" +
			"      - email\n" +
			"    has_custom_routing: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"users:preferences", "bulk-replace",
			"--user-id", "user_id",
			"--tenant-id", "tenant_id",
		)
	})
}

func TestUsersPreferencesBulkUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:preferences", "bulk-update",
			"--user-id", "user_id",
			"--topic", "{status: OPTED_IN, topic_id: 74Q4QGFBEX481DP6JRPMV751H4XT, custom_routing: [inbox, email], has_custom_routing: true}",
			"--topic", "{status: OPTED_OUT, topic_id: 5Q4QGFBEX481DP6JRPMV751H4YU, custom_routing: [direct_message], has_custom_routing: true}",
			"--tenant-id", "tenant_id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(usersPreferencesBulkUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:preferences", "bulk-update",
			"--user-id", "user_id",
			"--topic.status", "OPTED_IN",
			"--topic.topic-id", "74Q4QGFBEX481DP6JRPMV751H4XT",
			"--topic.custom-routing", "[inbox, email]",
			"--topic.has-custom-routing=true",
			"--topic.status", "OPTED_OUT",
			"--topic.topic-id", "5Q4QGFBEX481DP6JRPMV751H4YU",
			"--topic.custom-routing", "[direct_message]",
			"--topic.has-custom-routing=true",
			"--tenant-id", "tenant_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"topics:\n" +
			"  - status: OPTED_IN\n" +
			"    topic_id: 74Q4QGFBEX481DP6JRPMV751H4XT\n" +
			"    custom_routing:\n" +
			"      - inbox\n" +
			"      - email\n" +
			"    has_custom_routing: true\n" +
			"  - status: OPTED_OUT\n" +
			"    topic_id: 5Q4QGFBEX481DP6JRPMV751H4YU\n" +
			"    custom_routing:\n" +
			"      - direct_message\n" +
			"    has_custom_routing: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"users:preferences", "bulk-update",
			"--user-id", "user_id",
			"--tenant-id", "tenant_id",
		)
	})
}

func TestUsersPreferencesDeleteTopic(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:preferences", "delete-topic",
			"--user-id", "user_id",
			"--topic-id", "topic_id",
			"--tenant-id", "tenant_id",
		)
	})
}

func TestUsersPreferencesRetrieveTopic(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:preferences", "retrieve-topic",
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
			t,
			"--api-key", "string",
			"users:preferences", "update-or-create-topic",
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
			t,
			"--api-key", "string",
			"users:preferences", "update-or-create-topic",
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
			t, pipeData,
			"--api-key", "string",
			"users:preferences", "update-or-create-topic",
			"--user-id", "user_id",
			"--topic-id", "topic_id",
			"--tenant-id", "tenant_id",
		)
	})
}
