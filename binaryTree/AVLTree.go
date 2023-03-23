package binaryTree

import "fmt"

type AVLNode struct {
	data, height int
	left, right  *AVLNode
}

type AVLTree struct {
	root *AVLNode
}

func NewAVLTree(val int) *AVLTree {
	return &AVLTree{
		root: &AVLNode{
			data:   val,
			height: 1,
			left:   nil,
			right:  nil,
		},
	}
}

func (tree *AVLTree) Find(val int) *AVLNode {
	if tree.root == nil {
		return nil
	}
	return tree.root.find(val)
}
func (node *AVLNode) find(val int) *AVLNode {
	if node == nil {
		return nil
	}
	if node.data == val {
		return node
	} else if node.data < val {
		return node.left.find(val)
	} else {
		return node.right.find(val)
	}
}

func (tree *AVLTree) Insert(val int) {
	tree.root = tree.root.insert(val)
}

func (node *AVLNode) insert(val int) *AVLNode {
	if node == nil {
		return &AVLNode{data: val, height: 1}
	}
	if node.data == val {
		return node
	}
	//辅助变量，用于存储旋转后子树根节点
	var newTreeNode *AVLNode
	if val > node.data {
		node.right = node.right.insert(val)
		bf := node.BalanceFactor()
		if bf == -2 {
			// 在右子树中插入右子节点导致失衡，左旋
			if val > node.right.data {
				newTreeNode = LeftRotate(node)
			} else {
				// 在右子树左子节点插入导致失衡，右左旋转
				newTreeNode = RightLeftRotate(node)
			}
		}
	} else {
		node.left = node.left.insert(val)
		bf := node.BalanceFactor()
		if bf == 2 {
			if val < node.left.data {
				newTreeNode = RightRotate(node)
			} else {
				newTreeNode = LeftRightRotate(node)
			}
		}
	}
	if newTreeNode == nil {
		// 根节点没变，直接更新子树高度
		node.UpdateHeight()
		return node
	} else {
		newTreeNode.UpdateHeight()
		return newTreeNode
	}
}

// BalanceFactor 计算左右子树高度差
func (node *AVLNode) BalanceFactor() int {
	leftHeight, rightHeight := 0, 0
	if node.left != nil {
		leftHeight = node.left.height
	}
	if node.right != nil {
		rightHeight = node.right.height
	}
	return leftHeight - rightHeight
}

// UpdateHeight 更新节点高度
func (node *AVLNode) UpdateHeight() {
	if node == nil {
		return
	}
	//分别计算左右子树高度
	leftHeight, rightHeight := 0, 0
	if node.left != nil {
		leftHeight = node.left.height
	}
	if node.right != nil {
		rightHeight = node.right.height
	}

	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	node.height = maxHeight + 1

}

// RightRotate 右旋
func RightRotate(node *AVLNode) *AVLNode {
	pivot := node.left
	pivotR := pivot.right
	pivot.right = node
	node.left = pivotR

	node.UpdateHeight()
	pivot.UpdateHeight()
	return pivot
}

// LeftRotate 左旋
func LeftRotate(node *AVLNode) *AVLNode {
	pivot := node.right  //新插入的节点
	pivotL := pivot.left //暂存pivot左子树入口节点
	pivot.left = node    //左旋后最小不平衡子树根节点 node 变成 pivot 的左子节点
	node.right = pivotL  //而pivot原本的左子树节点要挂载到node节点的右子树上

	node.UpdateHeight()
	pivot.UpdateHeight()
	return pivot
}

func LeftRightRotate(node *AVLNode) *AVLNode {
	node.left = LeftRotate(node.left)
	return RightRotate(node)
}

func RightLeftRotate(node *AVLNode) *AVLNode {
	node.right = RightRotate(node.right)
	return LeftRotate(node)
}

func (tree *AVLTree) Traverse() {
	tree.root.traverse()
}
func (node *AVLNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	fmt.Printf("%d(%d) ", node.data, node.height)
	node.right.traverse()
}
