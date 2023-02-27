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

func New() *SingleLink {
	singleLink := new(SingleLink)
	singleLink.size = 0
	return singleLink
}

func (singleLink *SingleLink) GetSize() int {
	return singleLink.size
}

func (singleLink *SingleLink) IsEmpty() bool {
	return singleLink.GetSize() == 0
}

func (singleLink *SingleLink) Append(data interface{}) {
	if singleLink.head == nil {
		singleLink.head = &Node{data: data}
		singleLink.size++
		return
	}
	current := singleLink.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{data: data}
	singleLink.size++
}

func (singleLink *SingleLink) InsertFirst(data interface{}) {
	if singleLink.head == nil {
		singleLink.head = &Node{data: data}
		singleLink.size++
		return
	}
	p := &Node{data: data, Next: singleLink.head}
	singleLink.head = p
	singleLink.size++
}

func (singleLink *SingleLink) RemoveData(data interface{}) bool {
	if singleLink.size == 0 {
		panic("single link is Empty!")
		return false
	}
	if singleLink.head.data == data {
		singleLink.head.data = nil
		singleLink.head = singleLink.head.Next
		singleLink.size--
		return true
	}
	current := singleLink.head
	for current.Next != nil {
		if current.Next.data == data {
			current.Next = current.Next.Next
			singleLink.size--
			return true
		}
		current = current.Next
	}
	return false
}

func (singleLink *SingleLink) SearchByValue(data interface{}) bool {
	current := singleLink.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.Next
	}
	return false
}

func (singleLink *SingleLink) GetItem(pos int) interface{} {
	if pos <= 0 || pos > singleLink.GetSize() {
		panic("pos is out of range!")
	}
	index := pos
	current := singleLink.head
	for index > 1 {
		current = current.Next
		index--
	}
	return current.data
}

func (singleLink *SingleLink) Print() {
	current := singleLink.head
	if current == nil {
		panic("single link is nil!")
	}
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.Next
	}
}
