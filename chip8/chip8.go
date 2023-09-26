package chip8

type Chip8 struct {
	Memory       [TotalMemory]byte
	Registers    [NumberOfRegisters]byte
	Pc           uint16
	I            uint16
	Stack        [StackLevels]uint16
	COpcode      opcode
	Instructions map[uint16]func()
}

func InitChip8() (*Chip8, error) {
	c8 := &Chip8{
		Memory:       [TotalMemory]byte{},
		Registers:    [NumberOfRegisters]byte{},
		Pc:           PCStartAddress,
		Stack:        [StackLevels]uint16{},
		Instructions: map[uint16]func(){},
	}

	c8.Instructions[0x00E0] = c8.I00E0
	c8.Instructions[0x00EE] = c8.I00EE
	// ...

	// Iterate through map
	/*for key := range c8.Instructions {
        fmt.Printf("Key: %d\n", key)
    }*/

	return c8, nil
}
