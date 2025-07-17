package objectutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjectKeys(t *testing.T) {
	t.Run("Object - get keys map of strings", func(t *testing.T) {
		inputMap := map[string]int{
			"key1": 100,
			"key2": 200,
			"key3": 300,
		}

		expected := []string{"key1", "key2", "key3"}

		results := Keys(inputMap)
		for _, result := range results {
			assert.Contains(t, expected, result)
		}
	})

	t.Run("Object - get keys map of slices", func(t *testing.T) {
		inputMap := map[string][]int{
			"key4": {5, 6, 7},
			"key3": {4},
			"key1": {100, 200, 300},
			"key5": {8, 9, 10},
			"key2": {100},
		}

		expected := []string{"key1", "key2", "key3", "key4", "key5"}

		results := Keys(inputMap)
		for _, result := range results {
			assert.Contains(t, expected, result)
		}
	})

	t.Run("Object - get keys map of structs", func(t *testing.T) {
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

		expected := []string{"key1", "key2", "key3"}

		results := Keys(inputMap)
		for _, result := range results {
			assert.Contains(t, expected, result)
		}
	})
}
