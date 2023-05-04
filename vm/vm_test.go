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
	testCase.AddStep(MakePUSH(5))
	testCase.AddStep(MakePUSH(13))
	testCase.AddStep(MakePUSH(7))
	testCase.AddStep(MakeADD())
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakeADD())
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
	testCase.AddStep(MakePUSH(11))
	testCase.AddStep(MakeSUB())
	testCase.AddStackTest(0, 0xffffffffffffffff)
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

func TestARI(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakeDUP())
	testCase.AddStep(MakeINC())
	testCase.AddStep(MakeADD())
	testCase.AddStep(MakePUSH(4))
	testCase.AddStep(MakeDEC())
	testCase.AddStep(MakeADD())
	testCase.AddStep(MakePUSH(6))
	testCase.AddStep(MakeEQ())    // [1]
	testCase.AddStep(MakeDUP())   // [1 1]
	testCase.AddStep(MakePUSH(2)) // [1 1 2]
	testCase.AddStep(MakeEQ())    // [1 0]
	testCase.AddStep(MakeDUP())   // [1 0 0]
	testCase.AddStep(MakePUSH(2)) // [1 0 0 2]
	testCase.AddStep(MakeLT())    // [1 0 1]
	testCase.AddStep(MakeDUP())   // [1 0 1 1]
	testCase.AddStep(MakePUSH(1)) // [1 0 1 1 1]
	testCase.AddStep(MakeLT())    // [1 0 1 0]
	testCase.AddStep(MakeDUP())   // [1 0 1 0 0]
	testCase.AddStep(MakeINC())   // [1 0 1 0 1]
	testCase.AddStep(MakeDUP())   // [1 0 1 0 1 1]
	testCase.AddStep(MakePUSH(0)) // [1 0 1 0 1 1 0]
	testCase.AddStep(MakeGT())    // [1 0 1 0 1 1]
	testCase.AddStep(MakeDUP())   // [1 0 1 0 1 1 1]
	testCase.AddStep(MakePUSH(1)) // [1 0 1 0 1 1 1 1]
	testCase.AddStep(MakeGT())    // [1 0 1 0 1 1 0]
	testCase.AddStep(MakeDUP())   // [1 0 1 0 1 1 0 0]
	testCase.AddStep(MakePUSH(5)) // [1 0 1 0 1 1 0 0 5]
	testCase.AddStep(MakeADD())   // [1 0 1 0 1 1 0 5]
	testCase.AddStackTest(0, 1)
	testCase.AddStackTest(1, 0)
	testCase.AddStackTest(2, 1)
	testCase.AddStackTest(3, 0)
	testCase.AddStackTest(4, 1)
	testCase.AddStackTest(5, 1)
	testCase.AddStackTest(6, 0)
	testCase.AddStackTest(7, 5)
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

func TestLOAD8(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.UpdateMemoryAddress(0, 0xfd)
	testCase.UpdateMemoryAddress(0xfd, 0xff)
	testCase.AddStep(MakePUSH(0))
	testCase.AddStep(MakeLOAD8())
	testCase.AddStep(MakeDUP())
	testCase.AddStep(MakeLOAD8())
	testCase.AddStackTest(0, 0xfd)
	testCase.AddStackTest(1, 0xff)
	testCase.Assert()
}

func TestSTORE8(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(0xee))
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakeSTORE8())
	testCase.AddStep(MakePUSH(2023))
	testCase.AddStackTest(0, 2023)
	testCase.AddMemoryTest(0, 0x00)
	testCase.AddMemoryTest(1, 0xee)
	testCase.AddMemoryTest(2, 0x00)
	testCase.Assert()
}

