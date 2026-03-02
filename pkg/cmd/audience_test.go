// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/trycourier/courier-cli/v3/internal/mocktest"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
)

func TestAudiencesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audiences", "retrieve",
		"--api-key", "string",
		"--audience-id", "audience_id",
	)
}

func TestAudiencesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audiences", "update",
		"--api-key", "string",
		"--audience-id", "audience_id",
		"--description", "description",
		"--filter", "{filters: [{operator: operator, filters: [], path: path, value: value}]}",
		"--name", "name",
		"--operator", "AND",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(audiencesUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"audiences", "update",
		"--audience-id", "audience_id",
		"--description", "description",
		"--filter.filters", "[{operator: operator, filters: [], path: path, value: value}]",
		"--name", "name",
		"--operator", "AND",
	)
}

func TestAudiencesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audiences", "list",
		"--api-key", "string",
		"--cursor", "cursor",
	)
}

func TestAudiencesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audiences", "delete",
		"--api-key", "string",
		"--audience-id", "audience_id",
	)
}

func TestAudiencesListMembers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"audiences", "list-members",
		"--api-key", "string",
		"--audience-id", "audience_id",
		"--cursor", "cursor",
	)
}
