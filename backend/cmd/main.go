package main

import (
	"log"

	"github.com/tarm/serial"
)

//const serialportname = "/dev/tty.usbserial-AC019QP9"

func main() {
	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte{0x07})
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%x", buf[:n])
}
