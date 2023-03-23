package queue

import "testing"

func TestArrayQueue(t *testing.T) {
	var aQueue = NewArrayQueue()
	if aQueue.IsEmpty() != false {
		t.Error("IsEmpty failed")
	}
	aQueue.EnQueue(1)
	aQueue.EnQueue(2)
	aQueue.EnQueue(3)
	aQueue.EnQueue(4)
	aQueue.EnQueue(5)

	for i := 1; i < 6; i++ {
		if aQueue.DeQueue() != i {
			t.Error("TestQueue failed", i)
		}
	}
}
