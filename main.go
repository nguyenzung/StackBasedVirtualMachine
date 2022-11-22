package main

import (
	"fmt"

	"github.com/nguyenzung/StackBasedVirtualMachine/vm"
)

func main() {
	fmt.Println("Virtual Machine")

	var a, b uint64 = 12, 18
	fmt.Println(a - b)
	a = 0
	b = 1
	fmt.Println(^a, a-b)

	myVM := vm.MakeVM(8 * 10000000)
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
	myVM.AddInstruction(vm.MakePUSH(16))  // 14

	myVM.DebugRom()
	myVM.StartVM()
	myVM.DebugMemory()
	myVM.DebugStack()

}
