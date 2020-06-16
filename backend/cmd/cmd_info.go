package main

import (
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/serialport"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/urfave/cli"
)

func CmdInfo(ctx *cli.Context) error {
	log := logrus.New()
	config := &serial.Config{Name: "COM3", Baud: 9600}
	p, err := serialport.MakeFakePort(config)
	if err != nil {
		return err
	}
	d := serialport.MakeSerialPortDriver(p, log)
	drillIDCommand := serialport.MakeCommand(0x04, 20)
	response, err := d.SendCommand(drillIDCommand)
	if err != nil {
		return err
	}
	drillID := message.DrillID{}
	err = drillID.Unmarshal(response)
	if err != nil {
		return err
	}
	return nil
}
