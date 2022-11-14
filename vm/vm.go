package vm

var (
	defaulRomSize   uint32 = 100000 * 8 // Each instruction takes 8 bytes
	minRunningSpace uint32 = 100000 * 8
)

type VM struct {
	memory []uint8
	rom    []Instruction
	ip     uint32
}

func MakeVM(memorySize uint32) *VM {
	if memorySize >= defaulRomSize+minRunningSpace {
		vm := &VM{memory: make([]uint8, memorySize), rom: make([]Instruction, defaulRomSize), ip: 0}
		return vm
	}
	return nil
}

func (vm *VM) ClearRom() {
	
}

func (vm *VM) AddInstruction(instruction Instruction) {
	vm.rom = append(vm.rom, instruction)
	if len(vm.rom) > int(defaulRomSize) {
		panic("Exceed ROM size")
	}
}

func (vm *VM) StartVM() {

}
