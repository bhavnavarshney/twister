// +build integration

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/serialport"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/tarm/serial"
)

func TestSendHeader(t *testing.T) {
	log := logrus.New()
	config := &serial.Config{Name: "/dev/tty.usbserial-AC019QP9", Baud: 9600}
	p, err := serial.OpenPort((config))
	assert.NoError(t, err)
	defer p.Close()
	d := serialport.MakeSerialPortDriver(p, log)
	response := d.SendKeepAlive()
	assert.True(t, response)
}

func TestReadConfig(t *testing.T) {
	log := logrus.New()
	config := &serial.Config{Name: "/dev/tty.usbserial-AC019QP9", Baud: 9600}
	p, err := serial.OpenPort((config))
	assert.NoError(t, err)
	defer p.Close()
	d := serialport.MakeSerialPortDriver(p, log)

	bulkParamCommand := serialport.MakeCommand(message.BulkParamReceiveMsg, 20)
	response, err := d.SendCommand(bulkParamCommand)
	t.Logf("%X", response)
	assert.NoError(t, err)

	drillType := message.DrillType{}
	err = drillType.Unmarshal(response)
	assert.NoError(t, err)
	log.Printf("Response Hex: %X", drillType.ToByte())
	log.Printf("Response ASCII: %s", drillType.ToString())
}

func TestSerial(t *testing.T) {
	config := &serial.Config{Name: "/dev/tty.usbserial-AC019QP9", Baud: 9600}
	p, err := serial.OpenPort(config)
	defer p.Close()
	if err != nil {
		fmt.Println(err)
	}
	_, err = p.Write([]byte{0x21, 0x15, 0x0E, 0x0A})
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 15)
	time.Sleep(1 * time.Second)
	p.Flush()
	p.Read(buf)

	fmt.Println(buf)
}
