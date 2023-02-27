package queue

const MaxSize int = 100

type ArrayQueue struct {
	front, rear int
	data        [MaxSize]interface{}
}

func NewArrayQueue() *ArrayQueue {
	var arrayQueue = new(ArrayQueue)
	arrayQueue.front = 0
	arrayQueue.rear = 0
	return arrayQueue
}

func (arrayQueue *ArrayQueue) IsEmpty() bool {
	if arrayQueue.front == arrayQueue.rear {
		return true
	}
	return false
}

func (arrayQueue *ArrayQueue) EnQueue(item interface{}) {
	if (arrayQueue.rear+1)%MaxSize == arrayQueue.front {
		return
	}
	arrayQueue.data[arrayQueue.rear] = item
	arrayQueue.rear = (arrayQueue.rear + 1) % MaxSize
}

func (arrayQueue *ArrayQueue) DeQueue() interface{} {
	if arrayQueue.front == arrayQueue.rear {
		return nil
	}
	var x = arrayQueue.data[arrayQueue.front]
	arrayQueue.front = (arrayQueue.front + 1) % MaxSize
	return x
}
