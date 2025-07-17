package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapWithIndex(t *testing.T) {
	actual := MapWithIndex([]int{5, 4, 3, 2, 1}, func(i int, t int) int {
		return i + t
	})
	expected := []int{5, 5, 5, 5, 5}
	assert.Equal(t, expected, actual)
}
