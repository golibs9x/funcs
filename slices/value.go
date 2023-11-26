package slices

import (
	"reflect"
)

func Value[V any](sl *[]V, k int, def V) (v V) {
	sv := *sl
	rv := reflect.ValueOf(sv)

	if rv.Len() <= k {
		v = def
	} else {
		v = sv[k]
	}

	return
}
