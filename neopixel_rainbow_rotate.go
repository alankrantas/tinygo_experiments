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
	LED_PIN          = machine.Pin(2)
	LED_NUM          = 12
	LED_MAX_LEVEL    = 64 // max 255
	LED_ROTATE_DELAY = 50 // ms
)

func main() {

	LED_PIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	neoPixel := ws2812.New(LED_PIN)

	var leds [LED_NUM]color.RGBA
	ledChange := uint8(LED_MAX_LEVEL / (LED_NUM / 3))
	ledIndex := [3]uint8{0, uint8(LED_NUM / 3), uint8(LED_NUM / 3 * 2)}

	for i := range leds {

		led_color := [3]uint8{0, 0, 0}
		for j := range ledIndex {
			if abs(int8(uint8(i)-ledIndex[j])) <= ledIndex[1] {
				led_color[j] = LED_MAX_LEVEL - abs(int8(uint8(i)-ledIndex[j]))*ledChange
			}
		}
		if uint8(i) >= ledIndex[2] {
			led_color[0] = LED_MAX_LEVEL - (LED_NUM-uint8(i))*ledChange
		}
		leds[i] = color.RGBA{R: led_color[0], G: led_color[1], B: led_color[2]}

	}

	neoPixel.WriteColors(leds[:])
	time.Sleep(time.Millisecond * 10)

	for {

		var ledsTmp []color.RGBA

		ledsTmp = leds[1:LED_NUM]
		ledsTmp = append(ledsTmp, leds[0])
		copy(leds[:], ledsTmp[:])

		neoPixel.WriteColors(leds[:])
		time.Sleep(time.Millisecond * LED_ROTATE_DELAY)

	}

}

func abs(x int8) uint8 {
	if x < 0 {
		return uint8(-x)
	} else {
		return uint8(x)
	}
}
