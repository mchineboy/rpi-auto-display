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
	/*
		font, err := truetype.Parse(goregular.TTF)

		if err != nil {
			panic(err)
		}

		face := truetype.NewFace(font, &truetype.Options{
			Size: 10,
		})
		AutoInt.Screen.SetFontFace(face)
	*/
	wtime, htime := AutoInt.Screen.MeasureString(time)

	log.Printf("width %f, height %f", wtime, htime)
	AutoInt.Screen.Rotate(gg.Radians(90))
	log.Printf("x: %f, y: %f\n", float64(AutoInt.Display.Height)-htime, float64((AutoInt.Display.Width-14)*-1))
	AutoInt.Screen.DrawString(time, (float64(AutoInt.Display.Height)-htime)/2, float64((AutoInt.Display.Width-14)*-1))
	AutoInt.Screen.Stroke()
	log.Printf("Update clock: %s\n", time)
}
