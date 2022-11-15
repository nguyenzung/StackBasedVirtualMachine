package vm

import (
	"fmt"
	"time"
)

type CPU struct {
	vm    *VM
	stack *Stack
	ip    uint32
}

func MakeCPU(vm *VM) *CPU {
	cpu := CPU{vm: vm, stack: MakeStack(), ip: 0}
	return &cpu
}

func (cpu *CPU) Run() {
	steps := 10
	counter := 0
	idleTime := 1
	for {
		instruction := cpu.fetch()
		opcode, operand := cpu.decode(instruction)
		cpu.exec(opcode, operand)
		if counter == steps {
			break
		} else {
			counter++
		}
		if idleTime > 0 {
			time.Sleep(time.Millisecond * time.Duration(idleTime))
		}
	}
}

func (cpu *CPU) fetch() uint64 {
	instruction := cpu.vm.LoadInstruction(cpu.ip)
	cpu.ip += 1
	return instruction
}

func (cpu *CPU) decode(instruction uint64) (uint8, uint64) {
	opcode := uint8(instruction >> 56)
	operand := instruction << 8
	operand = operand >> 8
	return opcode, operand
}

func (cpu *CPU) exec(opcode uint8, operand uint64) {
	fmt.Println("Exec instruction", opcode, operand)

	switch opcode {
	case PUSH:
		cpu.processPush(operand)
	case ADD:
		cpu.processAdd()
	default:
		break
	}
}

func (cpu *CPU) processPush(value uint64) {
	cpu.stack.Push(value)
}

func (cpu *CPU) processAdd() {
	a := cpu.stack.Pop()
	b := cpu.stack.Pop()
	cpu.stack.Push(a + b)
}
