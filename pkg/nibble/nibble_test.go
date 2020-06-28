// +build unit

package nibble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreak(t *testing.T) {
	high, low := Break(0xAF)
	assert.Equal(t, byte(0x0A), high)
	assert.Equal(t, byte(0x0F), low)
}

func TestAssemble(t *testing.T) {
	assembled := Assemble(byte(0x0A), byte(0x0F))
	assert.Equal(t, byte(0xAF), assembled)
}

func TestAssembleSlice(t *testing.T) {
	input := []byte{0x0F, 0x0F, 0x03, 0x0A, 0x00, 0x01}
	expected := []byte{0xFF, 0x3A, 0x01}
	assert.Equal(t, expected, AssembleSlice(input))
}

func TestBreakSlice(t *testing.T) {
	input := []byte{0xFF, 0x3A, 0x01}
	expected := []byte{0x0F, 0x0F, 0x03, 0x0A, 0x00, 0x01}
	assert.Equal(t, expected, BreakSlice(input))
}
