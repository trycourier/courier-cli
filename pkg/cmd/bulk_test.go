// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestBulkAddUsers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk", "add-users",
			"--job-id", "job_id",
			"--user", "{data: {}, preferences: {categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}, profile: {foo: bar}, recipient: recipient, to: {account_id: account_id, context: {tenant_id: tenant_id}, data: {foo: bar}, email: email, list_id: list_id, locale: locale, phone_number: phone_number, preferences: {notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, templateId: templateId}, tenant_id: tenant_id, user_id: user_id}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(bulkAddUsers)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk", "add-users",
			"--job-id", "job_id",
			"--user.data", "{}",
			"--user.preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
			"--user.profile", "{foo: bar}",
			"--user.recipient", "recipient",
			"--user.to", "{account_id: account_id, context: {tenant_id: tenant_id}, data: {foo: bar}, email: email, list_id: list_id, locale: locale, phone_number: phone_number, preferences: {notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, templateId: templateId}, tenant_id: tenant_id, user_id: user_id}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"users:\n" +
			"  - data: {}\n" +
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
			"    profile:\n" +
			"      foo: bar\n" +
			"    recipient: recipient\n" +
			"    to:\n" +
			"      account_id: account_id\n" +
			"      context:\n" +
			"        tenant_id: tenant_id\n" +
			"      data:\n" +
			"        foo: bar\n" +
			"      email: email\n" +
			"      list_id: list_id\n" +
			"      locale: locale\n" +
			"      phone_number: phone_number\n" +
			"      preferences:\n" +
			"        notifications:\n" +
			"          foo:\n" +
			"            status: OPTED_IN\n" +
			"            channel_preferences:\n" +
			"              - channel: direct_message\n" +
			"            rules:\n" +
			"              - until: until\n" +
			"                start: start\n" +
			"            source: subscription\n" +
			"        categories:\n" +
			"          foo:\n" +
			"            status: OPTED_IN\n" +
			"            channel_preferences:\n" +
			"              - channel: direct_message\n" +
			"            rules:\n" +
			"              - until: until\n" +
			"                start: start\n" +
			"            source: subscription\n" +
			"        templateId: templateId\n" +
			"      tenant_id: tenant_id\n" +
			"      user_id: user_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"bulk", "add-users",
			"--job-id", "job_id",
		)
	})
}

func TestBulkCreateJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk", "create-job",
			"--message", "{event: event, brand: brand, content: {body: body, title: title}, data: {foo: bar}, locale: {foo: {foo: bar}}, override: {foo: bar}, template: template}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(bulkCreateJob)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk", "create-job",
			"--message.event", "event",
			"--message.brand", "brand",
			"--message.content", "{body: body, title: title}",
			"--message.data", "{foo: bar}",
			"--message.locale", "{foo: {foo: bar}}",
			"--message.override", "{foo: bar}",
			"--message.template", "template",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"message:\n" +
			"  event: event\n" +
			"  brand: brand\n" +
			"  content:\n" +
			"    body: body\n" +
			"    title: title\n" +
			"  data:\n" +
			"    foo: bar\n" +
			"  locale:\n" +
			"    foo:\n" +
			"      foo: bar\n" +
			"  override:\n" +
			"    foo: bar\n" +
			"  template: template\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"bulk", "create-job",
		)
	})
}

func TestBulkListUsers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk", "list-users",
			"--job-id", "job_id",
			"--cursor", "cursor",
		)
	})
}

func TestBulkRetrieveJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk", "retrieve-job",
			"--job-id", "job_id",
		)
	})
}

func TestBulkRunJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk", "run-job",
			"--job-id", "job_id",
		)
	})
}
