package auto_interface

import (
	"image/color"
	"time"

	"github.com/fogleman/gg"
)

func (AutoInt *AutoInterface) DrawClock() {

	AutoInt.Screen.SetFontFace(AutoInt.Fonts["race-10"])
	loc, _ := time.LoadLocation("Local")
	currentTime := time.Now().In(loc)
	time := currentTime.Format("Mon, Jan 2 3:04pm")
	AutoInt.Screen.SetColor(color.Black)

	_, htime := AutoInt.Screen.MeasureString(time)
	AutoInt.Screen.Rotate(gg.Radians(90))
	AutoInt.Screen.DrawStringAnchored(time, (float64(AutoInt.Display.Height) - htime),
		float64((AutoInt.Display.Width)*-1), .9, 1)
	AutoInt.Screen.Stroke()
	AutoInt.Screen.Rotate(gg.Radians(0)) // Reset rotate once completed
}
