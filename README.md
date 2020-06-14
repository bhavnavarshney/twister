# TorqueCalibrationGo

A demo application in Go

## Serial Port

We use https://godoc.org/github.com/tarm/serial for our serial communications.

Currently all ports are opened with 8 data bits, 1 stop bit, no parity, no hardware flow control, and no software flow control.

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