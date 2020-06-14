package serialport

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/tarm/serial"
)

func TestOpenSerialPort(t *testing.T) {
	log := logrus.New()
	c := &serial.Config{Name: "COM3", Baud: 9600}
	p, err := MakeFakePort(c)
	assert.NoError(t, err)
	d := MakeSerialPortDriver(p, log)
	retVal := d.SendKeepAlive()
	assert.True(t, retVal)
}
