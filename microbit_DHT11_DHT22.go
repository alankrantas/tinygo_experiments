// Read humidity and temperature data from DHT11/DHT22.
//
// Due to some reason, this only works on BBC micro:bit.

package main

import (
	"machine"
	"strconv"
	"time"
)

const (
	DHT11 = iota
	DHT22
)

func dhtMasure(pin machine.Pin, model uint8) (h int32, t int32, cks bool) {

	var data [40]bool
	var result [5]uint8

	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pin.Low()
	time.Sleep(time.Microsecond * 18000)
	pin.High()
	pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	time.Sleep(time.Microsecond * 20)

	for pin.Get() {
	}
	for !pin.Get() {
	}
	for pin.Get() {
	}

	for i := 0; i < 40; i++ {
		for pin.Get() {
		}
		for !pin.Get() {
		}
		time.Sleep(time.Microsecond * 28)
		if pin.Get() {
			data[i] = true
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 8; j++ {
			if data[8*i+j] {
				result[i] += 1 << (7 - j)
			}
		}
	}

	dtaSum := uint8(result[0] + result[1] + result[2] + result[3])
	cks = (dtaSum == result[4])

	if model == DHT11 {
		h = int32(int32(result[0])*100+int32(result[1])) * 10
		t = int32(int32(result[2])*100+int32(result[3])) * 10

	} else {
		tSign := int32(1)
		if result[2] >= (1 << 7) {
			result[2] -= (1 << 7)
			tSign = -1
		}
		h = int32(int32(result[0])*(1<<8)+int32(result[1])) * 100
		t = int32(int32(result[2])*(1<<8)+int32(result[3])) * 100 * tSign
	}

	return
}

func main() {

	for {

		h, t, cks := dhtMasure(machine.P2, DHT11)

		h_str := strconv.FormatFloat(float64(h)/1000, 'f', 0, 64) + "%"
		t_str := strconv.FormatFloat(float64(t)/1000, 'f', 2, 64) + " *C"
		println("Humidity:", h_str, "/ temperature:", t_str, "/ checksum ok:", cks)

		time.Sleep(time.Millisecond * 2000)

	}

}
