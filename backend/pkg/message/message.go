// Package Message contains functions which handle the building and receiving of messages to the drill.

package message

type DrillID struct {
	header   byte
	message  []byte
	checksum []byte
}

func (id *DrillID) Marshal() []byte {
	return append([]byte{id.header}, id.message...)
}
