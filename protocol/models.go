package protocol

import (
	"bytes"
)


type RGBA struct {
	R, G, B uint8
}

type Frame struct {
	Pixels []RGBA
	Delay  int16
}

type Animation []Frame


func (r RGBA) encode() []byte {
	return []byte{r.R, r.G, r.B}
}

// frame decode

func (f Frame) size() int {
	return 1 + (len(f.Pixels) * 3)
}


func (f Frame) encode() []byte {
	buf := bytes.NewBuffer([]byte{byte(f.Delay)})

	for _, p := range f.Pixels{
		buf.Write(p.encode())
	}

	return buf.Bytes()
}

// Animation decode

func (a Animation) size() int {
	if len(a) == 0 {
		return 0
	}
	return len(a) * a[0].size()
}

func (a Animation) Encode() []byte {
	buf := bytes.NewBuffer([]byte{byte(len(a))})
	buf.Grow(a.size())

	for _, f := range a {
		buf.Write(f.encode())
	}

	return buf.Bytes()
}