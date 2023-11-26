package slices

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
