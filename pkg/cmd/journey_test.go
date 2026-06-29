// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
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
			"--node", "{trigger_type: api-invoke, type: trigger, id: send-1, conditions: [string, string], schema: {foo: bar}}",
			"--enabled=true",
			"--state", "DRAFT",
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
			"  - trigger_type: api-invoke\n" +
			"    type: trigger\n" +
			"    id: send-1\n" +
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
			"journeys", "create",
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
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("cancelation_token: x")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"journeys", "cancel",
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
