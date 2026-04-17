package cmd

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/trycourier/courier-cli/v3/internal/jsonview"
)

func TestStreamOutput(t *testing.T) {
	t.Setenv("PAGER", "cat")
	err := streamOutput("stream test", func(w *os.File) error {
		_, writeErr := w.WriteString("Hello world\n")
		return writeErr
	})
	if err != nil {
		t.Errorf("streamOutput failed: %v", err)
	}
}

func TestWriteBinaryResponse(t *testing.T) {
	t.Run("write to explicit file", func(t *testing.T) {
		tmpDir := t.TempDir()
		outfile := tmpDir + "/output.txt"
		body := []byte("test content")
		resp := &http.Response{
			Body: io.NopCloser(bytes.NewReader(body)),
		}

		msg, err := writeBinaryResponse(resp, os.Stdout, outfile)

		require.NoError(t, err)
		assert.Contains(t, msg, outfile)

		content, err := os.ReadFile(outfile)
		require.NoError(t, err)
		assert.Equal(t, body, content)
	})

	t.Run("write to stdout", func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		body := []byte("stdout content")
		resp := &http.Response{
			Body: io.NopCloser(bytes.NewReader(body)),
		}
		msg, err := writeBinaryResponse(resp, &buf, "-")

		require.NoError(t, err)
		assert.Empty(t, msg)
		assert.Equal(t, body, buf.Bytes())
	})
}

func TestCreateDownloadFile(t *testing.T) {
	t.Run("uses Content-Disposition filename", func(t *testing.T) {
		t.Chdir(t.TempDir())

		resp := &http.Response{Header: http.Header{}}
		resp.Header.Set("Content-Disposition", `attachment; filename="report.csv"`)

		f, err := createDownloadFile(resp, []byte("a,b,c"))
		require.NoError(t, err)
		defer f.Close()
		assert.Equal(t, "report.csv", filepath.Base(f.Name()))
	})

	t.Run("does not clobber existing file with same name", func(t *testing.T) {
		t.Chdir(t.TempDir())

		resp := &http.Response{
			Header: http.Header{
				"Content-Disposition": []string{`attachment; filename="test.txt"`},
			},
		}
		file, err := createDownloadFile(resp, []byte("first"))
		require.NoError(t, err)
		defer file.Close()

		file2, err := createDownloadFile(resp, []byte("second"))
		require.NoError(t, err)
		defer file2.Close()
		assert.NotEqual(t, file.Name(), file2.Name(), "second file should have a different name")
		assert.Contains(t, filepath.Base(file2.Name()), "test")
	})

	t.Run("falls back to MIME sniffing when no Content-Disposition", func(t *testing.T) {
		t.Chdir(t.TempDir())

		resp := &http.Response{Header: http.Header{}}
		pngData := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}

		f, err := createDownloadFile(resp, pngData)
		require.NoError(t, err)
		defer f.Close()
		assert.Contains(t, f.Name(), ".png")
	})

	t.Run("directory traversal in filename is sanitized", func(t *testing.T) {
		t.Chdir(t.TempDir())

		resp := &http.Response{Header: http.Header{}}
		resp.Header.Set("Content-Disposition", `attachment; filename="../../../etc/passwd"`)

		f, err := createDownloadFile(resp, []byte("not really"))
		require.NoError(t, err)
		defer f.Close()
		assert.Equal(t, "passwd", filepath.Base(f.Name()))
	})
}

func TestValidateBaseURL(t *testing.T) {
	t.Parallel()

	t.Run("ValidHTTPS", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, ValidateBaseURL("https://api.example.com", "--base-url"))
	})

	t.Run("ValidHTTP", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, ValidateBaseURL("http://localhost:8080", "--base-url"))
	})

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, ValidateBaseURL("", "MY_BASE_URL"))
	})

	t.Run("MissingScheme", func(t *testing.T) {
		t.Parallel()
		err := ValidateBaseURL("localhost:8080", "MY_BASE_URL")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "MY_BASE_URL")
		assert.Contains(t, err.Error(), "missing a scheme")
	})

	t.Run("HostOnly", func(t *testing.T) {
		t.Parallel()
		err := ValidateBaseURL("api.example.com", "--base-url")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--base-url")
	})
}

