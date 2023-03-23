package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var s = NewStack()
	s.IsEmpty()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.IsEmpty()
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
