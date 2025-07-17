package sliceutil

func RemoveEmptyStrings(slice []string) []string {
	result := slice[:0] // Use the same underlying array
	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
