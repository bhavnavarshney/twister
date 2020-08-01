package twister

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/profile"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/serialport"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/tarm/serial"
)

type Drill struct {
	sync.Mutex
	FS               afero.Fs
	Quit             chan struct{}
	Driver           *serialport.Driver
	ID               string
	Profile          profile.Profile
	CalibratedOffset uint16 // Indicates the zero value for the sensor
	CurrentOffset    uint16 // Indicates the current zero value for the sensor, we need to poll this every second
	Info             string
	Log              *logrus.Logger
}

// Poll the current offset every second, and send it to the frontend
func (dr *Drill) PollCurrentOffset() {
	ticker := time.NewTicker(5 * time.Second)
	dr.Quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				offset, err := dr.GetCurrentOffset()
				if err != nil {
					dr.Log.Errorln(err)
				} else {
					dr.Log.Infoln(offset)
					//dr.runtime.Events.Emit("CurrentOffset", offset)
				}
			case <-dr.Quit:
				ticker.Stop()
			}
		}
	}()
}

func (dr *Drill) SaveProfile(filepath string) error {
	return profile.SaveProfile(&dr.Profile, filepath, dr.FS)
}

func (dr *Drill) LoadProfile(file string) (profile.Profile, error) {
	p, err := profile.LoadProfile(file)
	if err != nil {
		return profile.Profile{}, err
	}
	dr.Profile = *p
	return dr.Profile, nil
}

func (dr *Drill) GetCurrentOffset() (uint16, error) {
	dr.Mutex.Lock()
	defer dr.Mutex.Unlock()
	if dr.Driver == nil {
		return 0, errors.New("Port not open")
	}

	currentOffsetCommand := serialport.MakeCommand(message.CurrentOffsetMsg, message.CurrentOffsetMsgLen)
	response, err := dr.Driver.SendCommand(currentOffsetCommand)
	if err != nil {
		return 0, err
	}
	currentOffset := message.Offset{}
	err = currentOffset.Unmarshal(response)
	if err != nil {
		return 0, err
	}

	return currentOffset.ToUInt16(), nil
}

func (dr *Drill) Open(portName string) (string, error) {
	if dr.Driver != nil {
		dr.Driver.Port.Close()
	}
	config := &serial.Config{Name: portName, Baud: 9600, ReadTimeout: time.Second * 2}
	var mock bool
	if portName == "COM999" {
		mock = true
	}
	p, err := serialport.MakeSerialPort(config, mock)
	if err != nil {
		return "", errors.New("Serial port error. Please check your port number and connection.")
	}
	dr.Driver = serialport.MakeDriver(p, dr.Log)
	ok := dr.Driver.SendKeepAlive()
	if !ok {
		return "", errors.New("Drill not connected.")
	}
	log.Println("Port Opened")
	return "Opened", nil
}

func (dr *Drill) Close() (string, error) {
	if dr.Driver == nil {
		return "", errors.New("Port not open")
	}
	// Close quit to signal termination of polling goroutine
	close(dr.Quit)
	err := dr.Driver.Port.Close()
	if err != nil {
		return "", err
	}
	dr.Driver = nil
	return "Closed", nil
}

type Info struct {
	DrillID          string `json:"DrillID"`
	DrillType        string `json:"DrillType"`
	CurrentOffset    uint16 `json:"CurrentOffset"`
	CalibratedOffset uint16 `json:"CalibratedOffset"`
}

