package slices

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
