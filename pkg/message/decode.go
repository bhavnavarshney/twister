package message

import (
	"encoding/hex"
	"fmt"
)

func Decode(input []byte) ([]byte, error) {
	var hexStr string
	// Convert to string of Ascii char
	for i := range input {
		hexStr += fmt.Sprintf("%c", input[i])
	}
	// Convert from Ascii string to hex
	newBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode input %X with error: %w", input, err)
	}
	return newBytes, nil
}
