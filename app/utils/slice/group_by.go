package sliceutil

func GroupBy[T1 any, T2 comparable](entries []T1, fn func(T1) T2) map[T2][]T1 {
	result := make(map[T2][]T1, len(entries)/2)

	for _, entry := range entries {
		key := fn(entry)
		if _, ok := result[key]; !ok {
			result[key] = []T1{}
		}

		result[key] = append(result[key], entry)
	}

	return result
}
