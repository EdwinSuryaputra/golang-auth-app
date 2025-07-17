package sliceutil

func MapWithError[T any, M any](a []T, f func(T) (M, error)) ([]M, error) {
	n := make([]M, len(a))
	for i, e := range a {
		r, err := f(e)
		if err != nil {
			return make([]M, 0), err
		}
		n[i] = r
	}
	return n, nil
}
