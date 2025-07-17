package objectutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjectValues(t *testing.T) {
	t.Run("Object - get values map of strings", func(t *testing.T) {
		inputMap := map[string]int{
			"key1": 100,
			"key2": 200,
			"key3": 300,
		}

		expected := []int{100, 200, 300}

		results := Values(inputMap)
		for _, result := range results {
			assert.Contains(t, expected, result)
		}
	})

	t.Run("Object - get values map of slices", func(t *testing.T) {
		inputMap := map[string][]int{
			"key4": {5, 6, 7},
			"key3": {4},
			"key1": {100, 200, 300},
			"key5": {8, 9, 10},
			"key2": {100},
		}

		expected := [][]int{{5, 6, 7}, {4}, {100, 200, 300}, {8, 9, 10}, {100}}

		results := Values(inputMap)
		for _, result := range results {
			assert.Contains(t, expected, result)
		}
	})

	t.Run("Object - get values map of structs", func(t *testing.T) {
		type ExampleStruct struct {
			Name string
			Age  int
		}

		inputMap := map[string]*ExampleStruct{
			"key3": nil,
			"key1": {
				Name: "Edwin",
				Age:  15,
			},
			"key2": {
				Name: "Ben",
				Age:  21,
			},
		}

		expected := []*ExampleStruct{
			nil,
			{
				Name: "Edwin",
				Age:  15,
			}, {
				Name: "Ben",
				Age:  21,
			},
		}

		results := Values(inputMap)
		for _, result := range results {
			assert.Contains(t, expected, result)
		}
	})
}
