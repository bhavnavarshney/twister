# serialport

## Connection

The serial connection is a duplex protocol implemented over two wires. Therefore it is capable of sending and receiving simultaneously.

## KeepAlive

The motor responds to 0x70 with 0x70 when connected. This is checked every 0.5 seconds to indicate the motor is connected.

Scenarios
- Verify that the device is connected initially
- Send during idle periods to check that the device is still connected
- If the program is left on, we want to know the device is still connected

## Fake

A fake version of the serialport is available which replicates the behavior of the motor. Useful for testing and demo purposes.



