package serialport

import (
	"bytes"

	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

func MakeSerialPortDriver(p Port, log *logrus.Logger) *SerialPortDriver {
	return &SerialPortDriver{
		Port: p,
		Log:  log,
	}
}

// SerialPortDriver handles communications to the drill
type SerialPortDriver struct {
	Config        *serial.Config
	Port          Port
	MessageBuffer []Message
	Log           *logrus.Logger
}

func (sp *SerialPortDriver) Send(m Message) {
	sp.MessageBuffer = append(sp.MessageBuffer, m)
}

func (sp *SerialPortDriver) Run() {
	for {
		err := sp.SendMessage(sp.MessageBuffer[0])
		if err != nil {
			sp.Log.Printf("Error sending message")
		}
	}
}

func (sp *SerialPortDriver) SendMessage(m Message) error {
	err := sp.write(m.Marshal())
	if err != nil {
		return err
	}

	buf := make([]byte, 1)
	err = sp.read(buf)
	if err != nil {
		return err
	}

	// If there's an expected response, check it
	if len(m.Response()) > 0 && !bytes.Equal(m.Response(), buf) {
		// Retry, bump retry count
		return sp.SendMessage(m)
	}
	return nil
}

const keepAlive = 0x07

func (sp *SerialPortDriver) SendKeepAlive() bool {
	err := sp.write([]byte{keepAlive})
	if err != nil {
		return false
	}
	buf := make([]byte, 1)
	err = sp.read(buf)
	if err != nil {
		return false
	}
	return buf[0] == keepAlive
}

// Wrap write for easier debugging
func (sp *SerialPortDriver) write(out []byte) error {
	n, err := sp.Port.Write(out)
	sp.Log.Debugf("wrote %d bytes", n)
	return err
}

// Wrap read for easier debugging
func (sp *SerialPortDriver) read(b []byte) error {
	n, err := sp.Port.Read(b)
	sp.Log.Debugf("read %d bytes", n)
	if err != nil {
		sp.Log.Fatal(err)
	}
	return err
}

type Port interface {
	Write(out []byte) (int, error)
	Read([]byte) (int, error)
	Close() error
}

type Message interface {
	// Marshal serialises the message into an slice of bytes
	Marshal() []byte
	// Response returns an expected response
	Response() []byte
	// Retry returns the number of time to retry when sending a message
	Retry() int
	// Timeout in milliseconds
	Timeout() int
}
