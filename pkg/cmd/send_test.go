// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
	"github.com/trycourier/courier-cli/internal/requestflag"
)

func TestSendMessage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"send", "message",
		"--api-key", "string",
		"--message", "{brand_id: brand_id, channels: {foo: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [string], routing_method: all, timeouts: {channel: 0, provider: 0}}}, content: {body: body, title: title}, context: {tenant_id: tenant_id}, data: {foo: bar}, delay: {duration: 0, timezone: timezone, until: until}, expiry: {expires_in: string, expires_at: expires_at}, metadata: {event: event, tags: [string], trace_id: trace_id, utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, preferences: {subscription_topic_id: subscription_topic_id}, providers: {foo: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, timeouts: 0}}, routing: {channels: [string], method: all}, template: template_id, timeout: {channel: {foo: 0}, criteria: no-escalation, escalation: 0, message: 0, provider: {foo: 0}}, to: {account_id: account_id, context: {tenant_id: tenant_id}, data: {foo: bar}, email: email, list_id: list_id, locale: locale, phone_number: phone_number, preferences: {notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, templateId: templateId}, tenant_id: tenant_id, user_id: user_id}}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(sendMessage)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"send", "message",
		"--message.brand-id", "brand_id",
		"--message.channels", "{foo: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [string], routing_method: all, timeouts: {channel: 0, provider: 0}}}",
		"--message.content", "{body: body, title: title}",
		"--message.context", "{tenant_id: tenant_id}",
		"--message.data", "{foo: bar}",
		"--message.delay", "{duration: 0, timezone: timezone, until: until}",
		"--message.expiry", "{expires_in: string, expires_at: expires_at}",
		"--message.metadata", "{event: event, tags: [string], trace_id: trace_id, utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}",
		"--message.preferences", "{subscription_topic_id: subscription_topic_id}",
		"--message.providers", "{foo: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, timeouts: 0}}",
		"--message.routing", "{channels: [string], method: all}",
		"--message.template", "template_id",
		"--message.timeout", "{channel: {foo: 0}, criteria: no-escalation, escalation: 0, message: 0, provider: {foo: 0}}",
		"--message.to", "{account_id: account_id, context: {tenant_id: tenant_id}, data: {foo: bar}, email: email, list_id: list_id, locale: locale, phone_number: phone_number, preferences: {notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, templateId: templateId}, tenant_id: tenant_id, user_id: user_id}",
	)
}
