package ints

func Sum[V int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](a ...V) V {
	var s V

	for _, v := range a {
		s += v
	}

	return s
}
