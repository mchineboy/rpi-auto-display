package auto_interface

import (
	"fmt"
	"image/color"
	"log"

	"github.com/fogleman/gg"
)

func (AutoInt *AutoInterface) DrawGPS() {
	gpsstring := fmt.Sprintf("Lat: %0.6f Lon: %0.6f Ele: %0.0f'", AutoInt.Agps.TPV.Lat, AutoInt.Agps.TPV.Lon, AutoInt.Agps.TPV.Alt*3.281)
	log.Printf("%s\n", gpsstring)
	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.Rotate(gg.Radians(90))
	AutoInt.Screen.DrawStringAnchored(gpsstring, 0, -14, 0, 1)
	AutoInt.Screen.Stroke()
}
