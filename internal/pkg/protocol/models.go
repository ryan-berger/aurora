package protocol

import (
	"bytes"
)

type RGBA struct {
	R, G, B uint8
}

func (r RGBA) encode() []byte {
	return []byte{r.B, r.G, r.R}
}

type Frame struct {
	Pixels []RGBA
	Delay  int
}

func (f Frame) size() int {
	return 1 + (len(f.Pixels) * 3)
}

func (f Frame) Encode(numLEDs int) []byte {
	buf := bytes.NewBuffer([]byte{})

	for i := 0; i < numLEDs; i++ {
		var p RGBA
		if i < len(f.Pixels) {
			p = f.Pixels[i]
		}
		buf.Write(p.encode())
	}

	return buf.Bytes()
}
