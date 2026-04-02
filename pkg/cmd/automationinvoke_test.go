// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestAutomationsInvokeInvokeAdHoc(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"automations:invoke", "invoke-ad-hoc",
			"--automation", "{steps: [{action: delay, duration: duration, until: 20240408T080910.123}, {action: send, brand: brand, data: {foo: bar}, profile: {foo: bar}, recipient: recipient, template: 64TP5HKPFTM8VTK1Y75SJDQX9JK0}], cancelation_token: delay-send--user-yes--abc-123}",
			"--brand", "brand",
			"--data", "{name: bar}",
			"--profile", "{tenant_id: bar}",
			"--recipient", "user-yes",
			"--template", "template",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(automationsInvokeInvokeAdHoc)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"automations:invoke", "invoke-ad-hoc",
			"--automation.steps", "[{action: delay, duration: duration, until: 20240408T080910.123}, {action: send, brand: brand, data: {foo: bar}, profile: {foo: bar}, recipient: recipient, template: 64TP5HKPFTM8VTK1Y75SJDQX9JK0}]",
			"--automation.cancelation-token", "delay-send--user-yes--abc-123",
			"--brand", "brand",
			"--data", "{name: bar}",
			"--profile", "{tenant_id: bar}",
			"--recipient", "user-yes",
			"--template", "template",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"automation:\n" +
			"  steps:\n" +
			"    - action: delay\n" +
			"      duration: duration\n" +
			"      until: 20240408T080910.123\n" +
			"    - action: send\n" +
			"      brand: brand\n" +
			"      data:\n" +
			"        foo: bar\n" +
			"      profile:\n" +
			"        foo: bar\n" +
			"      recipient: recipient\n" +
			"      template: 64TP5HKPFTM8VTK1Y75SJDQX9JK0\n" +
			"  cancelation_token: delay-send--user-yes--abc-123\n" +
			"brand: brand\n" +
			"data:\n" +
			"  name: bar\n" +
			"profile:\n" +
			"  tenant_id: bar\n" +
			"recipient: user-yes\n" +
			"template: template\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"automations:invoke", "invoke-ad-hoc",
		)
	})
}

func TestAutomationsInvokeInvokeByTemplate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"automations:invoke", "invoke-by-template",
			"--template-id", "templateId",
			"--recipient", "recipient",
			"--brand", "brand",
			"--data", "{foo: bar}",
			"--profile", "{foo: bar}",
			"--template", "template",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"recipient: recipient\n" +
			"brand: brand\n" +
			"data:\n" +
			"  foo: bar\n" +
			"profile:\n" +
			"  foo: bar\n" +
			"template: template\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"automations:invoke", "invoke-by-template",
			"--template-id", "templateId",
		)
	})
}
