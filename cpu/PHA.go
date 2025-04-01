package cpu

// PushA stack[SP] = A
func PushA(cpu *CPU, operand1 uint8, operand2 uint8) {
	cpu.Push(cpu.A)
}
