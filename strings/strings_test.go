package strings

import (
	"fmt"
	"testing"
)

type stringTest struct {
	value    any
	expected string
}

var stringNull = "null"
var stringEmpty = ""
var intOne = 1
var string1 = "1"
var stringOne = "one"
var string22 = "22"
var string1_One = "1_one"
var boolTrue = true
var boolFalse = false
var stringTrue = "true"
var stringFalse = "false"
var floatFloat = 3.14
var stringFloat = "3.14"
var slice = []any{stringOne, nil, &intOne, &stringOne}
var stringSlice = fmt.Sprintf(`["%s",%s,%d,"%s"]`, stringOne, stringNull, intOne, stringOne)
var mapany = map[int]any{1: stringOne, 2: nil, 5: &intOne, 10: &stringOne}
var stringMapany = fmt.Sprintf(`{"1":"%s","10":"%s","2":%s,"5":%d}`, stringOne, stringOne, stringNull, intOne)

var stringTests = []stringTest{
	{intOne, string1},
	{string22, string22},
	{stringOne, stringOne},
	{string1_One, string1_One},
	{nil, stringNull},
	{boolTrue, stringTrue},
	{boolFalse, stringFalse},
	{floatFloat, stringFloat},
	{slice, stringSlice},
	{mapany, stringMapany},
	{&intOne, string1},
	{&string22, string22},
	{&stringOne, stringOne},
	{&string1_One, string1_One},
	{&boolTrue, stringTrue},
	{&boolFalse, stringFalse},
	{&floatFloat, stringFloat},
	{&slice, stringSlice},
	{&mapany, stringMapany},
}

func TestString(t *testing.T) {
	for _, test := range stringTests {
		v := String(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %s", v, test.expected)
		}
	}
}
