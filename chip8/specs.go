package chip8

import "time"

const (
	TotalMemory       = 4096
	MaxCapacityForROM = 0xFFF - 0x200
	NumberOfRegisters = 16
	PCStartAddress    = 0x200
	StackLevels       = 16
	Frequency         = time.Second / time.Duration(60)
	WidthScreen       = 64
	HeightScreen      = 32
)
