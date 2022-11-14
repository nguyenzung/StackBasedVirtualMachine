package vm

import "fmt"

var (
	defaulRomSize   uint32 = 20000 * 8 // Each instruction takes 8 bytes
	codeSegmentSize uint32 = 200000 * 8
	dataSegmentSize uint32 = 2000000 * 8
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
	fmt.Println(len(vm.rom))
	if len(vm.rom) > int(defaulRomSize) {
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
	fmt.Println("Load ROM")
	for i := 0; i < len(vm.rom); i++ {
		for j := 0; j < 8; j++ {
			vm.memory[i*8+j] = 0xff & uint8(vm.rom[i]>>((7-j)*8))
		}
	}
}

func (vm *VM) LoadInstruction(index uint32) uint64 {
	var code uint64 = 0
	var i uint32
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
