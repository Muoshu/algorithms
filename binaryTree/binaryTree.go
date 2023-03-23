package binaryTree

import (
	"fmt"
)

type comparable func(c1 interface{}, c2 interface{}) bool

type BinaryTree struct {
	Node    interface{}
	left    *BinaryTree
	right   *BinaryTree
	lessFun comparable
}

func NewBTree(compareFun comparable) *BinaryTree {
	return &BinaryTree{
		Node:    nil,
		lessFun: compareFun,
		left:    nil,
		right:   nil,
	}

}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
	if tree.Node == nil {
		return nil
	}
	if tree.Node == value {
		return tree
	}
	if tree.lessFun(value, tree.Node) {
		return tree.left.Search(value)
	} else {
		return tree.right.Search(value)
	}
}

func (tree *BinaryTree) Insert(value interface{}) {
	if tree.Node == nil {
		tree.Node = value
		tree.left = NewBTree(tree.lessFun)
		tree.right = NewBTree(tree.lessFun)
		return
	}
	if tree.lessFun(value, tree.Node) {
		tree.left.Insert(value)
	} else {
		tree.right.Insert(value)
	}
}

func (tree *BinaryTree) Max() interface{} {
	if tree.Node == nil || tree.right.Node == nil {
		return tree.Node
	}
	return tree.right.Max()
}

func (tree *BinaryTree) Min() interface{} {
	if tree.Node == nil || tree.left.Node == nil {
		return tree.Node
	}
	return tree.left.Min()
}

func (tree *BinaryTree) DeleteNode(val interface{}) bool {
	if tree == nil || tree.Node == nil {
		return false
	}
	p := tree
	var pp *BinaryTree = nil //tree的父节点
	//找到待删除节点
	for p != nil && p.Node != val {
		pp = p
		if p.lessFun(p.Node, val) {
			p = p.right
		} else {
			p = p.left
		}
	}
	if p == nil {
		return false
	}
	var child *BinaryTree
	if p.left.Node != nil && p.right.Node != nil {

		minp := p.right
		minpp := p
		if minp.left.Node != nil {

			minpp = minp
			minp = minp.left
		}
		p.Node = minp.Node
		p = minp
		pp = minpp
	}

	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}
	if pp == nil {
		fmt.Println(1)
		tree.Node = nil
	} else if pp.left == p {
		pp.left = child
	} else if pp.right == p {
		pp.right = child
	}
	return true
}

func (tree *BinaryTree) PreOrder() {
	if tree.Node != nil {
		tree.left.PreOrder()
		tree.right.PreOrder()
	}
}

func (tree *BinaryTree) InOrder() {
	if tree.Node != nil {
		tree.left.InOrder()
		fmt.Printf("%d\t", tree.Node)
		tree.right.InOrder()
	}
}

func (tree *BinaryTree) InOrderRecur() {
	var s = new([]BinaryTree)
	p := tree
	for p != nil || len(*s) > 0 {
		if p != nil {
			*s = append(*s, *p)
			p = p.left
		} else {
			t := (*s)[len(*s)-1]
			*s = (*s)[:len(*s)-1]
			if t.Node != nil {
				fmt.Printf("%d\t", t.Node)
			}
			p = t.right
		}
	}
}

func (tree *BinaryTree) PostOrder() {
	if tree.Node != nil {
		tree.left.PostOrder()
		tree.right.PostOrder()
		fmt.Printf("%d\t", tree.Node)
	}
}

func (tree *BinaryTree) PostOrderRecur() {
	st := new([]BinaryTree)
	var r *BinaryTree
	p := tree
	for p != nil || len(*st) > 0 {
		if p.Node != nil {
			*st = append(*st, *p)
			p = p.left
		} else {
			*p = (*st)[len(*st)-1]
			if p.right.Node != nil && p.right != r {
				p = p.right
				*st = append(*st, *p)
				p = p.left
			} else {
				cur := (*st)[len(*st)-1]
				*st = (*st)[:len(*st)-1]
				fmt.Printf("%d\t", cur.Node)
				r = p
				p = nil
			}

		}

	}
}

func (tree *BinaryTree) LevelOrder() {
	var queue *[]BinaryTree
	queue = new([]BinaryTree)

	*queue = append(*queue, *tree)
	for len(*queue) > 0 {
		cur := (*queue)[0]
		*queue = (*queue)[1:]
		if cur.Node != nil {
			fmt.Printf("%d\t", cur.Node)
		}
		if cur.left != nil {
			*queue = append(*queue, *(cur.left))
		}
		if cur.right != nil {
			*queue = append(*queue, *(cur.right))
		}
	}
}
