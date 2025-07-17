package sliceutil

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {
	t.Run("Distinct - String", func(t *testing.T) {
		arr := []string{"1", "2", "3", "1", "2", "3", "1", "2", "3", "1", "2", "3"}
		expected := []string{"1", "2", "3"}

		actual := Distinct(arr)

		sort.Strings(expected)
		sort.Strings(actual)

		assert.NotNil(t, actual)
		assert.EqualValues(t, expected, actual)
	})

	t.Run("Distinct - Int", func(t *testing.T) {
		arr := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
		expected := []int{1, 2, 3}

		actual := Distinct(arr)

		sort.Ints(expected)
		sort.Ints(actual)

		assert.NotNil(t, actual)
		assert.EqualValues(t, expected, actual)
	})

	t.Run("Distinct - Float", func(t *testing.T) {
		arr := []float64{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3}
		expected := []float64{1, 2, 3}

		actual := Distinct(arr)

		sort.Float64s(expected)
		sort.Float64s(actual)

		assert.NotNil(t, actual)
		assert.EqualValues(t, expected, actual)
	})
}
