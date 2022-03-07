package main

import "github.com/mchineboy/rpi-auto-display/auto_interface"

func main() {
	AutoInt := auto_interface.New()
	AutoInt.ClearScreen()
}
