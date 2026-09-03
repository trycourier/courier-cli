// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v5/internal/mocktest"
	"github.com/trycourier/courier-cli/v5/internal/requestflag"
)

func TestTenantsTemplatesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants:templates", "retrieve",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
		)
	})
}

func TestTenantsTemplatesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants:templates", "list",
			"--tenant-id", "tenant_id",
			"--cursor", "cursor",
			"--limit", "0",
		)
	})
}

func TestTenantsTemplatesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants:templates", "delete",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
		)
	})
}

func TestTenantsTemplatesPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants:templates", "publish",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
			"--version", "latest",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("version: latest")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tenants:templates", "publish",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
		)
	})
}

func TestTenantsTemplatesReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants:templates", "replace",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
			"--template", "{content: {elements: [{channels: [string], if: if, loop: loop, ref: ref, channel: email, elements: [{channels: [string], if: if, loop: loop, ref: ref, align: left, bold: bold, color: color, content: 'Hello {{name}}', font_size: font_size, format: markdown, italic: italic, line_height: line_height, locales: {foo: {content: content}}, strikethrough: strikethrough, text_style: text, underline: underline, type: text}], font_size: font_size, line_height: line_height, padding: padding, raw: {foo: bar}, type: channel}], version: '2022-01-01'}, channels: {foo: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [string], routing_method: all, timeouts: {channel: 0, provider: 0}}}, providers: {foo: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, timeouts: 0}}, routing: {channels: [email], method: single}}",
			"--published=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(tenantsTemplatesReplace)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants:templates", "replace",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
			"--template.content", "{elements: [{channels: [string], if: if, loop: loop, ref: ref, channel: email, elements: [{channels: [string], if: if, loop: loop, ref: ref, align: left, bold: bold, color: color, content: 'Hello {{name}}', font_size: font_size, format: markdown, italic: italic, line_height: line_height, locales: {foo: {content: content}}, strikethrough: strikethrough, text_style: text, underline: underline, type: text}], font_size: font_size, line_height: line_height, padding: padding, raw: {foo: bar}, type: channel}], version: '2022-01-01'}",
			"--template.channels", "{foo: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [string], routing_method: all, timeouts: {channel: 0, provider: 0}}}",
			"--template.providers", "{foo: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, timeouts: 0}}",
			"--template.routing", "{channels: [email], method: single}",
			"--published=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"template:\n" +
			"  content:\n" +
			"    elements:\n" +
			"      - channels:\n" +
			"          - string\n" +
			"        if: if\n" +
			"        loop: loop\n" +
			"        ref: ref\n" +
			"        channel: email\n" +
			"        elements:\n" +
			"          - channels:\n" +
			"              - string\n" +
			"            if: if\n" +
			"            loop: loop\n" +
			"            ref: ref\n" +
			"            align: left\n" +
			"            bold: bold\n" +
			"            color: color\n" +
			"            content: Hello {{name}}\n" +
			"            font_size: font_size\n" +
			"            format: markdown\n" +
			"            italic: italic\n" +
			"            line_height: line_height\n" +
			"            locales:\n" +
			"              foo:\n" +
			"                content: content\n" +
			"            strikethrough: strikethrough\n" +
			"            text_style: text\n" +
			"            underline: underline\n" +
			"            type: text\n" +
			"        font_size: font_size\n" +
			"        line_height: line_height\n" +
			"        padding: padding\n" +
			"        raw:\n" +
			"          foo: bar\n" +
			"        type: channel\n" +
			"    version: '2022-01-01'\n" +
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
			"      - email\n" +
			"    method: single\n" +
			"published: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tenants:templates", "replace",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
		)
	})
}
