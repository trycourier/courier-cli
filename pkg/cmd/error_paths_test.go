package cmd

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func cliProject(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok)
	return filepath.Join(filepath.Dir(filename), "..", "..", "cmd", "...")
}

func runCLIExpectError(t *testing.T, flags ...string) string {
	t.Helper()
	project := cliProject(t)
	args := []string{"run", project}
	args = append(args, flags...)
	cmd := exec.Command("go", args...)
	var out strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	assert.Error(t, err, "expected CLI to return an error")
	return out.String()
}

// ---------------------------------------------------------------------------
// Missing required flags
// ---------------------------------------------------------------------------

func TestError_MissingRequiredFlag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		command     string
		subcommand  string
		expectedMsg string
	}{
		{"audience", "audiences", "retrieve", "audience-id"},
		{"brand", "brands", "retrieve", "brand-id"},
		{"list", "lists", "retrieve", "list-id"},
		{"message", "messages", "retrieve", "message-id"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			output := runCLIExpectError(t,
				tt.command, tt.subcommand,
				"--api-key", "test-key",
				"--format", "raw",
			)
			assert.Contains(t, output, tt.expectedMsg)
		})
	}
}

// ---------------------------------------------------------------------------
// Invalid format flag
// ---------------------------------------------------------------------------

func TestError_InvalidFormat(t *testing.T) {
	t.Parallel()
	output := runCLIExpectError(t,
		"--format", "invalid-format",
		"brands", "list",
		"--api-key", "test-key",
	)
	assert.Contains(t, output, "invalid-format")
}

// ---------------------------------------------------------------------------
// Extra unexpected arguments
// ---------------------------------------------------------------------------

func TestError_ExtraArguments(t *testing.T) {
	t.Parallel()
	output := runCLIExpectError(t,
		"audiences", "retrieve",
		"--api-key", "test-key",
		"--audience-id", "test-id",
		"--format", "raw",
		"--base-url", "http://localhost:1",
		"extra-arg-1", "extra-arg-2",
	)
	assert.Contains(t, output, "Unexpected extra arguments")
}

// ---------------------------------------------------------------------------
// API error simulation with httptest
// ---------------------------------------------------------------------------

func TestError_API404(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"message":"not found","type":"error"}`)
	}))
	defer srv.Close()

	output := runCLIExpectError(t,
		"brands", "retrieve",
		"--brand-id", "nonexistent",
		"--api-key", "test-key",
		"--base-url", srv.URL,
		"--format", "raw",
	)
	assert.Contains(t, output, "not found")
}

func TestError_API401(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"message":"unauthorized","type":"error"}`)
	}))
	defer srv.Close()

	output := runCLIExpectError(t,
		"brands", "list",
		"--api-key", "bad-key",
		"--base-url", srv.URL,
		"--format", "raw",
	)
	assert.Contains(t, output, "unauthorized")
}

func TestError_API500(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message":"internal server error","type":"error"}`)
	}))
	defer srv.Close()

	output := runCLIExpectError(t,
		"audiences", "list",
		"--api-key", "test-key",
		"--base-url", srv.URL,
		"--format", "raw",
	)
	assert.Contains(t, output, "internal server error")
}

func TestError_ConnectionRefused(t *testing.T) {
	t.Parallel()

	output := runCLIExpectError(t,
		"brands", "list",
		"--api-key", "test-key",
		"--base-url", "http://localhost:1",
		"--format", "raw",
	)
	assert.NotEmpty(t, output, "connection refused should produce error output")
}

// ---------------------------------------------------------------------------
// Unknown subcommand
// ---------------------------------------------------------------------------

func TestError_UnknownSubcommand(t *testing.T) {
	t.Parallel()

	project := cliProject(t)
	args := []string{"run", project, "nonexistent-command"}

	cmd := exec.Command("go", args...)
	var out strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	assert.Error(t, err)
	output := out.String()
	assert.True(t,
		strings.Contains(output, "not found") || strings.Contains(output, "Did you mean"),
		"unknown command should mention 'not found' or suggest alternatives: %s", output,
	)
}

// ---------------------------------------------------------------------------
// Help flag doesn't error
// ---------------------------------------------------------------------------

func TestNoError_HelpFlag(t *testing.T) {
	t.Parallel()
	err := Command.Run(context.Background(), []string{"courier", "brands", "--help"})
	assert.NoError(t, err)
}

func TestNoError_SubcommandHelp(t *testing.T) {
	t.Parallel()
	err := Command.Run(context.Background(), []string{"courier", "audiences", "retrieve", "--help"})
	assert.NoError(t, err)
}

// ---------------------------------------------------------------------------
// API error with format-error flag
// ---------------------------------------------------------------------------

func TestError_FormatErrorFlag(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"message":"bad request","type":"invalid_request_error"}`)
	}))
	defer srv.Close()

	output := runCLIExpectError(t,
		"brands", "list",
		"--api-key", "test-key",
		"--base-url", srv.URL,
		"--format", "raw",
		"--format-error", "raw",
	)
	assert.Contains(t, output, "bad request")
}
