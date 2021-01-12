// Generate rainbow colors on NeoPixels (WS2812) and rotate them
// (based on Adafruit's version with smoother cycle effect)
//
// TinyGo ws2812 drivers required

package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const (
	LED_PIN         = machine.D2
	LED_NUM         = 12
	LED_MAX_LEVEL   = 64 // max 255
	LED_CYCLE_DELAY = 5  // ms
)

var neoPixel ws2812.Device
var leds []color.RGBA = make([]color.RGBA, LED_NUM)

func main() {

	LED_PIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	neoPixel = ws2812.New(LED_PIN)

	var cycle uint8
	for {
		rainbowCycle(cycle)
		cycle++
		time.Sleep(time.Millisecond * LED_CYCLE_DELAY)
	}

}

func wheel(pos uint8) color.RGBA {
	var r, g, b uint8
	if pos < 0 || pos > 255 {
		r = 0
		g = 0
		b = 0
	} else if pos < 85 {
		r = 255 - pos*3
		g = pos * 3
		b = 0
	} else if pos < 170 {
		pos -= 85
		r = 0
		g = 255 - pos*3
		b = pos * 3
	} else {
		pos -= 170
		r = pos * 3
		g = 0
		b = 255 - pos*3
	}
	r = uint8(float32(r) * float32(LED_MAX_LEVEL) / 255)
	g = uint8(float32(g) * float32(LED_MAX_LEVEL) / 255)
	b = uint8(float32(b) * float32(LED_MAX_LEVEL) / 255)
	return color.RGBA{R: r, G: g, B: b}
}

func rainbowCycle(cycle uint8) {
	for i := range leds {
		rcIndex := int(i*256/LED_NUM) + int(cycle)
		leds[i] = wheel(uint8(rcIndex))
	}
	neoPixel.WriteColors(leds)
}
