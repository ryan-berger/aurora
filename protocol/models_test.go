package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	a        Animation
	expected []byte
}{
	{a: Animation{}, expected: []byte{0}},
	{
		a: Animation{
			{
				Delay: 2,
				Pixels: []RGBA{
					{R: 1, G: 4, B: 4},
				},
			},
		},
		expected: []byte{1, 2, 1, 4, 4},
	},
}

func TestDecode(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, test.a.Encode())
	}
}
