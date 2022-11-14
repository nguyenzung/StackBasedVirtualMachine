package vm

/*
*	Each instruction has size 64 bits
*	First 8 bits used for opcode
*	Next 56 bits used for operand
 */
var (
	PUSH   uint8 = 0x01
	ADD    uint8 = 0x04
	SUB    uint8 = 0x05
	MUL    uint8 = 0x06
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
	SHR    uint8 = 0x13
	DUP    uint8 = 0x38
	SWAP   uint8 = 0x39
	LOAD   uint8 = 0x40
	STORE  uint8 = 0x41
	SLOAD  uint8 = 0x60
	SSTORE uint8 = 0x61
	CALL   uint8 = 0x80
	RET    uint8 = 0x81
)

/*
*	PUSH constant
 */
func MakePush(value uint64) uint64 {
	value = value << 8
	value = value >> 8
	var opcode uint64 = uint64(PUSH)
	opcode = opcode << 56
	value = value | opcode
	return value
}

func MakeAdd() uint64 {
	var opcode uint64 = uint64(ADD)
	opcode = opcode << 56
	return opcode
}

func MakeSub() uint64 {
	var opcode uint64 = uint64(SUB)
	opcode = opcode << 56
	return opcode
}
