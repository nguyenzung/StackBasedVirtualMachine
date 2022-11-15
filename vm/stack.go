package vm

type Stack struct {
	data  [20000]uint64
	index uint32
}

func MakeStack() *Stack {
	stack := Stack{index: 0}
	return &stack
}

func (stack *Stack) Empty() bool {
	return stack.index == 0
}

func (stack *Stack) Push(value uint64) {
	stack.data[stack.index] = value
	stack.index++
}

func (stack *Stack) Top() uint64 {
	return stack.data[stack.index-1]
}

func (stack *Stack) Pop() uint64 {
	result := stack.data[stack.index-1]
	stack.index -= 1
	return result
}
