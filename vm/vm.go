package vm

import "fmt"

var (
	defaulRomSize   uint32 = 50000 * 8 // Each instruction takes 8 bytes
	codeSegmentSize uint32 = 550000 * 8
	dataSegmentSize uint32 = 16000000
)

type VM struct {
	memory []uint8
	rom    []uint64
	cpu    *CPU
}

func MakeVM(memorySize uint32) *VM {
	if memorySize >= defaulRomSize+codeSegmentSize+dataSegmentSize {
		vm := &VM{memory: make([]uint8, memorySize), rom: make([]uint64, 0)}
		cpu := MakeCPU(vm)
		vm.cpu = cpu
		return vm
	}
	return nil
}

func (vm *VM) AddInstruction(instruction uint64) {
	vm.rom = append(vm.rom, instruction)
	// fmt.Println(len(vm.rom))
	if len(vm.rom)*8 > int(defaulRomSize) {
		panic("Exceed ROM size")
	}
}

func (vm *VM) FlashRom(rom []uint64) {
	vm.rom = rom
}

func (vm *VM) StartVM() {
	vm.loadRom()
	vm.cpu.Run()
}

func (vm *VM) loadRom() {
	for i := 0; i < len(vm.rom); i++ {
		for j := 0; j < 8; j++ {
			vm.memory[i*8+j] = 0xff & uint8(vm.rom[i]>>((7-j)*8))
		}
	}
}

func (vm *VM) getDataSegment() uint32 {
	return defaulRomSize + codeSegmentSize
}

func (vm *VM) LoadInstruction(index uint64) uint64 {
	var code uint64 = 0
	var i uint64
	for i = 0; i < 8; i++ {
		code = code << 8
		code = code | uint64(vm.memory[index*8+i])
	}
	return code
}

func (vm *VM) DebugRom() {
	for i := 0; i < len(vm.rom); i++ {
		fmt.Printf(" %b", vm.rom[i])
	}
	fmt.Println()
}

func (vm *VM) DebugMemory() {
	for i := 0; i < 8*10; i++ {
		fmt.Printf(" %d", vm.memory[i])
	}
}

func (vm *VM) DebugStack() {
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d", vm.cpu.stack.data[i])
	}
}
