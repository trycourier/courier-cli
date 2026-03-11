package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestFormatJSON_JSONFormat(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"name":"Alice","age":30}`)
	result, err := formatJSON(os.Stdout, "test", data, "json", "")
	require.NoError(t, err)
	assert.Contains(t, string(result), "name")
	assert.Contains(t, string(result), "Alice")
	assert.Contains(t, string(result), "age")
	assert.Contains(t, string(result), "30")
}

func TestFormatJSON_JSONLFormat(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"name":"Alice","age":30}`)
	// Use a non-terminal writer to avoid colors
	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()

	result, err := formatJSON(w, "test", data, "jsonl", "")
	require.NoError(t, err)

	output := string(result)
	assert.True(t, strings.HasSuffix(output, "\n"), "jsonl output should end with newline")
	// jsonl should be one line (no extra newlines except the trailing one)
	lines := strings.Split(strings.TrimSuffix(output, "\n"), "\n")
	assert.Len(t, lines, 1, "jsonl should be a single line")
	assert.Contains(t, output, "name")
	assert.Contains(t, output, "Alice")
}

func TestFormatJSON_RawFormat(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"name":"Alice"}`)
	result, err := formatJSON(os.Stdout, "test", data, "raw", "")
	require.NoError(t, err)
	assert.Equal(t, `{"name":"Alice"}`+"\n", string(result))
}

func TestFormatJSON_PrettyFormat(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"name":"Alice"}`)
	result, err := formatJSON(os.Stdout, "test", data, "pretty", "")
	require.NoError(t, err)
	output := string(result)
	assert.Contains(t, output, "name:")
	assert.Contains(t, output, "Alice")
}

func TestFormatJSON_YAMLFormat(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"name":"Alice","age":30}`)
	r, w, _ := os.Pipe()
	defer r.Close()

	_, err := formatJSON(w, "test", data, "yaml", "")
	w.Close()
	require.NoError(t, err)

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()
	assert.Contains(t, output, "name:")
	assert.Contains(t, output, "Alice")
	assert.Contains(t, output, "age:")
}

func TestFormatJSON_AutoFormat(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"key":"value"}`)
	resultAuto, err := formatJSON(os.Stdout, "test", data, "auto", "")
	require.NoError(t, err)
	resultJSON, err := formatJSON(os.Stdout, "test", data, "json", "")
	require.NoError(t, err)
	assert.Equal(t, resultJSON, resultAuto, "auto should produce same output as json")
}

func TestFormatJSON_InvalidFormat(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"key":"value"}`)
	_, err := formatJSON(os.Stdout, "test", data, "nonexistent", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid format")
}

func TestFormatJSON_CaseInsensitive(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"key":"value"}`)
	result1, err := formatJSON(os.Stdout, "test", data, "JSON", "")
	require.NoError(t, err)
	result2, err := formatJSON(os.Stdout, "test", data, "json", "")
	require.NoError(t, err)
	assert.Equal(t, result1, result2)
}

func TestFormatJSON_Transform(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"user":{"name":"Alice"},"count":5}`)

	t.Run("jsonl format applies transform", func(t *testing.T) {
		r, w, _ := os.Pipe()
		defer r.Close()
		result, err := formatJSON(w, "test", data, "jsonl", "user.name")
		require.NoError(t, err)
		assert.Equal(t, "\"Alice\"\n", string(result))
	})

	t.Run("raw format skips transform", func(t *testing.T) {
		result, err := formatJSON(os.Stdout, "test", data, "raw", "user.name")
		require.NoError(t, err)
		// raw format explicitly skips transform, so full data is returned
		assert.Contains(t, string(result), "user")
		assert.Contains(t, string(result), "count")
	})

	t.Run("non-existent transform keeps original", func(t *testing.T) {
		r, w, _ := os.Pipe()
		defer r.Close()
		result, err := formatJSON(w, "test", data, "jsonl", "nonexistent")
		require.NoError(t, err)
		// When transform target doesn't exist, original data is used
		assert.Contains(t, string(result), "user")
	})
}

func TestShowJSON_Transform(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"outer":{"inner":"deep"}}`)

	t.Run("raw format skips transform", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "showjson-raw-*")
		require.NoError(t, err)
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		err = ShowJSON(tmpFile, "test", data, "raw", "outer.inner")
		require.NoError(t, err)

		tmpFile.Seek(0, 0)
		content, err := os.ReadFile(tmpFile.Name())
		require.NoError(t, err)
		// raw format explicitly skips transforms
		assert.Contains(t, string(content), "outer")
	})

	t.Run("jsonl format applies transform", func(t *testing.T) {
		r, w, _ := os.Pipe()
		defer r.Close()

		err := ShowJSON(w, "test", data, "jsonl", "outer.inner")
		w.Close()
		require.NoError(t, err)

		var buf bytes.Buffer
		buf.ReadFrom(r)
		assert.Equal(t, "\"deep\"\n", buf.String(), "transform should extract nested value")
	})
}

func TestCountTerminalLines(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		data     string
		width    int
		expected int
	}{
		{"empty", "", 80, 0},
		{"single line", "hello", 80, 0},
		{"two lines", "line1\nline2", 80, 1},
		{"wrapped at width boundary", strings.Repeat("x", 160), 80, 1},
		{"multiple newlines", "a\nb\nc\nd", 80, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := countTerminalLines([]byte(tt.data), tt.width)
			assert.Equal(t, tt.expected, lines)
		})
	}
}

func TestShouldUseColors(t *testing.T) {
	t.Run("FORCE_COLOR=1", func(t *testing.T) {
		t.Setenv("FORCE_COLOR", "1")
		assert.True(t, shouldUseColors(os.Stdout))
	})

	t.Run("FORCE_COLOR=0", func(t *testing.T) {
		t.Setenv("FORCE_COLOR", "0")
		assert.False(t, shouldUseColors(os.Stdout))
	})

	t.Run("non-terminal writer", func(t *testing.T) {
		t.Setenv("FORCE_COLOR", "")
		os.Unsetenv("FORCE_COLOR")
		var buf bytes.Buffer
		assert.False(t, shouldUseColors(&buf))
	})
}

func TestIsTerminal(t *testing.T) {
	t.Parallel()

	t.Run("pipe is not terminal", func(t *testing.T) {
		r, w, _ := os.Pipe()
		defer r.Close()
		defer w.Close()
		assert.False(t, isTerminal(w))
	})

	t.Run("bytes buffer is not terminal", func(t *testing.T) {
		var buf bytes.Buffer
		assert.False(t, isTerminal(&buf))
	})
}
