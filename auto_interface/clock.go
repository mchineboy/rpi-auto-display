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
	time := currentTime.Format("Jan 2 15:04")
	AutoInt.Screen.SetColor(color.Black)
	wtime, htime := AutoInt.Screen.MeasureString(time)
	AutoInt.Screen.Rotate(gg.Radians(90))
	fmt.Printf("%f, %f\n", float64(AutoInt.Display.Height)-htime, float64(AutoInt.Display.Width)-wtime)
	AutoInt.Screen.DrawString(time, float64(AutoInt.Display.Width)-wtime, float64(AutoInt.Display.Height)-htime-8)
	AutoInt.Screen.Stroke()
	log.Printf("Update clock: %s\n", time)
}
