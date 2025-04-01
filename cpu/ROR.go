package cpu

// RotateRight C = A[0]; A = A >> 1; A[7] = C
func RotateRight(cpu *CPU, value uint8, address uint16, accumulator bool) {
	result := value >> 1
	if cpu.C {
		result |= 0x80
	} else {
		result &= 0x7F
	}

	cpu.C = value&0x01 != 0
	cpu.Z = result == 0
	cpu.N = result&0x80 != 0

	if accumulator {
		cpu.A = result
	} else {
		cpu.Mapper.Write(cpu, address, value)
		cpu.Wait()
		cpu.Mapper.Write(cpu, address, result)
		cpu.Wait()
	}
}

func RotateRightA(cpu *CPU, operand1 uint8, operand2 uint8) {
	RotateRight(cpu, cpu.A, 0, true)
}

func RotateRightZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2)
	RotateRight(cpu, value(), address, false)
}

func RotateRightZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2)
	RotateRight(cpu, value(), address, false)
}

func RotateRightABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2)
	RotateRight(cpu, value(), address, false)
}

func RotateRightABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2)
	RotateRight(cpu, value(), address, false)
}
