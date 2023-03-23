package stack

import (
	"fmt"
	"testing"
)

func compare(x, y interface{}) bool {
	return x.(int) < y.(int)
}
func TestMinStack(t *testing.T) {
	var mStack = NewMinStack(compare)
	mStack.Push(1)
	mStack.Push(2)
	mStack.Push(3)
	mStack.Push(-1)

	fmt.Println(mStack.GetMin())
	mStack.Pop()
	mStack.Pop()
	fmt.Println(mStack.GetMin())
}
