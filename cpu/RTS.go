package cpu

// ReturnFromSubroutine PC = stack[SP] + 1
func ReturnFromSubroutine(cpu *CPU, operand1 uint8, operand2 uint8) {
	high := cpu.Pop()
	cpu.Wait()
	low := cpu.Pop()
	cpu.Wait()

	cpu.PC = uint16(low) | (uint16(high) << 8)
	cpu.Wait()
	//cpu.PC++
	cpu.Wait()
}
