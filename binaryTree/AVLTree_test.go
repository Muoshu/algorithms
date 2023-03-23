package binaryTree

import "testing"

func TestAVLTree(t *testing.T) {
	avlTree := NewAVLTree(3)
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(4)
	avlTree.Insert(5)
	avlTree.Insert(6)
	avlTree.Insert(7)
	avlTree.Insert(10)
	avlTree.Insert(9)
	avlTree.Insert(8)

	avlTree.Traverse()
}
