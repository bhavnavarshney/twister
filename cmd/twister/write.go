package main

import (
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/serialport"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/urfave/cli/v2"
)

func Write(c *cli.Context) error {
	log := logrus.New()
	config := &serial.Config{Name: c.String("port"), Baud: c.Int("baud")}
	p := serialport.MakeSerialPort(config, c.Bool("mock"))
	defer p.Close()
	d := serialport.MakeDriver(p, log)
	data := new([24 * 4]byte)
	torqueData := message.MakeTorqueData(*data)
	err := d.SendMessage(torqueData)
	if err != nil {
		return err
	}
	log.Printf("Successfully Sent Hex: %X", data)
	return nil
}
