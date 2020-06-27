package serialport

import (
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

func MakeFakePort(config *serial.Config) (*FakePort, error) {
	return &FakePort{
		config: config,
		Log:    logrus.New(),
	}, nil
}

// FakePort implements the Port interface and is used for testing without a real hardware device
type FakePort struct {
	writeLog   []byte
	readBuffer []byte
	config     *serial.Config
	Log        *logrus.Logger
}

var responseMap = map[byte][]byte{
	KeepAlive:                   {0x07},
	message.DrillIDMsg:          {0x21, 0x04, 0x35, 0x37, 0x33, 0x32, 0x34, 0x37, 0x34, 0x37, 0x33, 0x39, 0x33, 0x38, 0x33, 0x35, 0x33, 0x30, 0x30, 0x45},
	message.BulkParamReceiveMsg: {0xFF, 0x21, 0x15, 0x30, 0x30, 0x30, 0x46, 0x30, 0x30, 0x31, 0x45, 0x30, 0x30, 0x32, 0x38, 0x30, 0x30, 0x33, 0x32, 0x30, 0x30, 0x33, 0x43, 0x30, 0x30, 0x34, 0x36, 0x30, 0x30, 0x35, 0x30, 0x30, 0x30, 0x35, 0x41, 0x30, 0x30, 0x36, 0x34, 0x30, 0x30, 0x36, 0x45, 0x30, 0x30, 0x37, 0x38, 0x30, 0x30, 0x38, 0x32, 0x30, 0x30, 0x30, 0x46, 0x30, 0x30, 0x31, 0x45, 0x30, 0x30, 0x32, 0x38, 0x30, 0x30, 0x33, 0x32, 0x30, 0x30, 0x33, 0x43, 0x30, 0x30, 0x34, 0x36, 0x30, 0x30, 0x35, 0x30, 0x30, 0x30, 0x35, 0x41, 0x30, 0x30, 0x36, 0x34, 0x30, 0x30, 0x36, 0x45, 0x30, 0x30, 0x37, 0x38, 0x30, 0x30, 0x38, 0x32, 0x30, 0x30, 0x32, 0x44, 0x30, 0x30, 0x35, 0x41, 0x30, 0x30, 0x37, 0x38, 0x30, 0x30, 0x39, 0x36, 0x30, 0x30, 0x42, 0x33, 0x30, 0x30, 0x44, 0x30, 0x30, 0x30, 0x45, 0x44, 0x30, 0x31, 0x30, 0x41, 0x30, 0x31, 0x32, 0x37, 0x30, 0x31, 0x34, 0x34, 0x30, 0x31, 0x36, 0x31, 0x30, 0x31, 0x37, 0x45, 0x30, 0x30, 0x32, 0x44, 0x30, 0x30, 0x35, 0x41, 0x30, 0x30, 0x37, 0x38, 0x30, 0x30, 0x39, 0x36, 0x30, 0x30, 0x42, 0x33, 0x30, 0x30, 0x44, 0x30, 0x30, 0x30, 0x45, 0x44, 0x30, 0x31, 0x30, 0x41, 0x30, 0x31, 0x32, 0x37, 0x30, 0x31, 0x34, 0x34, 0x30, 0x31, 0x36, 0x31, 0x30, 0x31, 0x37, 0x45, 0x33, 0x30},
	message.BulkParamSendMsg:    {message.OkStatus},
}

func (mp *FakePort) Write(out []byte) (int, error) {
	if len(out) == 0 {
		return 0, nil
	}
	mp.Log.Println("Writing to mock port")
	mp.writeLog = append(mp.writeLog, out...)
	// we always read the second byte to see what to do
	mp.readBuffer = append(mp.readBuffer, responseMap[out[1]]...)
	mp.Log.Printf("Adding %X to read buffer", responseMap[out[1]])
	return len(out), nil
}

// Todo: Extend for multiple bytes
func (mp *FakePort) Read(b []byte) (int, error) {
	var numBytesCopied int
	mp.Log.Println("Reading from Mock Port")

	if len(mp.readBuffer) > 0 {
		mp.Log.Printf("Read HEX: %X", mp.readBuffer)
		numBytesCopied = copy(b, mp.readBuffer)
		// Remove from readbuffer
		mp.readBuffer = mp.readBuffer[numBytesCopied:]
	}

	return numBytesCopied, nil
}

func (mp *FakePort) Close() error {
	mp.Log.Println("Closing Mock Port")
	return nil
}
