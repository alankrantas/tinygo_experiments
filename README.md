# TinyGo Experiments

Some experiment codes written in [TinyGo](https://tinygo.org/), the Go compilier for small places (like microcontrollers).

```golang
package main

import (
	"machine"
	"time"
)

func main() {

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.Set(!led.Get())
		time.Sleep(time.Millisecond * 500)
	}

}
```

## Links for machine packages of some boards

* [Arduino Uno](https://tinygo.org/microcontrollers/machine/arduino/) / [Nano](https://tinygo.org/microcontrollers/machine/arduino-nano/)
* [Arduino Nano 33 IoT](https://tinygo.org/microcontrollers/machine/arduino-nano33/)
* [Bluepill](https://tinygo.org/microcontrollers/machine/bluepill/)
* [BBC micro:bit v1](https://tinygo.org/microcontrollers/machine/microbit/) / [v2](https://tinygo.org/microcontrollers/machine/microbit-v2/)
* [ESP8266 NodeMCU](https://tinygo.org/microcontrollers/machine/nodemcu/) / [D1 mini](https://tinygo.org/microcontrollers/machine/d1mini/)

Note:

1. Bluepills may have a built-in LED at ```PC13``` or ```PB12```.

2. BBC micro:bit v2 has 

* a buzzer at ```P27```
* a capacitive touch (on the front logo) at ```P26``` (no need to pull up and pin.Get() returns ```false``` when touched)
* a mic at ```P29``` and its enable pin/indicator LED at ```P28```. See [this script](https://github.com/alankrantas/tinygo_experiments/blob/master/microbit_v2_mic_level.go) for demostration.
* the ```LED``` pin is just a dummy
* the onboard LSM303AGR is on the internal I2C bus (```I2C0```). For external I2C devices, you need to use the following pins:

```golang
machine.I2C1.Configure(machine.I2CConfig{
	SCL:       machine.P19, // machine.SCL1_PIN
	SDA:       machine.P20, // machine.SDA1_PIN
	Frequency: machine.TWI_FREQ_400KHZ,
})
```

## Download TinyGo drivers

```
go get -u tinygo.org/x/drivers
```

After Go 1.16 you need to use Go Modules to set the module path (module is the collection of packages, including package ```main```) and thus be able to find downloaded external packages:

```
go mod init <project name>
go mod tidy
```

Use ```go mod tidy``` after you add or remove an external package.
