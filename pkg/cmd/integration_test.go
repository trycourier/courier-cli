package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func apiKey(t *testing.T) string {
	t.Helper()
	key := os.Getenv("COURIER_API_KEY")
	if key == "" {
		t.Skip("COURIER_API_KEY not set; skipping integration test")
	}
	return key
}

func runCLI(t *testing.T, flags ...string) string {
	t.Helper()
	key := apiKey(t)

	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok, "Could not get caller file path")
	project := filepath.Join(filepath.Dir(filename), "..", "..", "cmd", "...")

	args := []string{"run", project, "--format", "raw"}
	args = append(args, flags...)

	hasKey := false
	for _, f := range flags {
		if f == "--api-key" {
			hasKey = true
			break
		}
	}
	if !hasKey {
		args = append(args, "--api-key", key)
	}

	t.Logf("courier %s", strings.Join(args[4:], " "))

	cmd := exec.Command("go", args...)
	var out strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	require.NoError(t, err, "CLI command failed:\n%s", out.String())
	return out.String()
}

// ---------------------------------------------------------------------------
// Read-only list operations (safe against any workspace)
// ---------------------------------------------------------------------------

func TestIntegration_BrandsList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	output := runCLI(t, "brands", "list")
	assert.Contains(t, output, "{")
}

func TestIntegration_MessagesList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	output := runCLI(t, "messages", "list")
	assert.Contains(t, output, "{")
}

func TestIntegration_AudiencesList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	output := runCLI(t, "audiences", "list")
	assert.Contains(t, output, "{")
}

func TestIntegration_NotificationsList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	output := runCLI(t, "notifications", "list")
	assert.Contains(t, output, "{")
}

func TestIntegration_TenantsList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	output := runCLI(t, "tenants", "list")
	assert.Contains(t, output, "{")
}

func TestIntegration_ListsList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	output := runCLI(t, "lists", "list")
	assert.Contains(t, output, "{")
}

func TestIntegration_AuditEventsList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	output := runCLI(t, "audit-events", "list")
	assert.Contains(t, output, "{")
}

// ---------------------------------------------------------------------------
// Brand CRUD cycle: create → retrieve → update → delete
// ---------------------------------------------------------------------------

func TestIntegration_BrandCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	key := apiKey(t)
	_ = key

	brandID := fmt.Sprintf("cli-test-%d", time.Now().UnixMilli())
	brandName := "CLI Integration Test Brand"

	// Create (settings is required by the API)
	out := runCLI(t, "brands", "create", "--name", brandName, "--id", brandID,
		"--settings", `{colors: {primary: "#9D3789", secondary: "#9D3789"}}`)
	assert.Contains(t, out, brandID)

	// Retrieve
	out = runCLI(t, "brands", "retrieve", "--brand-id", brandID)
	assert.Contains(t, out, brandName)

	// Update (settings also required on update)
	updatedName := "CLI Integration Test Brand Updated"
	out = runCLI(t, "brands", "update", "--brand-id", brandID, "--name", updatedName,
		"--settings", `{colors: {primary: "#9D3789", secondary: "#9D3789"}}`)
	assert.Contains(t, out, updatedName)

	// Delete
	runCLI(t, "brands", "delete", "--brand-id", brandID)
}

// ---------------------------------------------------------------------------
// List (resource) CRUD cycle: update (upsert) → retrieve → delete
// ---------------------------------------------------------------------------

func TestIntegration_ListCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	listID := fmt.Sprintf("cli-test-%d", time.Now().UnixMilli())
	listName := "CLI Integration Test List"

	// Upsert
	out := runCLI(t, "lists", "update", "--list-id", listID, "--name", listName)
	_ = out

	// Retrieve
	out = runCLI(t, "lists", "retrieve", "--list-id", listID)
	assert.Contains(t, out, listName)

	// Delete
	runCLI(t, "lists", "delete", "--list-id", listID)
}

// ---------------------------------------------------------------------------
// Profile create/merge → retrieve
// ---------------------------------------------------------------------------

func TestIntegration_ProfileCreateAndRetrieve(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	userID := fmt.Sprintf("cli-test-user-%d", time.Now().UnixMilli())
	profile := `{email: "test@example.com"}`

	// Create/merge (profiles create uses --profile)
	out := runCLI(t, "profiles", "create", "--user-id", userID, "--profile", profile)
	_ = out

	// Retrieve
	out = runCLI(t, "profiles", "retrieve", "--user-id", userID)
	assert.Contains(t, out, "test@example.com")

	// Cleanup: delete
	runCLI(t, "profiles", "delete", "--user-id", userID)
}

// ---------------------------------------------------------------------------
// Send a message (enqueues even without configured channels)
// ---------------------------------------------------------------------------

func TestIntegration_SendMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	userID := fmt.Sprintf("cli-test-send-%d", time.Now().UnixMilli())
	msg := fmt.Sprintf(`{to: {user_id: %q}, content: {title: "Test", body: "Integration test message"}}`, userID)

	out := runCLI(t, "send", "message", "--message", msg)
	assert.Contains(t, out, "requestId")
}

// ---------------------------------------------------------------------------
// Auth: issue token
// ---------------------------------------------------------------------------

func TestIntegration_AuthIssueToken(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	userID := fmt.Sprintf("cli-test-auth-%d", time.Now().UnixMilli())

	out := runCLI(t, "auth", "issue-token",
		"--scope", fmt.Sprintf("user_id:%s read:preferences", userID),
		"--expires-in", "1 day",
	)
	assert.Contains(t, out, "token")
}

// ---------------------------------------------------------------------------
// Error path: retrieve nonexistent resource should fail
// ---------------------------------------------------------------------------

func TestIntegration_RetrieveNonexistentBrand(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	key := apiKey(t)

	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok)
	project := filepath.Join(filepath.Dir(filename), "..", "..", "cmd", "...")

	cmd := exec.Command("go", "run", project, "--format", "raw",
		"brands", "retrieve",
		"--brand-id", "nonexistent-brand-id-xyz-12345",
		"--api-key", key,
	)
	var out strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	assert.Error(t, err, "retrieving a nonexistent brand should return an error")
}

// ---------------------------------------------------------------------------
// Tenant CRUD cycle: update (upsert) → retrieve → delete
// ---------------------------------------------------------------------------

func TestIntegration_TenantCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	tenantID := fmt.Sprintf("cli-test-tenant-%d", time.Now().UnixMilli())
	tenantName := "CLI Integration Test Tenant"

	// Create/upsert
	out := runCLI(t, "tenants", "update", "--tenant-id", tenantID, "--name", tenantName)
	_ = out

	// Retrieve
	out = runCLI(t, "tenants", "retrieve", "--tenant-id", tenantID)
	assert.Contains(t, out, tenantName)

	// Delete
	runCLI(t, "tenants", "delete", "--tenant-id", tenantID)
}

func TestIntegration_RetrieveNonexistentList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	key := apiKey(t)

	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok)
	project := filepath.Join(filepath.Dir(filename), "..", "..", "cmd", "...")

	cmd := exec.Command("go", "run", project, "--format", "raw",
		"lists", "retrieve",
		"--list-id", "nonexistent-list-id-xyz-12345",
		"--api-key", key,
	)
	var out strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	assert.Error(t, err, "retrieving a nonexistent list should return an error")
}
