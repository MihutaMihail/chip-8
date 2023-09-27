package main

import (
	"chip-8/chip8"
	"fmt"
)

func main() {
	c8, err := chip8.InitChip8()
	if err != nil {
		panic(err)
	}

	err = c8.LoadROM("assets/1-chip8-logo.ch8")
	if err != nil {
		panic(err)
	}
	fmt.Println(c8.Memory)
}
