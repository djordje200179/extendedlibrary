package priorityqueue

import (
	"container/heap"
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

type Priority bool

const (
	SmallerFirst Priority = false
	BiggerFirst  Priority = true
)

type PriorityQueue[T any] struct {
	hs heapSlice[T]
}

func New[T any](priority Priority) *PriorityQueue[T] {
	return NewWithCapacity[T](priority, 0)
}

func NewWithCapacity[T any](priority Priority, initialCapacity int) *PriorityQueue[T] {
	pq := new(PriorityQueue[T])

	pq.hs.slice = make([]item[T], 0, initialCapacity)
	pq.hs.priority = priority

	return pq
}

func (pq *PriorityQueue[T]) Push(value T, priority int) {
	item := item[T]{
		value:    value,
		priority: priority,
	}

	heap.Push(&pq.hs, item)
}

func (pq *PriorityQueue[T]) Pop() T {
	item := heap.Pop(&pq.hs).(item[T])
	return item.value
}

func (pq *PriorityQueue[T]) Peek() T {
	return pq.hs.slice[0].value
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.hs.Len() == 0
}

func (pq *PriorityQueue[T]) ForEach(callback functions.ParamCallback[T]) {
	for !pq.Empty() {
		callback(pq.Pop())
	}
}
