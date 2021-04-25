// draw 5x5 icons quickly on BBC micro:bits

package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
)

type MatrixDisplay struct {
	display microbitmatrix.Device
	leds    [5][5]uint8
	enable  bool
}

func (m *MatrixDisplay) Init(rotation uint8) {
	defer func() {
		go m.Display()
	}()
	m.enable = true
	m.display = microbitmatrix.New()
	m.display.Configure(microbitmatrix.Config{Rotation: rotation})
}

func (m *MatrixDisplay) SetImage(leds [5][5]uint8) {
	m.leds = leds
}

func (m *MatrixDisplay) Display() {
	for m.enable {
		for i := int16(0); i < 5; i++ {
			for j := int16(0); j < 5; j++ {
				if m.leds[i][j] > 0 {
					m.display.SetPixel(i, j, color.RGBA{255, 255, 255, 255})
					continue
				}
				m.display.SetPixel(i, j, color.RGBA{0, 0, 0, 0})
			}
		}
		m.display.Display()
	}
}

func (m *MatrixDisplay) Clear() {
	var ledsEmpty [5][5]uint8
	m.leds = ledsEmpty
}

func (m *MatrixDisplay) Deinit() {
	m.enable = false
}

func main() {

	matrix := MatrixDisplay{}
	matrix.Init(0)
	matrix.Clear()
	time.Sleep(time.Millisecond * 100)

	for {
		matrix.SetImage([5][5]uint8{ // heart
			{0, 1, 0, 1, 0},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 0},
		})
		time.Sleep(time.Millisecond * 750)
		matrix.Clear()
		time.Sleep(time.Millisecond * 250)

		matrix.SetImage([5][5]uint8{ // smile face
			{0, 1, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 0, 0, 0},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		})
		time.Sleep(time.Millisecond * 750)
		matrix.Clear()
		time.Sleep(time.Millisecond * 250)

		matrix.SetImage([5][5]uint8{ // music note
			{0, 0, 1, 1, 0},
			{0, 0, 1, 1, 1},
			{0, 0, 1, 0, 1},
			{1, 1, 1, 0, 0},
			{1, 1, 1, 0, 0},
		})
		time.Sleep(time.Millisecond * 750)
		matrix.Clear()
		time.Sleep(time.Millisecond * 250)
	}

}
