package cpu

// StoreAAndX memory = A & X
func StoreAAndX(cpu *CPU, address uint16, isIndexed bool) {
	//if isIndexed {
	//	value := cpu.A & cpu.X & (uint8(address>>8) + 1)
	//	cpu.Mapper.Write(cpu, address, value)
	//} else {
	cpu.Mapper.Write(cpu, address, cpu.A&cpu.X)
	//}
	cpu.Wait()
}

func StoreAAndXZP(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := ZeroPage(cpu, operand1, operand2)
	StoreAAndX(cpu, address, false)
}

func StoreAAndXZPY(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := ZeroPageY(cpu, operand1, operand2)
	StoreAAndX(cpu, address, false)
}

func StoreAAndXABS(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := Absolute(cpu, operand1, operand2)
	StoreAAndX(cpu, address, false)
}

func StoreAAndXINDX(cpu *CPU, operand1 uint8, operand2 uint8) {
	_, address := IndexedIndirect(cpu, operand1, operand2)
	StoreAAndX(cpu, address, true)
}
