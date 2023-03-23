package Link

import "fmt"

type node struct {
	Pre  *node
	data interface{}
	Next *node
}

type DoubleLink struct {
	head *node
	size int
}

func NewDoubleLink() *DoubleLink {
	dLink := new(DoubleLink)
	dLink.size = 0
	return dLink
}

func (dLink *DoubleLink) GetSize() int {
	return dLink.size
}

func (dLink *DoubleLink) Append(data interface{}) {
	if dLink.size == 0 {
		dLink.head = &node{data: data, Next: nil, Pre: nil}
		dLink.size++
		return
	}
	current := dLink.head
	for current.Next != nil {
		current = current.Next
	}
	temp := &node{
		Pre:  current,
		data: data,
		Next: nil,
	}
	current.Next = temp
	dLink.size++
}

func (dLink *DoubleLink) Insert(data interface{}, index int) {
	if index < 0 || index > dLink.GetSize() {
		panic("index out of range!")
	}
	if dLink.GetSize() == 0 { //insert after double link initial
		temp := &node{data: data}
		dLink.head = temp
		dLink.size++
		return
	}
	if index == 1 { //insert to head,so head pointer need move
		temp := &node{Pre: nil, data: data, Next: dLink.head}
		dLink.head = temp
		dLink.size++
		return
	}
	current := dLink.head
	for i := 1; i < index; i++ {
		current = current.Next
	}
	temp := &node{Pre: current.Pre, data: data, Next: current}
	current = temp
	dLink.size++
}

func (dLink *DoubleLink) SearchByValue(data interface{}) bool {
	current := dLink.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.Next
	}
	return false
}

func (dLink *DoubleLink) RemoveData(data interface{}) bool {
	current := dLink.head
	for current != nil {
		if current.data == data {
			current.Pre.Next = current.Next
			dLink.size--
			return true
		}
		current = current.Next
	}
	return false
}

func (dLink *DoubleLink) GetItem(index int) interface{} {
	if index < 0 || index > dLink.size {
		panic("index out of range!")
	}
	current := dLink.head
	for i := 1; i < index; i++ {
		current = current.Next
	}
	return current.data
}

func (dLink *DoubleLink) Print() {
	current := dLink.head
	if current == nil {
		panic("single link is nil!")
	}
	for current != nil {
		fmt.Println(current.data, " ")
		current = current.Next
	}

}
