// I2C device scanner. Only looks for ACKs for device address.
package main

import (
	"machine"
	"strconv"
	"time"
)

func formatAddress(value uint16) string {
	return strconv.FormatInt(int64(value), 16)
}

func main() {
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{})

	empty := []byte{}
	for {

		for address := uint16(0); address <= uint16(127); address++ {
			err := i2c.Tx(address, empty, empty)
			if err != nil {
				println("I2C device not found on: 0x" + formatAddress(address))
			} else {
				println("\nFound I2C device on: 0x" + formatAddress(address) + "\n")
			}
		}

		println("----------------")

		time.Sleep(time.Second * 5)
	}
}
