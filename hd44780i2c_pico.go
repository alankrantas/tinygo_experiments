package main

import (
	"machine"
	"strconv"
	"time"

	"tinygo.org/x/drivers/hd44780i2c"
)

func main() {

  // pin 27/26 are on I2C1 bus
	machine.I2C1.Configure(machine.I2CConfig{
		SCL:       machine.GPIO27,
		SDA:       machine.GPIO26,
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	lcd := hd44780i2c.New(machine.I2C1, 0x27)

	lcd.Configure(hd44780i2c.Config{
		Width:       16,
		Height:      2,
		CursorOn:    true,
		CursorBlink: true,
	})

	lcd.Print([]byte(" TinyGo\n  LCD Test "))

	lcd.CreateCharacter(0x0, []byte{0x00, 0x11, 0x0E, 0x1F, 0x15, 0x1F, 0x1F, 0x1F})
	lcd.Print([]byte{0x0})

	time.Sleep(time.Millisecond * 7000)

	for i := 0; i < 5; i++ {
		lcd.BacklightOn(false)
		time.Sleep(time.Millisecond * 250)
		lcd.BacklightOn(true)
		time.Sleep(time.Millisecond * 250)
	}

	lcd.CursorOn(false)
	lcd.CursorBlink(false)

	i := 0
	for {

		lcd.ClearDisplay()
		lcd.SetCursor(2, 1)
		lcd.Print([]byte(strconv.FormatInt(int64(i), 10)))
		i++
		time.Sleep(time.Millisecond * 100)

	}
}
