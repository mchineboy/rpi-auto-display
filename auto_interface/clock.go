package auto_interface

import (
	"fmt"
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
	AutoInt.Screen.Rotate(gg.Radians(90))
	fmt.Printf("%f, %f\n", float64(AutoInt.Display.Height)-htime, float64(AutoInt.Display.Width)-wtime)
	AutoInt.Screen.DrawStringAnchored(time, float64(AutoInt.Display.Width/2)-wtime-8, float64(AutoInt.Display.Height/2)-htime-8, .5, .5)
	AutoInt.Screen.Stroke()
	log.Printf("Update clock: %s\n", time)
}
