package cpu

// DecrementY Y--
func DecrementY(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Y--

	cpu.Z = cpu.Y == 0
	cpu.N = cpu.Y&0x80 != 0
}
