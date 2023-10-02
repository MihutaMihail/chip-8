package game

import (
	"chip-8/chip8"
	"chip-8/window"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Game struct {
	c8     *chip8.Chip8
	window *pixelgl.Window
}

// Initializes a new Chip-8 game instance
func NewGame() (*Game, error) {
	myGame := new(Game)
	var err error

	// Init Chip-8
	myGame.c8, err = chip8.InitChip8()
	if err != nil {
		return nil, err
	}

	// Cfg file
	cfgPixel := pixelgl.WindowConfig{
		Title:  "Chip-8 Emulator",
		Bounds: pixel.R(0, 0, window.ScreenWidth, window.ScreenHeight),
	}

	// Pixel Window
	myGame.window, err = pixelgl.NewWindow(cfgPixel)
	if err != nil {
		return nil, err
	}

	return myGame, nil
}

// Loads the ROM and runs game loop
func (myGame *Game) RunGame() {
	if err := myGame.c8.LoadROM("assets/1-chip8-logo.ch8"); err != nil {
		panic(err)
	}

	go myGame.gameLoop()
	myGame.updateWindow()
}

// Runs the game loop by periodically executing the Chip-8 emulator cycle
func (myApp *Game) gameLoop() {
	clock := time.NewTicker(chip8.Frequency)

	for !myApp.c8.IsClosed() {
		<-clock.C
		myApp.c8.CycleEmulator()
	}
}

// Updating the game's window content periodically based on the FrameBuffer
func (myGame *Game) updateWindow() {
	clock := time.NewTicker(chip8.Frequency)

	for !myGame.c8.IsClosed() {
		<-clock.C

		if myGame.c8.MustDraw {
			myGame.c8.MustDraw = false
			window.ToDraw(myGame.c8.GetFrameBuffer(), myGame.window)
		}

		myGame.window.Update()
	}
}
