package vm

type Stack []uint64

func (stack *Stack) Empty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(value uint64) {
	*stack = append(*stack, value)
}

func (stack *Stack) Pop() uint64 {
	index := len(*stack) - 1
	value := (*stack)[len(*stack)-1]
	*stack = (*stack)[:index]
	return value
}
