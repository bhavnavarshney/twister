package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	input := []byte{15, 7}
	expected := byte(233)
	calcChecksum := Checksum(input)
	assert.Equal(t, expected, calcChecksum)
}

func TestChecksumCommand(t *testing.T) {
	input := []byte{0x05}
	expected := byte(0xFA)
	calcChecksum := Checksum(input)
	assert.Equal(t, expected, calcChecksum)
}
func TestChecksumLong(t *testing.T) {
	input := []byte{0x04, 0x57, 0x32, 0x47, 0x47, 0x39, 0x38, 0x35, 0x30}
	expected := byte(0xE9)
	calcChecksum := Checksum(input)
	assert.Equal(t, expected, calcChecksum)
}
