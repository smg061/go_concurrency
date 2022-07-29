package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func concurrent_balance() {
	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex
	// print out starting values
	fmt.Printf("Initial Account Balance: %d.00\n", bankBalance)
	// define weekly revenue
	incomes := []Income {
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},

	}
	wg.Add(len(incomes))
	// loop through 52 weeks and print out how much is made; keep running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp:= bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d, you earned %d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	// print out final balance

	wg.Wait()

	fmt.Printf("Final balance is %d.00", bankBalance)
}