// draw 5x5 icons quickly on BBC micro:bits

package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
)

var (
	display microbitmatrix.Device
	leds    [5][5]uint8
)

/*
To draw a icon, simply set the 5x5 led matrix:

leds = [5][5]uint8{
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
}

1 = light up and 0 = no light.
Once set, the icon will stay on in the matrix until you change it.
*/

func main() {

	setDisplay()
	go displayRoutine()
	sleep(100) // wait a bit to ensure the goroutine runs

	for {

		leds = [5][5]uint8{ // heart
			{0, 1, 0, 1, 0},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 0},
		}
		sleep(500)
		displayClear()
		sleep(500)

		leds = [5][5]uint8{ // smile face
			{0, 1, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 0, 0, 0},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		}
		sleep(500)
		displayClear()
		sleep(500)

		leds = [5][5]uint8{ // musical note
			{0, 0, 1, 1, 0},
			{0, 0, 1, 1, 1},
			{0, 0, 1, 0, 1},
			{1, 1, 1, 0, 0},
			{1, 1, 1, 0, 0},
		}
		sleep(500)
		displayClear()
		sleep(500)
	}

}

// utility functions

func setDisplay() {
	display = microbitmatrix.New()
	display.Configure(microbitmatrix.Config{Rotation: 0})
}

func displayRoutine() { // goroutine
	for {
		for i := int16(0); i < 5; i++ {
			for j := int16(0); j < 5; j++ {
				if leds[i][j] > 0 {
					display.SetPixel(i, j, color.RGBA{255, 255, 255, 255})
				} else {
					display.SetPixel(i, j, color.RGBA{0, 0, 0, 0})
				}
			}
		}
		display.Display()
	}
}

func displayClear() {
	var ledsEmpty [5][5]uint8
	leds = ledsEmpty
}

func sleep(delay uint16) {
	time.Sleep(time.Millisecond * time.Duration(delay))
}
