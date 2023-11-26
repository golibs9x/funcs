package slices

func Find[V comparable](sl *[]V, v V) (k int) {
	k = -1

	for slk, slv := range *sl {
		if v == slv {
			k = slk

			return
		}
	}

	return
}
