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

type Drill struct {
	driver           *serialport.Driver
	profile          profile.Profile
	calibratedOffset uint16
	currentOffset    uint16
	info             string
}

func (dr *Drill) GetInfo() string {
	// ID
	// Type
	// Current Offset
	// Calibrated Offset
	return ""
}

func (dr *Drill) WriteParam(p map[string]interface{}) (string, error) {
	id := byte(p["ID"].(float64))
	torque := uint16(p["Torque"].(float64))
	ad := uint16(p["AD"].(float64))
	int16Data := message.FromUInt16([]uint16{torque, ad})
	var arr [4]byte
	copy(arr[:], int16Data)
	torqueParam := message.MakeTorqueParam(id, arr)
	err := dr.driver.SendMessage(torqueParam)
	if err != nil {
		return "", err
	}
	return "OK", nil
}

func (dr *Drill) GetProfile() profile.Profile {
	readProfileCommand := serialport.MakeCommand(message.BulkParamReceiveMsg, message.BulkParamReceiveMsgLen)
	response, err := dr.driver.SendCommand(readProfileCommand)
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
	return *profile
}

func (dr *Drill) WriteProfile(p []interface{}) (string, error) {
	for i := range p {
		dr.profile.Fields[i].Torque = uint16(p[i].(map[string]interface{})["Torque"].(float64))
		dr.profile.Fields[i].AD = uint16(p[i].(map[string]interface{})["AD"].(float64))
	}
	data := dr.profile.MarshalBytes()
	int16Data := message.FromUInt16(data[:])
	var arr [96]byte
	copy(arr[:], int16Data)
	torqueData := message.MakeTorqueData(arr)
	err := dr.driver.SendMessage(torqueData)
	if err != nil {
		return "", err
	}
	return "OK", nil
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

	log := logrus.New()
	config := &serial.Config{Name: "/dev/tty.usbserial-AC019QP9", Baud: 9600}
	mock := true
	p := serialport.MakeSerialPort(config, mock)
	d := serialport.MakeDriver(p, log)
	defer p.Close()

	app.Bind(&Drill{
		driver: d,
	})
	app.Run()
}
