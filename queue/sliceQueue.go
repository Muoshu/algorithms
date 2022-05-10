package queue

import "fmt"

const MaxSize int = 100

type SliceQueue struct {
	front, rear int
	data        [MaxSize]interface{}
}

func NewSliceQueue() *SliceQueue {
	var sliceQueue *SliceQueue = new(SliceQueue)
	sliceQueue.front = 0
	sliceQueue.rear = 0
	return sliceQueue
}

func (sliceQueue *SliceQueue) Empty() bool {
	if sliceQueue.rear == sliceQueue.front {
		return true
	}
	return false
}

func (sliceQueue *SliceQueue) EnQueue(item interface{}) {
	if (sliceQueue.rear+1)%MaxSize == sliceQueue.front {
		fmt.Println("队列满")
		return
	}
	sliceQueue.data[sliceQueue.rear] = item
	sliceQueue.rear = (sliceQueue.rear + 1) % MaxSize
}

func (sliceQueue *SliceQueue) DeQueue() interface{} {
	if sliceQueue.rear == sliceQueue.front {
		fmt.Println("队列空")
		return nil
	}
	var x = sliceQueue.data[sliceQueue.front]
	sliceQueue.front = (sliceQueue.front + 1) % MaxSize
	return x
}
