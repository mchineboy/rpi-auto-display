package auto_interface

import (
	"fmt"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

type FontPack struct {
	File  string
	Sizes []float64
	Short string
}

var fonts = []FontPack{
	{
		File:  "/fonts/digital.ttf",
		Sizes: []float64{10, 12, 14, 16, 18, 36, 48, 72},
		Short: "digital",
	},
	{
		File:  "/fonts/race.ttf",
		Sizes: []float64{10, 11, 12, 14, 16, 18, 36, 48, 72},
		Short: "race",
	},
	{
		File:  "/fonts/roboto.ttf",
		Sizes: []float64{10, 11, 12, 14, 16, 36, 48},
		Short: "roboto",
	},
	{
		File:  "",
		Sizes: []float64{10, 11, 12, 13, 14, 16, 36, 48},
		Short: "default",
	},
}

func (AutoInt *AutoInterface) LoadFonts() {
	AutoInt.Fonts = map[string]font.Face{}
	for _, font := range fonts {
		var file []byte
		var ff *truetype.Font
		var err error
		if font.File != "" {
			file, err = os.ReadFile(font.File)
			if err != nil {
				panic(err)
			}
			ff, err = truetype.Parse(file)
			if err != nil {
				panic(err)
			}
		}
		for _, size := range font.Sizes {
			if ff != nil {
				AutoInt.Fonts[fmt.Sprintf("%s-%d", font.Short, int(size))] = truetype.NewFace(ff, &truetype.Options{
					Size: size,
				})
				continue
			}
			fontWidth := (size * 6) / 13
			AutoInt.Fonts[fmt.Sprintf("%s-%d", font.Short, int(size))] = &basicfont.Face{
				Advance: int(fontWidth + 1),
				Width:   int(fontWidth),
				Height:  int(size),
				Ascent:  int(size - 2),
				Descent: 2,
				Mask:    basicfont.Face7x13.Mask,
				Ranges:  basicfont.Face7x13.Ranges,
			}

		}
	}
	log.Printf("%+v", AutoInt.Fonts)
}
