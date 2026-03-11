package cmd

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDownloadFile(t *testing.T) {
	t.Parallel()

	t.Run("uses Content-Disposition filename", func(t *testing.T) {
		t.Parallel()
		dir := t.TempDir()
		origDir, _ := os.Getwd()
		require.NoError(t, os.Chdir(dir))
		defer os.Chdir(origDir)

		resp := &http.Response{Header: http.Header{}}
		resp.Header.Set("Content-Disposition", `attachment; filename="report.csv"`)

		f, err := createDownloadFile(resp, []byte("a,b,c"))
		require.NoError(t, err)
		defer f.Close()
		defer os.Remove(f.Name())
		assert.Equal(t, "report.csv", filepath.Base(f.Name()))
	})

	t.Run("falls back to MIME sniffing when no Content-Disposition", func(t *testing.T) {
		t.Parallel()
		dir := t.TempDir()
		origDir, _ := os.Getwd()
		require.NoError(t, os.Chdir(dir))
		defer os.Chdir(origDir)

		resp := &http.Response{Header: http.Header{}}
		pngData := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}

		f, err := createDownloadFile(resp, pngData)
		require.NoError(t, err)
		defer f.Close()
		defer os.Remove(f.Name())
		assert.Contains(t, f.Name(), ".png")
	})

	t.Run("directory traversal in filename is sanitized", func(t *testing.T) {
		t.Parallel()
		dir := t.TempDir()
		origDir, _ := os.Getwd()
		require.NoError(t, os.Chdir(dir))
		defer os.Chdir(origDir)

		resp := &http.Response{Header: http.Header{}}
		resp.Header.Set("Content-Disposition", `attachment; filename="../../../etc/passwd"`)

		f, err := createDownloadFile(resp, []byte("not really"))
		require.NoError(t, err)
		defer f.Close()
		defer os.Remove(f.Name())
		assert.Equal(t, "passwd", filepath.Base(f.Name()))
	})
}
