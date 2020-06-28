// +build unit

package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	input := []byte{0x35, 0x37, 0x33, 0x32, 0x34, 0x37, 0x34, 0x37, 0x33, 0x39, 0x33, 0x38, 0x33, 0x35, 0x33, 0x30}
	expected := []byte{0x57, 0x32, 0x47, 0x47, 0x39, 0x38, 0x35, 0x30}
	result, err := Decode(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestDecodeWithZeros(t *testing.T) {
	input := make([]byte, 6)
	var expected []byte
	result, err := Decode(input)
	assert.EqualError(t, err, "failed to decode input 000000000000 with error: encoding/hex: invalid byte: U+0000")
	assert.Equal(t, expected, result)
}