func TestFormatJSON(t *testing.T) {
	t.Parallel()

	t.Run("RawWithTransform", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`{"id":"abc123","name":"test"}`)
		formatted, err := formatJSON(os.Stdout, "test", res, "raw", "id", false)
		require.NoError(t, err)
		require.Equal(t, `"abc123"`+"\n", string(formatted))
	})

	t.Run("RawWithoutTransform", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`{"id":"abc123","name":"test"}`)
		formatted, err := formatJSON(os.Stdout, "test", res, "raw", "", false)
		require.NoError(t, err)
		require.Equal(t, `{"id":"abc123","name":"test"}`+"\n", string(formatted))
	})

	t.Run("RawWithNestedTransform", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`{"data":{"items":[1,2,3]}}`)
		formatted, err := formatJSON(os.Stdout, "test", res, "raw", "data.items", false)
		require.NoError(t, err)
		require.Equal(t, "[1,2,3]\n", string(formatted))
	})

	t.Run("RawWithNonexistentTransform", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`{"id":"abc123"}`)
		formatted, err := formatJSON(os.Stdout, "test", res, "raw", "missing", false)
		require.NoError(t, err)
		// Transform path doesn't exist, so original result is returned
		require.Equal(t, `{"id":"abc123"}`+"\n", string(formatted))
	})

	t.Run("RawOutputString", func(t *testing.T) {
		t.Parallel()

		res := gjson.Parse(`{"id":"abc123","name":"test"}`)
		formatted, err := formatJSON(os.Stdout, "test", res, "json", "id", true)
		require.NoError(t, err)
		require.Equal(t, "abc123\n", string(formatted))
	})

	t.Run("RawOutputNonString", func(t *testing.T) {
		t.Parallel()

		// --raw-output has no effect on non-string values
		res := gjson.Parse(`{"count":42}`)
		formatted, err := formatJSON(os.Stdout, "test", res, "raw", "count", true)
		require.NoError(t, err)
		require.Equal(t, "42\n", string(formatted))
	})

	t.Run("RawOutputObject", func(t *testing.T) {
		t.Parallel()

		// --raw-output has no effect on objects
		res := gjson.Parse(`{"nested":{"a":1}}`)
		formatted, err := formatJSON(os.Stdout, "test", res, "raw", "nested", true)
		require.NoError(t, err)
		require.Equal(t, `{"a":1}`+"\n", string(formatted))
	})
}

func TestShowJSONIterator(t *testing.T) {
	t.Parallel()

	t.Run("RawMultipleItems", func(t *testing.T) {
		t.Parallel()

		iter := &sliceIterator[map[string]any]{items: []map[string]any{
			{"id": "abc", "name": "first"},
			{"id": "def", "name": "second"},
		}}
		captured := captureShowJSONIterator(t, iter, "raw", "", -1)
		assert.Equal(t, `{"id":"abc","name":"first"}`+"\n"+`{"id":"def","name":"second"}`+"\n", captured)
	})

	t.Run("RawWithTransform", func(t *testing.T) {
		t.Parallel()

		iter := &sliceIterator[map[string]any]{items: []map[string]any{
			{"id": "abc", "name": "first"},
			{"id": "def", "name": "second"},
		}}
		captured := captureShowJSONIterator(t, iter, "raw", "id", -1)
		assert.Equal(t, `"abc"`+"\n"+`"def"`+"\n", captured)
	})

	t.Run("LimitItems", func(t *testing.T) {
		t.Parallel()

		iter := &sliceIterator[map[string]any]{items: []map[string]any{
			{"id": "abc"},
			{"id": "def"},
			{"id": "ghi"},
		}}
		captured := captureShowJSONIterator(t, iter, "raw", "", 2)
		assert.Equal(t, `{"id":"abc"}`+"\n"+`{"id":"def"}`+"\n", captured)
	})
}

