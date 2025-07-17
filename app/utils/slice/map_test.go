package sliceutil

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	type scenario[T any, S any] struct {
		description string
		input       []T
		expected    []S
	}

	for _, scenario := range []scenario[int, string]{
		{
			description: "Normal Case",
			input:       []int{1, 2, 3, 4, 5},
			expected:    []string{"1", "2", "3", "4", "5"},
		},
		{
			description: "Zero Case",
			input:       []int{},
			expected:    []string{},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			actual := Map(scenario.input, func(t int) string {
				return strconv.Itoa(t)
			})
			expected := scenario.expected
			assert.Equal(t, expected, actual)
		})
	}
}
