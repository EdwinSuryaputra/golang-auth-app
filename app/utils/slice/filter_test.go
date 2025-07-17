package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	type scenario[T any, S any] struct {
		description string
		input       []T
		expected    []S
	}

	for _, scenario := range []scenario[int, int]{
		{
			description: "Normal Case",
			input:       []int{1, 2, 3, 4, 5},
			expected:    []int{3, 4, 5},
		},
		{
			description: "Zero Case",
			input:       []int{},
			expected:    []int{},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			actual := Filter(scenario.input, func(t int) bool {
				return t > 2
			})

			expected := scenario.expected
			assert.Equal(t, expected, actual)
		})
	}
}
