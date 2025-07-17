package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	t.Run("Find - slices of int", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := 4

		result := Find(input, func(num int) bool {
			return num == 4
		})

		assert.Equal(t, expected, *result)
	})

	type ExampleStruct struct {
		Name string
		Age  int
	}

	t.Run("Find - slices of struct", func(t *testing.T) {
		input := []*ExampleStruct{
			{
				Name: "Edwin",
				Age:  15,
			},
			{
				Name: "Ben",
				Age:  21,
			},
		}

		expected := &ExampleStruct{
			Name: "Edwin",
			Age:  15,
		}

		result := Find(input, func(example *ExampleStruct) bool {
			return example.Name == expected.Name && example.Age == expected.Age
		})

		assert.Equal(t, expected, *result)
	})

	t.Run("Find - not found case", func(t *testing.T) {
		input := []*ExampleStruct{
			{
				Name: "Edwin",
				Age:  15,
			},
			{
				Name: "Ben",
				Age:  21,
			},
		}

		expected := &ExampleStruct{
			Name: "Alwin",
			Age:  99,
		}

		result := Find(input, func(example *ExampleStruct) bool {
			return example.Name == expected.Name && example.Age == expected.Age
		})

		assert.Nil(t, result)
	})

	t.Run("Find - nil case", func(t *testing.T) {
		var input []*ExampleStruct = nil

		result := Find(input, func(example *ExampleStruct) bool {
			return example.Name == "Edwin" && example.Age == 15
		})

		assert.Nil(t, result)
	})

	t.Run("Find - nil callback case", func(t *testing.T) {
		input := []*ExampleStruct{
			{
				Name: "Edwin",
				Age:  15,
			},
			{
				Name: "Ben",
				Age:  21,
			},
		}

		result := Find(input, nil)
		assert.Nil(t, result)
	})
}
