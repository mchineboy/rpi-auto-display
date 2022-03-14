package auto_interface

import "github.com/stianeikeland/go-rpio/v4"

func (AutoInt *AutoInterface) SetLed() {
	rpio.WritePin(rpio.Pin(13), rpio.High)
	rpio.WritePin(rpio.Pin(19), rpio.High)
	rpio.WritePin(rpio.Pin(26), rpio.High)
}
