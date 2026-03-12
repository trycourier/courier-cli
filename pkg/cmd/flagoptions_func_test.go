package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/trycourier/courier-cli/v3/internal/apiquery"
	"github.com/trycourier/courier-cli/v3/internal/requestflag"
	"github.com/urfave/cli/v3"
)

func TestFlagOptions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		flags      []cli.Flag
		setFlags   map[string]string
		bodyType   BodyContentType
		stdinInUse bool
		wantEmpty  bool
	}{
		{
			name:      "empty body, no flags",
			flags:     []cli.Flag{},
			bodyType:  EmptyBody,
			wantEmpty: true,
		},
		{
			name:     "query flag produces options",
			flags:    []cli.Flag{&requestflag.Flag[string]{Name: "cursor", QueryPath: "cursor"}},
			setFlags: map[string]string{"cursor": "abc123"},
			bodyType: EmptyBody,
		},
		{
			name:     "header flag produces options",
			flags:    []cli.Flag{&requestflag.Flag[string]{Name: "idempotency-key", HeaderPath: "Idempotency-Key"}},
			setFlags: map[string]string{"idempotency-key": "unique-key-123"},
			bodyType: EmptyBody,
		},
		{
			name:     "application/json body",
			flags:    []cli.Flag{&requestflag.Flag[string]{Name: "name", BodyPath: "name"}},
			setFlags: map[string]string{"name": "test-name"},
			bodyType: ApplicationJSON,
		},
		{
			name:     "application/octet-stream with string body",
			flags:    []cli.Flag{&requestflag.Flag[string]{Name: "data", BodyRoot: true}},
			setFlags: map[string]string{"data": "raw-binary-content"},
			bodyType: ApplicationOctetStream,
		},
		{
			name:     "debug mode adds middleware",
			flags:    []cli.Flag{&cli.BoolFlag{Name: "debug"}},
			setFlags: map[string]string{"debug": "true"},
			bodyType: EmptyBody,
		},
		{
			name:     "multipart form encoded",
			flags:    []cli.Flag{&requestflag.Flag[string]{Name: "field", BodyPath: "field"}},
			setFlags: map[string]string{"field": "value"},
			bodyType: MultipartFormEncoded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cli.Command{Name: "test", Flags: tt.flags}
			for k, v := range tt.setFlags {
				cmd.Set(k, v)
			}
			opts, err := flagOptions(cmd, apiquery.NestedQueryFormatBrackets, apiquery.ArrayQueryFormatComma, tt.bodyType, tt.stdinInUse)
			require.NoError(t, err)
			if tt.wantEmpty {
				assert.Empty(t, opts)
			} else {
				assert.NotEmpty(t, opts)
			}
		})
	}
}

func TestGuessExtension(t *testing.T) {
	t.Parallel()

	t.Run("JPEG", func(t *testing.T) {
		data := []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46}
		assert.Equal(t, ".jpg", guessExtension(data))
	})

	t.Run("PNG", func(t *testing.T) {
		data := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
		assert.Equal(t, ".png", guessExtension(data))
	})

	t.Run("GIF", func(t *testing.T) {
		data := []byte("GIF89a" + "\x00\x00\x00\x00")
		assert.Equal(t, ".gif", guessExtension(data))
	})

	t.Run("PDF", func(t *testing.T) {
		data := []byte("%PDF-1.4 test content here")
		assert.Equal(t, ".pdf", guessExtension(data))
	})

	t.Run("ZIP", func(t *testing.T) {
		data := []byte{0x50, 0x4B, 0x03, 0x04, 0x00, 0x00, 0x00, 0x00}
		assert.Equal(t, ".zip", guessExtension(data))
	})

	t.Run("plain text returns a text extension", func(t *testing.T) {
		data := []byte("Hello, World! This is plain text.")
		ext := guessExtension(data)
		assert.NotEmpty(t, ext)
		// text/plain may return .asc or .txt depending on the platform
		assert.True(t, ext == ".txt" || ext == ".asc", "expected .txt or .asc, got %s", ext)
	})

	t.Run("HTML returns an HTML extension", func(t *testing.T) {
		data := []byte("<html><body>Hello</body></html>")
		ext := guessExtension(data)
		assert.True(t, ext == ".html" || ext == ".htm", "expected .html or .htm, got %s", ext)
	})

	t.Run("JSON returns a text extension", func(t *testing.T) {
		data := []byte(`{"key":"value"}`)
		ext := guessExtension(data)
		assert.NotEmpty(t, ext)
	})
}

func TestGuessExtension_Binary(t *testing.T) {
	t.Parallel()

	// Random binary data that doesn't match known signatures
	data := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}
	ext := guessExtension(data)
	assert.NotEmpty(t, ext, "should return some extension for unknown binary")
}

