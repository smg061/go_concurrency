package main

import (
	"errors"
	"fmt"
	"math"
)
type PQNode [T any] struct {
	Value T
	Priority int
}

func NewPQNode[T any](value T, priority int) *PQNode [T] {
 n := &PQNode[T]{
	Value:value,
	Priority:priority,
 }

 return n
}
type PriorityQueue[T any] struct {
	data []PQNode[T]
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data: make([]PQNode[T], 0),
	}
}

func (p *PriorityQueue[T]) swap(idx1, idx2 int) []PQNode[T] {
if(len(p.data)> 0){
	temp := p.data[idx1]
	p.data[idx1] = p.data[idx2]
	p.data[idx2] = temp
}
 return p.data
}
func (p *PriorityQueue[T]) bubbleUp() int {
	idx := (len(p.data) - 1)
	for idx > 0 {
		parentIdx:= int(math.Floor((float64(idx)- 1)/2))
		if(p.data[parentIdx].Priority > p.data[idx].Priority) {
			p.swap(idx, parentIdx)
		} else {
			break
		}
	}
	return 0
}

func (p *PriorityQueue[T]) Enqueue (node PQNode[T]) []PQNode[T] {
	p.data = append(p.data, node)
	p.bubbleUp()
	return p.data
}

func (p *PriorityQueue[T]) bubbleDown() {
	parentIdx := 0
	length := len(p.data)// this._data.length;
	elementPriority := p.data[0].Priority
	for {
		leftChildIdx := (2 * parentIdx) + 1;
		rightChildIdx := (2 * parentIdx) + 2;
		leftChildPriority := -1;
		rightChildPriority := -1;
		idxToSwap := -1;
		if (leftChildIdx < length) {
			leftChildPriority = p.data[leftChildIdx].Priority;
			if (leftChildPriority < elementPriority) {
				idxToSwap = leftChildIdx;
			}
		}
		if rightChildIdx < length {
			rightChildPriority = p.data[rightChildIdx].Priority;

			if rightChildPriority < elementPriority && idxToSwap == -1 || rightChildPriority < leftChildPriority && idxToSwap != -1 {
				idxToSwap = rightChildIdx;
			}

		}
		if idxToSwap == -1 {
			break
		}
		// swap with planned element
		p.swap(parentIdx, idxToSwap);
		parentIdx = idxToSwap;
	}
}

func (p *PriorityQueue[T]) Dequeue() (*PQNode[T], error) {
	var poppedNode *PQNode[T]= nil
	if(len(p.data) <= 0) {
		return nil,   errors.New("the queue is empty")
	}
	p.swap(0, len(p.data) - 1)
	poppedNode = &p.data[0]
	p.data = p.data[1:]
	if len(p.data) > 0 {
		p.bubbleDown()
	}
	return poppedNode, nil
}

func (p *PriorityQueue[T]) print() {
	for _, n := range p.data {
		fmt.Printf("Node %d with priority %d\n", n.Value, n.Priority)
	}
}

