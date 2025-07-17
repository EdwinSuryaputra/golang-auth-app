package sliceutil

func AssociateBy[T1 any, T2 comparable](entries []T1, fn func(T1) T2) map[T2]T1 {
	result := make(map[T2]T1, len(entries))

	for _, entry := range entries {
		result[fn(entry)] = entry
	}

	return result
}
