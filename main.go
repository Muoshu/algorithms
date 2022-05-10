package main

import (
	"algorithms/queue"
	"fmt"
)

func main() {
	//var stack = stack.New()
	//
	//for i := 0; i < 10; i++ {
	//	stack.Push(i)
	//}
	//for i := 10; i > 0; i-- {
	//	fmt.Println(stack.Pop())
	//}

	var que = queue.NewSliceQueue()
	for i := 0; i < 10; i++ {
		que.EnQueue(i)
	}
	for i := 0; i < 10; i++ {
		item := que.DeQueue()
		fmt.Println(item)
	}
}
