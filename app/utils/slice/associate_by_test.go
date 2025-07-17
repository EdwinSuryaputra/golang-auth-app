package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssociateBy(t *testing.T) {
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

	t.Run("AssociateBy - String", func(t *testing.T) {
		expected := map[string]str{
			"1": {"1", 1, 1.0},
			"2": {"2", 2, 2.0},
			"3": {"3", 3, 3.0},
		}

		actual := AssociateBy(entries, func(s str) string { return s.String })
		assert.NotNil(t, actual)
		assert.EqualValues(t, expected, actual)
	})

	t.Run("AssociateBy - Int", func(t *testing.T) {
		expected := map[int]str{
			1: {"1", 1, 1.0},
			2: {"2", 2, 2.0},
			3: {"3", 3, 3.0},
		}

		actual := AssociateBy(entries, func(s str) int { return s.Int })
		assert.NotNil(t, actual)
		assert.EqualValues(t, expected, actual)
	})

	t.Run("AssociateBy - Float", func(t *testing.T) {
		expected := map[float64]str{
			1.0: {"1", 1, 1.0},
			2.0: {"2", 2, 2.0},
			3.0: {"3", 3, 3.0},
		}

		actual := AssociateBy(entries, func(s str) float64 { return s.Float })
		assert.NotNil(t, actual)
		assert.EqualValues(t, expected, actual)
	})
}
