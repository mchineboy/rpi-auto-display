package auto_interface

import (
	"image/color"
)

func (AutoInt *AutoInterface) DrawX() {
	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.SetLineWidth(2.0)
	AutoInt.Screen.DrawLine(0.0, 16-(float64(AutoInt.Display.Width)), float64(AutoInt.Display.Height), 16-(float64(AutoInt.Display.Width)))
	AutoInt.Screen.Stroke()
}
