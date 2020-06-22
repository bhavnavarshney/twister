package main

import (
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/serialport"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/urfave/cli/v2"
)

func CmdInfo(c *cli.Context) error {
	log := logrus.New()
	config := &serial.Config{Name: c.String("port"), Baud: c.Int("baud")}
	p, err := serialport.MakeFakePort(config)
	if err != nil {
		return err
	}
	d := serialport.MakeSerialPortDriver(p, log)

	drillTypeCommand := serialport.MakeCommand(message.DrillTypeMsg, message.DrillTypeMsgLen)
	response, err := d.SendCommand(drillTypeCommand)
	if err != nil {
		return err
	}

	drillType := message.DrillType{}
	err = drillType.Unmarshal(response)
	log.Printf("Response Hex: %X", drillType.ToByte())
	log.Printf("Response ASCII: %s", drillType.ToString())

	if err != nil {
		return err
	}
	return nil
}
