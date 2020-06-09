package serial

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

type SerialPort struct {
	Name string
	Baud int
	Port *serial.Port
	Log  logrus.Logger
}

func (s *SerialPort) Open() {
	c := &serial.Config{Name: "COM3", Baud: 9600}
	p, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	s.Port = p
}

func (s *SerialPort) Write() {
	n, err := s.Port.Write([]byte{0x07})
	s.Log.Println("wrote %d bytes", n)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SerialPort) Read() {
	buf := make([]byte, 1)
	n, err := s.Port.Read(buf)
	s.Log.Println("read %d bytes", n)
	if err != nil {
		log.Fatal(err)
	}
}
