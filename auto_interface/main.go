package auto_interface

import (
	"image/color"
	"log"
	"time"

	"github.com/fogleman/gg"
	"github.com/mchineboy/rpi-auto-display/auto_gps"
	"github.com/stianeikeland/go-rpio/v4"
	"go.riyazali.net/epd"
	"golang.org/x/image/font"
)

type AutoInterface struct {
	Display  *epd.EPD
	Screen   *gg.Context
	TimeZone string
	Agps     *auto_gps.AutoGps
	Fonts    map[string]font.Face
}

type ReadablePinPatch struct{ rpio.Pin }

func (pin ReadablePinPatch) Read() uint8 { return uint8(pin.Pin.Read()) }

func init() {
	//start the GPIO controller
	if err := rpio.Open(); err != nil {
		log.Printf("[FATAL] failed to start gpio: %v", err)
	}

	// Enable SPI on SPI0
	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		log.Printf("[FATAL] failed to enable SPI: %v", err)
	}

	// configure SPI settings
	rpio.SpiSpeed(4_000_000)
	rpio.SpiMode(0, 0)

	rpio.Pin(17).Mode(rpio.Output)
	rpio.Pin(25).Mode(rpio.Output)
	rpio.Pin(8).Mode(rpio.Output)
	rpio.Pin(24).Mode(rpio.Input)

	// Red
	rpio.Pin(13).Mode(rpio.Output)
	// Green
	rpio.Pin(19).Mode(rpio.Output)
	// Blue
	rpio.Pin(26).Mode(rpio.Output)
}

func New(Agps *auto_gps.AutoGps) *AutoInterface {
	defer rpio.Close()
	AutoInt := &AutoInterface{Display: epd.New(rpio.Pin(17), rpio.Pin(25), rpio.Pin(8), ReadablePinPatch{rpio.Pin(24)}, rpio.SpiTransmit)}
	log.Printf("Width %d, Height %d\n", AutoInt.Display.Width, AutoInt.Display.Height)
	AutoInt.TimeZone = "US/Pacific"
	AutoInt.Display.Mode(epd.FullUpdate)
	AutoInt.ClearScreen()
	AutoInt.Display.Mode(epd.PartialUpdate)
	AutoInt.LoadFonts()
	AutoInt.Agps = Agps
	ticker := time.NewTicker(3 * time.Second)

	for {
		<-ticker.C
		AutoInt.Screen = gg.NewContext(AutoInt.Display.Width, AutoInt.Display.Height)
		AutoInt.Screen.SetColor(color.White)
		AutoInt.Screen.Clear()
		AutoInt.DrawClock()
		AutoInt.DrawX()
		AutoInt.DrawGPS()
		AutoInt.DrawSpeed()
		err := AutoInt.Display.Draw(AutoInt.Screen.Image())
		if err != nil {
			log.Printf("%+v\n", err)
		}
	}
}

func (AutoInt *AutoInterface) ClearScreen() {
	AutoInt.Display.Clear(color.White)
}
