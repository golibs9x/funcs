package string

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func String(s any) (v string) {
	e := reflect.TypeOf(s)
	
	if nil == e {
		v = "null"

		return
	}

	if reflect.Ptr == reflect.TypeOf(s).Kind() {
		e = reflect.TypeOf(s).Elem()
		sr := reflect.ValueOf(s).Elem()
		s = sr.Interface()
	}

	switch e.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = strconv.FormatInt(reflect.ValueOf(s).Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v = strconv.FormatUint(reflect.ValueOf(s).Uint(), 10)
	case reflect.Float32, reflect.Float64:
		v = strconv.FormatFloat(reflect.ValueOf(s).Float(), 'g', -1, 64)
	case reflect.String:
		v = s.(string)
	case reflect.Bool:
		v = strconv.FormatBool(s.(bool))
	case reflect.Array, reflect.Slice, reflect.Struct, reflect.Interface, reflect.Map:
		b, _ := json.Marshal(s)
		v = string(b)
	default:
		v = ""
	}

	return
}
