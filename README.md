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
* [BBC micro:bit v1](https://tinygo.org/microcontrollers/machine/microbit/) / [v2](https://github.com/tinygo-org/tinygo/blob/release/src/machine/board_microbit-v2.go)
* [ESP8266 NodeMCU](https://tinygo.org/microcontrollers/machine/nodemcu/) / [D1 mini](https://tinygo.org/microcontrollers/machine/d1mini/)

Note:

1. Bluepills may have a built-in LED at ```PC13``` or ```PB12```.

2. BBC micro:bit v2 has 

* a buzzer at ```P27``` or ```Pin(0)```
* a capacitive touch (on the front logo) at ```P26``` or ```Pin(36)```
* the onboard LSM303AGR is on the internal I2C bus, it can be used without additional configuration. For external I2C devices, you need to use the following pins:

```golang
machine.I2C1.Configure(machine.I2CConfig{
	SCL:       machine.P19,  // machine.SCL1_PIN
	SDA:       machine.P20,  // machine.SDA1_PIN
	// Frequency: machine.TWI_FREQ_400KHZ,
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




