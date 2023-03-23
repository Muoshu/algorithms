package stack

import (
	"fmt"
	"testing"
)

func TestQueueByStack(t *testing.T) {
	var qs = NewQueueByStack()
	qs.IsEmpty()

	qs.Push(1)
	qs.Push(2)
	qs.Push(3)
	qs.Push(4)
	fmt.Println(qs.Peek())
	fmt.Println(qs.Pop())
	fmt.Println(qs.Pop())
	fmt.Println(qs.Pop())
	fmt.Println(qs.Pop())

}
