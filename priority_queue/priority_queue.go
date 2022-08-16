package main

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

func (p *PriorityQueue[T]) swap(idx1, idx2 int) []PQNode[T]{
 temp := p.data[idx1]
 p.data[idx1] = p.data[idx2]
 p.data[idx2] = temp
 return p.data
}

func NewPriorityQueue[T any](data []PQNode[T]) *PriorityQueue[T] {
	return nil
}