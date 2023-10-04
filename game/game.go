package game

import (
	"chip-8/chip8"
	"chip-8/keypad"
	"chip-8/window"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	C8      *chip8.Chip8
	Window  *pixelgl.Window
	Keypad  keypad.Keypad
	Channel chan byte
}

// Initializes a new Chip-8 game instance
func NewGame() (*Game, error) {
	myGame := new(Game)
	myGame.Channel = make(chan byte)
	var err error

	// Init Chip-8
	myGame.C8, err = chip8.InitChip8(myGame.Channel)
	if err != nil {
		return nil, err
	}

	// Cfg file
	cfgPixel := pixelgl.WindowConfig{
		Title:  "Chip-8 Emulator",
		Bounds: pixel.R(0, 0, window.ScreenWidth, window.ScreenHeight),
	}

	// Pixel Window
	myGame.Window, err = pixelgl.NewWindow(cfgPixel)
	if err != nil {
		return nil, err
	}

	// Keypad
	cmdKeypad := make(keypad.Keypad)
	for key, value := range keypad.KeyboardToKeypad {
		newKey := value
		cmdKeypad[key] = func() {
			myGame.Channel <- newKey
		}
	}
	myGame.Keypad = cmdKeypad

	return myGame, nil
}

// Loads the ROM and runs game loop
func (myGame *Game) RunGame() {
	if err := myGame.C8.LoadROM("assets/6-keypad.ch8"); err != nil {
		panic(err)
	}

	go keypad.KeyHandler(myGame.Window, myGame.Keypad)
	go myGame.gameLoop()
	myGame.updateWindow()
}

// Runs the game loop by periodically executing the Chip-8 emulator cycle
func (myGame *Game) gameLoop() {
	ticker := time.NewTicker(chip8.Frequency)
	defer ticker.Stop()

	for !myGame.C8.IsClosed() {
		<-ticker.C
		myGame.C8.CycleEmulator()
	}
}

// Updating the game's window content periodically based on the FrameBuffer
func (myGame *Game) updateWindow() {
	ticker := time.NewTicker(chip8.Frequency)
	defer ticker.Stop()

	for !myGame.C8.IsClosed() {
		<-ticker.C

		if myGame.C8.MustDraw {
			myGame.C8.MustDraw = false
			window.ToDraw(myGame.C8.GetFrameBuffer(), myGame.Window)
		}

		myGame.Window.Update()
	}
}
