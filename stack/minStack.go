package stack

type comparable func(c1 interface{}, c2 interface{}) bool

type MinStack struct {
	st, min *Stack
	lessFun comparable
}

func NewMinStack(compareFun comparable) *MinStack {
	var mStack = new(MinStack)
	mStack.st = new(Stack)
	mStack.min = new(Stack)
	mStack.lessFun = compareFun
	return mStack
}

func (mStack *MinStack) Peek() interface{} {
	return mStack.st.Peek()
}

func (mStack *MinStack) Pop() interface{} {
	var x = mStack.st.Pop()
	if mStack.min.IsEmpty() {
		mStack.min.Push(x)
	} else if x == mStack.min.Peek() {
		mStack.min.Pop()
	}
	return x
}

// Push Todo
func (mStack *MinStack) Push(val interface{}) {
	if mStack.min.Peek() == nil {
		mStack.min.Push(val)
	} else if mStack.lessFun(val, mStack.min.Peek()) {
		mStack.min.Push(val)
	}
	mStack.st.Push(val)
}

func (mStack *MinStack) GetMin() interface{} {
	return mStack.min.Peek()
}
