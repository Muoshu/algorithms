package stack

type Item struct {
	item interface{}
	next *Item
}

type Stack struct {
	sp    *Item
	depth int
}

// NewStack New new a stack
func NewStack() *Stack {
	var stack = new(Stack)
	stack.depth = 0
	return stack
}

// Pop delete and return the top element
func (st *Stack) Pop() interface{} {
	if st.depth > 0 {
		item := st.sp.item
		st.sp = st.sp.next
		st.depth--
		return item
	}
	return nil
}

// Push an element to the stack
func (st *Stack) Push(item interface{}) {
	st.sp = &Item{item: item, next: st.sp}
	st.depth++
}

// Peek return the top element(not delete)
func (st *Stack) Peek() interface{} {
	if st.depth > 0 {
		return st.sp.item
	}
	return nil
}

func (st *Stack) IsEmpty() bool {
	return st.depth == 0
}
