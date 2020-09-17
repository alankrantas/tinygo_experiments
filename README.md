# TinyGo Experiments

Here are some codes I write as experiments for [TinyGo](https://tinygo.org/), the Go compilier for small places (like microcontrollers).

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
