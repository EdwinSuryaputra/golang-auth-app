package sliceutil

func Associate[T1 any, T2 comparable, T3 comparable](entries []T1, fn func(T1) (T2, T3)) map[T2]T3 {
	result := make(map[T2]T3, len(entries))

	for _, entry := range entries {
		k, v := fn(entry)
		result[k] = v
	}

	return result
}
