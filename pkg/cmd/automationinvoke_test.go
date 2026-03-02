// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
	"github.com/trycourier/courier-cli/internal/requestflag"
)

func TestAutomationsInvokeInvokeAdHoc(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"automations:invoke", "invoke-ad-hoc",
		"--api-key", "string",
		"--automation", "{steps: [{action: delay, duration: duration, until: 20240408T080910.123}, {action: send, brand: brand, data: {foo: bar}, profile: {foo: bar}, recipient: recipient, template: 64TP5HKPFTM8VTK1Y75SJDQX9JK0}], cancelation_token: delay-send--user-yes--abc-123}",
		"--brand", "brand",
		"--data", "{name: bar}",
		"--profile", "{tenant_id: bar}",
		"--recipient", "user-yes",
		"--template", "template",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(automationsInvokeInvokeAdHoc)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"automations:invoke", "invoke-ad-hoc",
		"--automation.steps", "[{action: delay, duration: duration, until: 20240408T080910.123}, {action: send, brand: brand, data: {foo: bar}, profile: {foo: bar}, recipient: recipient, template: 64TP5HKPFTM8VTK1Y75SJDQX9JK0}]",
		"--automation.cancelation-token", "delay-send--user-yes--abc-123",
		"--brand", "brand",
		"--data", "{name: bar}",
		"--profile", "{tenant_id: bar}",
		"--recipient", "user-yes",
		"--template", "template",
	)
}

func TestAutomationsInvokeInvokeByTemplate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"automations:invoke", "invoke-by-template",
		"--api-key", "string",
		"--template-id", "templateId",
		"--recipient", "recipient",
		"--brand", "brand",
		"--data", "{foo: bar}",
		"--profile", "{foo: bar}",
		"--template", "template",
	)
}
