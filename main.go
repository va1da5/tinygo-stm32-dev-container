package main

import (
	"machine"
	"math"
	"strconv"
	"time"

	"tinygo.org/x/drivers/bme280"
)

func roundFloat(val float32) float32 {
	return float32(math.Round(float64(val*100)) / 100)
}

func formatAddress(value uint16) string {
	return strconv.FormatInt(int64(value), 16)

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

	for {
		temp, _ := sensor.ReadTemperature()
		println("Temperature:", roundFloat(float32(temp)/1000), "°C")

		pressure, _ := sensor.ReadPressure()
		println("Pressure:", float32(pressure)/100000, "hPa")

		humidity, _ := sensor.ReadHumidity()
		println("Humidity:", humidity/100, "%")

		altitude, _ := sensor.ReadAltitude()
		println("Altitude:", altitude, "m")

		println("-------------------------------")

		time.Sleep(2 * time.Second)
	}
}
