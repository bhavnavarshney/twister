// +build unit

package serialport

import (
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/tarm/serial"
)

func TestOpenSerialPort(t *testing.T) {
	log, _ := test.NewNullLogger()
	c := &serial.Config{Name: "COM3", Baud: 9600}
	p, err := MakeFakePort(c)
	assert.NoError(t, err)
	d := MakeDriver(p, log)

	retVal := d.SendKeepAlive()
	assert.True(t, retVal)
}
