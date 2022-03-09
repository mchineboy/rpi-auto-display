package auto_interface

import (
	"image/color"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func (AutoInt *AutoInterface) DrawClock() {
	currentTime := time.Now()
	time := currentTime.Format("Mon, Jan 2 15:04")
	AutoInt.Screen.SetColor(color.Black)
	font, err := truetype.Parse(goregular.TTF)

	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(font, &truetype.Options{
		Size: 10,
	})
	AutoInt.Screen.SetFontFace(face)
	_, htime := AutoInt.Screen.MeasureString(time)
	AutoInt.Screen.Rotate(gg.Radians(90))
	AutoInt.Screen.DrawStringAnchored(time, (float64(AutoInt.Display.Height) - htime), float64((AutoInt.Display.Width)*-1), 1, 1)
	AutoInt.Screen.Stroke()
}
