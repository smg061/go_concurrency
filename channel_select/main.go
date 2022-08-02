package main

import (
	"fmt"
	"time"
)


func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from sever 2"
	}
}


func main() {

	fmt.Println("Select with channels")

	fmt.Println("--------------------------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)
	for {
		select {

		case s1 := <-channel1:
			fmt.Printf("Case one: %s \n", s1)
		case s2 := <- channel1:
			fmt.Printf("Case two: %s \n", s2)
		case s3 := <- channel2:
			fmt.Printf("Case three: %s \n", s3)
		case s4 := <- channel2:
			fmt.Printf("Case four: %s \n", s4)
		
		default:
			// avoiding deadlock
			return
		}

	}


}