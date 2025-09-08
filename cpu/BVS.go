package cpu

// BranchIfOverflowSet PC = PC + offset if V == 1
func BranchIfOverflowSet(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Relative(cpu, operand1, operand2, true)
	if cpu.V {
		cpu.Wait()
		if address&0xFF00 != cpu.PC&0xFF00 {
			cpu.Wait()
		}
		cpu.PC = address
	}
}
