package keypad

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
)

type Keypad map[pixelgl.Button]func()

var KeyboardToKeypad = map[pixelgl.Button]byte{
	pixelgl.Key1: 1,
	pixelgl.Key2: 2,
	pixelgl.Key3: 3,
	pixelgl.Key4: 0xC,
	pixelgl.KeyQ: 4,
	pixelgl.KeyW: 5,
	pixelgl.KeyE: 6,
	pixelgl.KeyR: 0xD,
	pixelgl.KeyA: 7,
	pixelgl.KeyS: 8,
	pixelgl.KeyD: 9,
	pixelgl.KeyF: 0xE,
	pixelgl.KeyZ: 0xA,
	pixelgl.KeyX: 0,
	pixelgl.KeyC: 0xB,
	pixelgl.KeyV: 0xF,
}

func KeyHandler(w *pixelgl.Window, keypad Keypad) {
	ticker := time.NewTicker(time.Second / 250)
	defer ticker.Stop()

	for {
		<-ticker.C

		for key, cmd := range keypad {
			if w.JustPressed(key) {
				cmd()
			}
		}
	}
}
