package vm

import (
	"fmt"
	"time"
)

type CPU struct {
	vm    *VM
	stack *Stack
	ip    uint64
	hlt   bool
}

func MakeCPU(vm *VM) *CPU {
	cpu := CPU{vm: vm, stack: MakeStack(), ip: 0}
	return &cpu
}

func (cpu *CPU) Run() {
	steps := 20
	counter := 0
	idleTime := 1
	for !cpu.hlt {
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
	case POP:
		cpu.processPop()
	case ADD:
		cpu.processAdd()
	case SUB:
		cpu.processSub()
	case MUL:
		cpu.processMul()
	case DIV:
		cpu.processDiv()
	case MOD:
		cpu.processMod()
	case DUP:
		cpu.processDup()
	case SWAP:
		cpu.processSwap()
	case JMP:
		cpu.processJmp()

	default:
		cpu.stop()
	}
}

func (cpu *CPU) processPush(value uint64) {
	cpu.stack.Push(value)
}

func (cpu *CPU) processPop() {
	cpu.stack.Pop()
}

func (cpu *CPU) processAdd() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(a + b)
}

func (cpu *CPU) processSub() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(a - b)
}

func (cpu *CPU) processMul() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(a * b)
}

func (cpu *CPU) processDiv() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(a / b)
}

func (cpu *CPU) processMod() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(a % b)
}

func (cpu *CPU) processDup() {
	a := cpu.stack.Top()
	cpu.stack.Push(a)
}

func (cpu *CPU) processSwap() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(b)
	cpu.stack.Push(a)
}

func (cpu *CPU) processJmp() {
	a := cpu.stack.Pop()
	cpu.setPC(a)
}

func (cpu *CPU) setPC(value uint64) {
	cpu.ip = value
}

func (cpu *CPU) stop() {
	cpu.hlt = true
}
