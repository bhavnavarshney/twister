package main

import (
	"fmt"
	"log"
	"os"

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
				Action: func(c *cli.Context) error {
					fmt.Println("motor info: ", c.Args().First())
					return nil
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
		Action: func(c *cli.Context) error {
			fmt.Println("twister calibration tool")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
