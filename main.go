package main

import (
	"fmt"

	"github.com/nguyenzung/StackBasedVirtualMachine/vm"
)

func main() {
	fmt.Println("Virtual Machine")

	myVM := vm.MakeVM(8 * 10000000)
	myVM.AddInstruction(vm.MakePUSH(15))
	myVM.AddInstruction(vm.MakePUSH(25))
	myVM.AddInstruction(vm.MakeADD())
	myVM.AddInstruction(vm.MakePUSH(60))
	myVM.AddInstruction(vm.MakeADD())
	myVM.DebugRom()

	myVM.StartVM()
	myVM.DebugMemory()
	myVM.DebugStack()

}
