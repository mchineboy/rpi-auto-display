package auto_interface

import (
	"fmt"
	"image/color"
)

func (AutoInt *AutoInterface) DrawSpeed() {
	AutoInt.Screen.SetColor(color.Black)

	AutoInt.Screen.SetFontFace(AutoInt.Fonts["race-36"])
	AutoInt.Screen.DrawStringAnchored(
		fmt.Sprintf("%02d", int(AutoInt.Agps.TPV.Speed*1.609)), 0, (float64(AutoInt.Display.Width) * -1), 0.5, 0.5)
	AutoInt.Screen.Stroke()
}
