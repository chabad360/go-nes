package cpu

import "testing"

type ResultState struct {
	A, X, Y, SP         uint8
	C, Z, I, D, B, V, N bool

	PC uint16
}

func CPUToResultState(cpu *CPU) ResultState {
	return ResultState{
		A: cpu.A, X: cpu.X, Y: cpu.Y, SP: cpu.SP,
		C: cpu.C, Z: cpu.Z, I: cpu.I, D: cpu.D,
		B: cpu.B, V: cpu.V, N: cpu.N,
		PC: cpu.PC,
	}
}

func CompareResultState(cpu *CPU, r ResultState) bool {
	if cpu.A != r.A || cpu.X != r.X || cpu.Y != r.Y || cpu.SP != r.SP ||
		cpu.C != r.C || cpu.Z != r.Z || cpu.I != r.I || cpu.D != r.D ||
		cpu.B != r.B || cpu.V != r.V || cpu.N != r.N {
		return false
	}
	if cpu.PC != r.PC {
		return false
	}
	//for i := 0; i < len(cpu.Memory); i++ {
	//	if cpu.Memory[i] != r.Memory[i] {
	//		return false
	//	}
	//}
	return true
}

func TestAddWithCarryIMM(t *testing.T) {
	tests := []struct {
		name     string
		cpu      *CPU
		operand1 uint8
		result   ResultState
	}{
		{
			name:     "No Carry",
			cpu:      &CPU{A: 0x00, C: false},
			operand1: 0x01,
			result:   ResultState{A: 0x01, C: false, Z: false, N: false},
		},
		{
			name:     "With Carry",
			cpu:      &CPU{A: 0x01, C: true},
			operand1: 0x01,
			result:   ResultState{A: 0x03, C: false, Z: false, N: false},
		},
		{
			name:     "Carried",
			cpu:      &CPU{A: 0xFF, C: false},
			operand1: 0x01,
			result:   ResultState{A: 0x00, C: true, Z: true, N: false},
		},
		{
			name:     "Signed Overflow",
			cpu:      &CPU{A: 0x7F, C: false},
			operand1: 0x01,
			result:   ResultState{A: 0x80, C: false, Z: false, N: true, V: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddWithCarryIMM(tt.cpu, tt.operand1, 0)
			if !CompareResultState(tt.cpu, tt.result) {
				t.Errorf("AddWithCarryIMM() = %#v, want %#v", CPUToResultState(tt.cpu), tt.result)
			}
		})
	}
}
