// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
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
			"--description", "description",
			"--filter", "{filters: [{operator: operator, filters: [], path: path, value: value}]}",
			"--name", "name",
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
			"--description", "description",
			"--filter.filters", "[{operator: operator, filters: [], path: path, value: value}]",
			"--name", "name",
			"--operator", "AND",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"description: description\n" +
			"filter:\n" +
			"  filters:\n" +
			"    - operator: operator\n" +
			"      filters: []\n" +
			"      path: path\n" +
			"      value: value\n" +
			"name: name\n" +
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
