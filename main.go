package main

import (
	"log"

	"github.com/mchineboy/rpi-auto-display/auto_gps"
	"github.com/mchineboy/rpi-auto-display/auto_interface"
)

func main() {
	log.Printf("%s\n", "Hello World!")
	Agps := auto_gps.New()
	go func() { Agps.Monitor() }()
	_ = auto_interface.New(Agps)
}
