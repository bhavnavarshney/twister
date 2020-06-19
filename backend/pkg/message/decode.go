package message

import (
	"encoding/hex"
	"fmt"
)

func Decode(input []byte) ([]byte, error) {
	var hexStr string
	for i := range input {
		hexStr += fmt.Sprintf("%c", input[i])
	}
	newBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return newBytes, nil
}
