// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestRoutingStrategiesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing-strategies", "create",
			"--name", "Email via SendGrid",
			"--routing", "{channels: [email], method: single}",
			"--channels", "{email: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [sendgrid, ses], routing_method: all, timeouts: {channel: 0, provider: 0}}}",
			"--description", "Routes email through sendgrid with SES failover",
			"--providers", "{sendgrid: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {}, timeouts: 0}}",
			"--tag", "[production, email]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(routingStrategiesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing-strategies", "create",
			"--name", "Email via SendGrid",
			"--routing.channels", "[email]",
			"--routing.method", "single",
			"--channels", "{email: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [sendgrid, ses], routing_method: all, timeouts: {channel: 0, provider: 0}}}",
			"--description", "Routes email through sendgrid with SES failover",
			"--providers", "{sendgrid: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {}, timeouts: 0}}",
			"--tag", "[production, email]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Email via SendGrid\n" +
			"routing:\n" +
			"  channels:\n" +
			"    - email\n" +
			"  method: single\n" +
			"channels:\n" +
			"  email:\n" +
			"    brand_id: brand_id\n" +
			"    if: if\n" +
			"    metadata:\n" +
			"      utm:\n" +
			"        campaign: campaign\n" +
			"        content: content\n" +
			"        medium: medium\n" +
			"        source: source\n" +
			"        term: term\n" +
			"    override:\n" +
			"      foo: bar\n" +
			"    providers:\n" +
			"      - sendgrid\n" +
			"      - ses\n" +
			"    routing_method: all\n" +
			"    timeouts:\n" +
			"      channel: 0\n" +
			"      provider: 0\n" +
			"description: Routes email through sendgrid with SES failover\n" +
			"providers:\n" +
			"  sendgrid:\n" +
			"    if: if\n" +
			"    metadata:\n" +
			"      utm:\n" +
			"        campaign: campaign\n" +
			"        content: content\n" +
			"        medium: medium\n" +
			"        source: source\n" +
			"        term: term\n" +
			"    override: {}\n" +
			"    timeouts: 0\n" +
			"tags:\n" +
			"  - production\n" +
			"  - email\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"routing-strategies", "create",
		)
	})
}

func TestRoutingStrategiesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing-strategies", "retrieve",
			"--id", "id",
		)
	})
}

func TestRoutingStrategiesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing-strategies", "list",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}

func TestRoutingStrategiesArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing-strategies", "archive",
			"--id", "id",
		)
	})
}

func TestRoutingStrategiesListNotifications(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing-strategies", "list-notifications",
			"--id", "id",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}

func TestRoutingStrategiesReplace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing-strategies", "replace",
			"--id", "id",
			"--name", "Email via SendGrid v2",
			"--routing", "{channels: [email], method: single}",
			"--channels", "{email: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [ses, sendgrid], routing_method: all, timeouts: {channel: 0, provider: 0}}}",
			"--description", "Updated routing with SES primary",
			"--providers", "{ses: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {}, timeouts: 0}}",
			"--tag", "[production, email, v2]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(routingStrategiesReplace)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing-strategies", "replace",
			"--id", "id",
			"--name", "Email via SendGrid v2",
			"--routing.channels", "[email]",
			"--routing.method", "single",
			"--channels", "{email: {brand_id: brand_id, if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {foo: bar}, providers: [ses, sendgrid], routing_method: all, timeouts: {channel: 0, provider: 0}}}",
			"--description", "Updated routing with SES primary",
			"--providers", "{ses: {if: if, metadata: {utm: {campaign: campaign, content: content, medium: medium, source: source, term: term}}, override: {}, timeouts: 0}}",
			"--tag", "[production, email, v2]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Email via SendGrid v2\n" +
			"routing:\n" +
			"  channels:\n" +
			"    - email\n" +
			"  method: single\n" +
			"channels:\n" +
			"  email:\n" +
			"    brand_id: brand_id\n" +
			"    if: if\n" +
			"    metadata:\n" +
			"      utm:\n" +
			"        campaign: campaign\n" +
			"        content: content\n" +
			"        medium: medium\n" +
			"        source: source\n" +
			"        term: term\n" +
			"    override:\n" +
			"      foo: bar\n" +
			"    providers:\n" +
			"      - ses\n" +
			"      - sendgrid\n" +
			"    routing_method: all\n" +
			"    timeouts:\n" +
			"      channel: 0\n" +
			"      provider: 0\n" +
			"description: Updated routing with SES primary\n" +
			"providers:\n" +
			"  ses:\n" +
			"    if: if\n" +
			"    metadata:\n" +
			"      utm:\n" +
			"        campaign: campaign\n" +
			"        content: content\n" +
			"        medium: medium\n" +
			"        source: source\n" +
			"        term: term\n" +
			"    override: {}\n" +
			"    timeouts: 0\n" +
			"tags:\n" +
			"  - production\n" +
			"  - email\n" +
			"  - v2\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"routing-strategies", "replace",
			"--id", "id",
		)
	})
}
