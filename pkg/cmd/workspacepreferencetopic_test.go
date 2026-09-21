// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v5/internal/mocktest"
	"github.com/trycourier/courier-cli/v5/internal/requestflag"
)

func TestWorkspacePreferencesTopicsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "create",
			"--section-id", "section_id",
			"--default-status", "OPTED_OUT",
			"--name", "Marketing",
			"--allowed-preference", "[snooze]",
			"--description", "description",
			"--digest", "{schedules: [{frequency: instant, day_of_month: 1, day_of_week: sunday, days_of_week: [sunday], disabled: true, is_default: true, schedule_id: schedule_id, time: time, timezone: timezone}], template_id: template_id, audience_id: audience_id, categories: [{category_key: category_key, limit: 1, retain: FIRST, sort_key: sort_key}], trigger_empty: true}",
			"--include-unsubscribe-header=true",
			"--routing-option", "[direct_message]",
			"--topic-data", "{foo: bar}",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(workspacePreferencesTopicsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "create",
			"--section-id", "section_id",
			"--default-status", "OPTED_OUT",
			"--name", "Marketing",
			"--allowed-preference", "[snooze]",
			"--description", "description",
			"--digest.schedules", "[{frequency: instant, day_of_month: 1, day_of_week: sunday, days_of_week: [sunday], disabled: true, is_default: true, schedule_id: schedule_id, time: time, timezone: timezone}]",
			"--digest.template-id", "template_id",
			"--digest.audience-id", "audience_id",
			"--digest.categories", "[{category_key: category_key, limit: 1, retain: FIRST, sort_key: sort_key}]",
			"--digest.trigger-empty=true",
			"--include-unsubscribe-header=true",
			"--routing-option", "[direct_message]",
			"--topic-data", "{foo: bar}",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"default_status: OPTED_OUT\n" +
			"name: Marketing\n" +
			"allowed_preferences:\n" +
			"  - snooze\n" +
			"description: description\n" +
			"digest:\n" +
			"  schedules:\n" +
			"    - frequency: instant\n" +
			"      day_of_month: 1\n" +
			"      day_of_week: sunday\n" +
			"      days_of_week:\n" +
			"        - sunday\n" +
			"      disabled: true\n" +
			"      is_default: true\n" +
			"      schedule_id: schedule_id\n" +
			"      time: time\n" +
			"      timezone: timezone\n" +
			"  template_id: template_id\n" +
			"  audience_id: audience_id\n" +
			"  categories:\n" +
			"    - category_key: category_key\n" +
			"      limit: 1\n" +
			"      retain: FIRST\n" +
			"      sort_key: sort_key\n" +
			"  trigger_empty: true\n" +
			"include_unsubscribe_header: true\n" +
			"routing_options:\n" +
			"  - direct_message\n" +
			"topic_data:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-preferences:topics", "create",
			"--section-id", "section_id",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestWorkspacePreferencesTopicsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "retrieve",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
		)
	})
}

func TestWorkspacePreferencesTopicsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "list",
			"--section-id", "section_id",
		)
	})
}

func TestWorkspacePreferencesTopicsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "archive",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
		)
	})
}

func TestWorkspacePreferencesTopicsDeleteDigest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "delete-digest",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
		)
	})
}

func TestWorkspacePreferencesTopicsReleaseDigest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "release-digest",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
			"--user-id", "user_01h1p2c3d4e5f6g7h8",
			"--tenant-id", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"user_id: user_01h1p2c3d4e5f6g7h8\n" +
			"tenant_id: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-preferences:topics", "release-digest",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
		)
	})
}

func TestWorkspacePreferencesTopicsReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "replace",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
			"--default-status", "OPTED_IN",
			"--name", "Product Updates",
			"--allowed-preference", "[channel_preferences]",
			"--description", "description",
			"--digest", "{template_id: template_id, audience_id: audience_id, categories: [{category_key: category_key, limit: 1, retain: FIRST, sort_key: sort_key}], schedules: [{frequency: instant, day_of_month: 1, day_of_week: sunday, days_of_week: [sunday], disabled: true, is_default: true, schedule_id: schedule_id, time: time, timezone: timezone}], trigger_empty: true}",
			"--include-unsubscribe-header=true",
			"--routing-option", "[email, inbox]",
			"--topic-data", "{foo: bar}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(workspacePreferencesTopicsReplace)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-preferences:topics", "replace",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
			"--default-status", "OPTED_IN",
			"--name", "Product Updates",
			"--allowed-preference", "[channel_preferences]",
			"--description", "description",
			"--digest.template-id", "template_id",
			"--digest.audience-id", "audience_id",
			"--digest.categories", "[{category_key: category_key, limit: 1, retain: FIRST, sort_key: sort_key}]",
			"--digest.schedules", "[{frequency: instant, day_of_month: 1, day_of_week: sunday, days_of_week: [sunday], disabled: true, is_default: true, schedule_id: schedule_id, time: time, timezone: timezone}]",
			"--digest.trigger-empty=true",
			"--include-unsubscribe-header=true",
			"--routing-option", "[email, inbox]",
			"--topic-data", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"default_status: OPTED_IN\n" +
			"name: Product Updates\n" +
			"allowed_preferences:\n" +
			"  - channel_preferences\n" +
			"description: description\n" +
			"digest:\n" +
			"  template_id: template_id\n" +
			"  audience_id: audience_id\n" +
			"  categories:\n" +
			"    - category_key: category_key\n" +
			"      limit: 1\n" +
			"      retain: FIRST\n" +
			"      sort_key: sort_key\n" +
			"  schedules:\n" +
			"    - frequency: instant\n" +
			"      day_of_month: 1\n" +
			"      day_of_week: sunday\n" +
			"      days_of_week:\n" +
			"        - sunday\n" +
			"      disabled: true\n" +
			"      is_default: true\n" +
			"      schedule_id: schedule_id\n" +
			"      time: time\n" +
			"      timezone: timezone\n" +
			"  trigger_empty: true\n" +
			"include_unsubscribe_header: true\n" +
			"routing_options:\n" +
			"  - email\n" +
			"  - inbox\n" +
			"topic_data:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-preferences:topics", "replace",
			"--section-id", "section_id",
			"--topic-id", "topic_id",
		)
	})
}
