package queue

const MAXSIZE int = 100

type DoubleEndQueue struct {
	front, tail int
	dequeue     [MAXSIZE]interface{}
}

func NewDeQueue() *DoubleEndQueue {
	var deq = new(DoubleEndQueue)
	deq.front = 0
	deq.tail = 0
	return deq
}

func (deq *DoubleEndQueue) IsEmpty() bool {
	return deq.front == deq.tail
}

func (deq *DoubleEndQueue) EnHead(val interface{}) {
	if (deq.front+1)%MAXSIZE == deq.tail {
		return
	}
	deq.dequeue[deq.front] = val
	deq.front = (deq.front + 1) % MAXSIZE
}

func (deq *DoubleEndQueue) DeHead() interface{} {
	if deq.front == deq.tail {
		return nil
	}
	val := deq.dequeue[deq.front]
	deq.front = (deq.front - 1 + MAXSIZE) % MAXSIZE
	return val
}

func (deq *DoubleEndQueue) DeTail() interface{} {
	if deq.front == deq.tail {
		return nil
	}
	val := deq.dequeue[deq.tail]
	deq.tail = (deq.tail + 1) % MAXSIZE
	return val
}

func (deq *DoubleEndQueue) Entail(val interface{}) {
	if deq.front == deq.tail {
		return
	}
	deq.dequeue[deq.tail] = val
	deq.tail = (deq.tail - 1 + MAXSIZE) % MAXSIZE

}
