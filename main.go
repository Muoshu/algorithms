package main

import (
	"algorithms/binaryTree"
	"fmt"
)

func compare(x, y interface{}) bool {
	return x.(int) < y.(int)
}

func main() {
	//var stack = stack.New()
	//
	//for i := 0; i < 10; i++ {
	//	stack.Push(i)
	//}
	//for i := 10; i > 0; i-- {
	//	fmt.Println(stack.Pop())
	//}

	//var que = queue.NewSliceQueue()
	//for i := 0; i < 10; i++ {
	//	que.EnQueue(i)
	//}
	//for i := 0; i < 10; i++ {
	//	item := que.DeQueue()
	//	fmt.Println(item)
	//}

	tree := binaryTree.New(compare)
	for i := 0; i < 10; i++ {
		tree.Insert(i)
	}
	findTree := tree.Search(2)
	if findTree.Node != 2 {
		fmt.Println("Err")
	}
}
