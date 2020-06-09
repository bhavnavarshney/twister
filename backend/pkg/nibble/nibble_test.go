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
