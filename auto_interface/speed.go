package auto_interface

import "fmt"

func (AutoInt *AutoInterface) DrawSpeed() {
	AutoInt.Screen.SetRGB(0, 0, 0)
	if err := AutoInt.Screen.LoadFontFace("/fonts/race.ttf", 36); err != nil {
		panic(err)
	}
	AutoInt.Screen.DrawStringAnchored(
		fmt.Sprintf("%02d", int(AutoInt.Agps.TPV.Speed*1.609)), 0, float64(AutoInt.Display.Width-50)*-1, 0.0, 0.0)
	AutoInt.Screen.Stroke()
}
