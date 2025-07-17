package sliceutil

func Filter[T any](entries []T, fn func(T) bool) []T {
	result := make([]T, 0, len(entries))

	for _, entry := range entries {
		if fn(entry) {
			result = append(result, entry)
		}
	}

	return result
}
