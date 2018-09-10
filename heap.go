package main

import (
	"errors"
	"math"
)

type Item *Worker

type MinHeap struct {
	maxSize  int // Maximum number of elements in the heap
	size     int // Size of the heap
	elements []Item
}

func InitHeap(maxSize int) *MinHeap {
	heap := &MinHeap{
		maxSize:  maxSize,
		size:     0,
		elements: make([]Item, maxSize),
	}

	return heap
}

func Parent(i int) int {
	return int(math.Floor(math.Abs((float64)((i - 1) / 2))))
}

func LChild(i int) int {
	return 2*i + 1
}

func RChild(i int) int {
	return 2*i + 2
}

func (heap *MinHeap) Swap(x, y int) {
	heap.elements[x], heap.elements[y] = heap.elements[y], heap.elements[x]

	heap.elements[x].index = x
	heap.elements[y].index = y
}

func (heap *MinHeap) SiftUp(position int) {
	for position > 0 && heap.elements[Parent(position)].priority > heap.elements[position].priority {
		heap.Swap(position, Parent(position))
		position = Parent(position)
	}

	heap.elements[position].index = position
}

func (heap *MinHeap) SiftDown(position int) {
	minIdx := position
	lc := LChild(position)
	if lc < heap.size && heap.elements[lc].priority < heap.elements[minIdx].priority {
		minIdx = lc
	}

	rc := RChild(position)
	if rc < heap.size && heap.elements[rc].priority < heap.elements[minIdx].priority {
		minIdx = rc
	}

	if minIdx != position {
		heap.Swap(position, minIdx)
		heap.SiftDown(minIdx)
	}

	heap.elements[position].index = position
}

func (heap *MinHeap) Insert(value Item) error {
	if heap.size >= heap.maxSize {
		return errors.New("Heap full")
	}

	heap.size++
	heap.elements[heap.size-1] = value
	heap.SiftUp(heap.size - 1)
	return nil
}

func (heap *MinHeap) ExtractMin() Item {
	res := heap.elements[0]

	heap.elements[0] = heap.elements[heap.size-1]
	heap.size--
	heap.SiftDown(0)
	return res
}

func (heap *MinHeap) Remove(position int) {
	heap.elements[position].priority = math.MinInt32
	heap.SiftUp(position)
	heap.ExtractMin()
}

func (heap *MinHeap) SetPriority(position int, priority int) {
	old := heap.elements[position].priority
	heap.elements[position].priority = priority

	if old < priority {
		heap.SiftDown(position)
	} else {
		heap.SiftUp(position)
	}
}
