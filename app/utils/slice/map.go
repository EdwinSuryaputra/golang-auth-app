package sliceutil

func Map[T1 any, T2 any](entries []T1, fn func(T1) T2) []T2 {
	result := make([]T2, len(entries))

	for i, entry := range entries {
		result[i] = fn(entry)
	}

	return result
}
