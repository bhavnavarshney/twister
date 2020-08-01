package serialport

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cuminandpaprika/TorqueCalibrationGo/pkg/message"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

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
	result := make(chan []byte, 1)
	errResp := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		_, err := sp.write(m.Marshal())
		if err != nil {
			errResp <- fmt.Errorf("error writing to port: %w", err)
		}
		//nolint:staticcheck
		ticker := time.Tick(20 * time.Millisecond)
		for {
			select {
			case <-ticker:
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
			case <-ctx.Done():
				return
			}

		}
	}(ctx)

	select {
	case errMsg := <-errResp:
		return errMsg
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
	result := make(chan []byte, m.ResponseLen())
	errResp := make(chan error, 1)
	var received []byte
	bytesReceived := 0
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		_, err := sp.write(m.Marshal())
		if err != nil {
			errResp <- fmt.Errorf("error writing to port: %w", err)
		}
		//nolint:staticcheck
		ticker := time.Tick(20 * time.Millisecond)
		for {
			select {
			case <-ticker:
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
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	select {
	case <-errResp:
		return nil, errors.New("error reading from serial port")
	case resultVal := <-result:
		sp.Log.Printf("Received expected response %X", resultVal)
		return resultVal, nil
	case <-ctx.Done():
		return nil, errors.New("timeout waiting for response")
	}
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

func (sp *Driver) SendKeepAlive() bool {
	_, err := sp.write([]byte{message.KeepAlive})
	if err != nil {
		return false
	}
	buf := make([]byte, 1)
	_, err = sp.read(buf)
	if err != nil {
		return false
	}
	return buf[0] == message.KeepAlive
}

// Wrap write for easier debugging
// Adds delay
func (sp *Driver) write(out []byte) (int, error) {
	n, err := sp.Port.Write(out)
	if err != nil {
		sp.Log.Warningln(err)
		return n, err
	}
	sp.Log.Printf("wrote %d bytes", n)
	sp.Log.Printf("wrote %X", out)

	return len(out), nil
}

// Wrap read for easier debugging
func (sp *Driver) read(b []byte) (int, error) {
	n, err := sp.Port.Read(b)
	if err != nil {
		sp.Log.Warningln(err)
		return n, err
	}
	sp.Log.Printf("read %d bytes", n)
	sp.Log.Printf("read %X", b)
	return n, err
}
