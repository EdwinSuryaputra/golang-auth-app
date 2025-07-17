package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8}

	assert.Equal(t, true, Contains(list, 1))
	assert.Equal(t, false, Contains(list, 9))
}
