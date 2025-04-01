package cpu

// TransferXToStack SP = X
func TransferXToStack(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.SP = cpu.X
}
