package vm

/*
*	Each instruction has size 64 bits
*	First 8 bits used for opcode
*	Next 56 bits used for operand
 */
var (
	POP    uint8 = 0x00
	PUSH   uint8 = 0x01
	ADD    uint8 = 0x04
	SUB    uint8 = 0x05
	MUL    uint8 = 0x06 // stack[i - 1] = stack[i] - stack[i-1]
	DIV    uint8 = 0x07
	AND    uint8 = 0x08
	OR     uint8 = 0x09
	NAND   uint8 = 0x0A
	XOR    uint8 = 0x0B
	NOT    uint8 = 0x0C
	LT     uint8 = 0x0D // stack[i - 1] < stack[i]
	GT     uint8 = 0x0E // stack[i - 1] > stack[i]
	LTE    uint8 = 0x0F // stack[i - 1] <= stack[i]
	GTE    uint8 = 0x10 // stack[i - 1] >= stack[i]
	EQ     uint8 = 0x11 // stack[i - 1] == stack[i]
	SHL    uint8 = 0x12 // Shift left
	SHR    uint8 = 0x13 // Shift left
	FLIP   uint8 = 0x14 // Flip all bits
	INC    uint8 = 0x20
	DEC    uint8 = 0x21
	DUP    uint8 = 0x38
	SWAP   uint8 = 0x39
	LOAD   uint8 = 0x40
	STORE  uint8 = 0x41
	SLOAD  uint8 = 0x60
	SSTORE uint8 = 0x61
	CALL   uint8 = 0x80
	RET    uint8 = 0x81
	TIME   uint8 = 0x86
	JMP    uint8 = 0xA0 // Unconditinal jump
	JN     uint8 = 0xA1 // Jump if negative
	JP     uint8 = 0xA2 // Jump if positive
	JZ     uint8 = 0xA3 // Jump if zero
	JLE    uint8 = 0xA4 // Jump if less or equal 0
	JGE    uint8 = 0xA5 // Jump if greater or equal 0
)

func MakePOP() uint64 {
	var opcode uint64 = uint64(POP)
	opcode = opcode << 56
	return opcode
}

func MakePUSH(value uint64) uint64 {
	value = value << 8
	value = value >> 8
	var opcode uint64 = uint64(PUSH)
	opcode = opcode << 56
	value = value | opcode
	return value
}

func MakeADD() uint64 {
	var opcode uint64 = uint64(ADD)
	opcode = opcode << 56
	return opcode
}

func MakeSUB() uint64 {
	var opcode uint64 = uint64(SUB)
	opcode = opcode << 56
	return opcode
}
