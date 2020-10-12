// An experimental TinyGo servo driver.
// Basically, it uses GoRoutine to simulate continuous PWM signals.
// However, this does not work on 8-bit AVR boards, probably due to their
// lack of hardware timer support.
// Servos also require sufficient power (5V and at least 100-200 mA each)
// to run smoothly.

package main

import (
	"machine"
	"time"
)

type Device struct {
	pin      machine.Pin
	attached bool
	angle    int16
	pulseMin int32
	pulseMax int32
}

func Attach(pin machine.Pin) Device {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return Device{pin: pin, attached: false, pulseMin: 500, pulseMax: 2500}
}

func (d *Device) ServoRoutine() {
	for {
		if d.attached {
			pulse := valueMapping(int32(d.angle), 0, 180, d.pulseMin, d.pulseMax)
			d.pin.High()
			time.Sleep(time.Microsecond * time.Duration(pulse))
			d.pin.Low()
			time.Sleep(time.Microsecond * time.Duration(20000.0-pulse))
		}
	}
}

func (d *Device) Write(angle int16) {
	if angle < 0 {
		angle = 0
	} else if angle > 180 {
		angle = 180
	}
	d.attached = true
	d.angle = angle
}

func (d *Device) Read() int16 {
	return d.angle
}

func (d *Device) Attached() bool {
	return d.attached
}

func (d *Device) Detach() {
	d.attached = false
}

func (d *Device) PulseRange(min, max int32) {
	d.pulseMin = min
	d.pulseMax = max
}

func valueMapping(value, min, max, newMin, newMax int32) float32 {
	scale := float32(value-min) / float32(max-min)
	return float32(newMin) + scale*float32(newMax-newMin)
}

func main() {

	servo := Attach(machine.D2)

	// enable "PWM" function
	go servo.ServoRoutine()

	for i := 0; i <= 2; i++ {
		servo.Write(0) // angle 0
		time.Sleep(time.Millisecond * 1000)

		servo.Write(90) // angle 90
		time.Sleep(time.Millisecond * 1000)

		servo.Write(180) // angle 180
		time.Sleep(time.Millisecond * 1000)
	}

	// stop "PWM"
	servo.Detach()

}
