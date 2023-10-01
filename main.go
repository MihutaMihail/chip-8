package main

import (
	"chip-8/game"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	myGame, err := game.NewGame()
	if err != nil {
		panic(err)
	}
	myGame.RunGame()
}
