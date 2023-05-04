package vm

import (
	"encoding/binary"
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
	idleTime := 0
	for !cpu.hlt {
		instruction := cpu.fetch()
		opcode, operand := cpu.decode(instruction)
		cpu.exec(opcode, operand)
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
	case INC:
		cpu.processINC()
	case DEC:
		cpu.processDEC()
	case SHL:
		cpu.processSHL()
	case SHR:
		cpu.processSHR()
	case DUP:
		cpu.processDup()
	case SWAP:
		cpu.processSwap()
	case EQ:
		cpu.processEQ()
	case LT:
		cpu.processLT()
	case GT:
		cpu.processGT()
	case LTE:
		cpu.processLTE()
	case GTE:
		cpu.processGTE()
	case LOAD:
		cpu.processLOAD()
	case STORE:
		cpu.processSTORE()
	case LOAD8:
		cpu.processLOAD8()
	case STORE8:
		cpu.processSTORE8()
	case JMP:
		cpu.processJmp()
	case JN:
		cpu.processJN()
	case JP:
		cpu.processJP()
	case JZ:
		cpu.processJZ()
	case JNZ:
		cpu.processJNZ()
	case JE:
		cpu.processJE()
	case JNE:
		cpu.processJNE()
	case JLT:
		cpu.processJLT()
	case JGT:
		cpu.processJGT()
	case JLE:
		cpu.processJLE()
	case JGE:
		cpu.processJGE()
	case TIME:
		cpu.processTIME()
	case CALL:
		cpu.processCALL(operand)
	case RET:
		cpu.processRET()
	case HLT:
		cpu.processHLT()
	case SPACE:
		cpu.processSPACE()
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

func (cpu *CPU) processINC() {
	a := cpu.stack.Pop()
	cpu.stack.Push(a + 1)
}

func (cpu *CPU) processDEC() {
	a := cpu.stack.Pop()
	cpu.stack.Push(a - 1)
}

func (cpu *CPU) processSHL() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(a << b)
}

func (cpu *CPU) processSHR() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	cpu.stack.Push(a >> b)
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

func (cpu *CPU) processEQ() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	var r uint64 = 0
	if a == b {
		r = 1
	}
	cpu.stack.Push(r)
}

func (cpu *CPU) processLT() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	var r uint64 = 0
	if a < b {
		r = 1
	}
	cpu.stack.Push(r)
}

func (cpu *CPU) processGT() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	var r uint64 = 0
	if a > b {
		r = 1
	}
	cpu.stack.Push(r)
}

func (cpu *CPU) processLTE() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	var r uint64 = 0
	if a <= b {
		r = 1
	}
	cpu.stack.Push(r)
}

func (cpu *CPU) processGTE() {
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	var r uint64 = 0
	if a >= b {
		r = 1
	}
	cpu.stack.Push(r)
}

func (cpu *CPU) processLOAD() {
	index := cpu.stack.Pop() + uint64(cpu.vm.getDataSegment())
	bytes := cpu.vm.memory[index : index+8]
	num := binary.LittleEndian.Uint64(bytes)
	cpu.stack.Push(num)
}

func (cpu *CPU) processSTORE() {
	index := cpu.stack.Pop() + uint64(cpu.vm.getDataSegment())
	value := cpu.stack.Pop()
	binary.LittleEndian.PutUint64(cpu.vm.memory[index:index+8], value)
}

func (cpu *CPU) processLOAD8() {
	index := cpu.stack.Pop() + uint64(cpu.vm.getDataSegment())
	num := uint64(cpu.vm.memory[index])
	cpu.stack.Push(num)
}

func (cpu *CPU) processSTORE8() {
	index := cpu.stack.Pop() + uint64(cpu.vm.getDataSegment())
	value := cpu.stack.Pop()
	cpu.vm.memory[index] = uint8(value & 0x00000000000000ff)
}

func (cpu *CPU) processJmp() {
	ip := cpu.stack.Pop()
	cpu.setPC(ip)
}

func (cpu *CPU) processJN() {
	ip := cpu.stack.Pop()
	a := cpu.stack.Pop()
	b := ((a << 1) >> 1)
	a = a >> 63
	if a == 1 && b > 0 {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processJP() {
	ip := cpu.stack.Pop()
	a := cpu.stack.Pop()
	b := ((a << 1) >> 1)
	a = a >> 63
	if a == 0 && b > 0 {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processJZ() {
	ip := cpu.stack.Pop()
	a := cpu.stack.Pop()
	if a == 0 {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processJNZ() {
	ip := cpu.stack.Pop()
	a := cpu.stack.Pop()
	if a != 0 {
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

func (cpu *CPU) processJLT() {
	ip := cpu.stack.Pop()
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	if a < b {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processJGT() {
	ip := cpu.stack.Pop()
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	if a > b {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processJLE() {
	ip := cpu.stack.Pop()
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	if a <= b {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processJGE() {
	ip := cpu.stack.Pop()
	b := cpu.stack.Pop()
	a := cpu.stack.Pop()
	if a >= b {
		cpu.setPC(ip)
	}
}

func (cpu *CPU) processTIME() {
	time := uint64(time.Now().UnixMilli())
	cpu.stack.Push(time)
}

func (cpu *CPU) processCALL(label uint64) {
	cpu.stack.SetupCall(cpu.ip)
	cpu.setPC(label)
}

func (cpu *CPU) processRET() {
	pc := cpu.stack.SetupReturn()
	cpu.setPC(pc)
}

func (cpu *CPU) processHLT() {
	cpu.hlt = true
}

func (cpu *CPU) processSPACE() {
	cpu.stack.Push(uint64(cpu.vm.getDataSegment()))
}

func (cpu *CPU) setPC(value uint64) {
	cpu.ip = value
}

func (cpu *CPU) stop() {
	cpu.hlt = true
}
