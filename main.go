package main

import (
	"chip-8/chip8"
)

func main() {
	c8, err := chip8.InitChip8()
	if err != nil {
		panic(err)
	}

	err = c8.LoadROM("assets/1-chip8-logo.ch8")
	//err = c8.LoadROM("assets/2-ibm-logo.ch8")
	if err != nil {
		panic(err)
	}

	// Test if instructions are being executed
	for i := 0; i < 100; i++ {
		c8.Cycle()
	}
}
