package main

import (
	"os"
	"os/signal"

	"github.com/tarm/serial"

	"github.com/ryan-berger/aurora/internal/pkg/animator"
	"github.com/ryan-berger/aurora/internal/pkg/protocol"
)

var frames = []protocol.Frame{
	{
		Delay: 1000,
		Pixels: []protocol.RGBA{
			{R: 255, G: 0, B: 0},
			{R: 0, G: 255, B: 0},
			{R: 0, G: 0, B: 255},
		},
	},
	{
		Delay: 1000,
		Pixels: []protocol.RGBA{
			{R: 0, G: 255, B: 0},
			{R: 0, G: 0, B: 255},
			{R: 255, G: 0, B: 0},
		},
	},
	{
		Delay: 1000,
		Pixels: []protocol.RGBA{
			{R: 0, G: 0, B: 255},
			{R: 255, G: 0, B: 0},
			{R: 0, G: 255, B: 0},
		},
	},
}

func main() {
	p, err := serial.OpenPort(&serial.Config{
		Name: "/dev/ttyACM0",
		Baud: 9600,
	})

	if err != nil {
		panic(err)
	}
	defer p.Close()

	closer := make(chan struct{}, 1)

	animator.NewAnimator(p, closer)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	<- c
}
