package auto_interface

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

func (AutoInt *AutoInterface) DrawGPS() {
	gpsstring := fmt.Sprintf("Lat: %0.4f Lon: %0.4f Ele: %0.4f", AutoInt.Agps.TPV.Lat, AutoInt.Agps.TPV.Lon, AutoInt.Agps.TPV.Alt)
	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.Rotate(gg.Radians(90))
	AutoInt.Screen.DrawStringAnchored(gpsstring, 0, -14, 0, 0)
	AutoInt.Screen.Stroke()
}