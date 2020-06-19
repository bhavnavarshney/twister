package message

import (
	"strconv"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/nibble"
)

func Encode(input []byte) ([]byte, error) {
	nibbled := nibble.BreakSlice(input)
	var converted string

	for i := range nibbled {
		converted += strconv.FormatInt(int64(nibbled[i]), 16)
	}
	return []byte(converted), nil
}

func ToUInt16(input []byte) []uint16 {
	assembledBytes := make([]uint16, len(input)/2)
	for i := range assembledBytes {
		assembledBytes[i] = uint16(input[i])<<8 | uint16(input[i+1])
	}
	return assembledBytes
}

func FromUInt16(input []uint16) []byte {
	disassembledBytes := make([]byte, len(input)*2)
	for i := range input {
		disassembledBytes[2*i] = byte(input[i] >> 8)
		disassembledBytes[2*i+1] = byte(input[i] & 0xFF)
	}
	return disassembledBytes
}
