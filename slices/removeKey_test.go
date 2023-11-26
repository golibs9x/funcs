package slices

import (
	"github.com/golibs9x/funcs/strings"
	"testing"
)

type removeKeyTest struct {
	slice    []interface{}
	key      int
	expected []interface{}
}

var removeKeyTests = []removeKeyTest{
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 0, []any{stringOne, nil, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 1, []any{boolTrue, nil, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 2, []any{boolTrue, stringOne, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 3, []any{boolTrue, stringOne, nil, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 4, []any{boolTrue, stringOne, nil, &stringTwo, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 5, []any{boolTrue, stringOne, nil, &stringTwo, &intOne}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 6, []any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 7, []any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}},
}

func TestRemoveKey(t *testing.T) {
	for _, test := range removeKeyTests {
		RemoveKey(&test.slice, test.key)

		if strings.String(test.slice) != strings.String(test.expected) {
			t.Errorf("Output %q not equal to expected %q", test.slice, test.expected)
		}
	}
}
