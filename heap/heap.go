package heap

import (
	"fmt"
)

type Heap[S ~[]T, T comparable] struct {
	data        S
	size        int
	compareFunc func(a, b T) bool
}

func New[S ~[]T, T comparable](data S, compareFn func(a, b T) bool) *Heap[S, T] {
	h := &Heap[S, T]{
		data:        data,
		size:        len(data),
		compareFunc: compareFn,
	}
	h.build()
	return h
}

func (h *Heap[S, T]) Insert(item ...T) {
	h.data = append(h.data, item...)
	h.size = len(h.data)
	h.build()
}

func (h *Heap[S, T]) Pop() (T, error) {
	if h.size == 0 {
		var zero T
		return zero, fmt.Errorf("empty heap, cannot pop element")
	}
	item := h.data[0]
	h.data = h.data[1:]
	h.size = len(h.data)
	h.heapify(h.size, 0)
	return item, nil
}

func (h *Heap[S, T]) Display() {
	fmt.Println(h.data)
	// h.prettyPrint(0, 0)
}

func (h *Heap[S, T]) Data() S {
	return h.data
}

func (h *Heap[S, T]) Size() int {
	return h.size
}

func (h *Heap[S, T]) build() {
	for i := h.size/2 - 1; i >= 0; i-- {
		h.heapify(h.size, i)
	}
}

func (h *Heap[S, T]) heapify(n, i int) {
	best := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.compareFunc(h.data[left], h.data[best]) {
		best = left
	}

	if right < n && h.compareFunc(h.data[right], h.data[best]) {
		best = right
	}

	if best != i {
		h.data[best], h.data[i] = h.data[i], h.data[best]
		h.heapify(n, best)
	}
}

// func (h *Heap[S, T]) prettyPrint(i, level int) {
// 	if i >= h.size {
// 		return
// 	}

// 	h.prettyPrint(i*2+2, level+1)

// 	if level == 0 {
// 		fmt.Println(h.data[i])
// 	} else {
// 		for i := 1; i < level; i++ {
// 			fmt.Print("\t")
// 		}
// 		fmt.Printf("---->%v\n", h.data[i])
// 	}
// }
