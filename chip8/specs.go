package chip8

const (
	TotalMemory       = 4096
	MemoryForROM      = 0xFFF - 0x200
	NumberOfRegisters = 16
	PCStartAddress    = 0x200
	StackLevels       = 16
)