# TorqueCalibrationGo

A demo application in Go

## Prerequisites

- Go1.14
- golangci-lint 1.29.0
- make
- NodeJS
- Chrome or some other variant of chromium must be installed to run the built application
- After cloning the repo, run `git update-index --assume-unchanged assets.go` to ignore changes to `assets.go` which slows down git significantly as it is a large file (~37Mb)

## Getting Started

### Building

`make build`

Builds the app into the `/build` folder.

### Frontend

`npm start` from the `frontend` directory

### Test

`make test`

### Lint

`make lint`

## Overview

The frontend is built in React using MaterialUI. It is integrated with the backend using [lorca](https://github.com/zserge/lorca).

## Serial Port

We use https://godoc.org/github.com/tarm/serial for our serial communications.

Currently all ports are opened with 8 data bits, 1 stop bit, no parity, no hardware flow control, and no software flow control.

### Fake Port

For better testing, a fake serial port has been implemented to enable testing. When COM Port 999 is selected, the fake serial port is used.