package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/serialport"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/urfave/cli"
)

//const serialportname = "/dev/tty.usbserial-AC019QP9"

func main() {
	app := &cli.App{
		Name:  "twister",
		Usage: "configuration tool for the WD-TMAX drill",
		Commands: []cli.Command{
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "reads the torque profile from the drill",
				Action: func(c *cli.Context) error {
					fmt.Println("reading drill to: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "write",
				Aliases: []string{"w"},
				Usage:   "writes the torque profile to the drill",
				Action: func(c *cli.Context) error {
					fmt.Println("writing drill profile to: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "add a task to the list",
				Action: func(ctx *cli.Context) error {
					return CmdInfo(ctx)
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
			&cli.StringFlag{
				Name:  "Port",
				Value: "COM1",
				Usage: "e.g COM1 or COM3 on windows or /dev/tty.usbserial-AC019QP9 on *nix",
			},
			&cli.StringFlag{
				Name:  "Baudrate",
				Value: "9600",
				Usage: "Serial port baudrate",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func CmdInfo(ctx *cli.Context) error {
	log := logrus.New()
	config := &serial.Config{Name: "COM3", Baud: 9600}
	p, err := serialport.MakeFakePort(config)
	if err != nil {
		return err
	}
	d := serialport.MakeSerialPortDriver(p, log)
	drillTypeCommand := serialport.MakeCommand(0x04, 20)
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
