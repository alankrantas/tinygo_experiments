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

type Servo struct {
	pin      machine.Pin
	Attached bool
	Angle    int16
	PulseMin int16
	PulseMax int16
}

func (s *Servo) Init(pin machine.Pin) {
	defer func() {
		go s.servoRoutine()
	}()
	s.pin = pin
	s.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	s.PulseMin = 600
	s.PulseMax = 2400
}

func (s *Servo) servoRoutine() {
	for {
		if s.Attached {
			pulse := valueMapping(int16(s.Angle), 0, 180, s.PulseMin, s.PulseMax)
			s.pin.High()
			time.Sleep(time.Microsecond * time.Duration(pulse))
			s.pin.Low()
			time.Sleep(time.Microsecond * time.Duration(20000.0-pulse))
		}
	}
}

func (s *Servo) Write(angle int16) {
	if angle < 0 {
		angle = 0
	} else if angle > 180 {
		angle = 180
	}
	s.Attached = true
	s.Angle = angle
}

func (s *Servo) Detach() {
	s.Attached = false
}

func valueMapping(value, min, max, newMin, newMax int16) float32 {
	scale := float32(value-min) / float32(max-min)
	return float32(newMin) + scale*float32(newMax-newMin)
}

func main() {

	servo := Servo{}
	servo.Init(machine.P0)

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
