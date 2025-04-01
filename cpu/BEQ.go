package cpu

// BranchIfEqual PC = PC + offset if Z == true
func BranchIfEqual(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Relative(cpu, operand1, operand2)
	if cpu.Z {
		cpu.Wait()
		cpu.PC = address
	}
}
