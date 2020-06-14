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
	writeBuffer []byte
	readBuffer  []byte
	config      *serial.Config
	Log         *logrus.Logger
}

var responseMap = map[byte][]byte{
	0x07: {0x07},
}

func (mp *FakePort) Write(out []byte) (int, error) {
	mp.Log.Println("Writing to mock port")
	mp.writeBuffer = append(mp.writeBuffer, out...)
	mp.readBuffer = append(mp.readBuffer, responseMap[out[0]]...)
	return len(out), nil
}

// Todo: Extend for multiple bytes
func (mp *FakePort) Read(b []byte) (int, error) {
	mp.Log.Println("Reading from Mock Port")
	b[0] = mp.readBuffer[0]
	return len(b), nil
}

func (mp *FakePort) Close() error {
	mp.Log.Debugln("Closing Mock Port")
	return nil
}
