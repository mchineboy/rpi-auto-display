package auto_interface

import (
	"fmt"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type FontPack struct {
	File  string
	Sizes []float64
	Short string
}

var fonts = []FontPack{
	{
		File:  "/fonts/digital.ttf",
		Sizes: []float64{10, 12, 14, 16, 36, 48},
		Short: "digital",
	},
	{
		File:  "/fonts/race.ttf",
		Sizes: []float64{10, 11, 12, 14, 16, 36, 48},
		Short: "race",
	},
	{
		File:  "/fonts/roboto.ttf",
		Sizes: []float64{10, 11, 12, 14, 16, 36, 48},
		Short: "roboto",
	},
}

func (AutoInt *AutoInterface) LoadFonts() {
	AutoInt.Fonts = map[string]font.Face{}
	for _, font := range fonts {
		file, err := os.ReadFile(font.File)
		if err != nil {
			panic(err)
		}
		ff, err := truetype.Parse(file)
		if err != nil {
			panic(err)
		}
		for _, size := range font.Sizes {
			AutoInt.Fonts[fmt.Sprintf("%s-%d", font.Short, int(size))] = truetype.NewFace(ff, &truetype.Options{
				Size: size,
			})
		}
	}
	log.Printf("%+v", AutoInt.Fonts)
}
