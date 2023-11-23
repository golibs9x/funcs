package slice

import (
	"testing"
)

var stringOne = "one"
var stringTwo = "two"
var stringFive = "five"
var stringTen = "ten"
var intOne = 1
var intTwo = 2
var intFive = 5
var intTen = 10
var boolTrue = true
var boolFalse = false

type valueIntTest struct {
	slice    []int
	index    int
	def      int
	expected int
}

var valueIntTests = []valueIntTest{
	{[]int{intOne, intTwo, intFive}, 0, intOne, intOne},
	{[]int{intOne, intTwo, intFive}, 1, intTwo, intTwo},
	{[]int{intOne, intTwo, intFive}, 2, intFive, intFive},
	{[]int{intOne, intTwo, intFive}, 4, intTen, intTen},
}

func TestValueInt(t *testing.T) {
	for _, test := range valueIntTests {
		v := Value(&test.slice, test.index, test.def)

		if v != test.expected {
			t.Errorf("Output %d not equal to expected %d", v, test.expected)
		}
	}
}

type valueStringTest struct {
	slice    []string
	index    int
	def      string
	expected string
}

var valueStringTests = []valueStringTest{
	{[]string{stringOne, stringTwo, stringFive}, 0, stringOne, stringOne},
	{[]string{stringOne, stringTwo, stringFive}, 1, stringTwo, stringTwo},
	{[]string{stringOne, stringTwo, stringFive}, 2, stringFive, stringFive},
	{[]string{stringOne, stringTwo, stringFive}, 4, stringTen, stringTen},
}

func TestValueString(t *testing.T) {
	for _, test := range valueStringTests {
		v := Value(&test.slice, test.index, test.def)

		if v != test.expected {
			t.Errorf("Output %s not equal to expected %s", v, test.expected)
		}
	}
}

type valueBoolTest struct {
	slice    []bool
	index    int
	def      bool
	expected bool
}

var valueBoolTests = []valueBoolTest{
	{[]bool{boolTrue, boolTrue, boolFalse}, 0, boolTrue, boolTrue},
	{[]bool{boolTrue, boolTrue, boolFalse}, 1, boolTrue, boolTrue},
	{[]bool{boolTrue, boolTrue, boolFalse}, 2, boolFalse, boolFalse},
	{[]bool{boolTrue, boolTrue, boolFalse}, 4, boolTrue, boolTrue},
}

func TestValueBool(t *testing.T) {
	for _, test := range valueBoolTests {
		v := Value(&test.slice, test.index, test.def)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %v", v, test.expected)
		}
	}
}

type valueAnyTest struct {
	slice    []interface{}
	index    int
	def      interface{}
	expected interface{}
}

var valueAnyTests = []valueAnyTest{
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 0, boolTrue, boolTrue},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 1, stringOne, stringOne},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 2, nil, nil},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 3, &stringTwo, &stringTwo},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 4, &intOne, &intOne},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 5, &boolTrue, &boolTrue},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 6, &boolTrue, &boolTrue},
	{[]any{boolTrue, stringOne, nil, &stringTwo, &intOne, &boolTrue}, 6, nil, nil},
}

func TestValueAny(t *testing.T) {
	for _, test := range valueAnyTests {
		v := Value(&test.slice, test.index, test.def)

		if v != test.expected {
			t.Errorf("Output %q not equal to expected %q", v, test.expected)
		}
	}
}
