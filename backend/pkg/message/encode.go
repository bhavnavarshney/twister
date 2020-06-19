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
