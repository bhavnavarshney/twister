// +build unit

package serialport

import (
	"testing"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/tarm/serial"
)

// func TestOpenSerialPort(t *testing.T) {
// 	log := logrus.New()
// 	c := &serial.Config{Name: "COM3", Baud: 9600}
// 	p, err := MakeFakePort(c)
// 	assert.NoError(t, err)
// 	d := MakeDriver(p, log)
// 	retVal := d.SendKeepAlive()
// 	assert.True(t, retVal)
// }

func TestSendMessage(t *testing.T) {
	log := logrus.New()
	message := MakeCommand(0x07, 1)
	c := &serial.Config{Name: "COM3", Baud: 9600}
	p, err := MakeFakePort(c)
	assert.NoError(t, err)
	d := MakeDriver(p, log)
	err = d.SendMessage(message)
	assert.NoError(t, err)
}

func TestSendCommand(t *testing.T) {
	log := logrus.New()
	message := MakeCommand(0x07, 1)
	c := &serial.Config{Name: "COM3", Baud: 9600}
	p, err := MakeFakePort(c)
	assert.NoError(t, err)
	d := MakeDriver(p, log)
	response, err := d.SendCommand(message)
	assert.NoError(t, err)
	t.Log(response)
}

func TestSendCommandRead(t *testing.T) {
	log := logrus.New()
	c := &serial.Config{Name: "COM3", Baud: 9600}
	p, err := MakeFakePort(c)
	assert.NoError(t, err)
	d := MakeDriver(p, log)

	expected := []uint8([]byte{0x0, 0xf, 0x0, 0x1e, 0x0, 0x28, 0x0, 0x32, 0x0, 0x3c, 0x0, 0x46, 0x0, 0x50, 0x0, 0x5a, 0x0, 0x64, 0x0, 0x6e, 0x0, 0x78, 0x0, 0x82, 0x0, 0xf, 0x0, 0x1e, 0x0, 0x28, 0x0, 0x32, 0x0, 0x3c, 0x0, 0x46, 0x0, 0x50, 0x0, 0x5a, 0x0, 0x64, 0x0, 0x6e, 0x0, 0x78, 0x0, 0x82, 0x0, 0x2d, 0x0, 0x5a, 0x0, 0x78, 0x0, 0x96, 0x0, 0xb3, 0x0, 0xd0, 0x0, 0xed, 0x1, 0xa, 0x1, 0x27, 0x1, 0x44, 0x1, 0x61, 0x1, 0x7e, 0x0, 0x2d, 0x0, 0x5a, 0x0, 0x78, 0x0, 0x96, 0x0, 0xb3, 0x0, 0xd0, 0x0, 0xed, 0x1, 0xa, 0x1, 0x27, 0x1, 0x44, 0x1, 0x61, 0x1, 0x7e})
	readProfileCommand := MakeCommand(message.BulkParamReceiveMsg, message.BulkParamReceiveMsgLen)
	response, err := d.SendCommand(readProfileCommand)
	assert.NoError(t, err)

	torqueData := message.TorqueData{}
	err = torqueData.Unmarshal(response)
	assert.NoError(t, err)
	assert.Equal(t, expected, torqueData.ToByte())
}
