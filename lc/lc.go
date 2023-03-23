package lc

type LinkNode struct {
	data int
	next *LinkNode
}

type MyLinkedList struct {
	head *LinkNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		&LinkNode{},
		0,
	}
}

func (linkList *MyLinkedList) Get(index int) int {
	if index < 0 || index >= linkList.size {
		return -1
	}
	current := linkList.head
	for i := 0; i <= index; i++ {
		current = current.next
	}
	return current.data
}

func (linkList *MyLinkedList) AddAtHead(val int) {
	linkList.AddAtIndex(0, val)
}

func (linkList *MyLinkedList) AddAtTail(val int) {
	linkList.AddAtIndex(linkList.size, val)
}

func (linkList *MyLinkedList) AddAtIndex(index int, val int) {
	if linkList.size < index {
		panic("index is out of range!")
		return
	}
	current := linkList.head
	max := max(index, 0)
	for i := 0; i < max; i++ {
		current = current.next
	}
	temp := &LinkNode{data: val, next: current.next}
	current.next = temp
	linkList.size++
}

func (linkList *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= linkList.size {
		//panic("index is out of range!")
		return
	}
	current := linkList.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.next = current.next.next

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//func (myLinkList *MyLinkedList) Get(index int) int {
//	if index < 0 || index > myLinkList.size {
//		panic("index out of range!")
//		return -1
//	}
//	if myLinkList.size == 0 {
//		panic("Link list is empty!")
//		return -1
//	}
//	current := myLinkList.head.next
//	for i := 0; i < index; i++ {
//		current = current.next
//	}
//	return current.data
//}
//
//func (myLinkList *MyLinkedList) AddAtHead(val int) {
//	if myLinkList.size == 0 {
//		myLinkList.head.next = &LinkNode{
//			pre:  myLinkList.head,
//			data: val,
//			next: nil,
//		}
//		myLinkList.size++
//		return
//	}
//	temp := &LinkNode{
//		pre:  myLinkList.head,
//		data: val,
//		next: myLinkList.head.next,
//	}
//	myLinkList.head.next.pre = temp
//	myLinkList.head.next = temp
//	myLinkList.size++
//}
//
//func (myLinkList *MyLinkedList) AddAtTail(val int) {
//	if myLinkList.size == 0 {
//		myLinkList.head.next = &LinkNode{
//			pre:  myLinkList.head,
//			data: val,
//			next: nil,
//		}
//		myLinkList.size++
//		return
//	}
//	current := myLinkList.head
//	for current.next != nil {
//		current = current.next
//	}
//	current.next = &LinkNode{pre: current, data: val, next: nil}
//	myLinkList.size++
//}
//
//func (myLinkList *MyLinkedList) AddAtIndex(index int, val int) {
//	if index > myLinkList.size {
//		panic("index is bigger than LinkList size!")
//	}
//	if index <= 0 {
//		myLinkList.AddAtHead(val)
//		return
//	}
//	if index == myLinkList.size {
//		myLinkList.AddAtTail(val)
//		return
//	}
//	current := myLinkList.head.next
//	for i := 0; i < index; i++ {
//		current = current.next
//	}
//	current.pre.next = &LinkNode{pre: current.pre, data: val, next: current}
//	myLinkList.size++
//}
//
//func (myLinkList *MyLinkedList) DeleteAtIndex(index int) {
//	if index < 0 || index > myLinkList.size {
//		panic("index out of range!")
//		return
//	}
//	if myLinkList.size == 0 {
//		panic("Link list is empty!")
//		return
//	}
//	if myLinkList.size == 1 {
//		myLinkList.head.next.pre = nil
//		myLinkList.head.next = nil
//		myLinkList.size--
//		return
//	}
//	current := myLinkList.head.next
//	if index == myLinkList.size {
//		for i := 0; i < index-1; i++ {
//			current = current.next
//		}
//		current.pre.next = nil
//		current.pre = nil
//		myLinkList.size--
//		return
//	}
//	if index == myLinkList.size-1 {
//		for i := 0; i < index; i++ {
//			current = current.next
//		}
//		current.pre.next = nil
//		current.pre = nil
//		myLinkList.size--
//		return
//	}
//	for i := 0; i < index; i++ {
//		current = current.next
//	}
//	current.pre.next = current.next
//	current.next.pre = current.pre
//
//	myLinkList.size--
//}
//
//func (myLinkList *MyLinkedList) Print() {
//	if myLinkList.size == 0 {
//		panic("list is empty!")
//		return
//	}
//	current := myLinkList.head.next
//	for current != nil {
//		fmt.Print(current.data, " ")
//		current = current.next
//	}
//	fmt.Println()
//}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if temp, ok := hash[target-nums[i]]; ok && i != temp {
			return []int{i, temp}
		}
	}
	return []int{0, 0}
}

func CanFormArray(arr []int, pieces [][]int) bool {
	hash := make(map[int]int, len(pieces))
	for i := 0; i < len(pieces); i++ {
		hash[pieces[i][0]] = i
	}
	i := 0
	for i < len(arr) {
		if index, ok := hash[arr[i]]; ok {
			for _, x := range pieces[index] {
				if x != arr[i] {
					return false
				}
				i++
			}
		} else {
			return false
		}
	}

	return true
}

//type ListNode struct {
//	Value int
//	Next  *ListNode
//}
//
//func addTwoNumber(l1, l2 *ListNode) (l3 *ListNode) {
//	temp := 0
//	l1Value := 0
//	l2Value := 0
//	l3Value := 0
//	var current *ListNode
//	for l1 != nil || l2 != nil {
//		if l1 != nil {
//			l1Value = l1.Value
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			l2Value = l2.Value
//			l2 = l2.Next
//		}
//		l3Value, temp = (l1Value+l2Value+temp)%10, (l1Value+l2Value+temp)/10
//		if l3 == nil {
//			l3 = &ListNode{l3Value, nil}
//			current = l3
//		} else {
//			current.Next = &ListNode{l3Value, nil}
//			current = current.Next
//		}
//	}
//	if temp == 1 {
//		current.Next = &ListNode{l3Value, nil}
//	}
//	return
//}
//
//func combine(n int, k int) (ans [][]int) {
//	var path []int
//	var dfs func(int, int, int)
//	dfs = func(n int, k int, start int) {
//		if len(path) == k {
//			temp := make([]int, len(path))
//			copy(temp, path)
//			ans = append(ans, temp)
//			//fmt.Println(len(ans), path)
//			//ans = append(ans, path)
//			return
//		}
//		for i := start; i <= n; i++ {
//			path = append(path, i)
//			dfs(n, k, i+1)
//			path = path[:len(path)-1]
//		}
//	}
//	dfs(n, k, 1) // n=4,k=2
//	return
//}
