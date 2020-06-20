package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	input := []byte{0x57, 0x32, 0x47, 0x47, 0x39, 0x38, 0x35, 0x30}
	expected := []byte{0x35, 0x37, 0x33, 0x32, 0x34, 0x37, 0x34, 0x37, 0x33, 0x39, 0x33, 0x38, 0x33, 0x35, 0x33, 0x30}
	result := Encode(input)
	assert.Equal(t, expected, result)
}

func TestToUInt16(t *testing.T) {
	input := []byte{0xAF, 0xEF}
	expected := []uint16{0xAFEF}
	result := ToUInt16(input)
	assert.Equal(t, expected, result)
}

func TestFromUInt16(t *testing.T) {
	input := []uint16{0xAFEF}
	expected := []byte{0xAF, 0xEF}
	result := FromUInt16(input)
	assert.Equal(t, expected, result)
}
