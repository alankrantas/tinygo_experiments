package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const (
	ledPin      = machine.D0 // NeoPixels pin
	ledNUm      = 12         // number of NeoPixels
	ledMaxLevel = 0.5        // brightness level of NeoPxels (0~1)
)

// NeoPixels struct build upon the WS2812 driver
type NeoPixels struct {
	neo        ws2812.Device
	colors     []color.RGBA
	brightness float32
}

func main() {
	var leds NeoPixels

	leds.Init(ledPin, ledNUm, ledMaxLevel)
	leds.Clear()

	cs := []color.RGBA{
		color.RGBA{R: 255, G: 0, B: 0},
		color.RGBA{R: 0, G: 255, B: 0},
		color.RGBA{R: 0, G: 0, B: 255},
		color.RGBA{R: 255, G: 255, B: 0},
		color.RGBA{R: 0, G: 255, B: 255},
		color.RGBA{R: 255, G: 0, B: 255},
		color.RGBA{R: 255, G: 255, B: 255},
		color.RGBA{R: 0, G: 0, B: 0},
	}

	// fill
	for _, c := range cs {
		leds.Fill(c)
		leds.Show()
		time.Sleep(time.Millisecond * 250)
	}

	// chase
	for _, c := range cs {
		for i := range leds.colors {
			leds.Set(i, c)
			leds.Show()
			time.Sleep(time.Millisecond * 25)
		}
	}

	// rotate
	leds.Rainbow(0)
	leds.Show()
	for i := 0; i < 50; i++ {
		if i < 25 {
			leds.Rotate(true)
		} else {
			leds.Rotate(false)
		}
		leds.Show()
		time.Sleep(time.Millisecond * 50)
	}

	// rainbow cycle
	var pos uint8
	for {
		leds.Rainbow(pos)
		leds.Show()
		time.Sleep(time.Microsecond * 2500)
		pos++
	}
}

// Init : initialize a NeoPixels struct
func (ws *NeoPixels) Init(pin machine.Pin, ledNum int, ledLevel float32) {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws.neo = ws2812.New(pin)
	ws.colors = make([]color.RGBA, ledNum)
	ws.SetBrightness(ledLevel)
}

// SetBrightness : set brightness level of NeoPixels
func (ws *NeoPixels) SetBrightness(ledLevel float32) {
	if ledLevel > 1 {
		ledLevel = 1
	} else if ledLevel < 0 {
		ledLevel = 0
	}
	ws.brightness = ledLevel
}

// Fill : fill a color to all NeoPixels
func (ws *NeoPixels) Fill(c color.RGBA) {
	for i := range ws.colors {
		ws.colors[i] = colorLevel(c, ws.brightness)
	}
}

// FillRange : fill a color to certain NeoPIxels
func (ws *NeoPixels) FillRange(c color.RGBA, start, end int) {
	for i := range ws.colors {
		if i >= start && i <= end {
			ws.colors[i] = colorLevel(c, ws.brightness)
		}
	}
}

// Set : set a color to a specific NeoPixel
func (ws *NeoPixels) Set(i int, c color.RGBA) {
	if i >= 0 && i < len(ws.colors) {
		ws.colors[i] = colorLevel(c, ws.brightness)
	}
}

// Rotate : rotate NeoPixels
func (ws *NeoPixels) Rotate(clockwise bool) {
	if clockwise {
		ws.colors = append(ws.colors[(len(ws.colors)-1):], ws.colors[:(len(ws.colors)-1)]...)
	} else {
		ws.colors = append(ws.colors[1:], ws.colors[0])
	}
}

// Rainbow : fill rainbow colors to NeoPixels
func (ws *NeoPixels) Rainbow(pos uint8) {
	for i := range ws.colors {
		rcIndex := int(i*256/len(ws.colors)) + int(pos)
		ws.colors[i] = colorLevel(wheel(uint8(rcIndex)), ws.brightness)
	}
}

// Clear : clear all NeoPixels
func (ws *NeoPixels) Clear() {
	ws.Fill(color.RGBA{R: 0, G: 0, B: 0})
}

// Show : write colors to NeoPixels to take effect
func (ws *NeoPixels) Show() {
	ws.neo.WriteColors(ws.colors)
}

// return a rainbow color in a specific position
func wheel(pos uint8) color.RGBA {
	var r, g, b uint8
	switch {
	case pos < 0 || pos > 255:
		r = 0
		g = 0
		b = 0
	case pos < 85:
		r = 255 - pos*3
		g = pos * 3
		b = 0
	case pos < 170:
		pos -= 85
		r = 0
		g = 255 - pos*3
		b = pos * 3
	default:
		pos -= 170
		r = pos * 3
		g = 0
		b = 255 - pos*3
	}
	return color.RGBA{R: r, G: g, B: b}
}

// adjust color brightness level
func colorLevel(c color.RGBA, level float32) color.RGBA {
	c.R = uint8(float32(c.R) * level)
	c.G = uint8(float32(c.G) * level)
	c.B = uint8(float32(c.B) * level)
	return c
}
