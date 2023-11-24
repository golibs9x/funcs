package ints

import "testing"

var int_0 int = 0
var int_1 int = 1
var int_2 int = 2
var int_10 int = 10
var int64_0 int64 = 0
var int64_1 int64 = 1
var int64_2 int64 = 2
var int64_10 int64 = 10
var string1 = "1"
var string2 = "2"
var string10 = "10"
var intFloat = 3.14
var stringFloat = "3.14"
var boolTrue = true
var boolFalse = false
var stringEmpty = ""
var stringString = "string"

type int64Test struct {
	value    any
	expected int64
}

var int64Tests = []int64Test{
	{int_0, int64_0},
	{int_1, int64_1},
	{int_2, int64_2},
	{int_10, int64_10},
	{int64_0, int64_0},
	{int64_1, int64_1},
	{int64_2, int64_2},
	{int64_10, int64_10},
	{string1, int64_1},
	{string2, int64_2},
	{string10, int64_10},
	{intFloat, int64_0},
	{stringFloat, int64_0},
	{boolTrue, int64_1},
	{boolFalse, int64_0},
	{stringEmpty, int64_0},
	{stringString, int64_0},
	{&int_0, int64_0},
	{&int_1, int64_1},
	{&int_2, int64_2},
	{&int_10, int64_10},
	{&int64_0, int64_0},
	{&int64_1, int64_1},
	{&int64_2, int64_2},
	{&int64_10, int64_10},
	{&string1, int64_1},
	{&string2, int64_2},
	{&string10, int64_10},
	{&intFloat, int64_0},
	{&stringFloat, int64_0},
	{&boolTrue, int64_1},
	{&boolFalse, int64_0},
	{&stringEmpty, int64_0},
	{&stringString, int64_0},
}

