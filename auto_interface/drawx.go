package auto_interface

import (
	"image/color"
	"log"
)

func (AutoInt *AutoInterface) DrawX() {
	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.SetLineWidth(2.0)
	AutoInt.Screen.SetLineCapButt()
	AutoInt.Screen.DrawLine(10, -10, 100, -10)
	//AutoInt.Screen.DrawLine(0.0, float64((AutoInt.Display.Width-14)*-1), float64(AutoInt.Display.Height), float64((AutoInt.Display.Width-14)*-1))
	log.Printf("%f %f %f %f\n", 0.0, float64((AutoInt.Display.Width-14)*-1), float64(AutoInt.Display.Height), float64((AutoInt.Display.Width-14)*-1))
	AutoInt.Screen.Stroke()
}
