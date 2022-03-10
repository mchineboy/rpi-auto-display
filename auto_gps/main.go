package auto_gps

import (
	"errors"
	"fmt"

	"github.com/amenzhinsky/go-gpsd"
)

type AutoGps struct {
	TPV *gpsd.TPV
}

func New() *AutoGps {
	agps := &AutoGps{}
	return agps
}

func (Agps *AutoGps) Monitor() {
	g, err := gpsd.Dial("locahost:2947")
	if err != nil {
		panic(err)
	}
	defer g.Close()

	if err := g.Stream(gpsd.WATCH_ENABLE|gpsd.WATCH_JSON, ""); err != nil {
		panic(err)
	}
	defer g.Stream(gpsd.WATCH_DISABLE, "")

	for v := range g.C() {
		switch t := v.(type) {
		case *gpsd.VERSION:
			fmt.Printf("GPSD Version: %s, Proto: %.0f.%.0f\n", t.Release, t.ProtoMajor, t.ProtoMinor)
		case *gpsd.DEVICES:
			if len(t.Devices) == 0 {
				panic(errors.New("no devices available"))
			}
			fmt.Println("Available devices:")
			for _, d := range t.Devices {
				fmt.Printf("\t%s\n", d.Path)
			}
		case *gpsd.TPV:
			Agps.TPV = t
		}
	}
}