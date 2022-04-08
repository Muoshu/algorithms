package Stack

type stackItem struct {
	item interface{}
	next *stackItem
}

type stack struct {
	sp    *stackItem
	depth uint64
}

// New new a stack
func New() *stack {
	var stack *stack = new(stack)
	stack.depth = 0
	return stack
}
