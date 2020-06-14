// Package Message contains functions which handle the building and receiving of messages to the drill.

package message

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
	checksum []byte
}

type DrillID struct {
	header   byte
	dataInfo byte
	message  [6]byte
	checksum []byte
}

func (id *DrillID) Marshal() []byte {
	return append([]byte{id.header}, id.message[:]...)
}

type Command struct {
	command byte
	retries int
	timeout int
}
