package main

import (
	"fmt"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/profile"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/serialport"
	"github.com/leaanthony/mewn"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"github.com/wailsapp/wails"
)

func basic() string {
	log := logrus.New()
	config := &serial.Config{Name: "/dev/tty.usbserial-AC019QP9", Baud: 9600}
	mock := true
	p := serialport.MakeSerialPort(config, mock)
	defer p.Close()
	d := serialport.MakeSerialPortDriver(p, log)

	readProfileCommand := serialport.MakeCommand(message.BulkParamReceiveMsg, message.BulkParamReceiveMsgLen)
	response, err := d.SendCommand(readProfileCommand)
	if err != nil {
		panic(err)
	}

	torqueData := message.TorqueData{}
	err = torqueData.Unmarshal(response)
	if err != nil {
		panic(err)
	}
	int16Data := message.ToUInt16(torqueData.ToByte())
	fmt.Printf("Torque Int16 Data: %d\n", int16Data)
	var profileArr [48]uint16
	copy(profileArr[:], int16Data)
	profile := profile.MakeProfile(profileArr)
	return fmt.Sprintln(profile)
}

func main() {

	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "twister",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Run()
}
