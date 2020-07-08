// Format temperature and pressure readings from BMP180
// and display them on a 128x64 SSD1306 OLED
//
// TinyGo drivers and TinyFont are required

package main

import (
	"image/color"
	"machine"
	"strconv"
	"time"

	"tinygo.org/x/drivers/bmp180"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freeserif"
)

func main() {

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	sensor := bmp180.New(machine.I2C0)
	sensor.Configure()

	connected := sensor.Connected()
	if !connected {
		println("BMP180 not detected")
		return
	}
	println("BMP180 detected")

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: ssd1306.Address_128_32,
		Width:   128,
		Height:  64,
	})

	c := color.RGBA{0xff, 0xff, 0xff, 0xff}
	display.ClearDisplay()

	for {

		temp, _ := sensor.ReadTemperature()
		pressure, _ := sensor.ReadPressure()
		println("Temperature:", temp, " / Pressure:", pressure)

		temp_str := strconv.FormatFloat(float64(temp)/1000, 'f', 1, 64) + " *C"
		pressure_str := strconv.FormatFloat(float64(pressure)/100000, 'f', 2, 64) + " hPa"

		display.ClearBuffer()
		tinyfont.WriteLine(&display, &freeserif.Bold12pt7b, 0, 24, []byte(temp_str), c)
		tinyfont.WriteLine(&display, &freeserif.Bold12pt7b, 0, 48, []byte(pressure_str), c)
		display.Display()

		time.Sleep(time.Millisecond * 100)
	}
}
