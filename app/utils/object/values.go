package objectutil

func Values[T1 comparable, T2 any](entries map[T1]T2) []T2 {
	result := make([]T2, 0, len(entries))

	for _, value := range entries {
		result = append(result, value)
	}

	return result
}
