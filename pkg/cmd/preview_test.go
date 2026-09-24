// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v5/internal/mocktest"
)

func TestPreviewsArchiveDeviceSet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"previews", "archive-device-set",
			"--device-set-id", "deviceSetId",
		)
	})
}

func TestPreviewsCreateDeviceSet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"previews", "create-device-set",
			"--device-id", "pvd_1w6dgafr3aaycvv9a8bm996pkc",
			"--name", "Mobile",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"device_ids:\n" +
			"  - pvd_1w6dgafr3aaycvv9a8bm996pkc\n" +
			"name: Mobile\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"previews", "create-device-set",
		)
	})
}

func TestPreviewsListDeviceSets(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"previews", "list-device-sets",
		)
	})
}

func TestPreviewsListDevices(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"previews", "list-devices",
		)
	})
}

func TestPreviewsRetrieveDeviceSet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"previews", "retrieve-device-set",
			"--device-set-id", "deviceSetId",
		)
	})
}

func TestPreviewsUpdateDeviceSet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"previews", "update-device-set",
			"--device-set-id", "deviceSetId",
			"--device-id", "pvd_1w6dgafr3aaycvv9a8bm996pkc",
			"--device-id", "pvd_34qvmj6p4dbqaa5mpys1ekt9jx",
			"--name", "Mobile and desktop",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"device_ids:\n" +
			"  - pvd_1w6dgafr3aaycvv9a8bm996pkc\n" +
			"  - pvd_34qvmj6p4dbqaa5mpys1ekt9jx\n" +
			"name: Mobile and desktop\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"previews", "update-device-set",
			"--device-set-id", "deviceSetId",
		)
	})
}
