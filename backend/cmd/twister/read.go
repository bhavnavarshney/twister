package main

import (
	"fmt"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/serialport"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/urfave/cli/v2"
)

func Read(c *cli.Context) error {
	log := logrus.New()
	var p serialport.Port
	var err error
	config := &serial.Config{Name: c.String("port"), Baud: c.Int("baud")}

	if c.Bool("mock") {
		log.Println("Starting in MOCK mode")
		p, err = serialport.MakeFakePort(config)
	} else {
		p, err = serial.OpenPort(config)
	}
	defer p.Close()
	d := serialport.MakeSerialPortDriver(p, log)

	drillTypeCommand := serialport.MakeCommand(message.BulkParamReceiveMsg, 20)
	response, err := d.SendCommand(drillTypeCommand)
	if err != nil {
		return err
	}

	torqueData := message.TorqueData{}
	err = torqueData.Unmarshal(response)
	fmt.Println(response)
	if err != nil {
		return err
	}
	return nil
}
