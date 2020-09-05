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
		rainbowRotate()
		time.Sleep(time.Millisecond * LED_ROTATE_DELAY)
	}

}

func rainbowSet() {
	ledChange := uint8(LED_MAX_LEVEL / (LED_NUM / 3))
	ledIndex := [3]uint8{0, uint8(LED_NUM / 3), uint8(LED_NUM / 3 * 2)}
	for i := range leds {
		ledColor := [3]uint8{0, 0, 0}
		for j := range ledIndex {
			if abs(int8(uint8(i)-ledIndex[j])) <= ledIndex[1] {
				ledColor[j] = LED_MAX_LEVEL - abs(int8(uint8(i)-ledIndex[j]))*ledChange
			}
		}
		if uint8(i) >= ledIndex[2] {
			ledColor[0] = LED_MAX_LEVEL - (LED_NUM-uint8(i))*ledChange
		}
		leds[i] = color.RGBA{R: ledColor[0], G: ledColor[1], B: ledColor[2]}
	}
	neoPixel.WriteColors(leds[:])
}

func rainbowRotate() {
	var ledsTmp []color.RGBA = leds[1:LED_NUM]
	ledsTmp = append(ledsTmp, leds[0])
	copy(leds[:], ledsTmp[:])
	neoPixel.WriteColors(leds[:])
}

func abs(x int8) uint8 {
	if x < 0 {
		return uint8(-x)
	} else {
		return uint8(x)
	}
}
