package protocol

import (
	"testing"
)

var tests = []struct {
	a        []Frame
	expected []byte
}{
	{a: []Frame{}, expected: []byte{}},
	{
		a: []Frame{
			{Delay: 2,
				Pixels: []RGBA{
					{R: 255, G: 0, B: 0},
				},
			},
		},
		expected: []byte{1, 2, 1, 255, 0, 0},
	},
	{
		a: []Frame{
			{
				Delay: 2,
				Pixels: []RGBA{
					{R: 1, G: 4, B: 4},
					{R: 1, G: 4, B: 4},
				},
			},
			{
				Delay: 2,
				Pixels: []RGBA{
					{R: 1, G: 4, B: 4},
					{R: 1, G: 4, B: 4},
				},
			},
		},
		expected: []byte{2, 2, 2, 1, 4, 4, 1, 4, 4, 2, 2, 1, 4, 4, 1, 4, 4},
	},
}

func TestDecode(t *testing.T) {}
