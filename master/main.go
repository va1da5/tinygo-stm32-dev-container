package main

import (
	"machine"
	"strconv"
	"time"
)

const address = uint16(0x70)

var empty = []byte{}

func formatAddress(value uint16) string {
	return "0x" + strconv.FormatInt(int64(value), 16)
}

func main() {
	println("REBOOT")
	i2c := machine.I2C0
	machine.I2C0.Configure(machine.I2CConfig{})

	formattedAddress := formatAddress(address)

	for {

		err := i2c.Tx(address, empty, empty)
		if err != nil {
			println("ADDR ACK  FAILED @ " + formattedAddress)
		} else {
			println("ADDR ACK OK @ " + formattedAddress)
		}

		err = i2c.Tx(address, []byte{0x11, 0x22, 0x33, 0x44, 0x55}, empty)
		if err != nil {
			println("DATA SEND FAILED @ " + formattedAddress)
		} else {
			println("DATA SEND OK @ " + formattedAddress)
		}

		println("----------------------")
		machine.I2C0.Configure(machine.I2CConfig{})
		time.Sleep(time.Second * 5)
	}
}
