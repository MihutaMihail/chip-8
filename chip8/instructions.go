package chip8

import "fmt"

// Clear Screen
func (c8 *Chip8) I00E0() {
	fmt.Println("I00E0")
}

// Return from subroutine
func (c8 *Chip8) I00EE() {
	fmt.Println("I00EE")
}

// Jumps to address NNN
func (c8 *Chip8) I1NNN() {
	fmt.Println("I1NNN")
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

// Sets VX to NN
func (c8 *Chip8) I6XNN() {
	fmt.Println("I6XNN")
}

// Adds NN to VX
func (c8 *Chip8) I7XNN() {
	fmt.Println("I7XNN")
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

// Sets I to address NNNN
func (c8 *Chip8) IANNN() {
	fmt.Println("IANNN")
}

// Jumps to address NNN plus V0
func (c8 *Chip8) IBNNN() {
	fmt.Println("IBNNN")
}

// Sets VX to result of AND on a random number (0 to 255) and NN
func (c_ *Chip8) ICXNN() {
	fmt.Println("ICXNN")
}

// Draws sprite at coordinate (VX, VY) width 8 pixels and height N pixels
func (c8 *Chip8) IDXYN() {
	fmt.Println("IDXYN")
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
