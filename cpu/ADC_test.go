package cpu

import "testing"

func TestAddWithCarry(t *testing.T) {
	tests := []struct {
		name   string
		cpu    *CPU
		memory uint8
		result ResultState
	}{
		{
			name:   "No Carry",
			cpu:    &CPU{A: 0x00, C: false},
			memory: 0x01,
			result: ResultState{A: 0x01, C: false, Z: false, N: false},
		},
		{
			name:   "With Carry",
			cpu:    &CPU{A: 0x01, C: true},
			memory: 0x01,
			result: ResultState{A: 0x03, C: false, Z: false, N: false},
		},
		{
			name:   "Carried",
			cpu:    &CPU{A: 0xFF, C: false},
			memory: 0x01,
			result: ResultState{A: 0x00, C: true, Z: true, N: false},
		},
		{
			name:   "Signed Overflow",
			cpu:    &CPU{A: 0x7F, C: false},
			memory: 0x01,
			result: ResultState{A: 0x80, C: false, Z: false, N: true, V: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddWithCarry(tt.cpu, tt.memory)
			if !CompareResultState(tt.cpu, tt.result) {
				t.Errorf("AddWithCarry() = %#v, want %#v", CPUToResultState(tt.cpu), tt.result)
			}
		})
	}
}
