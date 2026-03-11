// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
	"github.com/trycourier/courier-cli/internal/requestflag"
)

func TestTenantsTemplatesRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants:templates", "retrieve",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
		"--template-id", "template_id",
	)
}

func TestTenantsTemplatesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants:templates", "list",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
		"--cursor", "cursor",
		"--limit", "0",
	)
}

func TestTenantsTemplatesPublish(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants:templates", "publish",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
		"--template-id", "template_id",
		"--version", "version",
	)
}

func TestTenantsTemplatesReplace(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants:templates", "replace",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
		"--template-id", "template_id",
		"--template", "{content: {elements: [{channels: [string], if: if, loop: loop, ref: ref, type: text}], version: version}, channels: {foo: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [string], routing_method: all, timeouts: {channel: 0, provider: 0}}}, providers: {foo: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, timeouts: 0}}, routing: {channels: [string], method: all}}",
		"--published=true",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(tenantsTemplatesReplace)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants:templates", "replace",
		"--tenant-id", "tenant_id",
		"--template-id", "template_id",
		"--template.content", "{elements: [{channels: [string], if: if, loop: loop, ref: ref, type: text}], version: version}",
		"--template.channels", "{foo: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [string], routing_method: all, timeouts: {channel: 0, provider: 0}}}",
		"--template.providers", "{foo: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, timeouts: 0}}",
		"--template.routing", "{channels: [string], method: all}",
		"--published=true",
	)
}
