package animator

import (
	"math"

	"github.com/ryan-berger/aurora/internal/pkg/protocol"
)

type hsv struct {
	h, s, v int
}

func Lerp(numLights int, l, r protocol.RGBA) []protocol.RGBA {

	return nil
}

func rgbToHSV(c protocol.RGBA) hsv {
	r := float64(c.R) / 255 //RGB from 0 to 255
	g := float64(c.G) / 255
	b := float64(c.B) / 255

	min := math.Min(r, math.Min(g, b)) //Min. value of RGB
	max := math.Max(r, math.Max(g, b)) //Max. value of RGB
	del := max - min                   //Delta RGB value

	v := max

	var h, s float64

	if del == 0 { //This is a gray, no chroma...
		h = 0 //HSV results from 0 to 1
		s = 0
	} else { //Chromatic data...
		s = del / max

		delR := (((max - r) / 6) + (del / 2)) / del
		delG := (((max - g) / 6) + (del / 2)) / del
		delB := (((max - b) / 6) + (del / 2)) / del

		if r == max {
			h = delB - delG
		} else if g == max {
			h = (1 / 3) + delR - delB
		} else if b == max {
			h = (2 / 3) + delG - delR
		}

		if h < 0 {
			h += 1
		}
		if h > 1 {
			h -= 1
		}
	}
	return hsv{int(h), int(s), int(v)}
}