func (dr *Drill) GetInfo() (Info, error) {
	dr.Mutex.Lock()
	defer dr.Mutex.Unlock()
	drillTypeCommand := serialport.MakeCommand(message.DrillTypeMsg, message.DrillTypeMsgLen)
	response, err := dr.Driver.SendCommand(drillTypeCommand)
	if err != nil {
		return Info{}, err
	}
	drillType := message.DrillType{}
	err = drillType.Unmarshal(response)
	if err != nil {
		return Info{}, err
	}

	drillIDCommand := serialport.MakeCommand(message.DrillIDMsg, message.DrillIDMsgLen)
	response, err = dr.Driver.SendCommand(drillIDCommand)
	if err != nil {
		return Info{}, err
	}
	drillID := message.DrillID{}
	err = drillID.Unmarshal(response)
	if err != nil {
		return Info{}, err
	}
	dr.ID = drillID.ToString()

	currentOffsetCommand := serialport.MakeCommand(message.CurrentOffsetMsg, message.CurrentOffsetMsgLen)
	response, err = dr.Driver.SendCommand(currentOffsetCommand)
	if err != nil {
		return Info{}, err
	}
	currentOffset := message.Offset{}
	err = currentOffset.Unmarshal(response)
	if err != nil {
		return Info{}, err
	}

	calibratedOffsetCommand := serialport.MakeCommand(message.CalibratedOffsetMsg, message.CalibratedOffsetMsgLen)
	response, err = dr.Driver.SendCommand(calibratedOffsetCommand)
	if err != nil {
		return Info{}, err
	}
	calibratedOffset := message.Offset{}
	err = calibratedOffset.Unmarshal(response)
	if err != nil {
		return Info{}, err
	}

	// Poll for current offset every second
	dr.PollCurrentOffset()

	return Info{
		DrillType:        drillType.ToString(),
		DrillID:          drillID.ToString(),
		CurrentOffset:    currentOffset.ToUInt16(),
		CalibratedOffset: calibratedOffset.ToUInt16(),
	}, nil
}

func (dr *Drill) WriteParam(p map[string]interface{}) (string, error) {
	dr.Mutex.Lock()
	defer dr.Mutex.Unlock()
	if dr.Driver == nil {
		return "", errors.New("Port has not been opened")
	}
	id := byte(p["ID"].(float64))
	torque := uint16(p["Torque"].(float64))
	ad := uint16(p["AD"].(float64))
	int16Data := message.FromUInt16([]uint16{torque, ad})
	var arr [4]byte
	copy(arr[:], int16Data)
	torqueParam := message.MakeTorqueParam(id, arr)
	err := dr.Driver.SendMessage(torqueParam)
	if err != nil {
		return "", err
	}
	// Update saved param
	dr.Profile.Fields[id-1].AD = ad
	dr.Profile.Fields[id-1].Torque = torque
	return "OK", nil
}

func (dr *Drill) GetProfile() (profile.Profile, error) {
	dr.Mutex.Lock()
	defer dr.Mutex.Unlock()
	if dr.Driver == nil {
		return profile.Profile{}, errors.New("Port has not been opened")
	}
	readProfileCommand := serialport.MakeCommand(message.BulkParamReceiveMsg, message.BulkParamReceiveMsgLen)
	response, err := dr.Driver.SendCommand(readProfileCommand)
	if err != nil {
		return profile.Profile{}, err
	}

	torqueData := message.TorqueData{}
	err = torqueData.Unmarshal(response)
	if err != nil {
		return profile.Profile{}, err
	}
	fmt.Printf("Torque Data Hex: %X", torqueData.ToByte())
	int16Data := message.ToUInt16(torqueData.ToByte())
	fmt.Printf("Torque Int16 Data: %d\n", int16Data)
	var profileArr [48]uint16
	copy(profileArr[:], int16Data)
	profile := profile.MakeProfile(profileArr)
	dr.Profile = *profile
	return *profile, nil
}

func (dr *Drill) WriteProfile(p []interface{}) (string, error) {
	dr.Mutex.Lock()
	defer dr.Mutex.Unlock()
	for i := range p {
		dr.Profile.Fields[i].Torque = uint16(p[i].(map[string]interface{})["Torque"].(float64))
		dr.Profile.Fields[i].AD = uint16(p[i].(map[string]interface{})["AD"].(float64))
	}
	data := dr.Profile.MarshalBytes()
	int16Data := message.FromUInt16(data[:])
	var arr [96]byte
	copy(arr[:], int16Data)
	torqueData := message.MakeTorqueData(arr)
	err := dr.Driver.SendMessage(torqueData)
	if err != nil {
		return "", err
	}
	return "OK", nil
}
