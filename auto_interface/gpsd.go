package auto_interface

import (
	"fmt"
	"image/color"
	"log"

	"github.com/fogleman/gg"
)

var signalMode = []string{
	"unknown", "no fix", "2D", "3D",
}

func (AutoInt *AutoInterface) DrawGPS() {

	AutoInt.Screen.SetFontFace(AutoInt.Fonts["roboto-12"])
	gpsstring := fmt.Sprintf("Lat: %0.4f Lon: %0.4f Ele: %0.0f'", AutoInt.Agps.TPV.Lat, AutoInt.Agps.TPV.Lon, AutoInt.Agps.TPV.Alt*3.281)
	log.Printf("%s\n", gpsstring)
	AutoInt.Screen.SetColor(color.Black)
	//AutoInt.Screen.Rotate(gg.Radians(90))
	_, htime := AutoInt.Screen.MeasureString(gpsstring)
	AutoInt.Screen.DrawStringAnchored(gpsstring, (float64(AutoInt.Display.Height) - htime), -14, .95, 1)
	AutoInt.Screen.Stroke()

	signalstrength := fmt.Sprintf("%s %0.1fx %0.1fy", signalMode[int(AutoInt.Agps.TPV.Mode)], AutoInt.Agps.TPV.Epx*3.281, AutoInt.Agps.TPV.Epy*3.281)

	AutoInt.Screen.DrawStringAnchored(signalstrength, 0, float64((AutoInt.Display.Width)*-1), 0, 1)
	AutoInt.Screen.Stroke()

	locations := AutoInt.Agps.FindNearestTowns(AutoInt.Agps.TPV.Lat, AutoInt.Agps.TPV.Lon)
	for i, loc := range locations {
		AutoInt.Screen.DrawStringAnchored(
			loc, float64(AutoInt.Display.Height)/2, (float64(AutoInt.Display.Width-((i+2)*14)))*-1, 0, 1)

	}
	AutoInt.Screen.Rotate(gg.Radians(0)) // Reset rotate once completed
}
