package objectutil

func Keys[T1 comparable, T2 any](entries map[T1]T2) []T1 {
	result := make([]T1, 0, len(entries))

	for key := range entries {
		result = append(result, key)
	}

	return result
}
