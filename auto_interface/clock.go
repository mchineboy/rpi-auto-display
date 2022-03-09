package auto_interface

import (
	"image/color"
	"log"
	"time"

	"github.com/fogleman/gg"
)

func (AutoInt *AutoInterface) DrawClock() {
	currentTime := time.Now()
	time := currentTime.Format("Mon, Jan 2 15:04")
	AutoInt.Screen.SetColor(color.Black)
	wtime, htime := AutoInt.Screen.MeasureString(time)
	log.Printf("width %f, height %f", wtime, htime)
	AutoInt.Screen.Rotate(gg.Radians(90))
	log.Printf("x: %f, y: %f\n", float64(AutoInt.Display.Height/2), float64(AutoInt.Display.Width/2))
	AutoInt.Screen.DrawStringAnchored(time, float64(AutoInt.Display.Width)-wtime, float64(AutoInt.Display.Height)-14, 0, 0)
	AutoInt.Screen.Stroke()
	log.Printf("Update clock: %s\n", time)
}
