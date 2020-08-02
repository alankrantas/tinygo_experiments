// This is a test code for controlling servos with TinyGo.
// Basically, it uses GoRoutine to simulate continuous PWM signals.
// However, this does not work on 8-bit AVR boards, probably due to their
// lack of hardware timer issue.
// Servos also need sufficient power (5V and at least 100-200 mA each) to run smoothly.

package main

import (
	"machine"
	"time"
)

type Device struct {
	pin      machine.Pin
	on       bool
	angle    uint8
	pulseMin uint16
	pulseMax uint16
}

func Init(pin machine.Pin) Device {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return Device{pin: pin, on: false, pulseMin: 600, pulseMax: 2400}
}

func (d *Device) ServoRoutine() {
	for {
		if d.on {
			pulse := valueMapping(uint16(d.angle), 0, 180, d.pulseMin, d.pulseMax)
			d.pin.High()
			time.Sleep(time.Microsecond * time.Duration(pulse))
			d.pin.Low()
			time.Sleep(time.Millisecond * 20)
		}
	}
}

func (d *Device) Angle(angle uint8) {
	if angle < 0 {
		angle = 0
	} else if angle > 180 {
		angle = 180
	}
	d.on = true
	d.angle = angle
}

func (d *Device) PulseRange(min, max uint16) {
	d.pulseMin = min
	d.pulseMax = max
}

func (d *Device) Deinit() {
	d.on = false
}

func valueMapping(value, min, max, newMin, newMax uint16) float32 {
	scale := float32(value-min) / float32(max-min)
	return float32(newMin) + scale*float32(newMax-newMin)
}

func main() {

	servo1 := Init(machine.P8)
	servo2 := Init(machine.P12)

	go servo1.ServoRoutine()
	go servo2.ServoRoutine()

	for i := 0; i <= 2; i++ {
		servo1.Angle(0)
		servo2.Angle(0)
		time.Sleep(time.Millisecond * 1000)

		servo1.Angle(90)
		servo2.Angle(90)
		time.Sleep(time.Millisecond * 1000)

		servo1.Angle(180)
		servo2.Angle(180)
		time.Sleep(time.Millisecond * 1000)
	}

	// stop "PWM" signal (otherwise the servos may turn again when main() ended)
	servo1.Deinit()
	servo2.Deinit()

}
