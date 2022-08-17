package main

import (
	"time"
	 "fmt"
)
const sleepyTime = 1
func main() {
	queue := NewPriorityQueue[int]()
	queue.Enqueue(*NewPQNode(1,40))
	queue.print()
	fmt.Println("----------------------------------------------------------------")
	time.Sleep(sleepyTime * time.Second)
	queue.Enqueue(*NewPQNode(2,50))
	queue.print()
	fmt.Println("----------------------------------------------------------------")

	time.Sleep(sleepyTime * time.Second)

	queue.Enqueue(*NewPQNode(3,1))
	queue.print()
	fmt.Println("----------------------------------------------------------------")

	time.Sleep(sleepyTime * time.Second)

	queue.Dequeue()
	fmt.Println("----------------------------------------------------------------")
	queue.print()
	time.Sleep(sleepyTime * time.Second)

	fmt.Println("----------------------------------------------------------------")

	queue.Dequeue()
	queue.print()
	time.Sleep(sleepyTime * time.Second)

	fmt.Println("----------------------------------------------------------------")

	queue.Dequeue()
	queue.print()
	time.Sleep(sleepyTime * time.Second)

	fmt.Println("----------------------------------------------------------------")

}