package auto_interface

import "fmt"

func (AutoInt *AutoInterface) DrawSpeed() {
	AutoInt.Screen.SetRGB(0, 0, 0)

	AutoInt.Screen.SetFontFace(AutoInt.Fonts["race-36"])
	AutoInt.Screen.DrawStringAnchored(
		fmt.Sprintf("%02d", int(AutoInt.Agps.TPV.Speed*1.609)), 0, (float64(AutoInt.Display.Width/2) * -1), 0.0, 0.5)
	AutoInt.Screen.Stroke()
}