func TestJMP(t *testing.T) {
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(4))
	testCase.AddStep(MakeJMP())
	testCase.AddStep(MakePUSH(12))
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023))
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJN(t *testing.T) {
	var FAILED_LABEL uint64 = 8
	var PASSED_LABEL uint64 = 10
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(0))
	testCase.AddStep(MakePUSH(FAILED_LABEL)) // jump to the instruction 8
	testCase.AddStep(MakeJN())
	testCase.AddStep(MakePUSH(0xffffffffffffff))
	testCase.AddStep(MakePUSH(8))
	testCase.AddStep(MakeSHL())
	testCase.AddStep(MakePUSH(PASSED_LABEL)) // jump to the instruction 10
	testCase.AddStep(MakeJN())
	testCase.AddStep(MakePUSH(15))
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023))
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJP(t *testing.T) {
	var FAILED_LABEL uint64 = 8
	var PASSED_LABEL uint64 = 10
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(0))
	testCase.AddStep(MakePUSH(FAILED_LABEL)) // jump to the instruction 8
	testCase.AddStep(MakeJP())
	testCase.AddStep(MakePUSH(0xffffffffffffff))
	testCase.AddStep(MakePUSH(7))
	testCase.AddStep(MakeSHL())
	testCase.AddStep(MakePUSH(PASSED_LABEL)) // jump to the instruction 10
	testCase.AddStep(MakeJP())
	testCase.AddStep(MakePUSH(15))
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023))
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJZ(t *testing.T) {
	var FAILED_LABEL uint64 = 6
	var PASSED_LABEL uint64 = 8
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(FAILED_LABEL)) // jump to the instruction 6
	testCase.AddStep(MakeJZ())
	testCase.AddStep(MakePUSH(0))
	testCase.AddStep(MakePUSH(PASSED_LABEL)) // jump to the instruction 8
	testCase.AddStep(MakeJZ())
	testCase.AddStep(MakePUSH(15))
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023))
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJNZ(t *testing.T) {
	var FAILED_LABEL uint64 = 6
	var PASSED_LABEL uint64 = 8
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(0))
	testCase.AddStep(MakePUSH(FAILED_LABEL)) // jump to the instruction 6
	testCase.AddStep(MakeJNZ())
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(PASSED_LABEL)) // jump to the instruction 8
	testCase.AddStep(MakeJNZ())
	testCase.AddStep(MakePUSH(15))
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023))
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJE(t *testing.T) {
	var FAILED_LABEL uint64 = 8
	var PASSED_LABEL uint64 = 10
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(FAILED_LABEL)) // jump to the instruction 8
	testCase.AddStep(MakeJE())
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(PASSED_LABEL)) // jump to the instruction 10
	testCase.AddStep(MakeJE())
	testCase.AddStep(MakePUSH(15))
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023))
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJNE(t *testing.T) {
	var FAILED_LABEL uint64 = 8
	var PASSED_LABEL uint64 = 10
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(FAILED_LABEL)) // jump to the instruction 8
	testCase.AddStep(MakeJNE())
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(PASSED_LABEL)) // jump to the instruction 10
	testCase.AddStep(MakeJNE())
	testCase.AddStep(MakePUSH(15))
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023))
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJLT(t *testing.T) {
	var FAILED_LABEL uint64 = 12
	var PASSED_LABEL uint64 = 14
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(FAILED_LABEL))
	testCase.AddStep(MakeJLT())
	testCase.AddStep(MakePUSH(3))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(FAILED_LABEL))
	testCase.AddStep(MakeJLT())
	testCase.AddStep(MakePUSH(3))
	testCase.AddStep(MakePUSH(5))
	testCase.AddStep(MakePUSH(PASSED_LABEL))
	testCase.AddStep(MakeJLT())
	testCase.AddStep(MakePUSH(15)) // JMP to failed
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023)) // JMP to PASS
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJGT(t *testing.T) {
	var FAILED_LABEL uint64 = 12
	var PASSED_LABEL uint64 = 14
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(FAILED_LABEL))
	testCase.AddStep(MakeJGT())
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(3))
	testCase.AddStep(MakePUSH(FAILED_LABEL))
	testCase.AddStep(MakeJGT())
	testCase.AddStep(MakePUSH(8))
	testCase.AddStep(MakePUSH(5))
	testCase.AddStep(MakePUSH(PASSED_LABEL))
	testCase.AddStep(MakeJGT())
	testCase.AddStep(MakePUSH(15)) // JMP to failed
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023)) // JMP to PASS
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJLE_LESS(t *testing.T) {
	var FAILED_LABEL uint64 = 8
	var PASSED_LABEL uint64 = 10
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(FAILED_LABEL))
	testCase.AddStep(MakeJLE())
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(PASSED_LABEL))
	testCase.AddStep(MakeJLE())
	testCase.AddStep(MakePUSH(15)) // JMP to failed
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023)) // JMP to PASS
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJLE_EQUAL(t *testing.T) {
	var FAILED_LABEL uint64 = 8
	var PASSED_LABEL uint64 = 10
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(FAILED_LABEL))
	testCase.AddStep(MakeJLE())
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(PASSED_LABEL))
	testCase.AddStep(MakeJLE())
	testCase.AddStep(MakePUSH(15)) // JMP to failed
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023)) // JMP to PASS
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJGE_GREATER(t *testing.T) {
	var FAILED_LABEL uint64 = 8
	var PASSED_LABEL uint64 = 10
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(FAILED_LABEL))
	testCase.AddStep(MakeJGE())
	testCase.AddStep(MakePUSH(3))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(PASSED_LABEL))
	testCase.AddStep(MakeJGE())
	testCase.AddStep(MakePUSH(15)) // JMP to failed
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023)) // JMP to PASS
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
	testCase.Assert()
}

