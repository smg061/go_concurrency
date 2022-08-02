package main

import (
	"fmt"
	"strings"
)

func shout(ping, pong chan string) {

	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}


func main() {
	ping := make(chan string)
	pong := make(chan string)


	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")
	for {
		fmt.Print("-> ")
		var userInput string
		_, _ = fmt.Scanf("%s", &userInput)
		if strings.ToLower(userInput) == "q"{
			break
		}
		ping <- userInput
		// wait for a response
		response := <- pong

		fmt.Println("Response: ", response)

	}
	fmt.Println("All done; closing channels...")
	close(ping)
	close(pong)
}