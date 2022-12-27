package main

import (
	// "machine"

	"machine"
	"strconv"
	"time"
)

func formatAddress(value uint16) string {
	return strconv.FormatInt(int64(value), 16)

}

func main() {
	i2c := machine.I2C0
	machine.I2C0.Configure(machine.I2CConfig{})

	write := []byte{0x75}
	receive := make([]byte, 1)
	for {

		for address := uint16(0); address <= uint16(127); address++ {
			err := i2c.Tx(address, write, receive)
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
