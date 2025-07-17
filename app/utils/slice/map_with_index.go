package sliceutil

func MapWithIndex[T, M any](t []T, f func(int, T) M) []M {
	m := make([]M, len(t))
	for i, e := range t {
		r := f(i, e)
		m[i] = r
	}
	return m
}
