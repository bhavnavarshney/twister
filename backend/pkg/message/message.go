// Package Message contains functions which handle the building and receiving of messages to the drill.

package message

import (
	"fmt"
)

const (
	drillType    = 0x04
	drillID      = 0x05
	singleParam  = 0x06
	torqueOffset = 0x11
	readOffset   = 0x13
	bulkParam    = 0x14
)

type TorqueData struct {
	header   byte
	dataInfo byte
	message  [12]uint16
	checksum byte
}

func (td *TorqueData) Marshal() []byte {
	return []byte{}
}

func (td *TorqueData) Unmarshal(input []byte) error {
	td.header = input[0]
	td.dataInfo = input[1]
	td.checksum = input[len(input)-1]
	data, err := Decode(input[2:])
	for i := range td.message {
		td.message[i] = uint16(data[i])<<8 | uint16(data[i+1])
	}
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
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
