package vm

/*
*	Each instruction has size 64 bits
*	First 8 bits used for opcode
*	Next 56 bits used for operand
 */
var (
	POP     uint8 = 0x01
	PUSH    uint8 = 0x02
	ADD     uint8 = 0x04
	SUB     uint8 = 0x05 // stack[i - 1] = stack[i - 1] - stack[i]
	MUL     uint8 = 0x06
	DIV     uint8 = 0x07
	AND     uint8 = 0x08
	OR      uint8 = 0x09
	NAND    uint8 = 0x0A
	XOR     uint8 = 0x0B
	NOT     uint8 = 0x0C
	LT      uint8 = 0x0D // stack[i - 1] < stack[i]
	GT      uint8 = 0x0E // stack[i - 1] > stack[i]
	LTE     uint8 = 0x0F // stack[i - 1] <= stack[i]
	GTE     uint8 = 0x10 // stack[i - 1] >= stack[i]
	EQ      uint8 = 0x11 // stack[i - 1] == stack[i]
	SHL     uint8 = 0x12 // Shift left stack[i - 1] by stack[i] bits
	SHR     uint8 = 0x13 // Shift right stack[i - 1] by stack[i] bits
	INC     uint8 = 0x20
	DEC     uint8 = 0x21
	MOD     uint8 = 0x22
	POW     uint8 = 0x23
	IMUL    uint8 = 0x24
	IDIV    uint8 = 0x25
	DUP     uint8 = 0x38
	SWAP    uint8 = 0x39
	LOAD    uint8 = 0x40 // Load 8 bytes from memory that point by stack[i]
	STORE   uint8 = 0x41 // Store 8 bytes at stack[i-1] to the memory that point by stack[i]
	LOAD8   uint8 = 0x42 // Load 8 bytes from memory that point by stack[i]
	STORE8  uint8 = 0x43 // Store 8 bytes at stack[i-1] to the memory that point by stack[i]
	SLOAD   uint8 = 0x60
	SSTORE  uint8 = 0x61
	SLOAD8  uint8 = 0x62
	SSTORE8 uint8 = 0x63
	CALL    uint8 = 0x80
	RET     uint8 = 0x81
	HLT     uint8 = 0x85
	TIME    uint8 = 0x86
	SPACE   uint8 = 0x87 // Load available RAM index after ROM
	JMP     uint8 = 0xA0 // Unconditinal jump
	JN      uint8 = 0xA1 // Jump if negative
	JP      uint8 = 0xA2 // Jump if positive
	JZ      uint8 = 0xA3 // Jump if zero
	JE      uint8 = 0xA4 // Jump if equal
	JNE     uint8 = 0xA5 // Jump if not equal
	JLT     uint8 = 0xA6 // Jump to stack[i] if stack[i-2] less than stack[i-1]
	JGT     uint8 = 0xA7 // Jump to stack[i] if stack[i-2] greater than stack[i-1]
	JLE     uint8 = 0xA8 // Jump to stack[i] if stack[i-2] less or equal stack[i-1]
	JGE     uint8 = 0xA9 // Jump to stack[i] if stack[i-2] greater or equal stack[i-1]
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

func MakeMUL() uint64 {
	var opcode uint64 = uint64(MUL)
	opcode = opcode << 56
	return opcode
}

func MakeDIV() uint64 {
	var opcode uint64 = uint64(DIV)
	opcode = opcode << 56
	return opcode
}

func MakeMOD() uint64 {
	var opcode uint64 = uint64(MOD)
	opcode = opcode << 56
	return opcode
}

func MakeAND() uint64 {
	var opcode uint64 = uint64(AND)
	opcode = opcode << 56
	return opcode
}

func MakeOR() uint64 {
	var opcode uint64 = uint64(OR)
	opcode = opcode << 56
	return opcode
}

func MakeXOR() uint64 {
	var opcode uint64 = uint64(XOR)
	opcode = opcode << 56
	return opcode
}

func MakeNOT() uint64 {
	var opcode uint64 = uint64(NOT)
	opcode = opcode << 56
	return opcode
}

func MakeLT() uint64 {
	var opcode uint64 = uint64(LT)
	opcode = opcode << 56
	return opcode
}

func MakeGT() uint64 {
	var opcode uint64 = uint64(GT)
	opcode = opcode << 56
	return opcode
}

func MakeLTE() uint64 {
	var opcode uint64 = uint64(LTE)
	opcode = opcode << 56
	return opcode
}

func MakeGTE() uint64 {
	var opcode uint64 = uint64(GTE)
	opcode = opcode << 56
	return opcode
}

func MakeEQ() uint64 {
	var opcode uint64 = uint64(EQ)
	opcode = opcode << 56
	return opcode
}

func MakeSHL() uint64 {
	var opcode uint64 = uint64(SHL)
	opcode = opcode << 56
	return opcode
}

func MakeSHR() uint64 {
	var opcode uint64 = uint64(SHR)
	opcode = opcode << 56
	return opcode
}

func MakeINC() uint64 {
	var opcode uint64 = uint64(INC)
	opcode = opcode << 56
	return opcode
}

func MakeDEC() uint64 {
	var opcode uint64 = uint64(DEC)
	opcode = opcode << 56
	return opcode
}

func MakePOW() uint64 {
	var opcode uint64 = uint64(POW)
	opcode = opcode << 56
	return opcode
}

func MakeIMUL() uint64 {
	var opcode uint64 = uint64(IMUL)
	opcode = opcode << 56
	return opcode
}

func MakeIDIV() uint64 {
	var opcode uint64 = uint64(IDIV)
	opcode = opcode << 56
	return opcode
}

func MakeDUP() uint64 {
	var opcode uint64 = uint64(DUP)
	opcode = opcode << 56
	return opcode
}

func MakeSWAP() uint64 {
	var opcode uint64 = uint64(SWAP)
	opcode = opcode << 56
	return opcode
}

func MakeLOAD() uint64 {
	var opcode uint64 = uint64(LOAD)
	opcode = opcode << 56
	return opcode
}

func MakeSTORE() uint64 {
	var opcode uint64 = uint64(STORE)
	opcode = opcode << 56
	return opcode
}

func MakeSLOAD() uint64 {
	var opcode uint64 = uint64(SLOAD)
	opcode = opcode << 56
	return opcode
}

func MakeSSTORE() uint64 {
	var opcode uint64 = uint64(SSTORE)
	opcode = opcode << 56
	return opcode
}

func MakeCALL() uint64 {
	var opcode uint64 = uint64(CALL)
	opcode = opcode << 56
	return opcode
}

func MakeRET() uint64 {
	var opcode uint64 = uint64(RET)
	opcode = opcode << 56
	return opcode
}

func MakeHLT() uint64 {
	var opcode uint64 = uint64(HLT)
	opcode = opcode << 56
	return opcode
}

func MakeSPACE() uint64 {
	var opcode uint64 = uint64(SPACE)
	opcode = opcode << 56
	return opcode
}

func MakeTIME() uint64 {
	var opcode uint64 = uint64(TIME)
	opcode = opcode << 56
	return opcode
}

func MakeJMP() uint64 {
	var opcode uint64 = uint64(JMP)
	opcode = opcode << 56
	return opcode
}

func MakeJP() uint64 {
	var opcode uint64 = uint64(JP)
	opcode = opcode << 56
	return opcode
}

func MakeJZ() uint64 {
	var opcode uint64 = uint64(JZ)
	opcode = opcode << 56
	return opcode
}

func MakeJN() uint64 {
	var opcode uint64 = uint64(JN)
	opcode = opcode << 56
	return opcode
}

func MakeJLT() uint64 {
	var opcode uint64 = uint64(JLT)
	opcode = opcode << 56
	return opcode
}

func MakeJGT() uint64 {
	var opcode uint64 = uint64(JGT)
	opcode = opcode << 56
	return opcode
}

func MakeJE() uint64 {
	var opcode uint64 = uint64(JE)
	opcode = opcode << 56
	return opcode
}

func MakeJNE() uint64 {
	var opcode uint64 = uint64(JNE)
	opcode = opcode << 56
	return opcode
}

func MakeJLE() uint64 {
	var opcode uint64 = uint64(JLE)
	opcode = opcode << 56
	return opcode
}

func MakeJGE() uint64 {
	var opcode uint64 = uint64(JGE)
	opcode = opcode << 56
	return opcode
}
