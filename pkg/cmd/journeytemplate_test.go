// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestJourneysTemplatesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:templates", "create",
			"--template-id", "x",
			"--channel", "email",
			"--notification", "{brand: {id: id}, content: {elements: [{channels: [string], if: if, loop: loop, ref: ref, type: text}], version: '2022-01-01', scope: default}, name: Welcome email, subscription: {topic_id: topic_id}, tags: [string]}",
			"--provider-key", "x",
			"--state", "state",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(journeysTemplatesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:templates", "create",
			"--template-id", "x",
			"--channel", "email",
			"--notification.brand", "{id: id}",
			"--notification.content", "{elements: [{channels: [string], if: if, loop: loop, ref: ref, type: text}], version: '2022-01-01', scope: default}",
			"--notification.name", "Welcome email",
			"--notification.subscription", "{topic_id: topic_id}",
			"--notification.tags", "[string]",
			"--provider-key", "x",
			"--state", "state",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"channel: email\n" +
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
			"        type: text\n" +
			"    version: '2022-01-01'\n" +
			"    scope: default\n" +
			"  name: Welcome email\n" +
			"  subscription:\n" +
			"    topic_id: topic_id\n" +
			"  tags:\n" +
			"    - string\n" +
			"providerKey: x\n" +
			"state: state\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"journeys:templates", "create",
			"--template-id", "x",
		)
	})
}

func TestJourneysTemplatesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:templates", "retrieve",
			"--template-id", "x",
			"--notification-id", "x",
		)
	})
}

func TestJourneysTemplatesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:templates", "list",
			"--template-id", "x",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}

func TestJourneysTemplatesArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:templates", "archive",
			"--template-id", "x",
			"--notification-id", "x",
		)
	})
}

func TestJourneysTemplatesListVersions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:templates", "list-versions",
			"--template-id", "x",
			"--notification-id", "x",
		)
	})
}

func TestJourneysTemplatesPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:templates", "publish",
			"--template-id", "x",
			"--notification-id", "x",
			"--version", "v321669910225",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("version: v321669910225")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"journeys:templates", "publish",
			"--template-id", "x",
			"--notification-id", "x",
		)
	})
}

func TestJourneysTemplatesReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:templates", "replace",
			"--template-id", "x",
			"--notification-id", "x",
			"--notification", "{brand: {id: id}, content: {elements: [{channels: [string], if: if, loop: loop, ref: ref, type: text}], version: '2022-01-01', scope: default}, name: name, subscription: {topic_id: topic_id}, tags: [string]}",
			"--state", "state",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(journeysTemplatesReplace)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"journeys:templates", "replace",
			"--template-id", "x",
			"--notification-id", "x",
			"--notification.brand", "{id: id}",
			"--notification.content", "{elements: [{channels: [string], if: if, loop: loop, ref: ref, type: text}], version: '2022-01-01', scope: default}",
			"--notification.name", "name",
			"--notification.subscription", "{topic_id: topic_id}",
			"--notification.tags", "[string]",
			"--state", "state",
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
			"        type: text\n" +
			"    version: '2022-01-01'\n" +
			"    scope: default\n" +
			"  name: name\n" +
			"  subscription:\n" +
			"    topic_id: topic_id\n" +
			"  tags:\n" +
			"    - string\n" +
			"state: state\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"journeys:templates", "replace",
			"--template-id", "x",
			"--notification-id", "x",
		)
	})
}
