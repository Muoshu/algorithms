package Link

import "fmt"

type Node struct {
	data interface{}
	Next *Node
}

type SingleLink struct {
	head *Node
	size int
}

func NewSingleLink() *SingleLink {
	singleLink := new(SingleLink)
	singleLink.size = 0
	return singleLink
}

func (sLink *SingleLink) GetSize() int {
	return sLink.size
}

func (sLink *SingleLink) IsEmpty() bool {
	return sLink.GetSize() == 0
}

func (sLink *SingleLink) Append(data interface{}) {
	current := sLink.head
	if current == nil {
		sLink.head = &Node{
			data: data,
			Next: nil,
		}
		sLink.size++
		return
	}
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{data: data}
	sLink.size++
}

func (sLink *SingleLink) InsertFirst(data interface{}) {
	current := sLink.head
	if current == nil {
		sLink.head = &Node{
			data: data,
			Next: nil,
		}
		sLink.size++
		return
	}
	p := &Node{data: data, Next: sLink.head}
	sLink.head = p
	sLink.size++
}

func (sLink *SingleLink) RemoveData(data interface{}) bool {
	if sLink.size == 0 {
		panic("single link is Empty!")
		return false
	}

	current := sLink.head
	for current.Next != nil {
		if current.Next.data == data {
			current.Next = current.Next.Next
			sLink.size--
			return true
		}
		current = current.Next
	}
	return false
}

func (sLink *SingleLink) RemoveAtBeg() interface{} {
	if sLink.size == 0 {
		return nil
	}
	val := sLink.head.data
	sLink.size--
	sLink.head = sLink.head.Next
	return val
}

func (sLink *SingleLink) SearchByValue(data interface{}) bool {
	current := sLink.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.Next
	}
	return false
}

func (sLink *SingleLink) GetItem(pos int) interface{} {
	if pos <= 0 || pos > sLink.GetSize() {
		panic("pos is out of range!")
	}
	index := pos
	current := sLink.head
	for index > 1 {
		current = current.Next
		index--
	}
	return current.data
}

func (sLink *SingleLink) ReverseLink() {
	head := sLink.head
	var pre, next *Node
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

}

func (sLink *SingleLink) Print() {
	current := sLink.head
	if current == nil {
		panic("single link is nil!")
	}
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.Next
	}
}