func TestInt64(t *testing.T) {
	for _, test := range int64Tests {
		v := Int64(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}

type intTest struct {
	value    any
	expected int
}

var intTests = []intTest{
	{int_0, int_0},
	{int_1, int_1},
	{int_2, int_2},
	{int_10, int_10},
	{int64_0, int_0},
	{int64_1, int_1},
	{int64_2, int_2},
	{int64_10, int_10},
	{string1, int_1},
	{string2, int_2},
	{string10, int_10},
	{intFloat, int_0},
	{stringFloat, int_0},
	{boolTrue, int_1},
	{boolFalse, int_0},
	{stringEmpty, int_0},
	{stringString, int_0},
	{&int_0, int_0},
	{&int_1, int_1},
	{&int_2, int_2},
	{&int_10, int_10},
	{&int64_0, int_0},
	{&int64_1, int_1},
	{&int64_2, int_2},
	{&int64_10, int_10},
	{&string1, int_1},
	{&string2, int_2},
	{&string10, int_10},
	{&intFloat, int_0},
	{&stringFloat, int_0},
	{&boolTrue, int_1},
	{&boolFalse, int_0},
	{&stringEmpty, int_0},
	{&stringString, int_0},
}

func TestInt(t *testing.T) {
	for _, test := range intTests {
		v := Int(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}

type int8Test struct {
	value    any
	expected int8
}

var int8Tests = []int8Test{
	{int_0, int8(int_0)},
	{int_1, int8(int_1)},
	{int_2, int8(int_2)},
	{int_10, int8(int_10)},
	{int64_0, int8(int_0)},
	{int64_1, int8(int_1)},
	{int64_2, int8(int_2)},
	{int64_10, int8(int_10)},
	{string1, int8(int_1)},
	{string2, int8(int_2)},
	{string10, int8(int_10)},
	{intFloat, int8(int_0)},
	{stringFloat, int8(int_0)},
	{boolTrue, int8(int_1)},
	{boolFalse, int8(int_0)},
	{stringEmpty, int8(int_0)},
	{stringString, int8(int_0)},
	{&int_0, int8(int_0)},
	{&int_1, int8(int_1)},
	{&int_2, int8(int_2)},
	{&int_10, int8(int_10)},
	{&int64_0, int8(int_0)},
	{&int64_1, int8(int_1)},
	{&int64_2, int8(int_2)},
	{&int64_10, int8(int_10)},
	{&string1, int8(int_1)},
	{&string2, int8(int_2)},
	{&string10, int8(int_10)},
	{&intFloat, int8(int_0)},
	{&stringFloat, int8(int_0)},
	{&boolTrue, int8(int_1)},
	{&boolFalse, int8(int_0)},
	{&stringEmpty, int8(int_0)},
	{&stringString, int8(int_0)},
}

func TestInt8(t *testing.T) {
	for _, test := range int8Tests {
		v := Int8(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}

type int16Test struct {
	value    any
	expected int16
}

var int16Tests = []int16Test{
	{int_0, int16(int_0)},
	{int_1, int16(int_1)},
	{int_2, int16(int_2)},
	{int_10, int16(int_10)},
	{int64_0, int16(int_0)},
	{int64_1, int16(int_1)},
	{int64_2, int16(int_2)},
	{int64_10, int16(int_10)},
	{string1, int16(int_1)},
	{string2, int16(int_2)},
	{string10, int16(int_10)},
	{intFloat, int16(int_0)},
	{stringFloat, int16(int_0)},
	{boolTrue, int16(int_1)},
	{boolFalse, int16(int_0)},
	{stringEmpty, int16(int_0)},
	{stringString, int16(int_0)},
	{&int_0, int16(int_0)},
	{&int_1, int16(int_1)},
	{&int_2, int16(int_2)},
	{&int_10, int16(int_10)},
	{&int64_0, int16(int_0)},
	{&int64_1, int16(int_1)},
	{&int64_2, int16(int_2)},
	{&int64_10, int16(int_10)},
	{&string1, int16(int_1)},
	{&string2, int16(int_2)},
	{&string10, int16(int_10)},
	{&intFloat, int16(int_0)},
	{&stringFloat, int16(int_0)},
	{&boolTrue, int16(int_1)},
	{&boolFalse, int16(int_0)},
	{&stringEmpty, int16(int_0)},
	{&stringString, int16(int_0)},
}

func TestInt16(t *testing.T) {
	for _, test := range int16Tests {
		v := Int16(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}

type int32Test struct {
	value    any
	expected int32
}

var int32Tests = []int32Test{
	{int_0, int32(int_0)},
	{int_1, int32(int_1)},
	{int_2, int32(int_2)},
	{int_10, int32(int_10)},
	{int64_0, int32(int_0)},
	{int64_1, int32(int_1)},
	{int64_2, int32(int_2)},
	{int64_10, int32(int_10)},
	{string1, int32(int_1)},
	{string2, int32(int_2)},
	{string10, int32(int_10)},
	{intFloat, int32(int_0)},
	{stringFloat, int32(int_0)},
	{boolTrue, int32(int_1)},
	{boolFalse, int32(int_0)},
	{stringEmpty, int32(int_0)},
	{stringString, int32(int_0)},
	{&int_0, int32(int_0)},
	{&int_1, int32(int_1)},
	{&int_2, int32(int_2)},
	{&int_10, int32(int_10)},
	{&int64_0, int32(int_0)},
	{&int64_1, int32(int_1)},
	{&int64_2, int32(int_2)},
	{&int64_10, int32(int_10)},
	{&string1, int32(int_1)},
	{&string2, int32(int_2)},
	{&string10, int32(int_10)},
	{&intFloat, int32(int_0)},
	{&stringFloat, int32(int_0)},
	{&boolTrue, int32(int_1)},
	{&boolFalse, int32(int_0)},
	{&stringEmpty, int32(int_0)},
	{&stringString, int32(int_0)},
}

func TestInt32(t *testing.T) {
	for _, test := range int32Tests {
		v := Int32(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}

type uint64Test struct {
	value    any
	expected uint64
}

var uint64Tests = []uint64Test{
	{int_0, uint64(int64_0)},
	{int_1, uint64(int64_1)},
	{int_2, uint64(int64_2)},
	{int_10, uint64(int64_10)},
	{int64_0, uint64(int64_0)},
	{int64_1, uint64(int64_1)},
	{int64_2, uint64(int64_2)},
	{int64_10, uint64(int64_10)},
	{string1, uint64(int64_1)},
	{string2, uint64(int64_2)},
	{string10, uint64(int64_10)},
	{intFloat, uint64(int64_0)},
	{stringFloat, uint64(int64_0)},
	{boolTrue, uint64(int64_1)},
	{boolFalse, uint64(int64_0)},
	{stringEmpty, uint64(int64_0)},
	{stringString, uint64(int64_0)},
	{&int_0, uint64(int64_0)},
	{&int_1, uint64(int64_1)},
	{&int_2, uint64(int64_2)},
	{&int_10, uint64(int64_10)},
	{&int64_0, uint64(int64_0)},
	{&int64_1, uint64(int64_1)},
	{&int64_2, uint64(int64_2)},
	{&int64_10, uint64(int64_10)},
	{&string1, uint64(int64_1)},
	{&string2, uint64(int64_2)},
	{&string10, uint64(int64_10)},
	{&intFloat, uint64(int64_0)},
	{&stringFloat, uint64(int64_0)},
	{&boolTrue, uint64(int64_1)},
	{&boolFalse, uint64(int64_0)},
	{&stringEmpty, uint64(int64_0)},
	{&stringString, uint64(int64_0)},
}

func TestUint64(t *testing.T) {
	for _, test := range uint64Tests {
		v := Uint64(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}

type uintTest struct {
	value    any
	expected uint
}

var uintTests = []uintTest{
	{int_0, uint(int_0)},
	{int_1, uint(int_1)},
	{int_2, uint(int_2)},
	{int_10, uint(int_10)},
	{int64_0, uint(int_0)},
	{int64_1, uint(int_1)},
	{int64_2, uint(int_2)},
	{int64_10, uint(int_10)},
	{string1, uint(int_1)},
	{string2, uint(int_2)},
	{string10, uint(int_10)},
	{intFloat, uint(int_0)},
	{stringFloat, uint(int_0)},
	{boolTrue, uint(int_1)},
	{boolFalse, uint(int_0)},
	{stringEmpty, uint(int_0)},
	{stringString, uint(int_0)},
	{&int_0, uint(int_0)},
	{&int_1, uint(int_1)},
	{&int_2, uint(int_2)},
	{&int_10, uint(int_10)},
	{&int64_0, uint(int_0)},
	{&int64_1, uint(int_1)},
	{&int64_2, uint(int_2)},
	{&int64_10, uint(int_10)},
	{&string1, uint(int_1)},
	{&string2, uint(int_2)},
	{&string10, uint(int_10)},
	{&intFloat, uint(int_0)},
	{&stringFloat, uint(int_0)},
	{&boolTrue, uint(int_1)},
	{&boolFalse, uint(int_0)},
	{&stringEmpty, uint(int_0)},
	{&stringString, uint(int_0)},
}

func TestUint(t *testing.T) {
	for _, test := range uintTests {
		v := Uint(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}

type uint8Test struct {
	value    any
	expected uint8
}

var uint8Tests = []uint8Test{
	{int_0, uint8(int_0)},
	{int_1, uint8(int_1)},
	{int_2, uint8(int_2)},
	{int_10, uint8(int_10)},
	{int64_0, uint8(int_0)},
	{int64_1, uint8(int_1)},
	{int64_2, uint8(int_2)},
	{int64_10, uint8(int_10)},
	{string1, uint8(int_1)},
	{string2, uint8(int_2)},
	{string10, uint8(int_10)},
	{intFloat, uint8(int_0)},
	{stringFloat, uint8(int_0)},
	{boolTrue, uint8(int_1)},
	{boolFalse, uint8(int_0)},
	{stringEmpty, uint8(int_0)},
	{stringString, uint8(int_0)},
	{&int_0, uint8(int_0)},
	{&int_1, uint8(int_1)},
	{&int_2, uint8(int_2)},
	{&int_10, uint8(int_10)},
	{&int64_0, uint8(int_0)},
	{&int64_1, uint8(int_1)},
	{&int64_2, uint8(int_2)},
	{&int64_10, uint8(int_10)},
	{&string1, uint8(int_1)},
	{&string2, uint8(int_2)},
	{&string10, uint8(int_10)},
	{&intFloat, uint8(int_0)},
	{&stringFloat, uint8(int_0)},
	{&boolTrue, uint8(int_1)},
	{&boolFalse, uint8(int_0)},
	{&stringEmpty, uint8(int_0)},
	{&stringString, uint8(int_0)},
}

func TestUint8(t *testing.T) {
	for _, test := range uint8Tests {
		v := Uint8(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}

type uint16Test struct {
	value    any
	expected uint16
}

var uint16Tests = []uint16Test{
	{int_0, uint16(int_0)},
	{int_1, uint16(int_1)},
	{int_2, uint16(int_2)},
	{int_10, uint16(int_10)},
	{int64_0, uint16(int_0)},
	{int64_1, uint16(int_1)},
	{int64_2, uint16(int_2)},
	{int64_10, uint16(int_10)},
	{string1, uint16(int_1)},
	{string2, uint16(int_2)},
	{string10, uint16(int_10)},
	{intFloat, uint16(int_0)},
	{stringFloat, uint16(int_0)},
	{boolTrue, uint16(int_1)},
	{boolFalse, uint16(int_0)},
	{stringEmpty, uint16(int_0)},
	{stringString, uint16(int_0)},
	{&int_0, uint16(int_0)},
	{&int_1, uint16(int_1)},
	{&int_2, uint16(int_2)},
	{&int_10, uint16(int_10)},
	{&int64_0, uint16(int_0)},
	{&int64_1, uint16(int_1)},
	{&int64_2, uint16(int_2)},
	{&int64_10, uint16(int_10)},
	{&string1, uint16(int_1)},
	{&string2, uint16(int_2)},
	{&string10, uint16(int_10)},
	{&intFloat, uint16(int_0)},
	{&stringFloat, uint16(int_0)},
	{&boolTrue, uint16(int_1)},
	{&boolFalse, uint16(int_0)},
	{&stringEmpty, uint16(int_0)},
	{&stringString, uint16(int_0)},
}

func TestUint16(t *testing.T) {
	for _, test := range uint16Tests {
		v := Uint16(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}

type uint32Test struct {
	value    any
	expected uint32
}

var uint32Tests = []uint32Test{
	{int_0, uint32(int_0)},
	{int_1, uint32(int_1)},
	{int_2, uint32(int_2)},
	{int_10, uint32(int_10)},
	{int64_0, uint32(int_0)},
	{int64_1, uint32(int_1)},
	{int64_2, uint32(int_2)},
	{int64_10, uint32(int_10)},
	{string1, uint32(int_1)},
	{string2, uint32(int_2)},
	{string10, uint32(int_10)},
	{intFloat, uint32(int_0)},
	{stringFloat, uint32(int_0)},
	{boolTrue, uint32(int_1)},
	{boolFalse, uint32(int_0)},
	{stringEmpty, uint32(int_0)},
	{stringString, uint32(int_0)},
	{&int_0, uint32(int_0)},
	{&int_1, uint32(int_1)},
	{&int_2, uint32(int_2)},
	{&int_10, uint32(int_10)},
	{&int64_0, uint32(int_0)},
	{&int64_1, uint32(int_1)},
	{&int64_2, uint32(int_2)},
	{&int64_10, uint32(int_10)},
	{&string1, uint32(int_1)},
	{&string2, uint32(int_2)},
	{&string10, uint32(int_10)},
	{&intFloat, uint32(int_0)},
	{&stringFloat, uint32(int_0)},
	{&boolTrue, uint32(int_1)},
	{&boolFalse, uint32(int_0)},
	{&stringEmpty, uint32(int_0)},
	{&stringString, uint32(int_0)},
}

func TestUint32(t *testing.T) {
	for _, test := range uint32Tests {
		v := Uint32(test.value)

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}
