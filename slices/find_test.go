package slices

import (
	"testing"
)

type findTest struct {
	slice    []interface{}
	value    interface{}
	expected int
}

var findTests = []findTest{
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, boolTrue, 0},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, stringOne, 1},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, nil, 2},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, &stringTwo, 3},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, &intOne, 4},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, &boolTrue, 5},
	{[]any{boolTrue, nil, stringOne, nil, &stringTwo, &intOne, &boolTrue}, nil, 1},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue, stringOne}, stringOne, 1},
	{[]any{boolTrue, stringOne, &stringTwo, &stringTwo, nil, &stringTwo, &intOne, &boolTrue}, &stringTwo, 2},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, &stringFive, -1},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 0, -1},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 20, -1},
}

func TestFind(t *testing.T) {
	for _, test := range findTests {
		k := Find(&test.slice, test.value)

		if k != test.expected {
			t.Errorf("Output %d not equal to expected %d", k, test.expected)
		}
	}
}
