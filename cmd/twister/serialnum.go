package main

import (
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/serialport"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/urfave/cli/v2"
)

func CmdSerialNum(c *cli.Context) error {
	log := logrus.New()
	config := &serial.Config{Name: c.String("port"), Baud: c.Int("baud")}
	p, err := serialport.MakeSerialPort(config, c.Bool("mock"))
	if err != nil {
		return err
	}
	defer p.Close()
	d := serialport.MakeDriver(p, log)

	drillIDCommand := serialport.MakeCommand(message.DrillIDMsg, message.DrillIDMsgLen)
	response, err := d.SendCommand(drillIDCommand)
	if err != nil {
		return err
	}

	drillID := message.DrillID{}
	err = drillID.Unmarshal(response)
	log.Printf("Response Hex: %X", drillID.ToByte())
	log.Printf("Response ASCII: %s", drillID.ToString())

	if err != nil {
		return err
	}
	return nil
}
