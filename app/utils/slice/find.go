package sliceutil

func Find[T any](entries []T, fn func(T) bool) *T {
	if fn == nil {
		return nil
	}

	for _, entry := range entries {
		if fn(entry) {
			return &entry
		}
	}

	return nil
}
