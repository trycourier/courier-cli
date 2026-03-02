// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestInboundTrackEvent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"inbound", "track-event",
		"--api-key", "string",
		"--event", "New Order Placed",
		"--message-id", "4c62c457-b329-4bea-9bfc-17bba86c393f",
		"--properties", "{order_id: bar, total_orders: bar, last_order_id: bar}",
		"--type", "track",
		"--user-id", "1234",
	)
}
