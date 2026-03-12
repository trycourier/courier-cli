package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trycourier/courier-cli/v3/pkg/cmd"
	"github.com/urfave/cli/v3"
)

func TestPrepareForAutocomplete(t *testing.T) {
	t.Parallel()

	root := &cli.Command{
		Name: "test",
		Commands: []*cli.Command{
			{
				Name: "child1",
				Commands: []*cli.Command{
					{Name: "grandchild1"},
				},
			},
			{Name: "child2"},
		},
	}

	assert.False(t, root.SkipFlagParsing)
	assert.False(t, root.Commands[0].SkipFlagParsing)
	assert.False(t, root.Commands[0].Commands[0].SkipFlagParsing)
	assert.False(t, root.Commands[1].SkipFlagParsing)

	prepareForAutocomplete(root)

	assert.True(t, root.SkipFlagParsing, "root should have SkipFlagParsing set")
	assert.True(t, root.Commands[0].SkipFlagParsing, "child1 should have SkipFlagParsing set")
	assert.True(t, root.Commands[0].Commands[0].SkipFlagParsing, "grandchild1 should have SkipFlagParsing set")
	assert.True(t, root.Commands[1].SkipFlagParsing, "child2 should have SkipFlagParsing set")
}

func TestPrepareForAutocomplete_EmptyCommand(t *testing.T) {
	t.Parallel()

	root := &cli.Command{Name: "empty"}
	prepareForAutocomplete(root)
	assert.True(t, root.SkipFlagParsing)
}

func TestCommandErrorBuffer(t *testing.T) {
	t.Parallel()

	// Verify the CommandErrorBuffer exists and is usable
	cmd.CommandErrorBuffer.Reset()
	assert.Equal(t, 0, cmd.CommandErrorBuffer.Len())

	cmd.CommandErrorBuffer.WriteString("test error")
	assert.Equal(t, "test error", cmd.CommandErrorBuffer.String())
	cmd.CommandErrorBuffer.Reset()
}

func TestAppCommandIsInitialized(t *testing.T) {
	t.Parallel()

	app := cmd.Command
	assert.NotNil(t, app)
	assert.Equal(t, "courier", app.Name)
	assert.NotEmpty(t, app.Version)
	assert.NotEmpty(t, app.Commands)
}
