// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v5/internal/mocktest"
)

func TestNotificationsPreviewsRunsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications:previews:runs", "create",
			"--id", "id",
			"--data", "{foo: bar}",
			"--device-id", "pvd_1w6dgafr3aaycvv9a8bm996pkc",
			"--device-id", "pvd_34qvmj6p4dbqaa5mpys1ekt9jx",
			"--device-set-id", "device_set_id",
			"--locale", "locale",
			"--template-version", "draft",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data:\n" +
			"  foo: bar\n" +
			"device_ids:\n" +
			"  - pvd_1w6dgafr3aaycvv9a8bm996pkc\n" +
			"  - pvd_34qvmj6p4dbqaa5mpys1ekt9jx\n" +
			"device_set_id: device_set_id\n" +
			"locale: locale\n" +
			"template_version: draft\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"notifications:previews:runs", "create",
			"--id", "id",
			"--idempotency-key", "order-ORD-456-user-123",
			"--x-idempotency-expiration", "1785312000",
		)
	})
}

func TestNotificationsPreviewsRunsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications:previews:runs", "retrieve",
			"--id", "id",
			"--preview-run-id", "previewRunId",
		)
	})
}

func TestNotificationsPreviewsRunsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"notifications:previews:runs", "list",
			"--id", "id",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}
