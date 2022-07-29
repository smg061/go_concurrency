package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/fatih/color"
)
const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}
type PizzaOrder struct {
	pizzaNumber int
	message string
	success bool
}
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber > NumberOfPizzas {
		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
		}
	}
	delay := rand.Intn(5) + 1
	fmt.Printf("Received order number %d\n", pizzaNumber)
	rnd := rand.Intn(12) + 1
	msg := ""
	success := rnd > 5
	if success {
		pizzasMade++
	} else {
		pizzasFailed++
	}
	total++

	fmt.Printf("Making pizza order %d. It will take %d seconds\n", pizzaNumber, delay)
	time.Sleep(time.Duration(delay) * time.Second)

	if rnd <=2 {
		msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d\n", pizzaNumber)
	} else if rnd <= 4 {
		msg = fmt.Sprintf("*** The cook quit while making pizza #%d\n", pizzaNumber)
	} else {
		msg = fmt.Sprintf("Pizza order #%d is ready\n", pizzaNumber)
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
		message: msg,
		success: success,
	}
	

}
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	 i := 0
	// run forever or until we receive a quit notification
	// try to make pizza order
	for {
		// try to make a pizza
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we sent something to the data channel
			case pizzaMaker.data <- *currentPizza:
			case quitChan:= <- pizzaMaker.quit:
				// close channel
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <- ch
}

type Consumer struct {

}
func main () {

	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	// print out message

	color.Cyan("The pizzeria is open for business")
	color.Cyan("--------------------------------")
	// run the producer in the background
	pizzaJob := &Producer {
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)

	// create and run consumer

	for i:= range pizzaJob.data {

		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order %d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad")
			}
		} else {
			color.Cyan("Done making pizzas!")
			err := pizzaJob.Close()

			if err != nil {
				color.Red("Error closing channel! %s", err.Error)
			}
		}
	}
	// print out the ending message
	color.Cyan("--------------------------------------------------------")
	color.Cyan("Done for the day.")
	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day...")
	default:
		color.Green("It was an excellent day")
	}
}