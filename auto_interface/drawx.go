package auto_interface

import "image/color"

func (AutoInt *AutoInterface) DrawX() {
	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.DrawLine(0, 0, float64(AutoInt.Display.Width), float64(AutoInt.Display.Height*-1))
	AutoInt.Screen.Stroke()
}
