package cpu

// RotateLeft C = A[7]; A = A << 1; A[0] = C
func RotateLeft(cpu *CPU, value uint8, address uint16, accumulator bool) (result uint8) {
	result = value << 1
	if cpu.C {
		result |= 0x01
	} else {
		result &= 0xFE
	}

	cpu.C = value&0x80 != 0
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

	return result
}

func RotateLeftA(cpu *CPU, operand1 uint8, operand2 uint8) {
	RotateLeft(cpu, cpu.A, 0, true)
}

func RotateLeftZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPage(cpu, operand1, operand2, false)
	RotateLeft(cpu, value(), address, false)
}

func RotateLeftZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := ZeroPageX(cpu, operand1, operand2, false)
	RotateLeft(cpu, value(), address, false)
}

func RotateLeftABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := Absolute(cpu, operand1, operand2, false)
	RotateLeft(cpu, value(), address, false)
}

func RotateLeftABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	value, address := AbsoluteX(cpu, operand1, operand2, false)
	cpu.Wait()
	RotateLeft(cpu, value(), address, false)
}
