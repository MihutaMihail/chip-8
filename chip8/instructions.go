package chip8

import (
	"math/rand"
)

// Clears the window
func (c8 *Chip8) I00E0() {
	c8.FrameBuffer = [64 * 32]byte{}
	c8.MustDraw = true
}

// Return from a subroutine
func (c8 *Chip8) I00EE() {
	c8.Sp--
	c8.Pc = c8.Stack[c8.Sp]
}

// Jumps to address NNN
func (c8 *Chip8) I1NNN() {
	addr := c8.COpcode.NNN()
	c8.Pc = addr
}

// Calls subroutine at NNN
func (c8 *Chip8) I2NNN() {
	addr := c8.COpcode.NNN()
	c8.Stack[c8.Sp] = c8.Pc
	c8.Sp++
	c8.Pc = addr
}

// Skip next instruction if vX == NN
func (c8 *Chip8) I3XNN() {
	vX := c8.Registers[c8.COpcode.X()]
	_byte := c8.COpcode.NN()

	if vX == _byte {
		c8.Pc += 2
	}
}

// Skip next instruction if vX != NN
func (c8 *Chip8) I4XNN() {
	vX := c8.Registers[c8.COpcode.X()]
	_byte := c8.COpcode.NN()

	if vX != _byte {
		c8.Pc += 2
	}
}

// Skip the next instructions if vX == vY
func (c8 *Chip8) I5XY0() {
	vX := c8.Registers[c8.COpcode.X()]
	vY := c8.Registers[c8.COpcode.Y()]

	if vX == vY {
		c8.Pc += 2
	}
}

// Set vX = NN
func (c8 *Chip8) I6XNN() {
	_byte := c8.COpcode.NN()
	c8.Registers[c8.COpcode.X()] = _byte
}

// Add NN to vX
func (c8 *Chip8) I7XNN() {
	_byte := c8.COpcode.NN()
	c8.Registers[c8.COpcode.X()] += _byte
}

// Set vX = vY
func (c8 *Chip8) I8XY0() {
	c8.Registers[c8.COpcode.X()] = c8.Registers[c8.COpcode.Y()]
}

// Set vX to vX or vY (OR)
func (c8 *Chip8) I8XY1() {
	c8.Registers[c8.COpcode.X()] |= c8.Registers[c8.COpcode.Y()]
}

// Set vX to vX and vY (AND)
func (c8 *Chip8) I8XY2() {
	c8.Registers[c8.COpcode.X()] &= c8.Registers[c8.COpcode.Y()]
}

// Set vX to vX xor vY (XOR)
func (c8 *Chip8) I8XY3() {
	c8.Registers[c8.COpcode.X()] ^= c8.Registers[c8.COpcode.Y()]
}

// Add vY to Vx. Carry flag ---> vF = 1
func (c8 *Chip8) I8XY4() {
	vX := c8.Registers[c8.COpcode.X()]
	vY := c8.Registers[c8.COpcode.Y()]

	result := vX + vY

	c8.Registers[c8.COpcode.X()] = uint8(result)

	if result > 255 {
		c8.Registers[0xF] = 1 // Carry
	} else {
		c8.Registers[0xF] = 0 // !Carry
	}
}

// vY substracted from vX. Borrow ---> vF = 0
func (c8 *Chip8) I8XY5() {
	vX := c8.Registers[c8.COpcode.X()]
	vY := c8.Registers[c8.COpcode.Y()]

	if vX > vY {
		c8.Registers[0xF] = 1 // !Borrow
	} else {
		c8.Registers[0xF] = 0 // Borrow
	}

	c8.Registers[c8.COpcode.X()] = (vX - vY) & 0xFF
}

// Stores the least significant bit of VX in VF and then shifts VX to the right by 1
func (c8 *Chip8) I8XY6() {
	vXLowestBit := c8.Registers[c8.COpcode.X()] & 0x1

	c8.Registers[0xF] = vXLowestBit
	c8.Registers[c8.COpcode.X()] >>= 1
}

// Set vX to vY minus vX. vF set to 0 if borrow, if not 1
func (c8 *Chip8) I8XY7() {
	vX := c8.Registers[c8.COpcode.X()]
	vY := c8.Registers[c8.COpcode.Y()]

	if vY > vX {
		c8.Registers[0xF] = 1 // !Borrow
	} else {
		c8.Registers[0xF] = 0 // Borrow
	}

	result := vY - vX

	c8.Registers[c8.COpcode.X()] = uint8(result)
}

// Stores the most significant bit of VX in VF and then shifts VX to the left by 1
func (c8 *Chip8) I8XYE() {
	vXHighestBit := c8.Registers[c8.COpcode.X()] & 0x80

	c8.Registers[0xF] = vXHighestBit
	c8.Registers[c8.COpcode.X()] <<= 1
}

