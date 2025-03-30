package cpu

import (
	"bufio"
	"fmt"
	"go-nes/cart"
	"os"
	"regexp"
	"testing"
)

type TestLog struct {
	ResultState
	Exec

	Operation string
	Args      string
	Size      int
}

var LogRegexp = regexp.MustCompile(`([0-9A-F]{4}\s+[0-9A-F]{2}\s+[0-9A-F]{0,2}\s+[0-9A-F]{0,2})\s+([A-Z*]{3,4})\s+(.*?)\s+(A:([0-9A-F]{2})\s+X:([0-9A-F]{2})\s+Y:([0-9A-F]{2})\s+P:([0-9A-F]{2})\s+SP:([0-9A-F]{2})\s+PPU:\s*(\d+),\s*(\d+)\s+CYC:(\d+))`)

func ParseTestLog(logFile string) ([]TestLog, error) {
	file, err := os.Open(logFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var states []TestLog

	for scanner.Scan() {
		line := scanner.Text()
		var state TestLog

		// Parse the line to extract the PC, instruction, and state
		matches := LogRegexp.FindStringSubmatch(line)
		if matches == nil {
			return nil, fmt.Errorf("failed to parse line: %s", line)
		}

		// Extract the PC, opcode, operands, and operation
		n, err := fmt.Sscanf(matches[1], "%X  %X %X %X",
			&state.Exec.PC, &state.Exec.Opcode, &state.Exec.Operand1, &state.Exec.Operand2)
		if err != nil && n < 2 {
			return nil, fmt.Errorf("failed to parse instruction: %s", line)
		}

		state.Operation = matches[2]
		state.Args = matches[3]
		state.Size = n - 1

		flags := uint8(0)
		n, err = fmt.Sscanf(matches[4], "A:%X X:%X Y:%X P:%X SP:%X",
			&state.ResultState.A, &state.ResultState.X, &state.ResultState.Y,
			&flags, &state.ResultState.SP)
		if err != nil && n < 4 {
			return nil, fmt.Errorf("failed to parse result state: %s", line)
		}

		// Set the flags based on the parsed value
		state.ResultState.C = (flags & 0x01) != 0
		state.ResultState.Z = (flags & 0x02) != 0
		state.ResultState.I = (flags & 0x04) != 0
		state.ResultState.D = (flags & 0x08) != 0
		state.ResultState.B = (flags & 0x10) != 0
		state.ResultState.V = (flags & 0x40) != 0
		state.ResultState.N = (flags & 0x80) != 0

		states = append(states, state)
	}

	return states, nil
}

func TestCPU_NESTest(t *testing.T) {
	testlog, err := ParseTestLog("testdata/NESTest.log")
	if err != nil {
		t.Fatalf("Failed to parse test log: %v", err)
	}

	// Load the NESTest cart
	cart, err := cart.NewCartidge("testdata/NESTest.nes")
	if err != nil {
		t.Fatalf("Failed to load NESTest cart: %v", err)
	}

	clock := make(chan int)

	cpu := &CPU{
		Mapper: &NROM{
			PRGROM: cart.PRGROM,
			CHRROM: cart.CHRROM,
		},
		Clock: clock,
		Log:   true,
	}

	cpu.PowerOn()

	go func() {
		for i := 0; ; i++ {
			clock <- i
		}
	}()

	cpu.PC = 0xC000              // Start at the beginning of nestest
	for i := 0; i < 10000; i++ { // Run for a reasonable number of cycles
		exec := cpu.Decode()

		l := testlog[i]
		if _, exists := Instructions[exec.Opcode]; !exists {
			t.Errorf("Encountered unknown opcode at PC=%04X: %X\nExpected: %X (%s), Args: %X %X (%s)",
				exec.PC, exec.Opcode, l.Opcode, l.Operation, l.Operand1, l.Operand2, l.Args)
			t.FailNow()
		}
		if exec.PC != l.Exec.PC {
			t.Errorf("Decode failed: PC mismatch: expected %04X, got %04X", l.Exec.PC, exec.PC)
		}
		if exec.Opcode != l.Exec.Opcode {
			t.Errorf("Decode failed: Opcode mismatch: expected %02X %s, got %02X", l.Exec.Opcode, l.Operation, exec.Opcode)
		}

		cpu.Fetch(exec)

		if exec.Operand1 != l.Exec.Operand1 || exec.Operand2 != l.Exec.Operand2 {
			t.Errorf("Decode failed: Operand mismatch: expected %02X %02X (%s), got %02X %02X", l.Exec.Operand1, l.Exec.Operand2, l.Args, exec.Operand1, exec.Operand2)
		}

		cpu.Execute(exec)
	}

}
