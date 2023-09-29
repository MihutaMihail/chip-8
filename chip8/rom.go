package chip8

import (
	"errors"
	"os"
)

func (c8 *Chip8) LoadROM(filename string) error {
	return c8.loadFile(filename, MaxCapacityForROM, PCStartAddress, &c8.Memory)
}

func (c8 *Chip8) loadFile(filename string, maxCapacity int, startAddress int, memory *[TotalMemory]byte) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if len(file) > maxCapacity {
		errSize := "ROM : '" + filename + "' has exceeded the size limitations"
		return errors.New(errSize)
	}

	copy(memory[startAddress:], file[:])

	return nil
}
