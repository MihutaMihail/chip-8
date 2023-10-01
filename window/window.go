package window

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	SizePixel    = 16
	ScreenWidth  = SizePixel * 64
	ScreenHeight = SizePixel * 32
)

type FrameBuffer [64 * 32]byte

// Renders the content of the FrameBuffer onto the window
func ToDraw(buffer FrameBuffer, w *pixelgl.Window) {
	w.Clear(color.RGBA{R: 153, G: 102, B: 1, A: 255})
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(255, 1, 0)

	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {

			if *buffer.Get(x, 31-y) != 0 {
				imd.Push(pixel.V(SizePixel*float64(x), SizePixel*float64(y)))
				imd.Push(pixel.V(SizePixel*float64(x)+SizePixel, SizePixel*float64(y)+SizePixel))
				imd.Rectangle(0)
			}
		}
	}
	imd.Draw(w)
}

// Returns a pointer (memory address) at the coordinates (x,y) in the FrameBuffer
func (f *FrameBuffer) Get(x, y int) *byte {
	return &f[y*64+x]
}
