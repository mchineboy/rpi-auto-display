package auto_gps

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/amenzhinsky/go-gpsd"
	_ "github.com/briansorahan/spatialite"
)

type AutoGps struct {
	TPV     *gpsd.TPV
	Tz      string
	Spatial *sql.DB
}

func New() *AutoGps {
	agps := &AutoGps{}
	spatial, err := sql.Open("spatialite", "/tmp/locations.sqlite3")
	if err != nil {
		log.Printf("SQLite error? %+v", err)
	}
	agps.Spatial = spatial
	agps.BuildDatabase()
	return agps
}

func (Agps *AutoGps) Monitor() {
	g, err := gpsd.Dial("127.0.0.1:2947")
	if err != nil {
		log.Printf("%+v", err)
	}
	defer g.Close()

	if err := g.Stream(gpsd.WATCH_ENABLE|gpsd.WATCH_JSON, ""); err != nil {
		log.Printf("%+v", err)
	}
	defer g.Stream(gpsd.WATCH_DISABLE, "")

	for v := range g.C() {
		switch t := v.(type) {
		case *gpsd.VERSION:
			fmt.Printf("GPSD Version: %s, Proto: %.0f.%.0f\n", t.Release, t.ProtoMajor, t.ProtoMinor)
		case *gpsd.DEVICES:
			if len(t.Devices) == 0 {
				log.Printf("%+v", errors.New("no devices available"))
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
