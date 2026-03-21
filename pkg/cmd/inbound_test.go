// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
)

func TestInboundTrackEvent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"inbound", "track-event",
			"--event", "New Order Placed",
			"--message-id", "4c62c457-b329-4bea-9bfc-17bba86c393f",
			"--properties", "{order_id: bar, total_orders: bar, last_order_id: bar}",
			"--type", "track",
			"--user-id", "1234",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"event: New Order Placed\n" +
			"messageId: 4c62c457-b329-4bea-9bfc-17bba86c393f\n" +
			"properties:\n" +
			"  order_id: bar\n" +
			"  total_orders: bar\n" +
			"  last_order_id: bar\n" +
			"type: track\n" +
			"userId: '1234'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"inbound", "track-event",
		)
	})
}
