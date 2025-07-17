package publicfacingutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublicFacing(t *testing.T) {
	ids := []int32{1, 5}
	expected := []string{"2wxkBEJ3", "gnEBVE7A"}
	for idx, id := range ids {
		encoded, _ := Encode(id)
		assert.Equal(t, expected[idx], encoded)
	}

	encodedIds := []string{"BYQ7zqdv", "2wxkBEJ3"}
	expectedDecode := []int32{99, 1}
	for idx, id := range encodedIds {
		decoded, _ := Decode(id)
		assert.Equal(t, expectedDecode[idx], decoded)
	}
}
