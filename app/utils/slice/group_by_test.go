package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	type str struct {
		String string
		Int    int
		Float  float64
	}

	entries := []str{
		{"value_1", 1, 1.0},
		{"value_2", 2, 2.0},
		{"value_3", 3, 3.0},
		{"value_4", 3, 4.0},
		{"value_5", 2, 5.0},
		{"value_6", 1, 6.0},
		{"value_7", 1, 6.0},
	}

	t.Run("Default", func(t *testing.T) {
		actual := GroupBy(entries, func(e str) int { return e.Int })
		expected := map[int][]str{
			1: {{"value_1", 1, 1.0}, {"value_6", 1, 6.0}, {"value_7", 1, 6.0}},
			2: {{"value_2", 2, 2.0}, {"value_5", 2, 5.0}},
			3: {{"value_3", 3, 3.0}, {"value_4", 3, 4.0}},
		}
		assert.EqualValues(t, expected, actual)
	})
}
