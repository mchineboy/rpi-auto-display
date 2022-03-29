package auto_interface

import (
	"fmt"
	"image/color"
	"log"
)

func (AutoInt *AutoInterface) DrawSpeed() {
	AutoInt.Screen.SetColor(color.Black)

	AutoInt.Screen.SetFontFace(AutoInt.Fonts["digital-96"])
	log.Printf("%02d", int(AutoInt.Agps.TPV.Speed*1.609))
	AutoInt.Screen.DrawStringAnchored(
		fmt.Sprintf("%02d", int(AutoInt.Agps.TPV.Speed*1.609)), 8, (float64(AutoInt.Display.Width-7) * -1), 0.0, 1.0)
	AutoInt.Screen.Stroke()

	AutoInt.Screen.SetFontFace(AutoInt.Fonts["race-18"])
	compass := AutoInt.Agps.GetCompassDirection(AutoInt.Agps.TPV.Track)
	AutoInt.Screen.DrawStringAnchored(
		fmt.Sprintf("%6s", compass), 0, -20, 0, 0.0,
	)
	AutoInt.Screen.Stroke()
}
