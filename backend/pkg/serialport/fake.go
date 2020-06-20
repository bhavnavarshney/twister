package serialport

import (
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
	0x07: {0x07},
	0x70: {0x70},
	0x04: {0x21, 0x04, 0x35, 0x37, 0x33, 0x32, 0x34, 0x37, 0x34, 0x37, 0x33, 0x39, 0x33, 0x38, 0x33, 0x35, 0x33, 0x30, 0x30, 0x45},
}

func (mp *FakePort) Write(out []byte) (int, error) {
	mp.Log.Println("Writing to mock port")
	mp.writeLog = append(mp.writeLog, out...)
	// we always read the second byte to see what to do
	mp.readBuffer = append(mp.readBuffer, responseMap[out[1]]...)
	return len(out), nil
}

// Todo: Extend for multiple bytes
func (mp *FakePort) Read(b []byte) (int, error) {
	mp.Log.Println("Reading from Mock Port")

	if len(mp.readBuffer) > 0 {
		mp.Log.Printf("Read %X", mp.readBuffer)
		for i := range b {
			b[i] = mp.readBuffer[i]
		}
	}
	return len(b), nil
}

func (mp *FakePort) Close() error {
	mp.Log.Debugln("Closing Mock Port")
	return nil
}
