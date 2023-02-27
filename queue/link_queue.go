package queue

type Item struct {
	val  interface{}
	next *Item
}

type Queue struct {
	current *Item
	last    *Item
	depth   int64
}

func New() *Queue {
	var queue = new(Queue)
	queue.depth = 0
	return queue
}
func (queue *Queue) EnQueue(item interface{}) {
	if queue.depth == 0 {
		queue.current = &Item{val: item, next: nil}
		queue.last = queue.current
		queue.depth++
		return
	}
	q := &Item{val: item, next: nil}
	queue.last.next = q
	queue.last = q
	queue.depth++
}

func (queue *Queue) DeQueue() interface{} {
	if queue.depth == 0 {
		return nil
	}
	item := queue.current
	queue.current = queue.current.next
	queue.depth--
	return item.val
}
