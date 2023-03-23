package main

import (
	lc "algorithms/Lc"
	"fmt"
	"sort"
)

func compare(x, y interface{}) bool {
	return x.(int) < y.(int)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseRecur(head *ListNode, pre *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return ReverseRecur(next, head)
}

func Reverse(listNode *ListNode) *ListNode {
	var pre, next *ListNode
	cur := listNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func mergerTwoLink(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergerTwoLink(l1, l2.Next)
		return l2
	}
	l1.Next = mergerTwoLink(l1.Next, l2)
	return l1
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1, cur2 := l1, l2
	var v, t int
	for cur1 != nil && cur2 != nil {
		v = cur1.Val + cur2.Val + t
		if v > 9 {
			v = v - 10
			t = 1
		} else {
			t = 0
		}
		cur1.Val = v
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	for cur1 != nil && t == 1 {
		v = cur1.Val + t
		if v > 9 {
			v = v - 10
			t = 1
		} else {
			t = 0
		}
		cur1.Val = v
		cur1 = cur1.Next

	}
	if cur2 != nil {
		cur1 = l1
		for cur1.Next != nil {
			cur1 = cur1.Next
		}
		cur1.Next = cur2
		cur1 = cur1.Next
		for t == 1 && cur1 != nil {
			v = cur1.Val + t
			if v > 9 {
				v = v - 10
				t = 1
			} else {
				t = 0
			}
			cur1.Val = v
			cur1 = cur1.Next
		}
	}
	if t == 1 {
		x := l1
		for x.Next != nil {
			x = x.Next
		}
		x.Next = &ListNode{1, nil}
	}

	return l1
}

func inorderPrint(root *lc.TreeNode) {
	if root == nil {
		return
	}
	inorderPrint(root.Left)
	fmt.Printf("%d ", root.Val)
	inorderPrint(root.Right)
}

func (root *ListNode) Print() {
	cur := root
	for cur != nil {
		fmt.Printf("%d\t", cur.Val)
		cur = cur.Next
	}
}
func main() {
	root := new(lc.TreeNode)
	l1 := new(lc.TreeNode)
	l2 := new(lc.TreeNode)
	l3 := new(lc.TreeNode)
	l4 := new(lc.TreeNode)

	root.Val = 3
	root.Left = l1
	root.Right = l2

	l1.Val = 1
	l1.Left = nil
	l1.Right = nil

	l2.Val = 4
	l2.Left = l3
	l2.Right = nil

	l3.Val = 2
	l3.Left = nil
	l3.Right = nil

	l4.Val = 7
	l4.Left = nil
	l4.Right = nil

	lc.RecoverTree(root)
	inorderPrint(root)

}

func ValidAnagram(s1, s2 string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m1[s1[i]] += 1
	}
	for i := 0; i < len(s2); i++ {
		m2[s2[i]] += 1
	}

	for i := 0; i < len(s1); i++ {
		if m1[s1[i]] != m2[s1[i]] {
			return false
		}
	}
	for i := 0; i < len(s2); i++ {
		if m1[s2[i]] != m2[s2[i]] {
			return false
		}
	}

	return true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func exGCD(a, b int, x, y *int) int {
	if b == 0 {
		*x, *y = 1, 0
		return a
	}
	r := exGCD(b, a%b, x, y)
	temp := y
	*y = *x - (a/b)*(*y)
	x = temp
	return r
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	var h int
	for temp > 0 {
		h = h*10 + temp%10
		temp = temp / 10
	}
	return x == h
}

func BFSTravel(G [][]int) {
	n := len(G)
	visited := make([]bool, n)
	var queue []int
	for i := 0; i < n; i++ {
		if !visited[i] {
			BFS(G, i, &queue, visited)
		}
	}
}

func BFS(G [][]int, v int, queue *[]int, visited []bool) {
	visited[v] = true
	*queue = append(*queue, v)
	for len(*queue) != 0 {
		i := (*queue)[0]
		*queue = (*queue)[1:]
		for j := 0; j < len(G); j++ {
			if G[i][j] == 1 && !visited[j] {
				visited[j] = true
				*queue = append(*queue, j)
			}
		}
	}
}

func DFSTraverse(G [][]int) {
	n := len(G)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			DFS(G, i, visited)
		}
	}

}

func DFS(G [][]int, v int, visited []bool) {
	visited[v] = true
	fmt.Println(v)
	for i := 0; i < len(G); i++ {
		if G[v][i] == 1 && !visited[i] {
			DFS(G, i, visited)
		}
	}

}

func maxProduct(words []string) int {
	hash := make(map[int]int)
	var ans int
	for _, word := range words {
		mask, size := 0, len(word)
		for _, c := range word {
			mask |= 1 << (c - 'a')
		}
		hash[mask] = max(hash[mask], size)
		for hMask, hLen := range hash {
			if (mask & hMask) != 0 {
				ans = max(ans, size*hLen)
			}
		}
	}
	return ans
}

func power4(x int) bool {
	return x > 0 && (x&(x-1)) == 0 && (x&1431655765) != 0
}

func countBits(n int) []int {
	var ans []int
	for i := 0; i < n+1; i++ {
		count := 0
		temp := i & 126975
		for temp != 0 {
			count += temp & 1
			temp >>= 1
		}
		ans = append(ans, count)
	}
	return ans
}

func hammingDistance(x int, y int) int {
	temp, count := x^y, 0
	for temp != 0 {
		count += temp & 1
		fmt.Println(temp & 1)
		temp >>= 1
	}
	return count
}
func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans <<= 1
		ans += num & 1
		num >>= 1
	}
	return ans
}
func candy(ratings []int) int {
	var candy int
	pre := 1
	sort.Ints(ratings)
	fmt.Println(ratings)
	size := len(ratings)
	for i := 0; i < size; i++ {
		if i == 0 {
			candy = candy + 1
			continue
		}
		if size > 1 && i == size-1 {
			if ratings[size-2] == ratings[size-1] {
				candy = candy + 1
			} else {
				candy = candy + pre + 1
			}
			continue
		}

		if ratings[i] > ratings[i-1] {
			candy = candy + pre + 1
			pre = pre + 1
		} else {
			candy = candy + 1
			pre = 1
		}
	}
	return candy
}

func swap(a []int) []int {
	a[0], a[1] = a[1], a[0]
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
