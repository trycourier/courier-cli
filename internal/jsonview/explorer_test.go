package jsonview

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestNewTextView(t *testing.T) {
	t.Parallel()

	t.Run("valid string", func(t *testing.T) {
		data := gjson.Parse(`"hello world"`)
		tv, err := newTextView("root", data)
		require.NoError(t, err)
		assert.Equal(t, "root", tv.GetPath())
		assert.Equal(t, "hello world", tv.GetData().String())
	})

	t.Run("non-string type returns error", func(t *testing.T) {
		data := gjson.Parse(`42`)
		_, err := newTextView("root", data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid text JSON")
	})

	t.Run("non-existent returns error", func(t *testing.T) {
		data := gjson.Parse("")
		_, err := newTextView("root", data)
		assert.Error(t, err)
	})
}

func TestNewTableView_Object(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"name":"Alice","age":30}`)
	tv, err := newTableView("", data, false)
	require.NoError(t, err)
	assert.Equal(t, "", tv.GetPath())
	assert.True(t, tv.GetData().IsObject())
	assert.Len(t, tv.rowData, 2)
}

func TestNewTableView_Array(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`[1,2,3]`)
	tv, err := newTableView("items", data, false)
	require.NoError(t, err)
	assert.Equal(t, "items", tv.GetPath())
	assert.True(t, tv.GetData().IsArray())
	assert.Len(t, tv.rowData, 3)
}

func TestNewTableView_ArrayOfObjects(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`[{"id":1,"name":"a"},{"id":2,"name":"b"}]`)
	tv, err := newTableView("", data, false)
	require.NoError(t, err)
	assert.Len(t, tv.rowData, 2)
	assert.Len(t, tv.columns, 2)
}

func TestNewTableView_EmptyArray(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`[]`)
	tv, err := newTableView("", data, false)
	require.NoError(t, err)
	assert.Len(t, tv.rowData, 0)
}

func TestNewTableView_EmptyObject(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{}`)
	tv, err := newTableView("", data, false)
	require.NoError(t, err)
	assert.Len(t, tv.rowData, 0)
}

func TestNewTableView_InvalidJSON(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`"string value"`)
	_, err := newTableView("", data, false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid table JSON")
}

func TestNewTableView_NonExistent(t *testing.T) {
	t.Parallel()

	data := gjson.Parse("")
	_, err := newTableView("", data, false)
	assert.Error(t, err)
}

func TestNewView_DispatchesToCorrectType(t *testing.T) {
	t.Parallel()

	t.Run("string creates TextView", func(t *testing.T) {
		data := gjson.Parse(`"hello"`)
		view, err := newView("", data, false)
		require.NoError(t, err)
		_, ok := view.(*TextView)
		assert.True(t, ok, "expected *TextView, got %T", view)
	})

	t.Run("object creates TableView", func(t *testing.T) {
		data := gjson.Parse(`{"key":"val"}`)
		view, err := newView("", data, false)
		require.NoError(t, err)
		_, ok := view.(*TableView)
		assert.True(t, ok, "expected *TableView, got %T", view)
	})

	t.Run("array creates TableView", func(t *testing.T) {
		data := gjson.Parse(`[1,2]`)
		view, err := newView("", data, false)
		require.NoError(t, err)
		_, ok := view.(*TableView)
		assert.True(t, ok, "expected *TableView, got %T", view)
	})
}

func TestFormatValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		json     string
		raw      bool
		contains string
	}{
		{"string normal", `"hello"`, false, "hello"},
		{"number normal", `42`, false, "42"},
		{"object normal", `{"a":1}`, false, "{"},
		{"array normal", `[1,2]`, false, "2 items"},
		{"string raw", `"hello"`, true, "hello"},
		{"number raw", `42`, true, "42"},
		{"object raw", `{"a":1}`, true, `"a":1`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := gjson.Parse(tt.json)
			output := formatValue(data, tt.raw)
			assert.Contains(t, output, tt.contains)
		})
	}
}

func TestFormatObject(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		json     string
		contains []string
	}{
		{
			"simple object",
			`{"name":"Alice","age":30}`,
			[]string{"name:", "Alice", "age:", "30"},
		},
		{
			"nested object",
			`{"outer":{"inner":1}}`,
			[]string{"outer:{…}"},
		},
		{
			"nested array",
			`{"items":[1,2,3]}`,
			[]string{"items:[…]"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := gjson.Parse(tt.json)
			output := formatObject(data)
			for _, s := range tt.contains {
				assert.Contains(t, output, s)
			}
		})
	}
}

func TestFormatArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		json     string
		expected string
	}{
		{"empty", `[]`, "[]"},
		{"one item", `[1]`, "[...1 item...]"},
		{"multiple items", `[1,2,3]`, "[...3 items...]"},
		{"ten items", `[1,2,3,4,5,6,7,8,9,10]`, "[...10 items...]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := gjson.Parse(tt.json)
			output := formatArray(data)
			assert.Equal(t, tt.expected, output)
		})
	}
}

func TestIsArrayOfObjects(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		json     string
		expected bool
	}{
		{"all objects", `[{"a":1},{"b":2}]`, true},
		{"mixed", `[{"a":1},2]`, false},
		{"all primitives", `[1,2,3]`, false},
		{"empty", `[]`, false},
		{"single object", `[{"a":1}]`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			array := gjson.Parse(tt.json).Array()
			assert.Equal(t, tt.expected, isArrayOfObjects(array))
		})
	}
}

func TestCanNavigateInto(t *testing.T) {
	t.Parallel()

	viewer := &JSONViewer{}

	tests := []struct {
		name     string
		json     string
		expected bool
	}{
		{"non-empty object", `{"key":"val"}`, true},
		{"empty object", `{}`, false},
		{"non-empty array", `[1,2,3]`, true},
		{"empty array", `[]`, false},
		{"short string", `"hello"`, false},
		{"number", `42`, false},
		{"boolean", `true`, false},
		{"null", `null`, false},
		{"multiline string", `"line1\nline2\nline3"`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := gjson.Parse(tt.json)
			assert.Equal(t, tt.expected, viewer.canNavigateInto(data))
		})
	}
}

func TestNavigateForward_EmptyRowData(t *testing.T) {
	emptyArray := gjson.Parse("[]")
	view, err := newTableView("", emptyArray, false)
	require.NoError(t, err)

	viewer := &JSONViewer{stack: []JSONView{view}, root: "Test"}

	model, cmd := viewer.navigateForward()
	require.Equal(t, model, viewer, "expected same viewer model returned")
	require.Nil(t, cmd)

	require.Equal(t, 1, len(viewer.stack), "expected stack length 1, got %d", len(viewer.stack))
}

func TestSum(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, sum([]int{}))
	assert.Equal(t, 0, sum(nil))
	assert.Equal(t, 6, sum([]int{1, 2, 3}))
	assert.Equal(t, -3, sum([]int{-1, -2, 0}))
	assert.Equal(t, 100, sum([]int{100}))
}

func TestQuoteString(t *testing.T) {
	t.Parallel()

	result := quoteString("hello")
	assert.Contains(t, result, `"hello"`)

	result = quoteString(`say "hi"`)
	assert.Contains(t, result, `\"hi\"`)

	result = quoteString(`back\slash`)
	assert.Contains(t, result, `\\`)
}

func TestJSONViewerBuildTitle(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`{"a":1}`)
	view, _ := newView("items[0]", data, false)

	t.Run("with path", func(t *testing.T) {
		viewer := &JSONViewer{stack: []JSONView{view}, root: "Test"}
		title := viewer.buildTitle(view)
		assert.Contains(t, title, "Test")
		assert.Contains(t, title, "items[0]")
	})

	t.Run("raw mode", func(t *testing.T) {
		viewer := &JSONViewer{stack: []JSONView{view}, root: "Test", rawMode: true}
		title := viewer.buildTitle(view)
		assert.Contains(t, title, "(JSON)")
	})

	t.Run("root only", func(t *testing.T) {
		rootView, _ := newView("", data, false)
		viewer := &JSONViewer{stack: []JSONView{rootView}, root: "MyRoot"}
		title := viewer.buildTitle(rootView)
		assert.Equal(t, "MyRoot", title)
	})
}

func TestJSONViewerNavigateBack(t *testing.T) {
	t.Parallel()

	data1 := gjson.Parse(`{"a":1}`)
	data2 := gjson.Parse(`{"b":2}`)
	view1, _ := newView("", data1, false)
	view2, _ := newView("a", data2, false)

	viewer := &JSONViewer{stack: []JSONView{view1, view2}, root: "Test"}
	assert.Len(t, viewer.stack, 2)

	viewer.navigateBack()
	assert.Len(t, viewer.stack, 1)

	// Should not go below 1
	viewer.navigateBack()
	assert.Len(t, viewer.stack, 1)
}

