package auto_interface

import (
	"image/color"
	"time"

	"github.com/fogleman/gg"
)

func (AutoInt *AutoInterface) DrawClock() {
	currentTime := time.Now()
	time := currentTime.Format("Mon, Jan 2 15:04")
	AutoInt.Screen.SetColor(color.Black)

	_, htime := AutoInt.Screen.MeasureString(time)
	AutoInt.Screen.Rotate(gg.Radians(90))
	AutoInt.Screen.DrawStringAnchored(time, (float64(AutoInt.Display.Height) - htime), float64((AutoInt.Display.Width)*-1), .9, 1)
	AutoInt.Screen.Stroke()
}
