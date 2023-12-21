package ints

import "testing"

type sumTest struct {
	caseType     string
	valueInt     []int
	valueInt8    []int8
	valueInt16   []int16
	valueInt32   []int32
	valueInt64   []int64
	valueUint    []uint
	valueUint8   []uint8
	valueUint16  []uint16
	valueUint32  []uint32
	valueUint64  []uint64
	valueFloat32 []float32
	valueFloat64 []float64
	expected     any
}

var sumTests = []sumTest{
	{caseType: "int", valueInt: []int{}, expected: 0},
	{caseType: "int", valueInt: []int{1, 3}, expected: 4},
	{caseType: "int", valueInt: []int{1, 3, 10}, expected: 14},
	{caseType: "int8", valueInt8: []int8{}, expected: int8(0)},
	{caseType: "int8", valueInt8: []int8{1, 3}, expected: int8(4)},
	{caseType: "int8", valueInt8: []int8{1, 3, 10}, expected: int8(14)},
	{caseType: "int16", valueInt16: []int16{}, expected: int16(0)},
	{caseType: "int16", valueInt16: []int16{1, 3}, expected: int16(4)},
	{caseType: "int16", valueInt16: []int16{1, 3, 10}, expected: int16(14)},
	{caseType: "int32", valueInt32: []int32{}, expected: int32(0)},
	{caseType: "int32", valueInt32: []int32{1, 3}, expected: int32(4)},
	{caseType: "int32", valueInt32: []int32{1, 3, 10}, expected: int32(14)},
	{caseType: "int64", valueInt64: []int64{}, expected: int64(0)},
	{caseType: "int64", valueInt64: []int64{1, 3}, expected: int64(4)},
	{caseType: "int64", valueInt64: []int64{1, 3, 10}, expected: int64(14)},
	{caseType: "uint", valueUint: []uint{}, expected: uint(0)},
	{caseType: "uint", valueUint: []uint{1, 3}, expected: uint(4)},
	{caseType: "uint", valueUint: []uint{1, 3, 10}, expected: uint(14)},
	{caseType: "uint8", valueUint8: []uint8{}, expected: uint8(0)},
	{caseType: "uint8", valueUint8: []uint8{1, 3}, expected: uint8(4)},
	{caseType: "uint8", valueUint8: []uint8{1, 3, 10}, expected: uint8(14)},
	{caseType: "uint16", valueUint16: []uint16{}, expected: uint16(0)},
	{caseType: "uint16", valueUint16: []uint16{1, 3}, expected: uint16(4)},
	{caseType: "uint16", valueUint16: []uint16{1, 3, 10}, expected: uint16(14)},
	{caseType: "uint32", valueUint32: []uint32{}, expected: uint32(0)},
	{caseType: "uint32", valueUint32: []uint32{1, 3}, expected: uint32(4)},
	{caseType: "uint32", valueUint32: []uint32{1, 3, 10}, expected: uint32(14)},
	{caseType: "uint64", valueUint64: []uint64{}, expected: uint64(0)},
	{caseType: "uint64", valueUint64: []uint64{1, 3}, expected: uint64(4)},
	{caseType: "uint64", valueUint64: []uint64{1, 3, 10}, expected: uint64(14)},
	{caseType: "float32", valueFloat32: []float32{}, expected: float32(0)},
	{caseType: "float32", valueFloat32: []float32{0.7854, 1, 3}, expected: float32(4.7854)},
	{caseType: "float32", valueFloat32: []float32{1, 3, 10.131313131235183}, expected: float32(14.131313131235183)},
	{caseType: "float64", valueFloat64: []float64{}, expected: float64(0)},
	{caseType: "float64", valueFloat64: []float64{0.7854, 1, 3}, expected: float64(4.7854)},
	{caseType: "float64", valueFloat64: []float64{1, 3, 10.131313131235183}, expected: float64(14.131313131235183)},
}

func TestSum(t *testing.T) {
	var v any

	for _, test := range sumTests {
		v = 0

		switch test.caseType {
		case "int":
			v = Sum(test.valueInt...)
		case "int8":
			v = Sum(test.valueInt8...)
		case "int16":
			v = Sum(test.valueInt16...)
		case "int32":
			v = Sum(test.valueInt32...)
		case "int64":
			v = Sum(test.valueInt64...)
		case "uint":
			v = Sum(test.valueUint...)
		case "uint8":
			v = Sum(test.valueUint8...)
		case "uint16":
			v = Sum(test.valueUint16...)
		case "uint32":
			v = Sum(test.valueUint32...)
		case "uint64":
			v = Sum(test.valueUint64...)
		case "float32":
			v = Sum(test.valueFloat32...)
		case "float64":
			v = Sum(test.valueFloat64...)
		}

		if v != test.expected {
			t.Errorf("Output %v not equal to expected %d", v, test.expected)
		}
	}
}