func TestJSONViewerGetStyleForData(t *testing.T) {
	t.Parallel()

	viewer := &JSONViewer{}

	// These should not panic
	viewer.getStyleForData(gjson.Parse(`"string"`))
	viewer.getStyleForData(gjson.Parse(`[1,2]`))
	viewer.getStyleForData(gjson.Parse(`{"a":1}`))
	viewer.getStyleForData(gjson.Parse(`42`))
}

func TestBuildNavigationPath(t *testing.T) {
	t.Parallel()

	t.Run("array path", func(t *testing.T) {
		data := gjson.Parse(`[{"id":1},{"id":2}]`)
		tv, _ := newTableView("items", data, false)
		viewer := &JSONViewer{}
		path := viewer.buildNavigationPath(tv, 0)
		assert.Equal(t, "items[0]", path)
	})

	t.Run("object path", func(t *testing.T) {
		data := gjson.Parse(`{"name":"Alice","age":30}`)
		tv, _ := newTableView("user", data, false)
		viewer := &JSONViewer{}
		path := viewer.buildNavigationPath(tv, 0)
		assert.Contains(t, path, "user[")
	})
}

func TestGenericIterator(t *testing.T) {
	t.Parallel()

	items := []string{"a", "b", "c"}
	it := &sliceIterator[string]{items: items, index: -1}
	anyIt := genericToAnyIterator[string](it)

	var results []any
	for anyIt.Next() {
		results = append(results, anyIt.Current())
	}

	assert.NoError(t, anyIt.Err())
	assert.Equal(t, []any{"a", "b", "c"}, results)
}

func TestGenericIterator_Empty(t *testing.T) {
	t.Parallel()

	it := &sliceIterator[int]{items: []int{}, index: -1}
	anyIt := genericToAnyIterator[int](it)

	assert.False(t, anyIt.Next())
	assert.NoError(t, anyIt.Err())
}

// sliceIterator is a test helper implementing Iterator[T]
type sliceIterator[T any] struct {
	items   []T
	index   int
	current T
}

func (s *sliceIterator[T]) Next() bool {
	s.index++
	if s.index >= len(s.items) {
		return false
	}
	s.current = s.items[s.index]
	return true
}

func (s *sliceIterator[T]) Current() T {
	return s.current
}

func (s *sliceIterator[T]) Err() error {
	return nil
}

func TestTableViewResize(t *testing.T) {
	t.Parallel()

	data := gjson.Parse(`[{"id":1,"name":"a"},{"id":2,"name":"b"}]`)
	tv, err := newTableView("", data, false)
	require.NoError(t, err)

	// Should not panic
	tv.Resize(100, 50)
	assert.Equal(t, 100, tv.width)
	assert.Equal(t, 50, tv.height)

	tv.Resize(40, 20)
	assert.Equal(t, 40, tv.width)
	assert.Equal(t, 20, tv.height)
}

func TestFormatObjectKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		key      string
		val      gjson.Result
		contains string
	}{
		{"object value", "nested", gjson.Parse(`{"a":1}`), "{…}"},
		{"array value", "items", gjson.Parse(`[1,2]`), "[…]"},
		{"short string", "name", gjson.Parse(`"Bob"`), `"Bob"`},
		{"number", "count", gjson.Parse(`42`), "42"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := formatObjectKey(tt.key, tt.val)
			assert.Contains(t, output, tt.key)
			assert.Contains(t, output, tt.contains)
		})
	}
}

type rawJSONItem struct {
	raw string
}

func (r rawJSONItem) RawJSON() string { return r.raw }

func TestMarshalItemsToJSONArray_WithHasRawJSON(t *testing.T) {
	items := []any{
		rawJSONItem{raw: `{"id":1,"name":"alice"}`},
		rawJSONItem{raw: `{"id":2,"name":"bob"}`},
	}

	got, err := marshalItemsToJSONArray(items)
	require.NoError(t, err)
	require.JSONEq(t, `[{"id":1,"name":"alice"},{"id":2,"name":"bob"}]`, string(got))
}

func TestMarshalItemsToJSONArray_WithoutHasRawJSON(t *testing.T) {
	items := []any{
		map[string]any{"id": 1, "name": "alice"},
		map[string]any{"id": 2, "name": "bob"},
	}

	got, err := marshalItemsToJSONArray(items)
	require.NoError(t, err)
	require.JSONEq(t, `[{"id":1,"name":"alice"},{"id":2,"name":"bob"}]`, string(got))
}
