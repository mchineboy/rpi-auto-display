package auto_interface

import (
	"image/color"
	"log"
)

func (AutoInt *AutoInterface) DrawX() {
	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.SetLineWidth(1.0)
	AutoInt.Screen.DrawLine(0.0, float64((AutoInt.AutoInt.Display.Width - 14)), float64(AutoInt.AutoInt.Display.Height), float64((AutoInt.AutoInt.Display.Width - 14)))
	log.Printf("%f %f %f %f\n", 0.0, float64((AutoInt.AutoInt.Display.Width-14)*-1), float64(AutoInt.AutoInt.Display.Height), float64((AutoInt.AutoInt.Display.Width-14)*-1))
	AutoInt.Screen.Stroke()
	var cx, cy = float64(AutoInt.AutoInt.Display.Width) / 2, float64(AutoInt.AutoInt.Display.Height) / 2

	var s1 = "hello"
	var hs1, _ = AutoInt.Screen.MeasureString(s1)
	var s2 = "world"
	var hs2, ws2 = AutoInt.Screen.MeasureString(s2)

	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.DrawRectangle(cx-(hs2/2)-4, cy-(ws2/2)-6, hs2+8, ws2+6)
	AutoInt.Screen.Fill()

	AutoInt.Screen.SetColor(color.Black)
	AutoInt.Screen.DrawString(s1, cx-(hs1/2), cy-ws2-8)
	AutoInt.Screen.Stroke()

	AutoInt.Screen.SetColor(color.White)
	AutoInt.Screen.DrawString(s2, cx-(hs2/2), cy)
	AutoInt.Screen.Stroke()

	if e := AutoInt.Display.Draw(AutoInt.Screen.Image()); e != nil {
		log.Printf("[ERROR] failed to draw: %v\n", e)
		AutoInt.Display.Clear(color.White)
	}
}
