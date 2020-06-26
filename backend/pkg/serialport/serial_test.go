// +build unit

package serialport

import (
	"testing"

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
