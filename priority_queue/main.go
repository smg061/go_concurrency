package main

import (
	"time"
	 "fmt"
)
const sleepyTime = 1
func main() {
	queue := NewPriorityQueue[int]()
	queue.Enqueue(*NewPQNode(1,0))
	queue.print()
	fmt.Println("----------------------------------------------------------------")
	time.Sleep(sleepyTime * time.Second)
	queue.Enqueue(*NewPQNode(2,1))
	queue.print()
	fmt.Println("----------------------------------------------------------------")

	time.Sleep(sleepyTime * time.Second)

	queue.Enqueue(*NewPQNode(3,2))
	queue.print()
	fmt.Println("----------------------------------------------------------------")

	time.Sleep(sleepyTime * time.Second)

	queue.Dequeue()
	fmt.Println("----------------------------------------------------------------")

	queue.print()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.print()
}