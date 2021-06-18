// Read humidity/temperature from DHT11/22. Only works on BBC micro:bit V2.

package main

import (
	"errors"
	"machine"
	"runtime/interrupt"
	"time"
)

const (
	DHT11 = iota
	DHT22
)

// Device defines a dht sensor object
type Device struct {
	pin         machine.Pin
	Model       uint8
	Humidity    int32
	Temperature int32
	ChecksumOk  bool
}

// New returns a new device
func New(pin machine.Pin, model uint8) Device {
	return Device{pin: pin, Model: model}
}

// Measure reads from the device and calculate humidity/temperature values
func (d *Device) Measure() error {
	var data [40]bool

	d.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.pin.Low()
	time.Sleep(time.Millisecond * 18)
	d.pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	time.Sleep(time.Microsecond * 20)

	for d.pin.Get() {
	}
	for !d.pin.Get() {
	}
	for d.pin.Get() {
	}

	mask := interrupt.Disable()
	defer interrupt.Restore(mask)

	for i := uint8(0); i < 40; i++ {
		for !d.pin.Get() {
		}
		start := time.Now()
		for d.pin.Get() {
			if time.Since(start) > time.Microsecond*100 {
				return errors.New("sensor timeout")
			}
		}
		data[i] = time.Since(start) > time.Microsecond*40
	}

	var result [5]uint8

	for i := uint8(0); i < 5; i++ {
		for j := uint8(0); j < 8; j++ {
			if data[8*i+j] {
				result[i] += 1 << (7 - j)
			}
		}
	}

	dtaSum := uint8(result[0] + result[1] + result[2] + result[3])
	d.ChecksumOk = (dtaSum == result[4])

	if d.Model == DHT11 { // DHT11
		d.Humidity = (int32(result[0])*100 + int32(result[1])) * 10
		d.Temperature = (int32(result[2])*100 + int32(result[3])) * 10

	} else { // DHT22
		tSign := int32(1)
		if result[2] >= (1 << 7) { // negative temperature value
			result[2] -= (1 << 7)
			tSign = -1
		}
		d.Humidity = (int32(result[0])*(1<<8) + int32(result[1])) * 100
		d.Temperature = (int32(result[2])*(1<<8) + int32(result[3])) * 100 * tSign
	}
	return nil
}

func main() {

	sensor := New(machine.P0, DHT11)

	for {
		if err := sensor.Measure(); err != nil {
			println(err)
		}

		println("H:", sensor.Humidity,
			"% / T:", sensor.Temperature,
			"*C / Checksum ok:", sensor.ChecksumOk)
		time.Sleep(time.Millisecond * 2000)
	}

}
