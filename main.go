package main

import (
	"fmt"

	"github.com/nguyenzung/StackBasedVirtualMachine/vm"
)

func main() {
	fmt.Println("Virtual Machine")
	myVM := vm.MakeVM(8 * 10000000)

	// setup bootstrap program is stored on ROM
	myVM.AddInstruction(vm.MakePUSH(15))  // 0
	myVM.AddInstruction(vm.MakePUSH(25))  // 1
	myVM.AddInstruction(vm.MakeADD())     // 2
	myVM.AddInstruction(vm.MakePUSH(60))  // 3
	myVM.AddInstruction(vm.MakeADD())     // 4
	myVM.AddInstruction(vm.MakePUSH(120)) // 5
	myVM.AddInstruction(vm.MakeSUB())     // 6
	myVM.AddInstruction(vm.MakePUSH(20))  // 7
	myVM.AddInstruction(vm.MakeADD())     // 8
	myVM.AddInstruction(vm.MakePUSH(32))  // 9
	myVM.AddInstruction(vm.MakeDUP())     // 10
	myVM.AddInstruction(vm.MakePUSH(14))  // 11
	myVM.AddInstruction(vm.MakeJMP())     // 12
	myVM.AddInstruction(vm.MakePUSH(24))  // 13
	myVM.AddInstruction(vm.MakePUSH(15))  // 14
	myVM.AddInstruction(vm.MakePUSH(15))  // 15
	myVM.AddInstruction(vm.MakeXOR())     // 16
	myVM.AddInstruction(vm.MakePUSH(0))   // 17
	myVM.AddInstruction(vm.MakeAND())     // 18
	myVM.AddInstruction(vm.MakePUSH(1))   // 19
	myVM.AddInstruction(vm.MakeOR())      // 20
	myVM.AddInstruction(vm.MakeADD())     // 21
	myVM.AddInstruction(vm.MakeOR())      // 22
	myVM.AddInstruction(vm.MakePUSH(13))  // 23

	myVM.DebugRom()
	myVM.StartVM()
	myVM.DebugMemory()
	myVM.DebugStack()
}
