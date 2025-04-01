package cpu

// TransferXToStack SP = X
func TransferXToStack(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.SP = cpu.X

	cpu.Z = cpu.SP == 0
	cpu.N = cpu.SP&0x80 != 0
}
