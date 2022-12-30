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
	receive := make([]byte, 5)
	for {

		err := i2c.Tx(address, empty, empty)
		if err != nil {
			println("ADDR ACK  FAILED @ " + formattedAddress)
		} else {
			println("ADDR ACK OK @ " + formattedAddress)
		}

		err = i2c.Tx(address, []byte{0x77}, empty)
		if err != nil {
			println("DATA TX FAILED @ " + formattedAddress)
		} else {

			println("DATA TX OK @", formattedAddress)
		}

		err = i2c.Tx(address, empty, receive)
		if err != nil {
			println("DATA RX FAILED @ " + formattedAddress)
		} else {

			data := "| Data: "

			for i := 0; i < len(receive); i++ {
				data += formatAddress(uint16(receive[i])) + " "
			}

			println("DATA RX OK @", formattedAddress, data)
		}

		err = i2c.Tx(address, []byte{0x11, 0x22, 0x33, 0x44}, receive)
		if err != nil {
			println("DATA TXRX FAILED @ " + formattedAddress)
		} else {

			data := "| Data: "

			for i := 0; i < len(receive); i++ {
				data += formatAddress(uint16(receive[i])) + " "
			}

			println("DATA TXRX OK @", formattedAddress, data)
		}

		println("----------------------")
		machine.I2C0.Configure(machine.I2CConfig{})
		time.Sleep(time.Second * 5)
	}
}