func TestExploreFallback(t *testing.T) {
	t.Parallel()

	t.Run("ShowJSONFallsBackToJsonOnNonTTY", func(t *testing.T) {
		t.Parallel()

		// os.Pipe() produces a *os.File that isn't a terminal, so explore should fall back.
		r, w, err := os.Pipe()
		require.NoError(t, err)
		defer r.Close()

		var stderr bytes.Buffer
		res := gjson.Parse(`{"id":"abc"}`)
		err = ShowJSON(res, ShowJSONOpts{
			Format: "explore",
			Stderr: &stderr,
			Stdout: w,
			Title:  "test",
		})
		w.Close()
		require.NoError(t, err)

		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)
		assert.Contains(t, buf.String(), `"id"`)
		assert.Contains(t, buf.String(), `"abc"`)
	})

	t.Run("ShowJSONIteratorFallsBackToJsonOnNonTTY", func(t *testing.T) {
		t.Parallel()

		iter := &sliceIterator[map[string]any]{items: []map[string]any{
			{"id": "abc"},
		}}
		captured := captureShowJSONIterator(t, iter, "explore", "", -1)
		assert.Contains(t, captured, `"id"`)
		assert.Contains(t, captured, `"abc"`)
	})

	t.Run("ShowJSONWarnsWhenExplicitFormatOnNonTTY", func(t *testing.T) {
		t.Parallel()

		r, w, err := os.Pipe()
		require.NoError(t, err)
		defer r.Close()

		var stderr bytes.Buffer
		res := gjson.Parse(`{"id":"abc"}`)
		err = ShowJSON(res, ShowJSONOpts{
			ExplicitFormat: true,
			Format:         "explore",
			Stderr:         &stderr,
			Stdout:         w,
			Title:          "test",
		})
		w.Close()
		require.NoError(t, err)

		assert.Equal(t, warningExploreNotSupported, stderr.String())
	})

	t.Run("ShowJSONSilentWhenDefaultFormatOnNonTTY", func(t *testing.T) {
		t.Parallel()

		r, w, err := os.Pipe()
		require.NoError(t, err)
		defer r.Close()

		var stderr bytes.Buffer
		res := gjson.Parse(`{"id":"abc"}`)
		err = ShowJSON(res, ShowJSONOpts{
			Format: "explore",
			Stderr: &stderr,
			Stdout: w,
			Title:  "test",
		})
		w.Close()
		require.NoError(t, err)

		assert.Empty(t, stderr.String(), "no warning expected when format was not explicit")
	})
}

// sliceIterator is a simple iterator over a slice for testing.
type sliceIterator[T any] struct {
	index int
	items []T
}

func (it *sliceIterator[T]) Next() bool {
	it.index++
	return it.index <= len(it.items)
}

func (it *sliceIterator[T]) Current() T {
	return it.items[it.index-1]
}

func (it *sliceIterator[T]) Err() error {
	return nil
}

var _ jsonview.Iterator[any] = (*sliceIterator[any])(nil)

// captureShowJSONIterator runs ShowJSONIterator and captures the output written to a file.
func captureShowJSONIterator[T any](t *testing.T, iter jsonview.Iterator[T], format, transform string, itemsToDisplay int64) string {
	t.Helper()

	r, w, err := os.Pipe()
	require.NoError(t, err)
	defer r.Close()

	err = ShowJSONIterator(iter, itemsToDisplay, ShowJSONOpts{
		Format:    format,
		Stderr:    io.Discard,
		Stdout:    w,
		Title:     "test",
		Transform: transform,
	})
	w.Close()
	require.NoError(t, err)

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}
