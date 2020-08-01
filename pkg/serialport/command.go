package serialport

import (
	"time"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
)

func MakeCommand(dataInfo byte, responseLen int) *Command {
	return &Command{
		header:      0x21,
		dataInfo:    dataInfo,
		checkSum:    NibbledChecksum([]byte{dataInfo}),
		responseLen: responseLen,
	}
}

func NibbledChecksum(dataInfo []byte) [2]byte {
	var checksum = message.Checksum(dataInfo)
	var encodedChecksum [2]byte
	copy(encodedChecksum[:], message.Encode([]byte{checksum}))
	return encodedChecksum
}

type Command struct {
	header      byte
	dataInfo    byte
	checkSum    [2]byte
	responseLen int
	response    []byte
}

func (k *Command) Marshal() []byte {
	return []byte{k.header, k.dataInfo, k.checkSum[0], k.checkSum[1]}
}

func (k *Command) Response() []byte {
	return k.response
}

func (k *Command) ResponseLen() int {
	return k.responseLen
}

func (k *Command) Retry() int {
	return 3
}

func (k *Command) Timeout() time.Duration {
	return time.Second
}
