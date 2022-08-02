package main

import (
	"fmt"
	"time"
)


func listenToChan(c chan int) {
	for {
		i := <- c
		fmt.Println("Got ", i, " from channel ")
		time.Sleep(1 * time.Second)
	}
}


func main(){
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i:= 0; i<= 100; i++ {

		fmt.Println("sending ", i , " to channel ")
		ch <- i
		fmt.Println("sent ", i, " from channel ")
	}

	fmt.Println("Done!")
	close(ch)
}