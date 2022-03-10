package main

import (
	"github.com/mchineboy/rpi-auto-display/auto_gps"
	"github.com/mchineboy/rpi-auto-display/auto_interface"
)

func main() {
	Agps := auto_gps.New()
	go func() { Agps.Monitor() }()
	_ = auto_interface.New(Agps)
}
