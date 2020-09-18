// Generate rainbow colors on NeoPixels (WS2812) and rotate them
//
// TinyGo drivers is required

package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const (
	LED_PIN          = machine.D2
	LED_NUM          = 12
	LED_MAX_LEVEL    = 64 // max 255
	LED_ROTATE_DELAY = 50 // ms
)

var (
	neoPixel ws2812.Device
	leds     [LED_NUM]color.RGBA
)

func main() {

	LED_PIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	neoPixel = ws2812.New(LED_PIN)

	rainbowSet()

	for {
		rainbowRotate(true)
		time.Sleep(time.Millisecond * LED_ROTATE_DELAY)
	}

}

func rainbowSet() {
	ledChange := int16(LED_MAX_LEVEL / (LED_NUM / 3))
	ledIndex := [3]int16{0, int16(LED_NUM / 3), int16(LED_NUM / 3 * 2)}
	for i := range leds {
		ledColor := [3]int16{0, 0, 0}
		for j := range ledIndex {
			if abs(int16(i)-ledIndex[j]) <= ledIndex[1] {
				ledColor[j] = int16(LED_MAX_LEVEL) - abs(int16(i)-ledIndex[j])*ledChange
			}
		}
		if int16(i) >= ledIndex[2] {
			ledColor[0] = LED_MAX_LEVEL - int16(LED_NUM-i)*ledChange
		}
		for j := range ledIndex {
			if ledColor[j] > LED_MAX_LEVEL {
				ledColor[j] = LED_MAX_LEVEL
			} else if ledColor[j] < 0 {
				ledColor[j] = 0
			}
		}
		leds[i] = color.RGBA{R: uint8(ledColor[0]), G: uint8(ledColor[1]), B: uint8(ledColor[2])}
	}
	neoPixel.WriteColors(leds[:])
}

func rainbowRotate(clockwise bool) {
	var ledsTmp []color.RGBA
	if clockwise {
		ledsTmp = leds[(LED_NUM - 1):]
		for _, l := range leds[:(LED_NUM - 1)] {
			ledsTmp = append(ledsTmp, l)
		}
	} else {
		ledsTmp = append(leds[1:], leds[0])
	}
	copy(leds[:], ledsTmp[:])
	neoPixel.WriteColors(leds[:])
}

func abs(x int16) int16 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
