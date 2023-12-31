package chip8

import "chip-8/window"

type Chip8 struct {
	Memory       [TotalMemory]byte
	Registers    [NumberOfRegisters]byte
	Pc           uint16
	I            uint16
	Stack        [StackLevels]uint16
	Sp           byte
	COpcode      opcode
	FrameBuffer  window.FrameBuffer
	Instructions map[uint16]func()

	KeyPressed chan byte
	DelayTimer byte
	SoundTimer byte
	MustDraw   bool
	Quit       bool
}

// Initializing Chip-8
func InitChip8(keyPressed chan byte) (*Chip8, error) {
	c8 := &Chip8{
		Memory:       [TotalMemory]byte{},
		Registers:    [NumberOfRegisters]byte{},
		Pc:           PCStartAddress,
		Stack:        [StackLevels]uint16{},
		FrameBuffer:  window.FrameBuffer{},
		Instructions: map[uint16]func(){},
	}

	c8.KeyPressed = keyPressed
	c8.Instructions[0x00E0] = c8.I00E0
	c8.Instructions[0x00EE] = c8.I00EE
	c8.Instructions[0x1000] = c8.I1NNN
	c8.Instructions[0x2000] = c8.I2NNN
	c8.Instructions[0x3000] = c8.I3XNN
	c8.Instructions[0x4000] = c8.I4XNN
	c8.Instructions[0x5000] = c8.I5XY0
	c8.Instructions[0x6000] = c8.I6XNN
	c8.Instructions[0x7000] = c8.I7XNN
	c8.Instructions[0x8000] = c8.I8XY0
	c8.Instructions[0x8001] = c8.I8XY1
	c8.Instructions[0x8002] = c8.I8XY2
	c8.Instructions[0x8003] = c8.I8XY3
	c8.Instructions[0x8004] = c8.I8XY4
	c8.Instructions[0x8005] = c8.I8XY5
	c8.Instructions[0x8006] = c8.I8XY6
	c8.Instructions[0x8007] = c8.I8XY7
	c8.Instructions[0x800E] = c8.I8XYE
	c8.Instructions[0x9000] = c8.I9XY0
	c8.Instructions[0xA000] = c8.IANNN
	c8.Instructions[0xB000] = c8.IBNNN
	c8.Instructions[0xC000] = c8.ICXNN
	c8.Instructions[0xD000] = c8.IDXYN
	c8.Instructions[0xE09E] = c8.IEX9E
	c8.Instructions[0xE0A1] = c8.IEXA1
	c8.Instructions[0xF007] = c8.IFX07
	c8.Instructions[0xF00A] = c8.IFX0A
	c8.Instructions[0xF015] = c8.IFX15
	c8.Instructions[0xF018] = c8.IFX18
	c8.Instructions[0xF01E] = c8.IFX1E
	c8.Instructions[0xF029] = c8.IFX29
	c8.Instructions[0xF033] = c8.IFX33
	c8.Instructions[0xF055] = c8.IFX55
	c8.Instructions[0xF065] = c8.IFX65

	return c8, nil
}

// Get Chip-8 FrameBuffer
func (c8 *Chip8) GetFrameBuffer() window.FrameBuffer {
	return c8.FrameBuffer
}

// Check if Chip-8 is closed
func (c8 *Chip8) IsClosed() bool {
	return c8.Quit
}

// Delay timer that decreases every cycle
func (c8 *Chip8) countBackDelayTimer() {
	if c8.DelayTimer != 0 {
		c8.DelayTimer--
	}
}
