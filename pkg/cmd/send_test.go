// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestSendMessage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"send", "message",
			"--message", "{brand_id: brand_id, channels: {foo: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [string], routing_method: all, timeouts: {channel: 0, provider: 0}}}, content: {body: body, title: title}, context: {tenant_id: tenant_id}, data: {foo: bar}, delay: {duration: 0, timezone: timezone, until: until}, expiry: {expires_in: string, expires_at: expires_at}, metadata: {event: event, tags: [string], trace_id: trace_id, utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, preferences: {subscription_topic_id: subscription_topic_id}, providers: {foo: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, timeouts: 0}}, routing: {channels: [string], method: all}, template: template_id, timeout: {channel: {foo: 0}, criteria: no-escalation, escalation: 0, message: 0, provider: {foo: 0}}, to: {account_id: account_id, context: {tenant_id: tenant_id}, data: {foo: bar}, email: email, list_id: list_id, locale: locale, phone_number: phone_number, preferences: {notifications: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, categories: {foo: {status: OPTED_IN, channel_preferences: [{channel: direct_message}], rules: [{until: until, start: start}], source: subscription}}, templateId: templateId}, tenant_id: tenant_id, user_id: user_id}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(sendMessage)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"message:\n" +
			"  brand_id: brand_id\n" +
			"  channels:\n" +
			"    foo:\n" +
			"      brand_id: brand_id\n" +
			"      if: if\n" +
			"      metadata:\n" +
			"        utm:\n" +
			"          campaign: campaign\n" +
			"          content: content\n" +
			"          medium: medium\n" +
			"          source: source\n" +
			"          term: term\n" +
			"      override:\n" +
			"        foo: bar\n" +
			"      providers:\n" +
			"        - string\n" +
			"      routing_method: all\n" +
			"      timeouts:\n" +
			"        channel: 0\n" +
			"        provider: 0\n" +
			"  content:\n" +
			"    body: body\n" +
			"    title: title\n" +
			"  context:\n" +
			"    tenant_id: tenant_id\n" +
			"  data:\n" +
			"    foo: bar\n" +
			"  delay:\n" +
			"    duration: 0\n" +
			"    timezone: timezone\n" +
			"    until: until\n" +
			"  expiry:\n" +
			"    expires_in: string\n" +
			"    expires_at: expires_at\n" +
			"  metadata:\n" +
			"    event: event\n" +
			"    tags:\n" +
			"      - string\n" +
			"    trace_id: trace_id\n" +
			"    utm:\n" +
			"      campaign: campaign\n" +
			"      content: content\n" +
			"      medium: medium\n" +
			"      source: source\n" +
			"      term: term\n" +
			"  preferences:\n" +
			"    subscription_topic_id: subscription_topic_id\n" +
			"  providers:\n" +
			"    foo:\n" +
			"      if: if\n" +
			"      metadata:\n" +
			"        utm:\n" +
			"          campaign: campaign\n" +
			"          content: content\n" +
			"          medium: medium\n" +
			"          source: source\n" +
			"          term: term\n" +
			"      override:\n" +
			"        foo: bar\n" +
			"      timeouts: 0\n" +
			"  routing:\n" +
			"    channels:\n" +
			"      - string\n" +
			"    method: all\n" +
			"  template: template_id\n" +
			"  timeout:\n" +
			"    channel:\n" +
			"      foo: 0\n" +
			"    criteria: no-escalation\n" +
			"    escalation: 0\n" +
			"    message: 0\n" +
			"    provider:\n" +
			"      foo: 0\n" +
			"  to:\n" +
			"    account_id: account_id\n" +
			"    context:\n" +
			"      tenant_id: tenant_id\n" +
			"    data:\n" +
			"      foo: bar\n" +
			"    email: email\n" +
			"    list_id: list_id\n" +
			"    locale: locale\n" +
			"    phone_number: phone_number\n" +
			"    preferences:\n" +
			"      notifications:\n" +
			"        foo:\n" +
			"          status: OPTED_IN\n" +
			"          channel_preferences:\n" +
			"            - channel: direct_message\n" +
			"          rules:\n" +
			"            - until: until\n" +
			"              start: start\n" +
			"          source: subscription\n" +
			"      categories:\n" +
			"        foo:\n" +
			"          status: OPTED_IN\n" +
			"          channel_preferences:\n" +
			"            - channel: direct_message\n" +
			"          rules:\n" +
			"            - until: until\n" +
			"              start: start\n" +
			"          source: subscription\n" +
			"      templateId: templateId\n" +
			"    tenant_id: tenant_id\n" +
			"    user_id: user_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"send", "message",
		)
	})
}
