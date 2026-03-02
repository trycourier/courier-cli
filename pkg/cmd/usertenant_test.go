// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestUsersTenantsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:tenants", "list",
		"--api-key", "string",
		"--user-id", "user_id",
		"--cursor", "cursor",
		"--limit", "0",
	)
}

func TestUsersTenantsAddMultiple(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:tenants", "add-multiple",
		"--api-key", "string",
		"--user-id", "user_id",
		"--tenant", "{tenant_id: tenant_id, profile: {foo: bar}, type: user, user_id: user_id}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(usersTenantsAddMultiple)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:tenants", "add-multiple",
		"--user-id", "user_id",
		"--tenant.tenant-id", "tenant_id",
		"--tenant.profile", "{foo: bar}",
		"--tenant.type", "user",
		"--tenant.user-id", "user_id",
	)
}

func TestUsersTenantsAddSingle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:tenants", "add-single",
		"--api-key", "string",
		"--user-id", "user_id",
		"--tenant-id", "tenant_id",
		"--profile", "{foo: bar}",
	)
}

func TestUsersTenantsRemoveAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:tenants", "remove-all",
		"--api-key", "string",
		"--user-id", "user_id",
	)
}

func TestUsersTenantsRemoveSingle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"users:tenants", "remove-single",
		"--api-key", "string",
		"--user-id", "user_id",
		"--tenant-id", "tenant_id",
	)
}
