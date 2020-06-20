// Package Message contains functions which handle the building and receiving of messages to the drill.

package message

import (
	"fmt"
	"time"
)

const (
	HeaderByte          = 0x21
	OkStatus            = 0x47
	DrillTypeMsg        = 0x04
	DrillIDMsg          = 0x05
	SingleParamMsg      = 0x06
	TorqueOffsetMsg     = 0x11
	ReadOffsetMsg       = 0x13
	BulkParamSendMsg    = 0x14
	BulkParamReceiveMsg = 0x15
)

func MakeTorqueData(message [24 * 4]byte) *TorqueData {
	return &TorqueData{
		header:   HeaderByte,
		dataInfo: BulkParamSendMsg,
		message:  message,
	}

}

type TorqueData struct {
	header   byte
	dataInfo byte
	message  [24 * 4]byte
	checksum byte
}

func (td *TorqueData) Marshal() []byte {
	dataInfoAdded := append([]byte{td.dataInfo}, td.message[:]...)
	td.checksum = Checksum(dataInfoAdded)
	encodedData := Encode(append(dataInfoAdded, td.checksum))
	return encodedData
}

func (td *TorqueData) Retry() int {
	return 1
}

func (td *TorqueData) Timeout() time.Duration {
	return time.Second
}

func (td *TorqueData) Response() []byte {
	return []byte{OkStatus}
}

func (td *TorqueData) ResponseLen() int {
	return 1
}

func (td *TorqueData) Unmarshal(input []byte) error {
	td.header = input[0]
	td.dataInfo = input[1]
	td.checksum = input[len(input)-1]
	data, err := Decode(input[2:])
	if err != nil {
		return err
	}
	copy(td.message[:], data)
	fmt.Println(data)
	td.checksum = data[len(data)-1]
	return td.isValidChecksum()
}

func (t *TorqueData) isValidChecksum() error {
	calc := Checksum(append([]byte{t.dataInfo}, t.message[:]...))
	if calc != t.checksum {
		return fmt.Errorf("checksum mismatch, expected %x, received %x", calc, t.checksum)
	}
	return nil
}

func (t *TorqueData) ToByte() []byte {
	return t.message[:]
}

type DrillType struct {
	header   byte
	dataInfo byte
	message  [8]byte
	checksum byte
}

func (t *DrillType) Unmarshal(input []byte) error {
	t.header = input[0]
	t.dataInfo = input[1]
	data, err := Decode(input[2:])
	if err != nil {
		return err
	}
	copy(t.message[:], data)
	t.checksum = data[len(data)-1]
	return t.isValidChecksum()
}

func (t *DrillType) isValidChecksum() error {
	calc := Checksum(append([]byte{t.dataInfo}, t.message[:]...))
	if calc != t.checksum {
		return fmt.Errorf("checksum mismatch, expected %x, received %x", calc, t.checksum)
	}
	return nil
}

func (t *DrillType) ToByte() []byte {
	return t.message[:]
}

func (t *DrillType) ToString() string {
	var idStr string
	for i := range t.message {
		idStr += fmt.Sprintf("%c", t.message[i])
	}
	return idStr
}

type DrillID struct {
	header   byte
	dataInfo byte
	message  [6]byte
	checksum byte
}

func (id *DrillID) Unmarshal(input []byte) error {
	id.header = input[0]
	id.dataInfo = input[1]
	decodedData, err := Decode(input[2:])
	if err != nil {
		return err
	}
	copy(id.message[:], decodedData)
	id.checksum = decodedData[len(decodedData)-1]
	fmt.Printf("%X", id.checksum)
	return id.isValidChecksum()
}

func (id *DrillID) ToByte() []byte {
	return id.message[:]
}

func (id *DrillID) ToString() string {
	var idStr string
	for i := range id.message {
		idStr += fmt.Sprintf("%X", id.message[i])
	}
	return idStr
}

func Checksum(input []byte) byte {
	byteSum := byte(0xFF)
	for i := range input {
		byteSum = byte(byteSum - input[i])
	}
	return byteSum
}

func (id *DrillID) isValidChecksum() error {
	calc := Checksum(append([]byte{id.dataInfo}, id.message[:]...))
	if calc != id.checksum {
		return fmt.Errorf("checksum mismatch, expected %x, received %x", calc, id.checksum)
	}
	return nil
}

func (id *DrillID) Marshal() []byte {
	return append([]byte{id.header}, id.message[:]...)
}

type Command struct {
	command byte
	retries int
	timeout int
}
