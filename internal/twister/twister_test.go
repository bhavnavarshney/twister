package twister

import (
	"testing"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/serialport"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tarm/serial"
)

func TestDrill_SaveProfile(t *testing.T) {
	fs := afero.NewMemMapFs()
	dr := Drill{}
	fileName := "testFileName.csv"
	dr.FS = fs
	err := dr.SaveProfile(fileName)
	assert.NoError(t, err)
	_, err = afero.ReadFile(fs, fileName)
	assert.NoError(t, err)
}
func TestDrill_LoadProfile(t *testing.T) {
	dr := Drill{}
	testInput := `ID,Torque,TorqueAD
1,4,60
2,5,72
3,6,83
4,7,94
5,8,108
6,9,122
7,10,131
8,11,153
9,12,168
10,13,180
11,14,191
20,3,50
21,4,60
22,5,72
23,6,83
24,7,94`
	p, err := dr.LoadProfile(testInput)
	assert.NoError(t, err)
	assert.Equal(t, p, dr.Profile)
	assert.Equal(t, dr.Profile.Fields[0].AD, uint16(60))
	assert.Equal(t, dr.Profile.Fields[0].Torque, uint16(4))
	assert.Equal(t, dr.Profile.Fields[1].AD, uint16(72))
	assert.Equal(t, dr.Profile.Fields[1].Torque, uint16(5))

}

func TestDrill_WriteParamUpdatesProfile(t *testing.T) {
	dr := Drill{}
	dr.Log = logrus.New()
	config := &serial.Config{}
	p, err := serialport.MakeSerialPort(config, true)
	assert.NoError(t, err)
	dr.Driver = serialport.MakeDriver(p, dr.Log)
	param := map[string]interface{}{
		"ID":     float64(1),
		"Torque": float64(30),
		"AD":     float64(50),
	}
	_, err = dr.WriteParam(param)
	assert.NoError(t, err)
	assert.Equal(t, uint16(50), dr.Profile.Fields[0].AD)
	assert.Equal(t, uint16(30), dr.Profile.Fields[0].Torque)
}
