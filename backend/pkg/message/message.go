// Package Message contains functions which handle the building and receiving of messages to the drill.

package message

import (
	"errors"
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

type DrillType struct {
	header   byte
	dataInfo byte
	message  [8]byte
	checksum byte
}

func (id *DrillType) Unmarshal(input []byte) error {
	id.header = input[0]
	id.dataInfo = input[1]
	data, err := Decode(input[2:])
	if err != nil {
		return err
	}
	copy(id.message[:], data)
	id.checksum = input[len(input)-1]
	return nil
}

func (id *DrillType) ToByte() []byte {
	return id.message[:]
}

func (id *DrillType) ToString() string {
	var idStr string
	for i := range id.message {
		idStr += fmt.Sprintf("%c", id.message[i])
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
	data, err := Decode(input[2:])
	if err != nil {
		return err
	}
	copy(id.message[:], data)
	id.checksum = input[len(input)-1]
	return nil
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
	var byteSum byte
	for i := range input {
		byteSum += input[i]
	}
	return 0xFF - byteSum
}

func (id *DrillID) isValidChecksum() error {
	calc := Checksum(append([]byte{id.dataInfo}, id.message[:]...))
	if calc != id.checksum {
		return errors.New("Checksum mismatch")
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
