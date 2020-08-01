package serialport

import (
	"github.com/tarm/serial"
)

// MakeSerialPort returns a serial port with the provided config applied
func MakeSerialPort(config *serial.Config) (Port, error) {
	var p Port
	var err error
	p, err = serial.OpenPort(config)
	if err != nil {
		return nil, err
	}
	return p, nil
}

type Port interface {
	Write(out []byte) (int, error)
	Read([]byte) (int, error)
	Close() error
}
