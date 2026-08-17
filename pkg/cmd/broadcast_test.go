// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v4/internal/mocktest"
	"github.com/trycourier/courier-cli/v4/internal/requestflag"
)

func TestBroadcastsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "create",
			"--channel", "email",
			"--name", "Spring Sale Announcement",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"channel: email\n" +
			"name: Spring Sale Announcement\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"broadcasts", "create",
		)
	})
}

func TestBroadcastsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "retrieve",
			"--broadcast-id", "broadcastId",
		)
	})
}

func TestBroadcastsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "update",
			"--broadcast-id", "broadcastId",
			"--name", "Spring Sale Announcement (v2)",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: Spring Sale Announcement (v2)")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"broadcasts", "update",
			"--broadcast-id", "broadcastId",
		)
	})
}

func TestBroadcastsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "list",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}

func TestBroadcastsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "archive",
			"--broadcast-id", "broadcastId",
		)
	})
}

func TestBroadcastsCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "cancel",
			"--broadcast-id", "broadcastId",
		)
	})
}

func TestBroadcastsDuplicate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "duplicate",
			"--broadcast-id", "broadcastId",
		)
	})
}

func TestBroadcastsPutContent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "put-content",
			"--broadcast-id", "broadcastId",
			"--content", "{elements: [{channels: [string], if: if, loop: loop, ref: ref, type: meta}, {channels: [string], if: if, loop: loop, ref: ref, type: text}], version: '2022-01-01'}",
			"--state", "DRAFT",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(broadcastsPutContent)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "put-content",
			"--broadcast-id", "broadcastId",
			"--content.elements", "[{channels: [string], if: if, loop: loop, ref: ref, type: meta}, {channels: [string], if: if, loop: loop, ref: ref, type: text}]",
			"--content.version", "2022-01-01",
			"--state", "DRAFT",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content:\n" +
			"  elements:\n" +
			"    - channels:\n" +
			"        - string\n" +
			"      if: if\n" +
			"      loop: loop\n" +
			"      ref: ref\n" +
			"      type: meta\n" +
			"    - channels:\n" +
			"        - string\n" +
			"      if: if\n" +
			"      loop: loop\n" +
			"      ref: ref\n" +
			"      type: text\n" +
			"  version: '2022-01-01'\n" +
			"state: DRAFT\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"broadcasts", "put-content",
			"--broadcast-id", "broadcastId",
		)
	})
}

func TestBroadcastsRetrieveContent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "retrieve-content",
			"--broadcast-id", "broadcastId",
			"--version", "version",
		)
	})
}

func TestBroadcastsSchedule(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "schedule",
			"--broadcast-id", "broadcastId",
			"--recipient-id", "aud_01kx4h2jdafq8bk9amzvy6hbv0",
			"--recipient-type", "audience",
			"--scheduled-to", "2026-08-01T15:00:00",
			"--timezone", "America/New_York",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"recipient_id: aud_01kx4h2jdafq8bk9amzvy6hbv0\n" +
			"recipient_type: audience\n" +
			"scheduled_to: '2026-08-01T15:00:00'\n" +
			"timezone: America/New_York\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"broadcasts", "schedule",
			"--broadcast-id", "broadcastId",
		)
	})
}

func TestBroadcastsSend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"broadcasts", "send",
			"--broadcast-id", "broadcastId",
			"--recipient-id", "cool-customers",
			"--recipient-type", "list",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"recipient_id: cool-customers\n" +
			"recipient_type: list\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"broadcasts", "send",
			"--broadcast-id", "broadcastId",
		)
	})
}
