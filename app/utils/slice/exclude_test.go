package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExclude(t *testing.T) {
	t.Run("Exclude - slice of strings", func(t *testing.T) {
		firstSlice := []string{"test1", "test2", "test3", "test4", "test5"}
		secondSlice := []string{"test1", "test2", "test3"}
		expected := []string{"test4", "test5"}

		result := Exclude(firstSlice, secondSlice)
		assert.Equal(t, expected, result)
	})

	t.Run("Exclude - slices of difference order", func(t *testing.T) {
		firstSlice := []string{"test1", "test2", "test3", "test4"}
		secondSlice := []string{"test4", "test1", "test2", "test3"}
		expected := []string{}

		result := Exclude(firstSlice, secondSlice)
		assert.Equal(t, expected, result)
	})

	t.Run("Exclude - slices of struct", func(t *testing.T) {
		type ExampleStruct struct {
			Name string
			Age  int
		}

		firstSlice := []ExampleStruct{
			{
				Name: "Si Ganteng",
				Age:  19,
			},
			{
				Name: "Si Pintar",
				Age:  17,
			},
			{
				Name: "Si Belom Kaya",
				Age:  29,
			},
		}

		secondSlice := []ExampleStruct{
			{
				Name: "Si Belom Kaya",
				Age:  29,
			},
		}

		expected := []ExampleStruct{
			{
				Name: "Si Ganteng",
				Age:  19,
			},
			{
				Name: "Si Pintar",
				Age:  17,
			},
		}

		result := Exclude(firstSlice, secondSlice)
		assert.Equal(t, expected, result)
	})

	t.Run("Exclude - slices of number", func(t *testing.T) {
		firstSlice := []float64{1999956.0, 3996.0, 9999999999999999.0000000}
		secondSlice := []float64{1999956.0}
		expected := []float64{3996.0, 9999999999999999.0000000}

		result := Exclude(firstSlice, secondSlice)
		assert.Equal(t, expected, result)
	})
}
