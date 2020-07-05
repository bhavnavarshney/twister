package serialport

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

func MakeSerialPort(config *serial.Config, isFake bool) (Port, error) {
	var p Port
	var err error
	if isFake {
		log.Println("Starting in MOCK mode")
		p, err = MakeFakePort(config)
	} else {
		p, err = serial.OpenPort(config)
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func MakeDriver(p Port, log *logrus.Logger) *Driver {
	return &Driver{
		Port: p,
		Log:  log,
	}
}

// Driver handles communications to the drill
type Driver struct {
	Config *serial.Config
	Port   Port
	Log    *logrus.Logger
}

// Message just receive a single byte response
// POST
func (sp *Driver) SendMessage(m Message) error {
	_, err := sp.write(m.Marshal())
	if err != nil {
		return err
	}

	result := make(chan []byte, 1)
	errResp := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		for {
			buf := make([]byte, m.ResponseLen())
			numBytesRead, err := sp.read(buf)
			sp.Log.Printf("Read %d bytes from port", numBytesRead)
			sp.Log.Printf("Received response: %d", buf[:numBytesRead])
			if err != nil {
				errResp <- fmt.Errorf("error reading from port: %w", err)
				return
			}
			if bytes.Contains(buf, m.Response()) {
				result <- buf
				return
			} else {
				sp.Log.Warnf("Expected %X but received %X", m.Response(), buf)
			}
		}
	}(ctx)

	select {
	case <-result:
		sp.Log.Printf("Received expected response %X", m.Response())
		return nil
	case <-ctx.Done():
		return errors.New("timeout waiting for response")
	}
}

// Commands receive a full payload response
// GET
func (sp *Driver) SendCommand(m Message) ([]byte, error) {
	_, err := sp.write(m.Marshal())
	if err != nil {
		return nil, err
	}

	result := make(chan []byte, m.ResponseLen())
	errResp := make(chan error, 1)
	var received []byte
	bytesReceived := 0
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		for {
			buf := make([]byte, 50)
			numBytesRead, err := sp.read(buf)
			fmt.Printf("reading %d bytes", numBytesRead)
			if err != nil {
				errResp <- fmt.Errorf("error reading from port: %w", err)
				return
			}

			// already have header
			if len(received) > 0 && (received[0] == byte(0x21)) {
				bytesReceived += numBytesRead
				fmt.Println(received)
				received = append(received, buf[:numBytesRead]...)
			} else {
				// waiting for header
				if startIndex := bytes.Index(buf, []byte{0x21}); startIndex != -1 {
					fmt.Println("Header detected")
					receiveData := buf[startIndex:numBytesRead]
					bytesReceived += len(receiveData)
					received = append(received, receiveData...)
				}
			}

			if len(received) == m.ResponseLen() {
				result <- received
				return
			}
		}
	}(ctx)

	select {
	case <-errResp:
		return nil, errors.New("error reading from serial port")
	case <-result:
		sp.Log.Printf("Received expected response %X", m.Response())
		return received, nil
	case <-ctx.Done():
		return nil, errors.New("timeout waiting for response")
	}
}

func MakeCommand(dataInfo byte, responseLen int) *Command {
	return &Command{
		header:      0x21,
		dataInfo:    dataInfo,
		checkSum:    NibbledChecksum([]byte{dataInfo}),
		responseLen: responseLen,
	}
}

func NibbledChecksum(dataInfo []byte) [2]byte {
	var checksum = message.Checksum(dataInfo)
	var encodedChecksum [2]byte
	copy(encodedChecksum[:], message.Encode([]byte{checksum}))
	return encodedChecksum
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

func (sp *Driver) SendKeepAlive() bool {
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
// Adds delay
func (sp *Driver) write(out []byte) (int, error) {
	n, err := sp.Port.Write(out)
	if err != nil {
		sp.Log.Fatal(err)
		return n, err
	}
	sp.Log.Printf("wrote %d bytes", n)
	sp.Log.Printf("wrote %X", out)

	return len(out), nil
}

// Wrap read for easier debugging
func (sp *Driver) read(b []byte) (int, error) {
	n, err := sp.Port.Read(b)
	sp.Log.Printf("read %d bytes", n)
	sp.Log.Printf("read %X", b)
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
