package slices

import (
	"github.com/golibs9x/funcs/strings"
	"testing"
)

type removeValueTest struct {
	slice    []interface{}
	value    interface{}
	expected []interface{}
}

var removeValueTests = []removeValueTest{
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, boolTrue, []any{stringOne, nil, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, stringOne, []any{boolTrue, nil, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, nil, []any{boolTrue, stringOne, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, &stringTwo, []any{boolTrue, stringOne, nil, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, &intOne, []any{boolTrue, stringOne, nil, &stringTwo, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, &boolTrue, []any{boolTrue, stringOne, nil, &stringTwo, &intOne}},
	{[]any{boolTrue, nil, stringOne, nil, &stringTwo, &intOne, &boolTrue}, nil, []any{boolTrue, stringOne, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue, stringOne}, stringOne, []any{boolTrue, nil, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, &stringTwo, &stringTwo, nil, &stringTwo, &intOne, &boolTrue}, &stringTwo, []any{boolTrue, stringOne, nil, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, &stringFive, []any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 0, []any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 20, []any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}},
}

func TestRemoveValue(t *testing.T) {
	for _, test := range removeValueTests {
		RemoveValue(&test.slice, test.value)

		if strings.String(test.slice) != strings.String(test.expected) {
			t.Errorf("Output %q not equal to expected %q", test.slice, test.expected)
		}
	}
}
