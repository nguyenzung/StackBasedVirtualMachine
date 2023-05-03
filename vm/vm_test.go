package vm

import (
	"testing"
	"time"

	"bou.ke/monkey"
)

type TestCase struct {
	t           *testing.T
	vm          *VM
	stackValue  map[int]uint64
	memoryValue map[uint32]uint8
}

func MakeTestCase(t *testing.T) *TestCase {
	return &TestCase{t: t, vm: MakeVM(8 * 10000000), stackValue: make(map[int]uint64), memoryValue: make(map[uint32]uint8)}
}

func (testCase *TestCase) AddStep(value uint64) {
	testCase.vm.AddInstruction(value)
}

func (testCase *TestCase) AddStackTest(k int, v uint64) {
	testCase.stackValue[k] = v
}

func (testCase *TestCase) AddMemoryTest(k uint32, v uint8) {
	vm := testCase.vm
	address := vm.getDataSegment() + k
	testCase.memoryValue[address] = v
}

func (testCase *TestCase) UpdateMemoryAddress(k uint32, v uint8) {
	vm := testCase.vm
	address := vm.getDataSegment()
	vm.memory[address+k] = v
}

func (testCase *TestCase) Assert() {
	testCase.vm.StartVM()
	t := testCase.t
	stack := testCase.vm.cpu.stack
	mem := testCase.vm.memory
	for k, v := range testCase.stackValue {
		if stack.data[k] != v {
			t.Errorf("Error at stack item %d, Stack value: %d Expected value %d", k, stack.data[k], v)
		}
	}

	for k, v := range testCase.memoryValue {
		if mem[k] != v {
			t.Errorf("Error at Mem item %d, Mem value: %d Expected value %d", k, mem[k], v)
		}
	}
}

func TestAdd(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(5))  // 1
	testCase.AddStep(MakePUSH(13)) // 2
	testCase.AddStep(MakePUSH(7))  // 3
	testCase.AddStep(MakeADD())    // 4
	testCase.AddStep(MakePUSH(1))  // 5
	testCase.AddStep(MakeADD())    // 6
	testCase.AddStackTest(0, 5)
	testCase.AddStackTest(1, 21)
	testCase.Assert()
}

func TestSub(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(35))
	testCase.AddStep(MakePUSH(15))
	testCase.AddStep(MakeSUB())
	testCase.AddStep(MakePUSH(10))
	testCase.AddStep(MakeSUB())
	testCase.AddStackTest(0, 10)
	testCase.Assert()
}

func TestSHL(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(12))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakeSHL())
	testCase.AddStackTest(0, 12<<2)
	testCase.Assert()
}

func TestSHR(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(12))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakeSHR())
	testCase.AddStackTest(0, 12>>2)
	testCase.Assert()
}

func TestLOAD(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.UpdateMemoryAddress(0, 0xfd)
	testCase.UpdateMemoryAddress(1, 0xfe)
	testCase.UpdateMemoryAddress(6, 0xfc)
	testCase.UpdateMemoryAddress(7, 0xff)
	testCase.AddStep(MakePUSH(0))
	testCase.AddStep(MakeLOAD())
	testCase.AddStackTest(0, 18445618173802774269)
	testCase.Assert()
}

func TestSTORE(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(70931694131150589))
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakeSTORE())
	testCase.AddStep(MakePUSH(2023))
	testCase.AddStackTest(0, 2023)
	testCase.AddMemoryTest(0, 0x00)
	testCase.AddMemoryTest(1, 0xfd)
	testCase.AddMemoryTest(2, 0xfe)
	testCase.AddMemoryTest(3, 0x00)
	testCase.AddMemoryTest(4, 0x00)
	testCase.AddMemoryTest(5, 0x00)
	testCase.AddMemoryTest(6, 0x00)
	testCase.AddMemoryTest(7, 0xfc)
	testCase.AddMemoryTest(8, 0x00)
	testCase.AddMemoryTest(9, 0x00)
	testCase.Assert()
}

func TestTIME(t *testing.T) {
	Now := func() time.Time {
		return time.Date(2023, 04, 30, 20, 0, 0, 0, time.UTC)
	}
	monkey.Patch(time.Now, Now)
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(35))
	testCase.AddStep(MakeTIME())
	testCase.AddStackTest(0, 35)
	testCase.AddStackTest(1, uint64(Now().UnixMilli()))
	testCase.Assert()
}

func TestSPACE(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.AddStep(MakeSPACE())
	testCase.AddStackTest(0, uint64(defaulRomSize+codeSegmentSize))
	testCase.Assert()
}

func TestHLT(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(5))
	testCase.AddStep(MakePUSH(5))
	testCase.AddStep(MakePUSH(5))
	testCase.AddStep(MakeADD())
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakeADD())
	testCase.AddStackTest(0, 5)
	testCase.AddStackTest(1, 10)
	testCase.Assert()
}

// Calculate sum from 1 -> n
// Write the result to the ten
func TestSumFrom1ToN(t *testing.T) {

}
