// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestTenantsTemplatesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "tenants:templates", "retrieve",
			"--api-key", "string",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
		)
	})
}

func TestTenantsTemplatesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "tenants:templates", "list",
			"--api-key", "string",
			"--tenant-id", "tenant_id",
			"--cursor", "cursor",
			"--limit", "0",
		)
	})
}

func TestTenantsTemplatesPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "tenants:templates", "publish",
			"--api-key", "string",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
			"--version", "version",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("version: version")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "tenants:templates", "publish",
			"--api-key", "string",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
		)
	})
}

func TestTenantsTemplatesReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "tenants:templates", "replace",
			"--api-key", "string",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
			"--template", "{content: {elements: [{channels: [string], if: if, loop: loop, ref: ref, type: text}], version: version}, channels: {foo: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [string], routing_method: all, timeouts: {channel: 0, provider: 0}}}, providers: {foo: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, timeouts: 0}}, routing: {channels: [string], method: all}}",
			"--published=true",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(tenantsTemplatesReplace)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "tenants:templates", "replace",
			"--api-key", "string",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
			"--template.content", "{elements: [{channels: [string], if: if, loop: loop, ref: ref, type: text}], version: version}",
			"--template.channels", "{foo: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [string], routing_method: all, timeouts: {channel: 0, provider: 0}}}",
			"--template.providers", "{foo: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, timeouts: 0}}",
			"--template.routing", "{channels: [string], method: all}",
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
			"        type: text\n" +
			"    version: version\n" +
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
			"      - string\n" +
			"    method: all\n" +
			"published: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "tenants:templates", "replace",
			"--api-key", "string",
			"--tenant-id", "tenant_id",
			"--template-id", "template_id",
		)
	})
}
