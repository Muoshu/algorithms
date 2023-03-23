package binaryTree

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	var h = NewHeap[int]()

	x := []int{9, 45, 78, 32, 17, 65, 53, 87}
	fmt.Println(x)
	h.BuildMaxHeap(x)
	fmt.Println(x)
}
