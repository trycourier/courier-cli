// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
	"github.com/trycourier/courier-cli/internal/requestflag"
)

func TestListsSubscriptionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "lists:subscriptions", "list",
			"--api-key", "string",
			"--list-id", "list_id",
			"--cursor", "cursor",
		)
	})
}

func TestListsSubscriptionsAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "lists:subscriptions", "add",
			"--api-key", "string",
			"--list-id", "list_id",
			"--recipient", "{recipientId: recipientId, preferences: {categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(listsSubscriptionsAdd)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "lists:subscriptions", "add",
			"--api-key", "string",
			"--list-id", "list_id",
			"--recipient.recipient-id", "recipientId",
			"--recipient.preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"recipients:\n" +
			"  - recipientId: recipientId\n" +
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
			t, pipeData, "lists:subscriptions", "add",
			"--api-key", "string",
			"--list-id", "list_id",
		)
	})
}

func TestListsSubscriptionsSubscribe(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "lists:subscriptions", "subscribe",
			"--api-key", "string",
			"--list-id", "list_id",
			"--recipient", "{recipientId: recipientId, preferences: {categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(listsSubscriptionsSubscribe)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "lists:subscriptions", "subscribe",
			"--api-key", "string",
			"--list-id", "list_id",
			"--recipient.recipient-id", "recipientId",
			"--recipient.preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"recipients:\n" +
			"  - recipientId: recipientId\n" +
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
			t, pipeData, "lists:subscriptions", "subscribe",
			"--api-key", "string",
			"--list-id", "list_id",
		)
	})
}

func TestListsSubscriptionsSubscribeUser(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "lists:subscriptions", "subscribe-user",
			"--api-key", "string",
			"--list-id", "list_id",
			"--user-id", "user_id",
			"--preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(listsSubscriptionsSubscribeUser)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "lists:subscriptions", "subscribe-user",
			"--api-key", "string",
			"--list-id", "list_id",
			"--user-id", "user_id",
			"--preferences.categories", "{foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}",
			"--preferences.notifications", "{foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}",
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
			"    foo:\n" +
			"      status: OPTED_IN\n" +
			"      channel_preferences:\n" +
			"        - channel: direct_message\n" +
			"      rules:\n" +
			"        - until: until\n" +
			"          start: start\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "lists:subscriptions", "subscribe-user",
			"--api-key", "string",
			"--list-id", "list_id",
			"--user-id", "user_id",
		)
	})
}

func TestListsSubscriptionsUnsubscribeUser(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "lists:subscriptions", "unsubscribe-user",
			"--api-key", "string",
			"--list-id", "list_id",
			"--user-id", "user_id",
		)
	})
}
