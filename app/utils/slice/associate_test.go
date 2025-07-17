package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssociate(t *testing.T) {
	type str struct {
		String string
		Int    int
		Float  float64
	}

	entries := []str{
		{"1", 1, 1.0},
		{"2", 2, 2.0},
		{"3", 3, 3.0},
	}

	t.Run("Associate - Default", func(t *testing.T) {
		expected := map[string]int{
			"1": 1,
			"2": 2,
			"3": 3,
		}

		actual := Associate(entries, func(s str) (string, int) { return s.String, s.Int })
		assert.NotNil(t, actual)
		assert.EqualValues(t, expected, actual)
	})
}
