package main

import (
	"machine"
	"time"
)

// Mic represents BBC micro:bit V2's onboard mic
type Mic struct {
	run machine.Pin
	in  machine.ADC
}

// Init initializes mic pins
func (mic *Mic) Init() {
	mic.run = machine.MIC_LED
	mic.run.Configure(machine.PinConfig{Mode: machine.PinOutput})
	mic.run.High()
	machine.InitADC()
	mic.in = machine.ADC{Pin: machine.MIC}
	mic.in.Configure(machine.ADCConfig{})
}

// MeasureVolume measures mic level
func (mic *Mic) MeasureVolume(samples uint16, gain float32) uint16 {
	var min, max uint16
	for i := uint16(0); i < samples; i++ {
		value := mic.in.Get()
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	output := uint32(float32(max-min) * gain)
	if output > 65535 {
		return uint16(65535)
	}
	return uint16(output)
}

func main() {
	mic := Mic{}
	mic.Init()
	for {
		println(mic.MeasureVolume(1024, 6.0))
		time.Sleep(time.Millisecond * 50)
	}
}
