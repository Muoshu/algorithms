package binaryTree

import (
	"fmt"
	"testing"
)

func compare(x, y interface{}) bool {
	return x.(int) < y.(int)
}

func TestBinaryTree(t *testing.T) {
	var tree = NewBTree(compare)
	tree.Insert(15)
	tree.Insert(13)
	tree.Insert(20)
	tree.Insert(8)
	tree.Insert(14)
	tree.Insert(16)
	tree.Insert(26)

	tree.InOrder()
	tree.DeleteNode(26)
	fmt.Println()
	tree.InOrder()

}
