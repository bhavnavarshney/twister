# TorqueCalibrationGo

A demo application in Go

## Prerequisites

- Go1.14
- Wails `go get -u github.com/wailsapp/wails/cmd/wails`
- NodeJS

## Getting Started

### Frontend

`npm start` from the `backend/frontend` directory

### Backend

`wails serve` from the `backend` directory

### Building

`wails build` produces a binary output in the `build` folder called twister. This binary a single executable containing both the backend and the frontend, packaged together.

`wails build -p` packages the application on windows and macOS.

## Overview

The frontend is built in React using MaterialUI. It is integrated with the backend using [wails](https://wails.app/).

## Serial Port

We use https://godoc.org/github.com/tarm/serial for our serial communications.

Currently all ports are opened with 8 data bits, 1 stop bit, no parity, no hardware flow control, and no software flow control.

### Fake Port

For better testing, a fake serial port has been implemented to enable testing.

## CLI Reference

### Read Parameters
```
twister read <config_file>
```

### Write Parameters
```
twister write <config_file>
```

### Get motor info
```
twister info
```