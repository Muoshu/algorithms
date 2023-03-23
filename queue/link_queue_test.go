package queue

import (
	"testing"
)

func TestLinkQueue(t *testing.T) {
	var queue *Queue = New()

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)

	for i := 1; i < 6; i++ {
		item := queue.DeQueue()
		if item != i {
			t.Error("TestQueue failed...", i)
		}
	}

}
