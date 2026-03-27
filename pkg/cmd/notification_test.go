// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestNotificationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "create",
			"--notification", "{brand: {id: brand_abc}, content: {elements: [{channels: [string], if: if, loop: loop, ref: ref, channel: email, raw: {foo: bar}, type: channel}], version: '2022-01-01'}, name: Welcome Email, routing: {strategy_id: rs_123}, subscription: {topic_id: marketing}, tags: [onboarding, welcome]}",
			"--state", "DRAFT",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(notificationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "create",
			"--notification.brand", "{id: brand_abc}",
			"--notification.content", "{elements: [{channels: [string], if: if, loop: loop, ref: ref, channel: email, raw: {foo: bar}, type: channel}], version: '2022-01-01'}",
			"--notification.name", "Welcome Email",
			"--notification.routing", "{strategy_id: rs_123}",
			"--notification.subscription", "{topic_id: marketing}",
			"--notification.tags", "[onboarding, welcome]",
			"--state", "DRAFT",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"notification:\n" +
			"  brand:\n" +
			"    id: brand_abc\n" +
			"  content:\n" +
			"    elements:\n" +
			"      - channels:\n" +
			"          - string\n" +
			"        if: if\n" +
			"        loop: loop\n" +
			"        ref: ref\n" +
			"        channel: email\n" +
			"        raw:\n" +
			"          foo: bar\n" +
			"        type: channel\n" +
			"    version: '2022-01-01'\n" +
			"  name: Welcome Email\n" +
			"  routing:\n" +
			"    strategy_id: rs_123\n" +
			"  subscription:\n" +
			"    topic_id: marketing\n" +
			"  tags:\n" +
			"    - onboarding\n" +
			"    - welcome\n" +
			"state: DRAFT\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications", "create",
		)
	})
}

func TestNotificationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "retrieve",
			"--id", "id",
			"--version", "version",
		)
	})
}

func TestNotificationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "list",
			"--cursor", "cursor",
			"--event-id", "event_id",
			"--notes=true",
		)
	})
}

func TestNotificationsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "archive",
			"--id", "id",
		)
	})
}

func TestNotificationsListVersions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "list-versions",
			"--id", "id",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestNotificationsPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "publish",
			"--id", "id",
			"--version", "v321669910225",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("version: v321669910225")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications", "publish",
			"--id", "id",
		)
	})
}

func TestNotificationsReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "replace",
			"--id", "id",
			"--notification", "{brand: {id: id}, content: {elements: [{channels: [string], if: if, loop: loop, ref: ref, channel: email, raw: {foo: bar}, type: channel}], version: '2022-01-01'}, name: Updated Name, routing: {strategy_id: strategy_id}, subscription: {topic_id: topic_id}, tags: [updated]}",
			"--state", "PUBLISHED",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(notificationsReplace)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "replace",
			"--id", "id",
			"--notification.brand", "{id: id}",
			"--notification.content", "{elements: [{channels: [string], if: if, loop: loop, ref: ref, channel: email, raw: {foo: bar}, type: channel}], version: '2022-01-01'}",
			"--notification.name", "Updated Name",
			"--notification.routing", "{strategy_id: strategy_id}",
			"--notification.subscription", "{topic_id: topic_id}",
			"--notification.tags", "[updated]",
			"--state", "PUBLISHED",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"notification:\n" +
			"  brand:\n" +
			"    id: id\n" +
			"  content:\n" +
			"    elements:\n" +
			"      - channels:\n" +
			"          - string\n" +
			"        if: if\n" +
			"        loop: loop\n" +
			"        ref: ref\n" +
			"        channel: email\n" +
			"        raw:\n" +
			"          foo: bar\n" +
			"        type: channel\n" +
			"    version: '2022-01-01'\n" +
			"  name: Updated Name\n" +
			"  routing:\n" +
			"    strategy_id: strategy_id\n" +
			"  subscription:\n" +
			"    topic_id: topic_id\n" +
			"  tags:\n" +
			"    - updated\n" +
			"state: PUBLISHED\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications", "replace",
			"--id", "id",
		)
	})
}

func TestNotificationsRetrieveContent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications", "retrieve-content",
			"--id", "id",
		)
	})
}
