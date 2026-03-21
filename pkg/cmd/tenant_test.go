// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestTenantsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants", "retrieve",
			"--tenant-id", "tenant_id",
		)
	})
}

func TestTenantsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants", "update",
			"--tenant-id", "tenant_id",
			"--name", "name",
			"--brand-id", "brand_id",
			"--default-preferences", "{items: [{status: OPTED_OUT, custom_routing: [direct_message], has_custom_routing: true, id: id}]}",
			"--parent-tenant-id", "parent_tenant_id",
			"--properties", "{foo: bar}",
			"--user-profile", "{foo: bar}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(tenantsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants", "update",
			"--tenant-id", "tenant_id",
			"--name", "name",
			"--brand-id", "brand_id",
			"--default-preferences.items", "[{status: OPTED_OUT, custom_routing: [direct_message], has_custom_routing: true, id: id}]",
			"--parent-tenant-id", "parent_tenant_id",
			"--properties", "{foo: bar}",
			"--user-profile", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"brand_id: brand_id\n" +
			"default_preferences:\n" +
			"  items:\n" +
			"    - status: OPTED_OUT\n" +
			"      custom_routing:\n" +
			"        - direct_message\n" +
			"      has_custom_routing: true\n" +
			"      id: id\n" +
			"parent_tenant_id: parent_tenant_id\n" +
			"properties:\n" +
			"  foo: bar\n" +
			"user_profile:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tenants", "update",
			"--tenant-id", "tenant_id",
		)
	})
}

func TestTenantsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants", "list",
			"--cursor", "cursor",
			"--limit", "0",
			"--parent-tenant-id", "parent_tenant_id",
		)
	})
}

func TestTenantsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants", "delete",
			"--tenant-id", "tenant_id",
		)
	})
}

func TestTenantsListUsers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tenants", "list-users",
			"--tenant-id", "tenant_id",
			"--cursor", "cursor",
			"--limit", "0",
		)
	})
}
