package main

import (
	"algorithms/sort"
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

	var x = []int{1, 4, 3, 2, 9, 8, 6, 5, 7}
	x = sort.MergeSort(x)
	fmt.Println(x)

}
