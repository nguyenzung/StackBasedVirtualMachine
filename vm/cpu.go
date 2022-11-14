package vm

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

	instruction := cpu.fetch()
	cpu.exec(instruction)
}

func (cpu *CPU) fetch() uint64 {
	instruction := cpu.vm.LoadInstruction(cpu.ip)
	cpu.ip += 1
	return instruction
}

func (cpu *CPU) exec(instruction uint64) {
}
