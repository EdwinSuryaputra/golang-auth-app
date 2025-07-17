package sliceutil

func Distinct[T comparable](entries []T) []T {
	length := len(entries)
	isExistMap := make(map[T]bool, length)
	result := make([]T, 0, length)

	for _, entry := range entries {
		if _, ok := isExistMap[entry]; !ok {
			isExistMap[entry] = true
			result = append(result, entry)
		}
	}

	return result
}
