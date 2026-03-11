package jsonview

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestFormatJSON_InvalidJSON(t *testing.T) {
	t.Parallel()
	result := gjson.Parse("")
	output := formatJSON(result, 80)
	assert.Contains(t, output, "Invalid JSON")
}

func TestFormatResult_AllTypes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		json     string
		width    int
		contains []string
	}{
		{"string", `"hello"`, 200, []string{"hello"}},
		{"empty string", `""`, 200, []string{"(empty)"}},
		{"integer", `99`, 200, []string{"99"}},
		{"negative", `-5`, 200, []string{"-5"}},
		{"float", `1.5`, 200, []string{"1.5"}},
		{"true", `true`, 200, []string{"yes"}},
		{"false", `false`, 200, []string{"no"}},
		{"null", `null`, 200, []string{"null"}},
		{"object", `{"name":"Alice","age":30}`, 200, []string{"name:", "Alice", "age:", "30"}},
		{"empty object", `{}`, 200, []string{"(empty)"}},
		{"array", `["a","b","c"]`, 200, []string{"1.", "a", "2.", "b", "3.", "c"}},
		{"empty array", `[]`, 200, []string{"(none)"}},
		{"nested object", `{"user":{"name":"Bob"}}`, 200, []string{"user:", "name:", "Bob"}},
		{"array of objects", `[{"id":1},{"id":2}]`, 200, []string{"1.", "2.", "id:"}},
		{"narrow width doesn't panic", `"` + string(make([]byte, 200)) + `"`, 20, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gjson.Parse(tt.json)
			output := formatResult(result, 0, tt.width)
			if len(tt.contains) == 0 {
				assert.NotEmpty(t, output)
			}
			for _, s := range tt.contains {
				assert.Contains(t, output, s)
			}
		})
	}
}

func TestIsSingleLine(t *testing.T) {
	t.Parallel()

	assert.True(t, isSingleLine(gjson.Parse(`"hello"`), 0))
	assert.True(t, isSingleLine(gjson.Parse(`42`), 0))
	assert.True(t, isSingleLine(gjson.Parse(`true`), 0))
	assert.True(t, isSingleLine(gjson.Parse(`null`), 0))
	assert.False(t, isSingleLine(gjson.Parse(`{"key":"val"}`), 0))
	assert.False(t, isSingleLine(gjson.Parse(`[1,2,3]`), 0))
}

func TestGetIndent(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", getIndent(0))
	assert.Equal(t, "  ", getIndent(1))
	assert.Equal(t, "    ", getIndent(2))
	assert.Equal(t, "      ", getIndent(3))
}

func TestRenderJSON(t *testing.T) {
	t.Parallel()

	result := gjson.Parse(`{"status":"ok","count":5}`)
	output := RenderJSON("Test Title", result)
	assert.Contains(t, output, "Test Title")
	assert.Contains(t, output, "status:")
	assert.Contains(t, output, "ok")
	assert.Contains(t, output, "count:")
	assert.Contains(t, output, "5")
}

func TestRenderJSON_Array(t *testing.T) {
	t.Parallel()

	result := gjson.Parse(`[1,2,3]`)
	output := RenderJSON("Array Title", result)
	assert.Contains(t, output, "Array Title")
	assert.Contains(t, output, "1")
	assert.Contains(t, output, "2")
	assert.Contains(t, output, "3")
}
