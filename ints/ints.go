package ints

import (
	"reflect"
	"strconv"
)

func Int64(i any) (v int64) {
	e := reflect.TypeOf(i)

	if nil == e {
		v = 0

		return
	}

	if reflect.Ptr == reflect.TypeOf(i).Kind() {
		e = reflect.TypeOf(i).Elem()
		sr := reflect.ValueOf(i).Elem()
		i = sr.Interface()
	}

	switch e.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, _ = strconv.ParseInt(strconv.FormatInt(reflect.ValueOf(i).Int(), 10), 10, 64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, _ = strconv.ParseInt(strconv.FormatUint(reflect.ValueOf(i).Uint(), 10), 10, 64)
	//case reflect.Float32, reflect.Float64:
	//	v = strconv.FormatFloat(reflect.ValueOf(s).Float(), 'g', -1, 64)
	case reflect.String:
		v, _ = strconv.ParseInt(i.(string), 10, 64)
	case reflect.Bool:
		if true == i.(bool) {
			v = 1
		}
	default:
		v = 0
	}

	return
}

func Int(i any) (v int) {
	return int(Int64(i))
}

func Int8(i any) (v int8) {
	return int8(Int64(i))
}

func Int16(i any) (v int16) {
	return int16(Int64(i))
}

func Int32(i any) (v int32) {
	return int32(Int64(i))
}

func Uint64(i any) (v uint64) {
	e := reflect.TypeOf(i)

	if nil == e {
		v = 0

		return
	}

	if reflect.Ptr == reflect.TypeOf(i).Kind() {
		e = reflect.TypeOf(i).Elem()
		sr := reflect.ValueOf(i).Elem()
		i = sr.Interface()
	}

	switch e.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		strconv.FormatInt(reflect.ValueOf(i).Int(), 10)
		v, _ = strconv.ParseUint(strconv.FormatInt(reflect.ValueOf(i).Int(), 10), 10, 64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, _ = strconv.ParseUint(strconv.FormatUint(reflect.ValueOf(i).Uint(), 10), 10, 64)
	//case reflect.Float32, reflect.Float64:
	//	v = strconv.FormatFloat(reflect.ValueOf(s).Float(), 'g', -1, 64)
	case reflect.String:
		v, _ = strconv.ParseUint(i.(string), 10, 64)
	case reflect.Bool:
		if true == i.(bool) {
			v = 1
		}
	default:
		v = 0
	}

	return
}

func Uint(i any) (v uint) {
	return uint(Uint64(i))
}

func Uint8(i any) (v uint8) {
	return uint8(Uint64(i))
}

func Uint16(i any) (v uint16) {
	return uint16(Uint64(i))
}

func Uint32(i any) (v uint32) {
	return uint32(Uint64(i))
}
