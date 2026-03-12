// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestUsersTokensRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:tokens", "retrieve",
			"--api-key", "string",
			"--user-id", "user_id",
			"--token", "token",
		)
	})
}

func TestUsersTokensUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:tokens", "update",
			"--api-key", "string",
			"--user-id", "user_id",
			"--token", "token",
			"--patch", "{op: op, path: path, value: value}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(usersTokensUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "users:tokens", "update",
			"--api-key", "string",
			"--user-id", "user_id",
			"--token", "token",
			"--patch.op", "op",
			"--patch.path", "path",
			"--patch.value", "value",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"patch:\n" +
			"  - op: op\n" +
			"    path: path\n" +
			"    value: value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "users:tokens", "update",
			"--api-key", "string",
			"--user-id", "user_id",
			"--token", "token",
		)
	})
}

func TestUsersTokensList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:tokens", "list",
			"--api-key", "string",
			"--user-id", "user_id",
		)
	})
}

func TestUsersTokensDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:tokens", "delete",
			"--api-key", "string",
			"--user-id", "user_id",
			"--token", "token",
		)
	})
}

func TestUsersTokensAddMultiple(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:tokens", "add-multiple",
			"--api-key", "string",
			"--user-id", "user_id",
		)
	})
}

func TestUsersTokensAddSingle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:tokens", "add-single",
			"--api-key", "string",
			"--user-id", "user_id",
			"--token", "token",
			"--provider-key", "firebase-fcm",
			"--device", "{ad_id: ad_id, app_id: app_id, device_id: device_id, manufacturer: manufacturer, model: model, platform: platform}",
			"--expiry-date", "string",
			"--properties", "{}",
			"--tracking", "{ip: ip, lat: lat, long: long, os_version: os_version}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(usersTokensAddSingle)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "users:tokens", "add-single",
			"--api-key", "string",
			"--user-id", "user_id",
			"--token", "token",
			"--provider-key", "firebase-fcm",
			"--device.ad-id", "ad_id",
			"--device.app-id", "app_id",
			"--device.device-id", "device_id",
			"--device.manufacturer", "manufacturer",
			"--device.model", "model",
			"--device.platform", "platform",
			"--expiry-date", "string",
			"--properties", "{}",
			"--tracking.ip", "ip",
			"--tracking.lat", "lat",
			"--tracking.long", "long",
			"--tracking.os-version", "os_version",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"provider_key: firebase-fcm\n" +
			"device:\n" +
			"  ad_id: ad_id\n" +
			"  app_id: app_id\n" +
			"  device_id: device_id\n" +
			"  manufacturer: manufacturer\n" +
			"  model: model\n" +
			"  platform: platform\n" +
			"expiry_date: string\n" +
			"properties: {}\n" +
			"tracking:\n" +
			"  ip: ip\n" +
			"  lat: lat\n" +
			"  long: long\n" +
			"  os_version: os_version\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "users:tokens", "add-single",
			"--api-key", "string",
			"--user-id", "user_id",
			"--token", "token",
		)
	})
}
