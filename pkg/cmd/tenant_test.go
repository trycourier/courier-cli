// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/internal/mocktest"
	"github.com/trycourier/courier-cli/internal/requestflag"
)

func TestTenantsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants", "retrieve",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
	)
}

func TestTenantsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants", "update",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
		"--name", "name",
		"--brand-id", "brand_id",
		"--default-preferences", "{items: [{status: OPTED_OUT, custom_routing: [direct_message], has_custom_routing: true, id: id}]}",
		"--parent-tenant-id", "parent_tenant_id",
		"--properties", "{foo: bar}",
		"--user-profile", "{foo: bar}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(tenantsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants", "update",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
		"--name", "name",
		"--brand-id", "brand_id",
		"--default-preferences.items", "[{status: OPTED_OUT, custom_routing: [direct_message], has_custom_routing: true, id: id}]",
		"--parent-tenant-id", "parent_tenant_id",
		"--properties", "{foo: bar}",
		"--user-profile", "{foo: bar}",
	)
}

func TestTenantsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants", "list",
		"--api-key", "string",
		"--cursor", "cursor",
		"--limit", "0",
		"--parent-tenant-id", "parent_tenant_id",
	)
}

func TestTenantsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants", "delete",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
	)
}

func TestTenantsListUsers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"tenants", "list-users",
		"--api-key", "string",
		"--tenant-id", "tenant_id",
		"--cursor", "cursor",
		"--limit", "0",
	)
}
