package cpu

// StoreA memory = A
func StoreA(cpu *CPU, memory uint16) {
	cpu.Mapper.Write(cpu, memory, cpu.A)
	cpu.Wait()
}

func StoreAZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := ZeroPage(cpu, operand1, operand2)
	StoreA(cpu, address)
}

func StoreAZPX(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := ZeroPageX(cpu, operand1, operand2)
	StoreA(cpu, address)
}

func StoreAABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Absolute(cpu, operand1, operand2)
	StoreA(cpu, address)
}

func StoreAABSX(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := AbsoluteX(cpu, operand1, operand2)
	StoreA(cpu, address)
}

func StoreAABSY(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := AbsoluteY(cpu, operand1, operand2)
	StoreA(cpu, address)
}

func StoreAINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := IndexedIndirect(cpu, operand1, operand2)
	StoreA(cpu, address)
}

func StoreAINDY(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := IndirectIndexed(cpu, operand1, operand2)
	StoreA(cpu, address)
}
