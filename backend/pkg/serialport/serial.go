package serialport

import (
	"bytes"
	"errors"
	"time"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/nibble"
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

// func (sp *SerialPortDriver) Send(m Message) {
// 	sp.MessageBuffer = append(sp.MessageBuffer, m)
// }

// func (sp *SerialPortDriver) Run() {
// 	for {
// 		err := sp.SendMessage(sp.MessageBuffer[0])
// 		if err != nil {
// 			sp.Log.Printf("Error sending message")
// 		}
// 	}
// }

func (sp *SerialPortDriver) SendMessage(m Message) error {
	_, err := sp.write(m.Marshal())
	if err != nil {
		return err
	}

	buf := make([]byte, m.ResponseLen())
	numBytesRead, err := sp.read(buf)
	if err != nil {
		return err
	}
	sp.Log.Printf("Read %d bytes from port", numBytesRead)
	sp.Log.Printf("Received response: %d", buf)

	// If there's an expected response, check it
	if len(m.Response()) > 0 && !bytes.Equal(m.Response(), buf) {
		// Retry, bump retry count
		//return sp.SendMessage(m)
		return errors.New("Unable to send message")
	}
	return nil
}

func (sp *SerialPortDriver) SendCommand(m Message) ([]byte, error) {
	_, err := sp.write(m.Marshal())
	if err != nil {
		return nil, err
	}

	// Check if response length is zero
	buf := make([]byte, m.ResponseLen())
	_, err = sp.read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func MakeCommand(dataInfo byte, responseLen int) *Command {
	return &Command{
		header:      0x21,
		dataInfo:    dataInfo,
		checkSum:    NibbledChecksum([]byte{dataInfo}),
		responseLen: responseLen,
	}
}

func NibbledChecksum(input []byte) [2]byte {
	hi, low := nibble.Break(message.Checksum(input))
	return [2]byte{hi, low}
}

type Command struct {
	header      byte
	dataInfo    byte
	checkSum    [2]byte
	responseLen int
	response    []byte
}

func (k *Command) Marshal() []byte {
	return []byte{k.header, k.dataInfo, k.checkSum[0], k.checkSum[1]}
}

func (k *Command) Response() []byte {
	return k.response
}

func (k *Command) ResponseLen() int {
	return k.responseLen
}

func (k *Command) Retry() int {
	return 3
}

func (k *Command) Timeout() time.Duration {
	return time.Second
}

const KeepAlive = 0x07

func (sp *SerialPortDriver) SendKeepAlive() bool {
	_, err := sp.write([]byte{KeepAlive})
	if err != nil {
		return false
	}
	buf := make([]byte, 1)
	_, err = sp.read(buf)
	if err != nil {
		return false
	}
	return buf[0] == KeepAlive
}

// Wrap write for easier debugging
func (sp *SerialPortDriver) write(out []byte) (int, error) {
	n, err := sp.Port.Write(out)
	sp.Log.Debugf("wrote %d bytes", n)
	if err != nil {
		sp.Log.Fatal(err)
	}
	return n, err
}

// Wrap read for easier debugging
func (sp *SerialPortDriver) read(b []byte) (int, error) {
	n, err := sp.Port.Read(b)
	sp.Log.Debugf("read %d bytes", n)
	if err != nil {
		sp.Log.Fatal(err)
	}
	return n, err
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
	// Number of bytes expected
	ResponseLen() int
	// Retry returns the number of time to retry when sending a message
	Retry() int
	// Timeout in milliseconds
	Timeout() time.Duration
}
