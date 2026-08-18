// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v5/internal/mocktest"
	"github.com/trycourier/courier-cli/v5/internal/requestflag"
)

func TestAudiencesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"audiences", "retrieve",
			"--audience-id", "audience_id",
		)
	})
}

func TestAudiencesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"audiences", "update",
			"--audience-id", "audience_id",
			"--description", "Users located in the US",
			"--filter", "{filters: [{operator: EQ, filters: [], path: profile.location, value: US}], operator: AND}",
			"--name", "Engaged US Users",
			"--operator", "AND",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(audiencesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"audiences", "update",
			"--audience-id", "audience_id",
			"--description", "Users located in the US",
			"--filter.filters", "[{operator: EQ, filters: [], path: profile.location, value: US}]",
			"--filter.operator", "AND",
			"--name", "Engaged US Users",
			"--operator", "AND",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: Users located in the US\n" +
			"filter:\n" +
			"  filters:\n" +
			"    - operator: EQ\n" +
			"      filters: []\n" +
			"      path: profile.location\n" +
			"      value: US\n" +
			"  operator: AND\n" +
			"name: Engaged US Users\n" +
			"operator: AND\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"audiences", "update",
			"--audience-id", "audience_id",
		)
	})
}

func TestAudiencesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"audiences", "list",
			"--cursor", "cursor",
		)
	})
}

func TestAudiencesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"audiences", "delete",
			"--audience-id", "audience_id",
		)
	})
}

func TestAudiencesListMembers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"audiences", "list-members",
			"--audience-id", "audience_id",
			"--cursor", "cursor",
		)
	})
}
