package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v3"
)

func TestJaroDistance(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a, b     string
		expected float64
	}{
		{"", "", 1.0},
		{"", "abc", 0.0},
		{"abc", "", 0.0},
		{"abc", "abc", 1.0},
		{"martha", "marhta", 0.9444444444444444},
		{"kitten", "sitting", 0.746031746031746},
		{"a", "a", 1.0},
		{"a", "b", 0.0},
		{"ab", "ba", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.a+"_vs_"+tt.b, func(t *testing.T) {
			result := jaroDistance(tt.a, tt.b)
			assert.InDelta(t, tt.expected, result, 0.0001, "jaroDistance(%q, %q)", tt.a, tt.b)
		})
	}
}

func TestJaroDistanceSymmetry(t *testing.T) {
	t.Parallel()

	pairs := [][2]string{
		{"hello", "hallo"},
		{"abc", "xyz"},
		{"test", "tent"},
	}

	for _, pair := range pairs {
		t.Run(pair[0]+"_"+pair[1], func(t *testing.T) {
			assert.Equal(t, jaroDistance(pair[0], pair[1]), jaroDistance(pair[1], pair[0]))
		})
	}
}

func TestJaroWinkler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a, b string
		// jaroWinkler should always be >= jaroDistance for strings with common prefix
		expectHigherThanJaro bool
	}{
		{"martha", "marhta", true},
		{"dwayne", "duane", true},
		{"abc", "xyz", false},
		{"test", "test", false}, // identical strings; jaro already 1.0
	}

	for _, tt := range tests {
		t.Run(tt.a+"_vs_"+tt.b, func(t *testing.T) {
			jw := jaroWinkler(tt.a, tt.b)
			jd := jaroDistance(tt.a, tt.b)
			assert.GreaterOrEqual(t, jw, jd, "jaroWinkler should be >= jaroDistance")
			if tt.expectHigherThanJaro && jd > 0.7 {
				assert.Greater(t, jw, jd, "jaroWinkler should boost score for common prefix above threshold")
			}
		})
	}
}

func TestSuggestCommand(t *testing.T) {
	t.Parallel()

	commands := []*cli.Command{
		{Name: "retrieve"},
		{Name: "update"},
		{Name: "list"},
		{Name: "delete"},
	}

	tests := []struct {
		provided string
		contains string
	}{
		{"retrive", "retrieve"},
		{"retirev", "retrieve"},
		{"updat", "update"},
		{"delet", "delete"},
		{"lst", "list"},
	}

	for _, tt := range tests {
		t.Run(tt.provided, func(t *testing.T) {
			result := suggestCommand(commands, tt.provided)
			assert.Contains(t, result, tt.contains)
			assert.Contains(t, result, "Did you mean")
		})
	}
}

func TestSuggestCommandWithAliases(t *testing.T) {
	t.Parallel()

	commands := []*cli.Command{
		{Name: "generate", Aliases: []string{"gen", "g"}},
		{Name: "test"},
	}

	result := suggestCommand(commands, "generaet")
	assert.Contains(t, result, "Did you mean")
}

func TestSuggestCommand_NoCloseMatch(t *testing.T) {
	t.Parallel()

	commands := []*cli.Command{
		{Name: "retrieve"},
		{Name: "update"},
		{Name: "list"},
		{Name: "delete"},
	}

	// "xyzzy" has no meaningful similarity to any command; suggestCommand
	// still returns a "Did you mean" with whichever name scores highest,
	// even if the score is very low.
	result := suggestCommand(commands, "xyzzy")
	assert.Contains(t, result, "Did you mean")
}
