package main

import (
	"machine"
	"strconv"
	"time"

	"tinygo.org/x/drivers/bme280"
)

func formatFloat(value float32, precision int) string {
	return strconv.FormatFloat(float64(value), 'f', precision, 64)
}

func makeSeparator(length int, char byte) string {
	slice := make([]byte, length, length)

	for i := range slice {
		slice[i] = char
	}

	return string(slice)
}

func main() {
	machine.I2C0.Configure(machine.I2CConfig{})
	sensor := bme280.New(machine.I2C0)
	sensor.Configure()

	connected := sensor.Connected()
	if !connected {
		println("BME280 not detected")
		return
	}
	println("BME280 detected")

	var separator = makeSeparator(37, '~')

	for {
		temp, _ := sensor.ReadTemperature()
		println("Temperature:", formatFloat(float32(temp)/1000, 2), "°C")

		pressure, _ := sensor.ReadPressure()
		println("Pressure:", formatFloat(float32(pressure)/1000000, 2), "kPa, (", formatFloat(float32(pressure)/100000, 2), "hPa)")

		humidity, _ := sensor.ReadHumidity()
		println("Humidity:", formatFloat(float32(humidity)/100, 1), "%")

		altitude, _ := sensor.ReadAltitude()
		println("Altitude:", altitude, "m")

		println(separator)

		time.Sleep(2 * time.Second)
	}
}
