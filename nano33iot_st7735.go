package main

import (
	"image/color"
	"machine"
	"math"
	"strconv"
	"time"

	"tinygo.org/x/drivers/lsm6ds3"
	"tinygo.org/x/drivers/st7735"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/freeserif"
)

const (
	BLACK = iota
	WHITE
	RED
	GREEN
	BLUE
	YELLOW
	CYAN
	PURPLE
)

func main() {

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	accel := lsm6ds3.New(machine.I2C0)
	accel.Configure(lsm6ds3.Configuration{})

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 24000000,
	})

	rst := machine.D10
	dc := machine.D9
	cs := machine.D8
	bl := machine.D7
	display := st7735.New(machine.SPI0, rst, dc, cs, bl)
	display.Configure(st7735.Config{
		Width:        128,
		Height:       160,
		RowOffset:    2,
		ColumnOffset: 2,
	})

	mycolors := []color.RGBA{
		color.RGBA{0, 0, 0, 255},       // BLACK
		color.RGBA{255, 255, 255, 255}, // WHITE
		color.RGBA{255, 0, 0, 255},     // RED
		color.RGBA{0, 255, 0, 255},     // GREEN
		color.RGBA{0, 0, 255, 255},     // BLUE
		color.RGBA{255, 255, 0, 255},   // YELLOW
		color.RGBA{0, 255, 255, 255},   // CYAN
		color.RGBA{255, 0, 255, 255},   // PURPLE
	}

	display.FillScreen(mycolors[BLACK])

	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, 30, []byte("Pitch"), mycolors[YELLOW])
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, 80, []byte("Roll"), mycolors[YELLOW])
	tinyfont.WriteLine(&display, &freeserif.Bold9pt7b, 10, 130, []byte("Temperature"), mycolors[BLUE])

	var c1, c2 color.RGBA

	for {

		x, y, z := accel.ReadAcceleration()
		t, _ := accel.ReadTemperature()

		x_value := float64(x) / 1000000
		y_value := float64(y) / 1000000
		z_value := float64(z) / 1000000
		t_value := float64(t) / 1000

		pitch := math.Atan2(y_value, math.Sqrt(math.Pow(x_value, 2)+math.Pow(z_value, 2))) * (180 / math.Pi)
		roll := math.Atan2(x_value, math.Sqrt(math.Pow(y_value, 2)+math.Pow(z_value, 2))) * (180 / math.Pi)

		pitch_str := strconv.FormatFloat(pitch, 'f', 3, 64)
		roll_str := strconv.FormatFloat(roll, 'f', 3, 64)
		temp_str := strconv.FormatFloat(t_value, 'f', 1, 64) + " *C"

		c1 = mycolors[CYAN]
		c2 = mycolors[CYAN]
		if math.Abs(pitch) >= 10.0 {
			c1 = mycolors[RED]
		} else if math.Abs(pitch) >= 2.0 {
			c1 = mycolors[GREEN]
		}
		if math.Abs(roll) >= 10.0 {
			c2 = mycolors[RED]
		} else if math.Abs(roll) >= 2.0 {
			c2 = mycolors[GREEN]
		}

		display.FillRectangle(12, 50-12, 80, 12+2, mycolors[BLACK])
		display.FillRectangle(12, 100-12, 80, 12+2, mycolors[BLACK])
		display.FillRectangle(12, 150-6, 40, 6+2, mycolors[BLACK])
		tinyfont.WriteLine(&display, &freemono.Bold9pt7b, 12, 50, []byte(pitch_str), c1)
		tinyfont.WriteLine(&display, &freemono.Bold9pt7b, 12, 100, []byte(roll_str), c2)
		tinyfont.WriteLine(&display, &tinyfont.Org01, 12, 150, []byte(temp_str), mycolors[WHITE])

		time.Sleep(time.Millisecond * 100)

	}

}
