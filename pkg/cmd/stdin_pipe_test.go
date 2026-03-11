package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/trycourier/courier-cli/internal/apiquery"
	"github.com/trycourier/courier-cli/internal/requestflag"
	"github.com/urfave/cli/v3"
)

// pipeStdin replaces os.Stdin with a pipe, writes data to it, and closes the
// write end so the reader sees EOF. Returns a cleanup function that restores
// the original stdin.
func pipeStdin(t *testing.T, data string) func() {
	t.Helper()
	origStdin := os.Stdin
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdin = r
	_, err = w.WriteString(data)
	require.NoError(t, err)
	w.Close()
	return func() { os.Stdin = origStdin; r.Close() }
}

func TestFlagOptions_StdinJSON(t *testing.T) {
	cleanup := pipeStdin(t, `{"name": "from-stdin", "color": "blue"}`)
	defer cleanup()

	cmd := &cli.Command{
		Name: "test",
		Flags: []cli.Flag{
			&requestflag.Flag[string]{
				Name:     "color",
				BodyPath: "color",
			},
		},
	}
	// Flag not explicitly set; stdin provides both values
	opts, err := flagOptions(cmd, apiquery.NestedQueryFormatBrackets, apiquery.ArrayQueryFormatComma, ApplicationJSON, false)
	require.NoError(t, err)
	assert.NotEmpty(t, opts, "piped JSON should produce request options")
}

func TestFlagOptions_StdinYAML(t *testing.T) {
	cleanup := pipeStdin(t, "name: from-yaml\ncount: 3\n")
	defer cleanup()

	cmd := &cli.Command{
		Name:  "test",
		Flags: []cli.Flag{},
	}

	opts, err := flagOptions(cmd, apiquery.NestedQueryFormatBrackets, apiquery.ArrayQueryFormatComma, ApplicationJSON, false)
	require.NoError(t, err)
	assert.NotEmpty(t, opts, "piped YAML should produce request options")
}

func TestFlagOptions_StdinMergesWithFlags(t *testing.T) {
	cleanup := pipeStdin(t, `{"name": "from-stdin"}`)
	defer cleanup()

	colorFlag := &requestflag.Flag[string]{
		Name:     "color",
		BodyPath: "color",
	}

	cmd := &cli.Command{
		Name:  "test",
		Flags: []cli.Flag{colorFlag},
	}
	cmd.Set("color", "red")

	opts, err := flagOptions(cmd, apiquery.NestedQueryFormatBrackets, apiquery.ArrayQueryFormatComma, ApplicationJSON, false)
	require.NoError(t, err)
	assert.NotEmpty(t, opts, "merged stdin+flags should produce request options")
}

func TestFlagOptions_StdinSkippedWhenInUse(t *testing.T) {
	cleanup := pipeStdin(t, `{"name": "should-be-ignored"}`)
	defer cleanup()

	cmd := &cli.Command{
		Name:  "test",
		Flags: []cli.Flag{},
	}

	opts, err := flagOptions(cmd, apiquery.NestedQueryFormatBrackets, apiquery.ArrayQueryFormatComma, EmptyBody, true)
	require.NoError(t, err)
	// stdinInUse=true, so stdin data is not read; no body options produced
	assert.Empty(t, opts, "stdin should be ignored when stdinInUse is true")
}

func TestFlagOptions_StdinNonMapWithBodyFlags(t *testing.T) {
	cleanup := pipeStdin(t, `[1, 2, 3]`)
	defer cleanup()

	bodyFlag := &requestflag.Flag[string]{
		Name:     "name",
		BodyPath: "name",
	}

	cmd := &cli.Command{
		Name:  "test",
		Flags: []cli.Flag{bodyFlag},
	}
	cmd.Set("name", "test-name")

	_, err := flagOptions(cmd, apiquery.NestedQueryFormatBrackets, apiquery.ArrayQueryFormatComma, ApplicationJSON, false)
	assert.Error(t, err, "merging flags with a non-map stdin body should fail")
	assert.Contains(t, err.Error(), "Cannot merge flags with a body that is not a map")
}

func TestFlagOptions_StdinEmpty(t *testing.T) {
	cleanup := pipeStdin(t, "")
	defer cleanup()

	cmd := &cli.Command{
		Name:  "test",
		Flags: []cli.Flag{},
	}

	opts, err := flagOptions(cmd, apiquery.NestedQueryFormatBrackets, apiquery.ArrayQueryFormatComma, EmptyBody, false)
	require.NoError(t, err)
	assert.Empty(t, opts, "empty stdin should produce no options")
}
