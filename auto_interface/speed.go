package auto_interface

import (
	"fmt"
	"image/color"
	"log"
)

func (AutoInt *AutoInterface) DrawSpeed() {
	AutoInt.Screen.SetColor(color.Black)

	AutoInt.Screen.SetFontFace(AutoInt.Fonts["race-36"])
	log.Printf("%02d", int(AutoInt.Agps.TPV.Speed*1.609))
	AutoInt.Screen.DrawStringAnchored(
		fmt.Sprintf("%02d", int(AutoInt.Agps.TPV.Speed*1.609)), 1, (float64(AutoInt.Display.Width) * -1), 0.0, 1.0)
	AutoInt.Screen.Stroke()
}
