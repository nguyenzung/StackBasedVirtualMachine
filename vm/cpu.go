package vm

import (
	"math"
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
	steps := 30
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
	// fmt.Println("Exec instruction", opcode, operand)
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
	case AND:
		cpu.processAnd()
	case OR:
		cpu.processOr()
	case XOR:
		cpu.processXor()
	case NOT:
		cpu.processNot()
	case DUP:
		cpu.processDup()
	case SWAP:
		cpu.processSwap()
	case JMP:
		cpu.processJmp()
	case JN:
		cpu.processJN()
	case JP:
		cpu.processJP()
	case JE:
		cpu.processJE()
	case JNE:
		cpu.processJNE()
	case TIME:
		cpu.processTIME()

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

func (cpu *CPU) processAnd() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(a & b)
}

func (cpu *CPU) processOr() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(a | b)
}

func (cpu *CPU) processXor() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(a ^ b)
}

func (cpu *CPU) processNot() {
	a := cpu.stack.Pop()
	cpu.stack.Push(math.MaxUint64 ^ a)
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
	ip := cpu.stack.Pop()
	cpu.setPC(ip)
}

func (cpu *CPU) processJN() {
	ip := cpu.stack.Pop()
	a := cpu.stack.Pop()
	a = a >> 63
	if a == 1 {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processJP() {
	ip := cpu.stack.Pop()
	a := cpu.stack.Pop()
	a = a >> 63
	if a == 0 {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processJE() {
	ip := cpu.stack.Pop()
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	if a == b {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processJNE() {
	ip := cpu.stack.Pop()
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	if a != b {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processTIME() {
	time := uint64(time.Now().UnixMilli())
	cpu.stack.Push(time)
}

func (cpu *CPU) setPC(value uint64) {
	cpu.ip = value
}

func (cpu *CPU) stop() {
	cpu.hlt = true
}