// Skip next instruction if vX != vY
func (c8 *Chip8) I9XY0() {
	vX := c8.Registers[c8.COpcode.X()]
	vY := c8.Registers[c8.COpcode.Y()]

	if vX != vY {
		c8.Pc += 2
	}
}

// Set I to address NNN
func (c8 *Chip8) IANNN() {
	c8.I = c8.COpcode.NNN()
}

// Jumps to address NNN plus
func (c8 *Chip8) IBNNN() {
	c8.Pc = c8.COpcode.NNN() + uint16(c8.Registers[0])
}

// Set vX = random number (0 to 255) + NN
func (c8 *Chip8) ICXNN() {
	c8.Registers[c8.COpcode.X()] = uint8(rand.Intn(256)) & c8.COpcode.NN()
}

// Draws Sprite (vX, vY), Width  = 8 px / Height = N px
func (c8 *Chip8) IDXYN() {
	vX := c8.Registers[c8.COpcode.X()]
	vY := c8.Registers[c8.COpcode.Y()]
	hSprite := int(c8.COpcode.N())

	// Memory index (I) = Sprite data location
	i := int(c8.I)

	// Collision flag (vF) = 0 (no collision).
	c8.Registers[0xF] = 0

	for y := 0; y < hSprite; y++ {
		for x := 0; x < 8; x++ {

			_byte := c8.Memory[i+y]
			spritePx := _byte & (0x80 >> x)

			// Check Sprite px = 1 (ON)
			if spritePx != 0 {
				screenX := int(vX) + x
				screenY := int(vY) + y

				cellFrameBuffer := c8.FrameBuffer.Get(screenX, screenY)

				// If pixel(frameBuffer) = ON, collision flag (vF) = 1
				if *cellFrameBuffer == 1 {
					c8.Registers[0xF] = 1
				}

				// Use a XOR operation to toggle the pixel in the FrameBuffer
				*cellFrameBuffer ^= 1
			}
		}
	}

	// Tell Chip-8 to draw
	c8.MustDraw = true
}

// Skip next instruction if key stored in vX is pressed
func (c8 *Chip8) IEX9E() {
	key := c8.Registers[c8.COpcode.X()]
	select {
	case keyPressed := <-c8.KeyPressed:
		if keyPressed == key {
			c8.Pc += 2
		}
		return
	default:
		return
	}
}

// Skip next instruction if key stored in vX is not pressed
func (c8 *Chip8) IEXA1() {
	key := c8.Registers[c8.COpcode.X()]

	select {
	case keyPressed := <-c8.KeyPressed:
		if keyPressed == key {
			return
		}
		c8.Pc += 2
	default:
		c8.Pc += 2

	}
}

// Set vX = delay timer
func (c8 *Chip8) IFX07() {
	c8.Registers[c8.COpcode.X()] = c8.DelayTimer
}

// Kes press is awaited and then stores to vX (all instruction halted until next key event)
func (c8 *Chip8) IFX0A() {
	for {
		select {
		case key := <-c8.KeyPressed:
			if key <= NumberOfKeys {
				c8.Registers[c8.COpcode.X()] = key
				return
			}
		}
	}
}

// Set delay timer = vX
func (c8 *Chip8) IFX15() {
	c8.DelayTimer = c8.Registers[c8.COpcode.X()]
}

// Set the sound timer to vX
func (c8 *Chip8) IFX18() {
	c8.SoundTimer = c8.Registers[c8.COpcode.X()]
}

// Add vX to I (vF is not affected)
func (c8 *Chip8) IFX1E() {
	vX := c8.Registers[c8.COpcode.X()]

	c8.I += uint16(vX)
}

// Set I = location of sprite for character vX
// FontSize is 5 bytes
func (c8 *Chip8) IFX29() {
	vX := c8.Registers[c8.COpcode.X()]
	c8.I = FontStartAddress + uint16(FontSize*vX)
}

// Stores the binary coded decimal representation of vX
// (hundreds digit at location in I, tens digit at location I+1, ones digit at location I+2)
func (c8 *Chip8) IFX33() {
	vX := c8.Registers[c8.COpcode.X()]

	c8.Memory[c8.I] = (vX / 100) % 10
	c8.Memory[c8.I+1] = (vX / 10) % 10
	c8.Memory[c8.I+2] = vX % 10
}

// Stores from v0 to vX (including vX) starting at address I
func (c8 *Chip8) IFX55() {
	for a := 0; a <= int(c8.COpcode.X()); a++ {
		c8.Memory[c8.I+uint16(a)] = c8.Registers[a]
	}
}

// Reads from v0 to vX (including vX) with values from memory, starting at address I.
func (c8 *Chip8) IFX65() {
	for a := 0; a <= int(c8.COpcode.X()); a++ {
		c8.Registers[a] = c8.Memory[c8.I+uint16(a)]
	}
}
