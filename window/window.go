package window

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	SizePixel    = 16
	ScreenWidth  = SizePixel * 64
	ScreenHeight = SizePixel * 32
)

type FrameBuffer [64 * 32]byte

// Renders the content of the FrameBuffer onto the window
func ToDraw(buffer FrameBuffer, w *pixelgl.Window) {
	w.Clear(colornames.Black)
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 1, 1)

	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {

			if *buffer.Get(x, 31-y) != 0 {
				upperLeftX := SizePixel * float64(x)
				upperLeftY := SizePixel * float64(y)
				bottomRightX := upperLeftX + SizePixel
				bottomRightY := upperLeftY + SizePixel

				imd.Push(pixel.V(upperLeftX, upperLeftY))
				imd.Push(pixel.V(bottomRightX, bottomRightY))
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
