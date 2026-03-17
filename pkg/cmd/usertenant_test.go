// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestUsersTenantsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:tenants", "list",
			"--user-id", "user_id",
			"--cursor", "cursor",
			"--limit", "0",
		)
	})
}

func TestUsersTenantsAddMultiple(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:tenants", "add-multiple",
			"--user-id", "user_id",
			"--tenant", "{tenant_id: tenant_id, profile: {foo: bar}, type: user, user_id: user_id}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(usersTenantsAddMultiple)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:tenants", "add-multiple",
			"--user-id", "user_id",
			"--tenant.tenant-id", "tenant_id",
			"--tenant.profile", "{foo: bar}",
			"--tenant.type", "user",
			"--tenant.user-id", "user_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"tenants:\n" +
			"  - tenant_id: tenant_id\n" +
			"    profile:\n" +
			"      foo: bar\n" +
			"    type: user\n" +
			"    user_id: user_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"users:tenants", "add-multiple",
			"--user-id", "user_id",
		)
	})
}

func TestUsersTenantsAddSingle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:tenants", "add-single",
			"--user-id", "user_id",
			"--tenant-id", "tenant_id",
			"--profile", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"profile:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"users:tenants", "add-single",
			"--user-id", "user_id",
			"--tenant-id", "tenant_id",
		)
	})
}

func TestUsersTenantsRemoveAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:tenants", "remove-all",
			"--user-id", "user_id",
		)
	})
}

func TestUsersTenantsRemoveSingle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:tenants", "remove-single",
			"--user-id", "user_id",
			"--tenant-id", "tenant_id",
		)
	})
}
