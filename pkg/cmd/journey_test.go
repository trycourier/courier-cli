// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v5/internal/mocktest"
)

func TestJourneysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys", "create",
			"--name", "Welcome Journey",
			"--node", "{trigger_type: api-invoke, type: trigger, id: trigger-1, conditions: [string, string], schema: {foo: bar}}",
			"--node", "{message: {context: {tenant_id: x}, data: {foo: bar}, delay: {until: x, timezone: x}, template: nt_01kx4h2jdafq8bk9aftxak4b40, to: {email_override: x, phone_number_override: x, user_id_override: x}}, type: send, id: send-1, conditions: [string, string], experiment: {bucketingKey: x, variants: [{id: x, templateId: x, weight: 0, name: name}, {id: x, templateId: x, weight: 0, name: name}], id: x, name: name}}",
			"--node", "{type: exit, id: exit-1}",
			"--enabled=true",
			"--state", "DRAFT",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Welcome Journey\n" +
			"nodes:\n" +
			"  - trigger_type: api-invoke\n" +
			"    type: trigger\n" +
			"    id: trigger-1\n" +
			"    conditions:\n" +
			"      - string\n" +
			"      - string\n" +
			"    schema:\n" +
			"      foo: bar\n" +
			"  - message:\n" +
			"      context:\n" +
			"        tenant_id: x\n" +
			"      data:\n" +
			"        foo: bar\n" +
			"      delay:\n" +
			"        until: x\n" +
			"        timezone: x\n" +
			"      template: nt_01kx4h2jdafq8bk9aftxak4b40\n" +
			"      to:\n" +
			"        email_override: x\n" +
			"        phone_number_override: x\n" +
			"        user_id_override: x\n" +
			"    type: send\n" +
			"    id: send-1\n" +
			"    conditions:\n" +
			"      - string\n" +
			"      - string\n" +
			"    experiment:\n" +
			"      bucketingKey: x\n" +
			"      variants:\n" +
			"        - id: x\n" +
			"          templateId: x\n" +
			"          weight: 0\n" +
			"          name: name\n" +
			"        - id: x\n" +
			"          templateId: x\n" +
			"          weight: 0\n" +
			"          name: name\n" +
			"      id: x\n" +
			"      name: name\n" +
			"  - type: exit\n" +
			"    id: exit-1\n" +
			"enabled: true\n" +
			"state: DRAFT\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"journeys", "create",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestJourneysRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys", "retrieve",
			"--template-id", "x",
			"--version", "published",
		)
	})
}

func TestJourneysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys", "list",
			"--cursor", "cursor",
			"--version", "published",
		)
	})
}

func TestJourneysArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys", "archive",
			"--template-id", "x",
		)
	})
}

func TestJourneysCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys", "cancel",
			"--cancelation-token", "x",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("cancelation_token: x")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"journeys", "cancel",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestJourneysInvoke(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys", "invoke",
			"--template-id", "templateId",
			"--data", "{order_id: bar, amount: bar}",
			"--profile", "{foo: bar}",
			"--user-id", "user-123",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data:\n" +
			"  order_id: bar\n" +
			"  amount: bar\n" +
			"profile:\n" +
			"  foo: bar\n" +
			"user_id: user-123\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"journeys", "invoke",
			"--template-id", "templateId",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestJourneysListVersions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys", "list-versions",
			"--template-id", "x",
		)
	})
}

func TestJourneysPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys", "publish",
			"--template-id", "x",
			"--version", "v321669910225",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("version: v321669910225")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"journeys", "publish",
			"--template-id", "x",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestJourneysReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys", "replace",
			"--template-id", "x",
			"--name", "Welcome Journey v2",
			"--node", "{trigger_type: api-invoke, type: trigger, id: x, conditions: [string, string], schema: {foo: bar}}",
			"--enabled=true",
			"--state", "DRAFT",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Welcome Journey v2\n" +
			"nodes:\n" +
			"  - trigger_type: api-invoke\n" +
			"    type: trigger\n" +
			"    id: x\n" +
			"    conditions:\n" +
			"      - string\n" +
			"      - string\n" +
			"    schema:\n" +
			"      foo: bar\n" +
			"enabled: true\n" +
			"state: DRAFT\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"journeys", "replace",
			"--template-id", "x",
		)
	})
}
