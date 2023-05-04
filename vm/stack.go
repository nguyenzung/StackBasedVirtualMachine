package vm

import (
	"fmt"
)

const MAX_DEPTH = 20000

type Stack struct {
	data      [MAX_DEPTH]uint64
	index     uint32
	baseIndex uint32
}

func MakeStack() *Stack {
	stack := Stack{index: 0, baseIndex: 0}
	return &stack
}

func (stack *Stack) Empty() bool {
	return stack.index == 0
}

func (stack *Stack) Reset() {
	stack.index = 0
}

func (stack *Stack) Push(value uint64) {

	if stack.index < MAX_DEPTH-1 {
		stack.data[stack.index] = value
		stack.index++
	} else {
		fmt.Println("Err: ", stack.index, MAX_DEPTH)
		panic("Push run out of stack")
	}
}

func (stack *Stack) Top() uint64 {
	return stack.data[stack.index-1]
}

func (stack *Stack) Pop() uint64 {
	if stack.index > stack.baseIndex {
		result := stack.data[stack.index-1]
		stack.index -= 1
		return result
	} else {
		panic("Exceed stack bottom, cannot pop more")
	}
}

func (stack *Stack) SetupCall(retPC uint64) {
	numParams := stack.data[stack.index-1]
	if stack.index-stack.baseIndex < uint32(numParams)+1 {
		panic("Invalid function call")
	}
	// calldata := stack.data[stack.index-1-uint32(numParams) : stack.index-1]
	stack.Push(retPC)
	stack.Push(uint64(stack.baseIndex))
	stack.baseIndex = stack.index
	if stack.baseIndex+uint32(numParams) < MAX_DEPTH {
		stack.index = stack.baseIndex + uint32(numParams)
		copy(stack.data[stack.baseIndex:stack.baseIndex+uint32(numParams)], stack.data[stack.baseIndex-3-uint32(numParams):stack.baseIndex-3])
	} else {
		panic("Setup call run out of stack")
	}
	// fmt.Println("Setup calldata", calldata, stack.baseIndex, stack.index, stack.data[stack.baseIndex:stack.baseIndex+uint32(numParams)])
}

func (stack *Stack) SetupReturn() uint64 {
	// fmt.Println("Stack value", stack.data[:stack.index], stack.baseIndex, stack.index)
	var retValue uint64 = 0
	hasRet := stack.index > stack.baseIndex
	if hasRet {
		retValue = stack.Top()
	}
	baseIndex := stack.data[stack.baseIndex-1]
	pc := stack.data[stack.baseIndex-2]
	numParams := stack.data[stack.baseIndex-3]
	stack.index = stack.baseIndex - 3 - uint32(numParams)
	stack.baseIndex = uint32(baseIndex)
	if hasRet {
		stack.Push(retValue)
	}
	return pc
}
