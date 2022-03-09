package auto_interface

import (
	"image/color"
	"time"
)

func (AutoInt *AutoInterface) DrawClock() {
	currentTime := time.Now()
	time := currentTime.Format("Mon Jan 2 15:04:05")
	AutoInt.Screen.SetColor(color.Black)
	htime, wtime := AutoInt.Screen.MeasureString(time)
	AutoInt.Screen.DrawString(time, float64(AutoInt.Display.Height)-htime, float64(AutoInt.Display.Width)-wtime)
	AutoInt.Screen.Stroke()
}
