// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
	"github.com/trycourier/courier-cli/internal/requestflag"
)

func TestBulkAddUsers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bulk", "add-users",
		"--api-key", "string",
		"--job-id", "job_id",
		"--user", "{data: {}, preferences: {categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}, profile: {foo: bar}, recipient: recipient, to: {account_id: account_id, context: {tenant_id: tenant_id}, data: {foo: bar}, email: email, list_id: list_id, locale: locale, phone_number: phone_number, preferences: {notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, templateId: templateId}, tenant_id: tenant_id, user_id: user_id}}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(bulkAddUsers)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"bulk", "add-users",
		"--api-key", "string",
		"--job-id", "job_id",
		"--user.data", "{}",
		"--user.preferences", "{categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}, notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}]}}}",
		"--user.profile", "{foo: bar}",
		"--user.recipient", "recipient",
		"--user.to", "{account_id: account_id, context: {tenant_id: tenant_id}, data: {foo: bar}, email: email, list_id: list_id, locale: locale, phone_number: phone_number, preferences: {notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, templateId: templateId}, tenant_id: tenant_id, user_id: user_id}",
	)
}

func TestBulkCreateJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bulk", "create-job",
		"--api-key", "string",
		"--message", "{event: event, brand: brand, content: {body: body, title: title}, data: {foo: bar}, locale: {foo: {foo: bar}}, override: {foo: bar}, template: template}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(bulkCreateJob)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"bulk", "create-job",
		"--api-key", "string",
		"--message.event", "event",
		"--message.brand", "brand",
		"--message.content", "{body: body, title: title}",
		"--message.data", "{foo: bar}",
		"--message.locale", "{foo: {foo: bar}}",
		"--message.override", "{foo: bar}",
		"--message.template", "template",
	)
}

func TestBulkListUsers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bulk", "list-users",
		"--api-key", "string",
		"--job-id", "job_id",
		"--cursor", "cursor",
	)
}

func TestBulkRetrieveJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bulk", "retrieve-job",
		"--api-key", "string",
		"--job-id", "job_id",
	)
}

func TestBulkRunJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"bulk", "run-job",
		"--api-key", "string",
		"--job-id", "job_id",
	)
}
