package stack

type QStack struct {
	in  *Stack
	out *Stack
}

func NewQueueByStack() *QStack {
	var qs = new(QStack)
	qs.in = NewStack()
	qs.out = NewStack()
	return qs
}

func (qStack *QStack) Push(val interface{}) {
	qStack.in.Push(val)
}

func (qStack *QStack) Pop() interface{} {
	qStack.in2Out()
	return qStack.out.Pop()
}

func (qStack *QStack) Peek() interface{} {
	qStack.in2Out()
	return qStack.out.Peek()
}

func (qStack *QStack) in2Out() {
	if qStack.out.IsEmpty() {
		for !qStack.in.IsEmpty() {
			var x = qStack.in.Pop()
			qStack.out.Push(x)
		}
	}
}

func (qStack *QStack) IsEmpty() bool {
	return qStack.in.IsEmpty()
}
