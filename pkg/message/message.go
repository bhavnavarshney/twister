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
	KeepAlive           = 0x07
	CalibratedOffsetMsg = 0x11
	CurrentOffsetMsg    = 0x16
	BulkParamSendMsg    = 0x14
	BulkParamReceiveMsg = 0x15
)

const (
	DrillIDMsgLen          = 16
	CurrentOffsetMsgLen    = 8
	CalibratedOffsetMsgLen = 8
	DrillTypeMsgLen        = 20
	BulkParamReceiveMsgLen = 24*4*2 + 4
)

func MakeTorqueParam(id byte, message [4]byte) *TorqueParam {
	return &TorqueParam{
		header:   HeaderByte,
		dataInfo: SingleParamMsg,
		id:       id,
		message:  message,
	}
}

type TorqueParam struct {
	header   byte
	dataInfo byte
	id       byte
	message  [4]byte
	checksum byte
}

func (td *TorqueParam) Marshal() []byte {
	dataInfoAdded := append([]byte{td.dataInfo, td.id}, td.message[:]...)
	td.checksum = Checksum(dataInfoAdded)
	headers := []byte{td.header, td.dataInfo}
	payload := append([]byte{td.id}, td.message[:]...)
	encodedData := Encode(append(payload, td.checksum))
	return append(headers, encodedData...)
}

func (td *TorqueParam) Response() []byte {
	return []byte{OkStatus}
}

func (td *TorqueParam) ResponseLen() int {
	return 1
}

func (td *TorqueParam) Retry() int {
	return 1
}

func (td *TorqueParam) Timeout() time.Duration {
	return time.Second
}

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
	headers := []byte{td.header, td.dataInfo}
	encodedData := Encode(append(td.message[:], td.checksum))
	return append(headers, encodedData...)
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
	data, err := Decode(input[2:])
	if err != nil {
		return err
	}
	copy(td.message[:], data)
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

func (t *TorqueData) ToProfile() []byte {
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

type Offset struct {
	header   byte
	dataInfo byte
	message  [2]byte
	checksum byte
}

func (id *Offset) Unmarshal(input []byte) error {
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

func (id *Offset) isValidChecksum() error {
	calc := Checksum(append([]byte{id.dataInfo}, id.message[:]...))
	if calc != id.checksum {
		return fmt.Errorf("checksum mismatch, expected %x, received %x", calc, id.checksum)
	}
	return nil
}

func (id *Offset) ToUInt16() uint16 {
	return ToUInt16(id.message[:])[0]
}
