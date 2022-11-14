package main

import (
	"fmt"

	"github.com/nguyenzung/StackBasedVirtualMachine/vm"
)

func main() {
	fmt.Println("Virtual Machine")
	instruction := vm.MakePush(12)
	fmt.Printf("%64b\n", instruction)

	myVM := vm.MakeVM(8 * 10000000)
	myVM.AddInstruction(vm.MakePush(15))
	myVM.AddInstruction(vm.MakePush(25))
	myVM.DebugRom()

	myVM.StartVM()
	myVM.DebugMemory()

}
