package chip8

import "fmt"

// Clears the window
func (c8 *Chip8) I00E0() {
	c8.FrameBuffer = [64 * 32]byte{}
	c8.MustDraw = true
}

// Return from a subroutine
func (c8 *Chip8) I00EE() {
	fmt.Println("I00EE")
}

// Jumps to address NNN
func (c8 *Chip8) I1NNN() {
	addr := c8.COpcode.NNN()
	c8.Pc = addr
}

// Calls subroutine at NNN
func (c8 *Chip8) I2NNN() {
	fmt.Println("I2NNN")
}

// Skips next instruction if VX equals NN
func (c8 *Chip8) I3XNN() {
	fmt.Println("I3XNN")
}

// Skips next instruction if VX does not equal NN
func (c8 *Chip8) I4XNN() {
	fmt.Println("I4XNN")
}

// Skips the next instructions if VX equals VY
func (c8 *Chip8) I5XY0() {
	fmt.Println("I5XY0")
}

// Sets VX = NN
func (c8 *Chip8) I6XNN() {
	_byte := c8.COpcode.NN()
	c8.Registers[c8.COpcode.X()] = _byte
}

// Adds NN to VX
func (c8 *Chip8) I7XNN() {
	_byte := c8.COpcode.NN()
	c8.Registers[c8.COpcode.X()] += _byte
}

// Sets VX to the value of VY
func (c8 *Chip8) I8XY0() {
	fmt.Println("I8XY0")
}

// Sets VX to VX or VY (OR)
func (c8 *Chip8) I8XY1() {
	fmt.Println("")
}

// Sets VX to VX and VY (AND)
func (c8 *Chip8) I8XY2() {
	fmt.Println("I8XY2")
}

// Sets VX to VX xor VY (XOR)
func (c8 *Chip8) I8XY3() {
	fmt.Println("I8XY3")
}

// VY substracted from VX. VF set to 0 if carry, if not 1
func (c8 *Chip8) I8XY4() {
	fmt.Println("I8XY4")
}

// VY substracted from VY. VF set to 0 if borrow, if not 1
func (c8 *Chip8) I8XY5() {
	fmt.Println("I8XY5")
}

// Stores the least signification VX in VF and shifts VX right by 1
func (c8 *Chip8) I8XY6() {
	fmt.Println("I8XY6")
}

// Sets VX to VY minus VX. VF set to 0 if borrow, if not 1
func (c8 *Chip8) I8XY7() {
	fmt.Println("I8XY7")
}

// Stores the most signification VX in VF and shifts VX left by 1
func (c8 *Chip8) I8XYE() {
	fmt.Println("I8XYE")
}

// Skips next instruction if VX does not equal VY
func (c8 *Chip8) I9XY0() {
	fmt.Println("I9XY0")
}

// Sets I to address NNN
func (c8 *Chip8) IANNN() {
	c8.I = c8.COpcode.NNN()
}

// Jumps to address NNN plus V0
func (c8 *Chip8) IBNNN() {
	fmt.Println("IBNNN")
}

// Sets VX to result of AND on a random number (0 to 255) and NN
func (c_ *Chip8) ICXNN() {
	fmt.Println("ICXNN")
}

// Draws Sprite (VX, VY), Width  = 8 px / Height = N px
func (c8 *Chip8) IDXYN() {
	VX := c8.Registers[c8.COpcode.X()]
	VY := c8.Registers[c8.COpcode.Y()]
	hSprite := int(c8.COpcode.N())

	// Memory index (I) = Sprite data location
	i := int(c8.I)

	// Collision flag (VF) = 0 (no collision).
	c8.Registers[0xF] = 0

	for y := 0; y < hSprite; y++ {
		for x := 0; x < 8; x++ {

			_byte := c8.Memory[i+y]
			spritePx := _byte & (0x80 >> x)

			// Check Sprite px = 1 (ON)
			if spritePx != 0 {
				screenX := int(VX) + x
				screenY := int(VY) + y

				cellFrameBuffer := c8.FrameBuffer.Get(screenX, screenY)

				// If pixel(frameBuffer) = ON, collision flag (VF) = 1
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

// Skips next instruction if key stored in VX is pressed
func (c8 *Chip8) IEX9E() {
	fmt.Println("IEX9E")
}

// Skips next instruction if key stored in VX is not pressed
func (c8 *Chip8) IEXA1() {
	fmt.Println("IEXA1")
}

// Sets VX to value of delay timer
func (c8 *Chip8) IFX07() {
	fmt.Println("IFX07")
}

// Kes press is awaited and then stores to VX (all instruction halted until nex key event)
func (c8 *Chip8) IFX0A() {
	fmt.Println("IFX0A")
}

// Sets the delay timer to VX
func (c8 *Chip8) IFX15() {
	fmt.Println("IFX15")
}

// Sets the sound timer to VX
func (c8 *Chip8) IFX18() {
	fmt.Println("IFX18")
}

// Adds VX to I (VF is not affected)
func (c8 *Chip8) IFX1E() {
	fmt.Println("IFX1E")
}

// Sets I to location of sprite for the character in VX (characters 0-F (hexadecimal) are represented by 4x5 font)
func (c8 *Chip8) IFX29() {
	fmt.Println("IFX29")
}

// Stores the binary coded decimal representation of VX (hundreds digit at location in I, tens digit at location I+1, ones digit at location I+2)
func (c8 *Chip8) IFX33() {
	fmt.Println("IFX33")
}

// Stores from V0 to VX (including VX) starting at address I. Offset from I is increased by 1 for each value (I left unmodified)
func (c8 *Chip8) IFX55() {
	fmt.Println("IFX55")
}

// Fills from V0 to VX (including VX) with values fro memory, starting at address I. Offset from I is incread by 1 for each value read, (I left unmodified)
func (c8 *Chip8) IFX65() {
	fmt.Println("IFX65")
}