func TestJGE_EQUAL(t *testing.T) {
	var FAILED_LABEL uint64 = 8
	var PASSED_LABEL uint64 = 10
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(FAILED_LABEL))
	testCase.AddStep(MakeJGE())
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(PASSED_LABEL))
	testCase.AddStep(MakeJGE())
	testCase.AddStep(MakePUSH(15)) // JMP to failed
	testCase.AddStep(MakeHLT())
	testCase.AddStep(MakePUSH(2023)) // JMP to PASS
	testCase.AddStep(MakeHLT())
	testCase.AddStackTest(0, 2023)
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

func TestSetupFuncCall(t *testing.T) {
	var SUM_3_NUMBER_FUNCTION uint64 = 13
	var SUM_2_NUMBER_FUNCTION uint64 = 18
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(0))
	testCase.AddStep(MakePUSH(1))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakePUSH(3))
	testCase.AddStep(MakeCALL(SUM_3_NUMBER_FUNCTION))
	testCase.AddStep(MakePUSH(4))
	testCase.AddStep(MakePUSH(5))
	testCase.AddStep(MakePUSH(3))
	testCase.AddStep(MakeCALL(SUM_3_NUMBER_FUNCTION))
	testCase.AddStep(MakePUSH(3))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakeCALL(SUM_2_NUMBER_FUNCTION))
	testCase.AddStep(MakeHLT())

	testCase.AddStep(MakePUSH(2)) // Sum (a + b + c)
	testCase.AddStep(MakeCALL(SUM_2_NUMBER_FUNCTION))
	testCase.AddStep(MakePUSH(2))
	testCase.AddStep(MakeCALL(SUM_2_NUMBER_FUNCTION))
	testCase.AddStep(MakeRET())

	testCase.AddStep(MakeADD()) // Sum (a + b)
	testCase.AddStep(MakeRET())

	testCase.AddStackTest(0, 15)
	testCase.Assert()
}

// Calculate sum from 1 -> n
// Write the result to the ten
func TestSumFrom1ToN(t *testing.T) {
	var sumSlot uint64 = 0
	var iSlot uint64 = 8
	var nSlot uint64 = 9
	var FINISH_LABEL uint64 = 28

	testCase := MakeTestCase(t)

	/*Setup memory at 0 for sum variable*/
	testCase.AddStep(MakePUSH(0))
	testCase.AddStep(MakePUSH(sumSlot))
	testCase.AddStep(MakeSTORE())

	/*Setup memory at 8 for i variable*/
	testCase.AddStep(MakePUSH(0))
	testCase.AddStep(MakePUSH(iSlot))
	testCase.AddStep(MakeSTORE8())

	/*Setup memory at 9 for n variable*/
	testCase.AddStep(MakePUSH(100))
	testCase.AddStep(MakePUSH(nSlot))
	testCase.AddStep(MakeSTORE8())

	testCase.AddStep(MakePUSH(iSlot))
	testCase.AddStep(MakeLOAD8())
	testCase.AddStep(MakePUSH(nSlot))
	testCase.AddStep(MakeLOAD8())

	testCase.AddStep(MakePUSH(FINISH_LABEL))
	testCase.AddStep(MakeJGT())

	testCase.AddStep(MakePUSH(sumSlot))
	testCase.AddStep(MakeLOAD())
	testCase.AddStep(MakePUSH(iSlot))
	testCase.AddStep(MakeLOAD8())
	testCase.AddStep(MakeDUP())
	testCase.AddStep(MakeINC())
	testCase.AddStep(MakePUSH(iSlot))
	testCase.AddStep(MakeSTORE8())
	testCase.AddStep(MakeADD())
	testCase.AddStep(MakePUSH(sumSlot))
	testCase.AddStep(MakeSTORE())
	testCase.AddStep(MakePUSH(9))
	testCase.AddStep(MakeJMP())
	testCase.AddStep(MakeHLT())

	testCase.AddMemoryTest(0, 0xBA)
	testCase.AddMemoryTest(1, 0x13)
	testCase.AddMemoryTest(8, 101)
	testCase.AddMemoryTest(9, 100)
	testCase.AddMemoryTest(10, 0)
	testCase.Assert()
}
