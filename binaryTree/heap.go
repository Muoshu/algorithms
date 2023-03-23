package binaryTree

import "algorithms/constraints"

type Heap[T any] struct {
	heaps   []T
	lessFun func(a, b T) bool
}

func NewHeap[T constraints.Ordered]() *Heap[T] {
	less := func(a, b T) bool {
		return a < b
	}
	return &Heap[T]{lessFun: less}
}

// Push todo
func (h *Heap[T]) Push(val T) {
	h.heaps = append(h.heaps, val)
	h.up(len(h.heaps) - 1)
}
func (h *Heap[T]) swap(i, j int) {
	h.heaps[i], h.heaps[j] = h.heaps[j], h.heaps[i]
}

// Pop todo
func (h *Heap[T]) Pop() (t T) {
	if len(h.heaps) == 0 {
		return
	}
	t = h.heaps[0]
	if len(h.heaps) == 1 {
		h.heaps = nil
		return
	}

	h.swap(0, h.Size()-1)
	h.heaps = h.heaps[:len(h.heaps)-1]
	h.down(0)
	return
}

func (h *Heap[T]) Peek() (t T) {
	if h.Size() == 0 {
		return
	}
	t = h.heaps[0]
	return
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.heaps) == 0
}

func (h *Heap[T]) Size() int {
	return len(h.heaps)
}

func (h *Heap[T]) up(child int) {
	if child <= 0 {
		return
	}
	parent := (child - 1) >> 1
	if !h.lessFun(h.heaps[child], h.heaps[parent]) {
		return
	}
	h.swap(child, parent)
	h.up(parent)
}

func (h *Heap[T]) down(parent int) {
	idx := parent
	lChild, rChild := (parent<<1)+1, (parent<<1)+2
	if lChild < len(h.heaps) && h.lessFun(h.heaps[lChild], h.heaps[idx]) {
		idx = lChild
	}
	if rChild < len(h.heaps) && h.lessFun(h.heaps[rChild], h.heaps[idx]) {
		idx = rChild
	}
	if idx == parent {
		return
	}
	h.swap(idx, parent)
	h.down(idx)
}

func (h *Heap[T]) BuildMaxHeap(a []T) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		h.HeapAdjust(a, i, len(a))
	}
}

func (h *Heap[T]) HeapAdjust(a []T, k, length int) {
	l := k*2 + 1
	r := l + 1
	largest := k
	if l < length && h.lessFun(a[largest], a[l]) {
		largest = l
	}
	if r < length && h.lessFun(a[largest], a[r]) {
		largest = r

	}
	if k != largest {
		a[k], a[largest] = a[largest], a[k]
		h.HeapAdjust(a, largest, length)
	}

}

func (h *Heap[T]) HeapSort(x []T) {
	h.BuildMaxHeap(x)
	for i := len(x) - 1; i > 0; i-- {
		x[0], x[i] = x[i], x[0]
		h.HeapAdjust(x, 0, i)
	}
}
