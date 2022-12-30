package main

import (
	"machine"
	"strconv"
	"time"

	wrapper "slave/i2c"
)

const (
	slaveAddress uint32 = 0x70 // 7 bit address
)

var i2c = wrapper.I2C

func formatAddress(value int64) string {
	return "0x" + strconv.FormatInt(value, 16)
}

func main() {

	println("<Reboot>")

	machine.I2C0.Configure(machine.I2CConfig{})
	println("I2C0 UP")

	i2c.Configure(machine.I2C0, wrapper.I2CSlaveConfig{Address: slaveAddress})
	println("I2C0 SLAVE UP")

	i2c.SetInterrupt()
	println("I2C0 INT UP")

	data := []byte{0x00, 0x00, 0x00, 0x00, 0x00}

	for {

		if i2c.Buffered() > 0 {
			_, err := i2c.Read(data)
			if err != nil {
				println("Error reading data")
			}

			print("Data received:")
			for i := 0; i < len(data); i++ {
				print(" " + formatAddress(int64(data[i])))
			}
			println()
		}

		println("alive")

		// println(i2c.DebugSR())

		time.Sleep(time.Second * 1)

	}
}
