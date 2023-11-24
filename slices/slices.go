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

func RemoveKey[V any](sl *[]V, k int) {
	if k >= len(*sl) {
		return
	}

	var n V
	a := *sl
	copy(a[k:], a[k+1:])
	a[len(a)-1] = n
	a = a[:len(a)-1]
	*sl = a
}

func RemoveValue[V comparable](sl *[]V, v V) {
con:
	for {
		for slk, slv := range *sl {
			if v == slv {
				RemoveKey(sl, slk)
				continue con
			}
		}

		break
	}
}
