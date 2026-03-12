package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// rawJSONItem implements HasRawJSON for testing
type rawJSONItem struct {
	raw string
}

func (r rawJSONItem) RawJSON() string { return r.raw }

// testIterator is a simple in-memory iterator for testing
type testIterator[T any] struct {
	items   []T
	index   int
	current T
	err     error
}

func newTestIterator[T any](items ...T) *testIterator[T] {
	return &testIterator[T]{items: items, index: -1}
}

func (it *testIterator[T]) Next() bool {
	it.index++
	if it.index >= len(it.items) {
		return false
	}
	it.current = it.items[it.index]
	return true
}

func (it *testIterator[T]) Current() T { return it.current }
func (it *testIterator[T]) Err() error { return it.err }

func TestShowJSONIterator_Formats(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		format   string
		items    []rawJSONItem
		contains []string
		usePipe  bool // yaml writes directly to the writer, not via returned bytes
	}{
		{
			name:     "raw",
			format:   "raw",
			items:    []rawJSONItem{{`{"id":"1","name":"alpha"}`}, {`{"id":"2","name":"beta"}`}},
			contains: []string{`{"id":"1","name":"alpha"}`, `{"id":"2","name":"beta"}`},
		},
		{
			name:     "jsonl",
			format:   "jsonl",
			items:    []rawJSONItem{{`{"id":"1"}`}, {`{"id":"2"}`}, {`{"id":"3"}`}},
			contains: []string{`{"id":"1"}`, `{"id":"2"}`, `{"id":"3"}`},
		},
		{
			name:     "json",
			format:   "json",
			items:    []rawJSONItem{{`{"status":"ok"}`}},
			contains: []string{"status", "ok"},
		},
		{
			name:     "pretty",
			format:   "pretty",
			items:    []rawJSONItem{{`{"name":"Alice","age":30}`}},
			contains: []string{"name:", "Alice"},
		},
		{
			name:     "yaml",
			format:   "yaml",
			items:    []rawJSONItem{{`{"name":"Alice","age":30}`}},
			contains: []string{"name:", "Alice"},
			usePipe:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			iter := newTestIterator(tt.items...)

			var output string
			if tt.usePipe {
				r, w, err := os.Pipe()
				require.NoError(t, err)
				defer r.Close()
				err = ShowJSONIterator(w, "Test", iter, tt.format, "", 0)
				w.Close()
				require.NoError(t, err)
				buf := make([]byte, 8192)
				n, _ := r.Read(buf)
				output = string(buf[:n])
			} else {
				tmpFile, err := os.CreateTemp("", "iter-*")
				require.NoError(t, err)
				defer os.Remove(tmpFile.Name())
				defer tmpFile.Close()
				err = ShowJSONIterator(tmpFile, "Test", iter, tt.format, "", 0)
				require.NoError(t, err)
				tmpFile.Seek(0, 0)
				content, err := os.ReadFile(tmpFile.Name())
				require.NoError(t, err)
				output = string(content)
			}

			for _, s := range tt.contains {
				assert.Contains(t, output, s)
			}
		})
	}
}

func TestShowJSONIterator_EmptyIterator(t *testing.T) {
	t.Parallel()

	items := newTestIterator[rawJSONItem]()

	tmpFile, err := os.CreateTemp("", "iter-empty-*")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	err = ShowJSONIterator(tmpFile, "Test", items, "json", "", 0)
	require.NoError(t, err)

	tmpFile.Seek(0, 0)
	content, err := os.ReadFile(tmpFile.Name())
	require.NoError(t, err)
	assert.Empty(t, string(content), "empty iterator should produce no output")
}

func TestShowJSONIterator_WithTransform(t *testing.T) {
	t.Parallel()

	items := newTestIterator(
		rawJSONItem{`{"data":{"value":"deep"},"other":"skip"}`},
	)

	tmpFile, err := os.CreateTemp("", "iter-transform-*")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	err = ShowJSONIterator(tmpFile, "Test", items, "jsonl", "data.value", 0)
	require.NoError(t, err)

	tmpFile.Seek(0, 0)
	content, err := os.ReadFile(tmpFile.Name())
	require.NoError(t, err)

	output := string(content)
	assert.Equal(t, "\"deep\"\n", output, "transform should extract nested value")
}

func TestShowJSONIterator_NonRawJSONItem(t *testing.T) {
	t.Parallel()

	type plainItem struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	items := newTestIterator(
		plainItem{ID: "1", Name: "alpha"},
		plainItem{ID: "2", Name: "beta"},
	)

	tmpFile, err := os.CreateTemp("", "iter-plain-*")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	err = ShowJSONIterator(tmpFile, "Test", items, "raw", "", 0)
	require.NoError(t, err)

	tmpFile.Seek(0, 0)
	content, err := os.ReadFile(tmpFile.Name())
	require.NoError(t, err)

	output := string(content)
	assert.Contains(t, output, "alpha")
	assert.Contains(t, output, "beta")
}

func TestShowJSONIterator_IteratorError(t *testing.T) {
	t.Parallel()

	items := newTestIterator(
		rawJSONItem{`{"id":"1"}`},
	)
	items.err = assert.AnError

	tmpFile, err := os.CreateTemp("", "iter-err-*")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	err = ShowJSONIterator(tmpFile, "Test", items, "raw", "", 0)
	// After iterating all items, iter.Err() should be returned
	assert.Error(t, err)
}
