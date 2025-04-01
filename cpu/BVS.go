package cpu

// BranchIfOverflowSet PC = PC + offset if V == 1
func BranchIfOverflowSet(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Relative(cpu, operand1, operand2)
	if cpu.V {
		cpu.Wait()
		cpu.PC = address
	}
}
