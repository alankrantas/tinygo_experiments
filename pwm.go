package main

import (
	"machine"
	"time"
)

func main() {

	pwm := machine.Timer1
	pin := machine.D9

	if err := pwm.Configure(machine.PWMConfig{}); err != nil {
		println("failed to configure PWM")
		return
	}

	channel, err := pwm.Channel(pin)
	if err != nil {
		println("failed to configure channel")
		return
	}

	for {

		for i := uint32(50); i >= 1; i-- {
			pwm.Set(channel, pwm.Top()/i)
			time.Sleep(time.Millisecond * 50)
		}

		for i := uint32(1); i <= 50; i++ {
			pwm.Set(channel, pwm.Top()/i)
			time.Sleep(time.Millisecond * 50)
		}

	}

}
