// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
)

func TestMessagesRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "retrieve",
		"--api-key", "string",
		"--message-id", "message_id",
	)
}

func TestMessagesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "list",
		"--api-key", "string",
		"--archived=true",
		"--cursor", "cursor",
		"--enqueued-after", "enqueued_after",
		"--event", "event",
		"--list", "list",
		"--message-id", "messageId",
		"--notification", "notification",
		"--provider", "string",
		"--recipient", "recipient",
		"--status", "string",
		"--tag", "string",
		"--tags", "tags",
		"--tenant-id", "tenant_id",
		"--trace-id", "traceId",
	)
}

func TestMessagesCancel(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "cancel",
		"--api-key", "string",
		"--message-id", "message_id",
	)
}

func TestMessagesContent(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "content",
		"--api-key", "string",
		"--message-id", "message_id",
	)
}

func TestMessagesHistory(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"messages", "history",
		"--api-key", "string",
		"--message-id", "message_id",
		"--type", "type",
	)
}
