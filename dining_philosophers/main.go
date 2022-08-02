package main

/*
	Five philosophers live in a house together;  they always dine together sitting in the same table
	They always eat a special type of spaghetti which requires two forks
	There are two forks next to each plate, which means that no two neighbors can be eating
	at the same time
*/

// variables - philosophers
// track who is eating and not
import (
	"fmt"
	"sync"
	"time"
)

// constants -

const hunger = 3

var wg sync.WaitGroup

var philosophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}

var sleepTime = 1 * time.Second

func diningProblem(philosopher string, forkLeft, forkRight *sync.Mutex) {
	defer wg.Done()

	// print a message
	fmt.Printf("%s is seated and wants to eat\n", philosopher)
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		// lock both forks
		fmt.Printf("%s is hungry\n", philosopher)
		forkLeft.Lock()
		fmt.Printf("\t%s picked up the fork to their left\n", philosopher)
		forkRight.Lock()
		fmt.Printf("\t%s picked up the fork to their right\n", philosopher)
		fmt.Println(philosopher, " has both forks and is eating...\n")
		time.Sleep(sleepTime)
		// unlock the mutexes
		forkLeft.Unlock()
		fmt.Printf("\t%s put down  the fork to their left\n", philosopher)
		fmt.Printf("\t%s put down  the fork to their right\n", philosopher)
		forkRight.Unlock()
	}

	fmt.Println(philosopher, " is satisfied")
	time.Sleep(sleepTime)
	fmt.Printf("%s has left the table \n", philosopher)
}
func main() {
	// print intro
	fmt.Println("The Dining philosopher's problem")
	fmt.Println("---------------------------------")
	wg.Add(len(philosophers))
	// I seriously do not understand this part 
	forkLeft := &sync.Mutex{}
	// spawn one go routine for each philosophers
	for i := 0; i < len(philosophers); i++ {
		// create mutex for right fork
		forkRight := &sync.Mutex{}
		go diningProblem(philosophers[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	wg.Wait()

	fmt.Println("The table is empty")
}
