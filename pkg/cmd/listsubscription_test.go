// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestListsSubscriptionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscriptions", "list",
			"--list-id", "list_id",
			"--cursor", "cursor",
		)
	})
}

func TestListsSubscriptionsAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscriptions", "add",
			"--list-id", "list_id",
			"--recipient", "{recipientId: user_abc, preferences: {categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}}",
			"--recipient", "{recipientId: user_def, preferences: {categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}}",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(listsSubscriptionsAdd)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscriptions", "add",
			"--list-id", "list_id",
			"--recipient.recipient-id", "user_abc",
			"--recipient.preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
			"--recipient.recipient-id", "user_def",
			"--recipient.preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"recipients:\n" +
			"  - recipientId: user_abc\n" +
			"    preferences:\n" +
			"      categories:\n" +
			"        foo:\n" +
			"          status: OPTED_IN\n" +
			"          channel_preferences:\n" +
			"            - channel: direct_message\n" +
			"          rules:\n" +
			"            - until: until\n" +
			"              start: start\n" +
			"      notifications:\n" +
			"        foo:\n" +
			"          status: OPTED_IN\n" +
			"          channel_preferences:\n" +
			"            - channel: direct_message\n" +
			"          rules:\n" +
			"            - until: until\n" +
			"              start: start\n" +
			"  - recipientId: user_def\n" +
			"    preferences:\n" +
			"      categories:\n" +
			"        foo:\n" +
			"          status: OPTED_IN\n" +
			"          channel_preferences:\n" +
			"            - channel: direct_message\n" +
			"          rules:\n" +
			"            - until: until\n" +
			"              start: start\n" +
			"      notifications:\n" +
			"        foo:\n" +
			"          status: OPTED_IN\n" +
			"          channel_preferences:\n" +
			"            - channel: direct_message\n" +
			"          rules:\n" +
			"            - until: until\n" +
			"              start: start\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"lists:subscriptions", "add",
			"--list-id", "list_id",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestListsSubscriptionsSubscribe(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscriptions", "subscribe",
			"--list-id", "list_id",
			"--recipient", "{recipientId: user_abc, preferences: {categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}}",
			"--recipient", "{recipientId: user_def, preferences: {categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(listsSubscriptionsSubscribe)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscriptions", "subscribe",
			"--list-id", "list_id",
			"--recipient.recipient-id", "user_abc",
			"--recipient.preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
			"--recipient.recipient-id", "user_def",
			"--recipient.preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"recipients:\n" +
			"  - recipientId: user_abc\n" +
			"    preferences:\n" +
			"      categories:\n" +
			"        foo:\n" +
			"          status: OPTED_IN\n" +
			"          channel_preferences:\n" +
			"            - channel: direct_message\n" +
			"          rules:\n" +
			"            - until: until\n" +
			"              start: start\n" +
			"      notifications:\n" +
			"        foo:\n" +
			"          status: OPTED_IN\n" +
			"          channel_preferences:\n" +
			"            - channel: direct_message\n" +
			"          rules:\n" +
			"            - until: until\n" +
			"              start: start\n" +
			"  - recipientId: user_def\n" +
			"    preferences:\n" +
			"      categories:\n" +
			"        foo:\n" +
			"          status: OPTED_IN\n" +
			"          channel_preferences:\n" +
			"            - channel: direct_message\n" +
			"          rules:\n" +
			"            - until: until\n" +
			"              start: start\n" +
			"      notifications:\n" +
			"        foo:\n" +
			"          status: OPTED_IN\n" +
			"          channel_preferences:\n" +
			"            - channel: direct_message\n" +
			"          rules:\n" +
			"            - until: until\n" +
			"              start: start\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"lists:subscriptions", "subscribe",
			"--list-id", "list_id",
		)
	})
}

func TestListsSubscriptionsSubscribeUser(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscriptions", "subscribe-user",
			"--list-id", "list_id",
			"--user-id", "user_id",
			"--preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {nt_01kx4h2jdafq8bk9aftxak4b40: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(listsSubscriptionsSubscribeUser)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscriptions", "subscribe-user",
			"--list-id", "list_id",
			"--user-id", "user_id",
			"--preferences.categories", "{foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}",
			"--preferences.notifications", "{nt_01kx4h2jdafq8bk9aftxak4b40: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"preferences:\n" +
			"  categories:\n" +
			"    foo:\n" +
			"      status: OPTED_IN\n" +
			"      channel_preferences:\n" +
			"        - channel: direct_message\n" +
			"      rules:\n" +
			"        - until: until\n" +
			"          start: start\n" +
			"  notifications:\n" +
			"    nt_01kx4h2jdafq8bk9aftxak4b40:\n" +
			"      status: OPTED_IN\n" +
			"      channel_preferences:\n" +
			"        - channel: direct_message\n" +
			"      rules:\n" +
			"        - until: until\n" +
			"          start: start\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"lists:subscriptions", "subscribe-user",
			"--list-id", "list_id",
			"--user-id", "user_id",
		)
	})
}

func TestListsSubscriptionsUnsubscribeUser(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"lists:subscriptions", "unsubscribe-user",
			"--list-id", "list_id",
			"--user-id", "user_id",
		)
	})
}
