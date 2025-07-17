package sliceutil

func Exclude[T comparable](firstEntries []T, secondEntries []T) []T {
	secondEntriesMap := make(map[T]T)
	for _, v := range secondEntries {
		secondEntriesMap[v] = v
	}

	result := []T{}
	for _, v := range firstEntries {
		_, isExist := secondEntriesMap[v]

		if !isExist {
			result = append(result, v)
		}
	}

	return result
}
