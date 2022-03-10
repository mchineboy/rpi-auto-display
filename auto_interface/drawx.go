package auto_interface

import (
	"image/color"
	"log"
)

func (AutoInt *AutoInterface) DrawX() {
	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.SetLineWidth(2.0)
	AutoInt.Screen.DrawLine(10, -10, 100, -200)
	//AutoInt.Screen.DrawLine(0.0, float64((AutoInt.Display.Width-14)*-1), float64(AutoInt.Display.Height), float64((AutoInt.Display.Width-14)*-1))
	log.Printf("%f %f %f %f\n", 0.0, float64((AutoInt.Display.Width-14)*-1), float64(AutoInt.Display.Height), float64((AutoInt.Display.Width-14)*-1))
	AutoInt.Screen.Stroke()
	var cx, cy = float64(AutoInt.Display.Width) / 2, float64(AutoInt.Display.Height) / 2

	var s1 = "hello"
	var hs1, _ = AutoInt.Screen.MeasureString(s1)
	var s2 = "world"
	var hs2, ws2 = AutoInt.Screen.MeasureString(s2)

	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.DrawRectangle(cx-(hs2/2)-4, (cy-(ws2/2)-6)*-1, hs2+8, (ws2+6)*-1)
	AutoInt.Screen.Fill()

	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.DrawString(s1, cx-(hs1/2), (cy-ws2-8)*-1)
	AutoInt.Screen.Stroke()

	AutoInt.Screen.SetColor(color.White)
	AutoInt.Screen.DrawString(s2, cx-(hs2/2), cy*-1)
	AutoInt.Screen.Stroke()

}
