package auto_interface

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.com/stianeikeland/go-rpio/v4"
	"go.riyazali.net/epd"
)

type AutoInterface struct {
	Display *epd.EPD
	Screen  *gg.Context
}

type ReadablePinPatch struct{ rpio.Pin }

func (pin ReadablePinPatch) Read() uint8 { return uint8(pin.Pin.Read()) }

func init() {
	//start the GPIO controller
	if err := rpio.Open(); err != nil {
		log.Fatalf("[FATAL] failed to start gpio: %v", err)
	}

	// Enable SPI on SPI0
	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		log.Fatalf("[FATAL] failed to enable SPI: %v", err)
	}

	// configure SPI settings
	rpio.SpiSpeed(4_000_000)
	rpio.SpiMode(0, 0)

	rpio.Pin(17).Mode(rpio.Output)
	rpio.Pin(25).Mode(rpio.Output)
	rpio.Pin(8).Mode(rpio.Output)
	rpio.Pin(24).Mode(rpio.Input)
}

func New() *AutoInterface {
	defer rpio.Close()
	AutoInt := &AutoInterface{Display: epd.New(rpio.Pin(17), rpio.Pin(25), rpio.Pin(8), ReadablePinPatch{rpio.Pin(24)}, rpio.SpiTransmit)}
	fmt.Printf("%d %d", AutoInt.Display.Height, AutoInt.Display.Width)
	AutoInt.Display.Mode(epd.PartialUpdate)
	AutoInt.Display.Clear(color.White)
	AutoInt.Screen = gg.NewContext(AutoInt.Display.Width, AutoInt.Display.Height)
	AutoInt.Screen.SetColor(color.White)
	AutoInt.Screen.Clear()

	ticker := time.NewTicker(5 * time.Second)

	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				AutoInt.DrawClock()
				AutoInt.Display.Draw(AutoInt.Screen.Image())
				AutoInt.Display.Sleep()
			}
		}
	}()
	return AutoInt
}

func (AutoInt *AutoInterface) ClearScreen() {
	AutoInt.Display.Clear(color.Opaque)
}
