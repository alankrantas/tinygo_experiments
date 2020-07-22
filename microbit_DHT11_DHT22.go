// Read humidity and temperature data from DHT11/DHT22.
//
// Due to some reason, this only works on BBC micro:bit.

package main

import (
	"machine"
	"time"
)

const (
	DHT11 = iota
	DHT22
)

type Device struct {
	pin         machine.Pin
	model       uint8
	humidity    int32
	temperature int32
	checksumOk  bool
}

func New(pin machine.Pin, model uint8) Device {
	return Device{pin: pin, model: model}
}

func (d *Device) Measure() (cks bool, err error) {

	var data [40]bool
	var result [5]uint8

	// initialize protocol
	time.Sleep(time.Millisecond * 1)
	d.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.pin.Low()
	time.Sleep(time.Millisecond * 18)
	d.pin.High()
	time.Sleep(time.Microsecond * 20)
	d.pin.Configure(machine.PinConfig{Mode: machine.PinInput})

	// wait for ready signal
	for d.pin.Get() {
	}
	for !d.pin.Get() {
	}
	for d.pin.Get() {
	}

	// read 40 bits from DHT
	for i := uint8(0); i < 40; i++ {
		for !d.pin.Get() {
		}
		time.Sleep(time.Microsecond * 28)
		if d.pin.Get() {
			data[i] = true
			if i < 39 {
				for d.pin.Get() {
				}
			}
		}
	}

	// convert bits to numbers
	for i := uint8(0); i < 5; i++ {
		for j := uint8(0); j < 8; j++ {
			if data[8*i+j] {
				result[i] += 1 << (7 - j)
			}
		}
	}

	// verify checksum
	dtaSum := uint8(result[0] + result[1] + result[2] + result[3])
	d.checksumOk = (dtaSum == result[4])
	cks = d.checksumOk

	// calculate data
	if d.model == DHT11 { // DHT11
		d.humidity = int32(int32(result[0])*100+int32(result[1])) * 10
		d.temperature = int32(int32(result[2])*100+int32(result[3])) * 10

	} else { // DHT22
		tSign := int32(1)
		if result[2] >= (1 << 7) { // negative temperature value
			result[2] -= (1 << 7)
			tSign = -1
		}
		d.humidity = int32(int32(result[0])*(1<<8)+int32(result[1])) * 100
		d.temperature = int32(int32(result[2])*(1<<8)+int32(result[3])) * 100 * tSign
	}

	err = nil
	return
}

func (d *Device) ReadHumidity() (int32, error) {
	return d.humidity, nil
}

func (d *Device) ReadTemperature() (int32, error) {
	return d.temperature, nil
}

func (d *Device) ReadChecksumOk() (bool, error) {
	return d.checksumOk, nil
}

func main() {

	dht := New(machine.P2, DHT11)

	for {

		cks, _ := dht.Measure()
		h, _ := dht.ReadHumidity()
		t, _ := dht.ReadTemperature()
		println("H:", h, "/ T:", t, "/ CKS ok:", cks)
		time.Sleep(time.Millisecond * 2000)

	}

}
