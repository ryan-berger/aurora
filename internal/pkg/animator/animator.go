package animator

import (
	"bytes"
	"io"
	"sync"
	"time"

	"github.com/ryan-berger/aurora/internal/pkg/protocol"
)

var preamble = []byte{0x0, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}

type Animator struct {
	numLED int

	frames     []protocol.Frame
	brightness int

	out    io.Writer
	mu     *sync.Mutex
	closer chan struct{}
}

func (a *Animator) getMessage(f protocol.Frame) []byte {
	msg := bytes.NewBuffer(preamble)
	msg.Write([]byte{0, 100})
	msg.Write(f.Encode(a.numLED))
	return msg.Bytes()
}

func (a *Animator) runWorker() {
	for {
		for _, f := range a.frames {
			select {
			case <-a.closer:
				return
			default:
				a.out.Write(a.getMessage(f))
				<-time.After(time.Duration(f.Delay) * time.Millisecond)
			}
		}
	}
}

func (a *Animator) SetFrames(frames []protocol.Frame) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.frames = frames
}

func NewAnimator(out io.Writer, closer chan struct{}) *Animator {
	a := &Animator{
		out:    out,
		closer: closer,
		frames: []protocol.Frame{{Pixels: []protocol.RGBA{{R: 255, G: 0, B: 0}}, Delay: 1000}},
		numLED: 168,
	}
	go a.runWorker()
	return a
}
