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
## Setup TinyGo

### Install [Golang](https://go.dev/dl/)

For Windows, simply download ```go<version>.windows-amd64.msi``` and run.

For Linux, find the correct target and

```bash
wget https://go.dev/dl/go{version}.linux-{architecture}.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go{version}.linux-{architecture}.tar.gz
```

Single-board computers with ARMv8 is ```arm64``` and ARMv7/v6 is ```arm6vl```.

### Install [TinyGo](https://github.com/tinygo-org/tinygo/releases)

For Windows, simply download ```tinygo<version>.windows-amd64.zip``` and unzip.

For Linux, find the correct target and

```bash
wget https://github.com/tinygo-org/tinygo/releases/download/v{version}/tinygo_{version}_{architecture}.deb
sudo dpkg -i tinygo_{version}_{architecture}.deb
```

### Setup $PATH

For the first time installing TinyGo, Windoes users have to add ```<path>\TinyGo\bin``` to $PATH.

For Linux:

```bash
sudo nano ~/.bashrc
```

Add the following two lines at the end

```bash
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/usr/local/tinygo/bin
```

Press ```Ctrl``` + ```X``` then ```Y``` then ```enter``` to save, reopen the terminal or apply the changes with

```bash
source ~/.bashrc
```

Check $PATH:

```bash
go version
tinygo version
```

### Install AVR Tools (Optional)

8-bit AVR boards require additional compiling and flashing tools.

For Windows users, download and unzip [AVR-GCC](https://blog.zakkemble.net/avr-gcc-builds/) then add ```<path>\avr-gcc-<version>-x64-windows\bin``` to $PATH, then reopen the terminal.

For Linux, run

```bash
sudo apt-get install gcc-avr avr-libc avrdude
```

### Install TinyGo Packages (Optional)

```bash
go get -u tinygo.org/x/drivers
go get -u tinygo.org/x/tinydraw
go get -u tinygo.org/x/tinyfont
go get -u tinygo.org/x/bluetooth
```

## Run TinyGo Script

### Setup Project

1. Create a directory. It is recommended to open it in [VS Code](https://code.visualstudio.com/download) with the [TinyGo extension](https://marketplace.visualstudio.com/items?itemName=tinygo.vscode-tinygo) installed.
2. Create ```main.go```.
3. Initialize the Go project with

```bash
go mod init <project name>
```

### Compile and Flash Script

Find out the [target name](https://tinygo.org/docs/reference/microcontrollers/) (some are not listed and has to lool up [here](https://github.com/tinygo-org/tinygo/tree/release/src/machine)) and port of your device.

To look up the port, run ```mode``` on Windows and ```dmesg | grep tty``` on Linux.

> For devices using UF2 firmware, you don't need to specify the port, although you'll have to make sure the device has entered bootloader mode.
>
> For most SAMD21/51 and nRF boards with UF2 firmwares, pressing the reset button twice. For Raspberry Pi Pico you'll need to hold down the BOOTSEL button and re-connect USB. The micro:bits has a separate chip as USB interface so it can be flashed directly.
>
>For Arduino Nano 33 IoT, you can install the UF2 bootloader with [this Arduino script](https://github.com/adafruit/uf2-samdx1/releases/download/v3.14.0/update-bootloader-nano33iot-v3.14.0.ino). (After that you can upgrade the boodloader with .uf2 files.)

If external drivers are used, update ```go.mod``` with

```bash
go mod tidy
```

Then run

```bash
tinygo flash -target <target> -port <port> main.go
# or
tinygo flash -target=<target> -port=<port> main.go
```

> For DigiSpark (ATTiny85), install the bootloader (On Windoes use [Zadig](https://github.com/micronucleus/micronucleus/tree/master/windows_driver_installer)) and download [micronucleus](https://github.com/micronucleus/micronucleus/tree/master/commandline). Add it to your $PATH. TinyGo will ask you to re-connect DigiSpark before flashing.
>
> For Bluepills, get a ST-LINK/V2 programmer and download/unzip [OpenOCD](https://freddiechopin.info/en/download/category/4-openocd), add the ```<path>\openocd-0.10.0\bin``` to $PATH. Port is not needed but *you may have to hold down reset button before flashing and release as soon as OpenOCD's first message appears*...

## Notes on Pins and I2C Bus

See the [machine package definitions](https://github.com/tinygo-org/tinygo/tree/release/src/machine) and board pinouts for the correct pin number and I2C/SPI bus. For some targets I2C. SPI and/or PWM may not be supported yet.

### micro:bit V2

For both V1 and V2, the onboard LSM303AGR is connected to the internal I2C bus (```I2C0```). However for external I2C devices, V1 uses the same bus while V2 uses ```I2C1```:

```golang
machine.I2C1.Configure(machine.I2CConfig{
	SCL:       machine.SCL1_PIN, // P19
	SDA:       machine.SDA1_PIN, // P20
	Frequency: machine.TWI_FREQ_400KHZ,
})
```

V2 also has the following additional pins:

* a buzzer at ```BUZZER```
* a capacitive touch (on the logo) at ```CAP_TOUCH``; do not need to be pulled up and ```pin.Get()``` returns ```false``` when touched)
* a mic at ```MIC``` and its enable pin/indicator LED at ```MIC_LED```. See [this script](https://github.com/alankrantas/tinygo_experiments/blob/master/microbit_v2_mic_level.go) for demostration.
* the ```LED``` pin is just a dummy

### Bluepills

The built-in LED may be at ```PC13``` or ```PB12```.

### Arduino Nano 33 BLE Sense

The board's onboard I2C sensors are connected to the internal ```I2C1``` bus which has to be powered up with ```P0_22``` and pulled up with ```P1_00```, although the TinyGo driver implementations already dealt with the issue.
