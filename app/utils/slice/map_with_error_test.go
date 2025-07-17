package sliceutil

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapWithError(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		actual, err := MapWithError([]string{"1", "2", "3"}, func(t string) (int, error) {
			return strconv.Atoi(t)
		})
		expected := []int{1, 2, 3}
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Error", func(t *testing.T) {
		actual, err := MapWithError([]string{"1", "2", "a"}, func(t string) (int, error) {
			return strconv.Atoi(t)
		})
		assert.Error(t, err)
		assert.Equal(t, 0, len(actual))
	})
}
