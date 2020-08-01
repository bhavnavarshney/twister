// +build unit

package serialport

import (
	"errors"
	"testing"
	"time"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/tarm/serial"
)

func TestSendMessage(t *testing.T) {
	log, _ := test.NewNullLogger()
	message := MakeCommand(message.SingleParamMsg, 1)
	c := &serial.Config{Name: "COM3", Baud: 9600}
	p, err := MakeFakePort(c)
	assert.NoError(t, err)
	d := MakeDriver(p, log)

	err = d.SendMessage(message)
	assert.NoError(t, err)
}

type TestPort struct {
}

func (t *TestPort) Write(out []byte) (int, error) {
	return len(out), nil
}

func (t *TestPort) Read(b []byte) (int, error) {
	time.Sleep(3 * time.Second)
	return 0, nil
}

func (t *TestPort) Close() error {
	return nil
}
func TestSendMessageWithTimeout(t *testing.T) {
	log, _ := test.NewNullLogger()
	message := MakeCommand(message.SingleParamMsg, 1)
	p := &TestPort{}
	d := MakeDriver(p, log)
	err := d.SendMessage(message)
	assert.Error(t, errors.New("timeout waiting for response"), err)
}

func TestSendCommandWithTimeout(t *testing.T) {
	log, _ := test.NewNullLogger()
	message := MakeCommand(message.SingleParamMsg, 1)
	p := &TestPort{}
	d := MakeDriver(p, log)
	_, err := d.SendCommand(message)
	assert.Error(t, errors.New("timeout waiting for response"), err)
}
