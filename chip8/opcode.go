package chip8

type opcode uint16

// Get opcode instruction
func (oc opcode) GetOpcodeInstruction() uint16 {
	switch (uint16(oc) & uint16(0xF000)) >> 12 {

	case uint16(0x0):
		return uint16(oc)

	case uint16(0x8):
		return uint16(oc) & uint16(0xF00F)

	case uint16(0xE):
		return uint16(oc) & uint16(0xF0FF)

	case uint16(0xF):
		return uint16(oc) & uint16(0xF0FF)

	default:
		return uint16(oc) & uint16(0xF000)
	}
}

// Executes all steps necessary in a row
func (c8 *Chip8) CycleEmulator() {
	c8.fetchOpcode()
	c8.executeOpcode()
}

// Fetch opcode from memory
func (c8 *Chip8) fetchOpcode() {
	c8.COpcode = opcode(uint16(c8.Memory[c8.Pc])<<8 | uint16(c8.Memory[c8.Pc+1]))
	c8.Pc += 2
}

// Get current opcode and execute its instruction
func (c8 *Chip8) executeOpcode() {
	keyInst := c8.COpcode.GetOpcodeInstruction()
	if inst, found := c8.Instructions[keyInst]; found {
		inst()
	}
}

// Get NNN parameter
func (oc opcode) NNN() uint16 {
	return uint16(oc) & uint16(0x0FFF)
}

// Get NN parameter
func (oc opcode) NN() uint8 {
	return uint8(oc) & uint8(0x00FF)
}

// Get N parameter
func (oc opcode) N() uint8 {
	return uint8(oc) & 0x000F
}

// Get X parameter
func (oc opcode) X() uint8 {
	return uint8(oc>>8) & 0x000F
}

// Get Y paramater
func (oc opcode) Y() uint8 {
	return uint8(oc>>4) & 0x000F
}
