package auto_interface

func (AutoInt *AutoInterface) DrawX() {
	AutoInt.Screen.SetRGBA(1, 1, 1, 1)
	AutoInt.Screen.SetLineWidth(1.0)
	AutoInt.Screen.DrawLine(0, float64((AutoInt.Display.Height - 14)), float64(AutoInt.Display.Width), float64((AutoInt.Display.Height - 14)))
	AutoInt.Screen.Stroke()
}
