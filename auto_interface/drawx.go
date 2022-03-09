package auto_interface

import "log"

func (AutoInt *AutoInterface) DrawX() {
	AutoInt.Screen.SetRGBA(0, 0, 0, 1)
	AutoInt.Screen.SetLineWidth(1.0)
	AutoInt.Screen.DrawLine(0, float64((AutoInt.Display.Height-14)*-1), float64(AutoInt.Display.Width), float64((AutoInt.Display.Height-14)*-1))
	log.Printf("%f %f %f %f\n", 0.0, float64((AutoInt.Display.Height-14)*-1), float64(AutoInt.Display.Width), float64((AutoInt.Display.Height-14)*-1))
	AutoInt.Screen.Stroke()
}
