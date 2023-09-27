package chip8

// Clear Screen
func (c8 *Chip8) I00E0() {
	// TO DO
}

// Return from subroutine
func (c8 *Chip8) I00EE() {
	// TO DO
}

// Jumps to address NNN
func (c8 *Chip8) I1NNN() {
	// TO DO
}

// Calls subroutine at NNN
func (c8 *Chip8) I2NNN() {
	// TO DO
}

// Skips next instruction if VX equals NN
func (c8 *Chip8) I3XNN() {
	// TO DO
}

// Skips next instruction if VX does not equal NN
func (c8 *Chip8) I4XNN() {
	// TO DO
}

// Skips the next instructions if VX equals VY
func (c8 *Chip8) I5XY0() {
	// TO DO
}

// Sets VX to NN
func (c8 *Chip8) I6XNN() {
	// TO DO
}

// Adds NN to VX
func (c8 *Chip8) I7XNN() {
	// TO DO
}

// Sets VX to the value of VY
func (c8 *Chip8) I8XY0() {
	// TO DO
}

// Sets VX to VX or VY (OR)
func (c8 *Chip8) I8XY1() {
	// TO DO
}

// Sets VX to VX and VY (AND)
func (c8 *Chip8) I8XY2() {
	// TO DO
}

// Sets VX to VX xor VY (XOR)
func (c8 *Chip8) I8XY3() {
	// TO DO
}

// VY substracted from VX. VF set to 0 if carry, if not 1
func (c8 *Chip8) I8XY4() {
	// TO DO
}

// VY substracted from VY. VF set to 0 if borrow, if not 1
func (c8 *Chip8) I8XY5() {
	// TO DO
}

// Stores the least signification VX in VF and shifts VX right by 1
func (c8 *Chip8) I8XY6() {
	// TO DO
}

// Sets VX to VY minus VX. VF set to 0 if borrow, if not 1
func (c8 *Chip8) I8XY7() {
	// TO DO
}

// Stores the most signification VX in VF and shifts VX left by 1
func (c8 *Chip8) I8XYE() {
	// TO DO
}

// Skips next instruction if VX does not equal VY
func (c8 *Chip8) I9XY0() {
	// TO DO
}

// Sets I to address NNNN
func (c8 *Chip8) IANNN() {
	// TO DO
}

// Jumps to address NNN plus V0
func (c8 *Chip8) IBNNN() {
	// TO DO
}

// Sets VX to result of AND on a random number (0 to 255) and NN
func (c_ *Chip8) ICXNN() {
	// TO DO
}

// Draws sprite at coordinate (VX, VY) width 8 pixels and height N pixels
func (c8 *Chip8) IDXYN() {
	// TO DO
}

// Skips next instruction if key stored in VX is pressed
func (c8 *Chip8) IEX9E() {
	// TO DO
}

// Skips next instruction if key stored in VX is not pressed
func (c8 *Chip8) IEXA1() {
	// TO DO
}

// Sets VX to value of delay timer
func (c8 *Chip8) IFX07() {
	// TO DO
}

// Kes press is awaited and then stores to VX (all instruction halted until nex key event)
func (c8 *Chip8) IFX0A() {
	// TO DO
}

// Sets the delay timer to VX
func (c8 *Chip8) IFX15() {
	// TO DO
}

// Sets the sound timer to VX
func (c8 *Chip8) IFX18() {
	// TO DO
}

// Adds VX to I (VF is not affected)
func (c8 *Chip8) IFX1E() {
	// TO DO
}

// Sets I to location of sprite for the character in VX (characters 0-F (hexadecimal) are represented by 4x5 font)
func (c8 *Chip8) IFX29() {
	// TO DO
}

// Stores the binary coded decimal representation of VX (hundreds digit at location in I, tens digit at location I+1, ones digit at location I+2)
func (c8 *Chip8) IFX33() {
	// TO DO
}

// Stores from V0 to VX (including VX) starting at address I. Offset from I is increased by 1 for each value (I left unmodified)
func (c8 *Chip8) IFX55() {
	// TO DO
}

// Fills from V0 to VX (including VX) with values fro memory, starting at address I. Offset from I is incread by 1 for each value read, (I left unmodified)
func (c8 *Chip8) IFX65() {
	// TO DO
}
