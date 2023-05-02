package vm

import (
	"fmt"
	"testing"
)

type TestCase struct {
	t           *testing.T
	vm          *VM
	stackValue  map[int]uint64
	memoryValue map[int]uint8
}

func MakeTestCase(t *testing.T) *TestCase {
	return &TestCase{t: t, vm: MakeVM(8 * 10000000), stackValue: make(map[int]uint64), memoryValue: make(map[int]uint8)}
}

func (testCase *TestCase) AddStep(value uint64) {
	testCase.vm.AddInstruction(value)
}

func (testCase *TestCase) AddStackTest(k int, v uint64) {
	testCase.stackValue[k] = v
}

func (testCase *TestCase) AddMemoryTest(k int, v uint8) {
	testCase.memoryValue[k] = v
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
			t.Errorf("Error at Mem item %d, Mem value: %d Expected value %d", k, stack.data[k], v)
		}
	}
}

func TestAdd(t *testing.T) {
	// tests :=
	fmt.Println("TestAdd")
	testCase := MakeTestCase(t)
	testCase.AddStep(MakePUSH(5))
	testCase.AddStep(MakePUSH(13))
	testCase.AddStep(MakePUSH(7))
	testCase.AddStep(MakeADD())
	testCase.AddStep(MakePUSH(7))
	testCase.AddStep(MakeADD())
	testCase.AddStackTest(0, 18)
	testCase.Assert()
}

func TestSub(t *testing.T) {
	// tests :=
	fmt.Println("Testsub")
}
