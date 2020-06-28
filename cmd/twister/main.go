package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//const serialportname = "/dev/tty.usbserial-AC019QP9"

func main() {
	app := &cli.App{
		Name:  "twister",
		Usage: "configuration tool for the WD-TMAX drill",
		Commands: []*cli.Command{
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "reads the torque profile from the drill",
				Action: func(c *cli.Context) error {
					return Read(c)
				},
			},
			{
				Name:    "write",
				Aliases: []string{"w"},
				Usage:   "writes the torque profile to the drill",
				Action: func(c *cli.Context) error {
					return Write(c)
				},
			},
			{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "add a task to the list",
				Action: func(ctx *cli.Context) error {
					return CmdSerialNum(ctx)
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
				Name:    "port",
				Aliases: []string{"p"},
				Value:   "COM1",
				Usage:   "e.g COM1 or COM3 on windows or /dev/tty.usbserial-AC019QP9 on *nix",
			},
			&cli.IntFlag{
				Name:    "baud",
				Aliases: []string{"b"},
				Value:   9600,
				Usage:   "Serial port baudrate",
			},
			&cli.BoolFlag{
				Name:    "mock",
				Aliases: []string{"m"},
				Usage:   "Enables test mode for mock testing without hardware",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Enables verbose logging",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
