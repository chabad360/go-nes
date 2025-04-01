package cpu

// BranchIfCarrySet PC = PC + offset
func BranchIfCarrySet(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Relative(cpu, operand1, operand2)
	if cpu.C {
		cpu.Wait()
		cpu.PC = address
	}
}
