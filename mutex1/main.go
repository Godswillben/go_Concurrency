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

func main(){
	// variable for banking
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Intial bank balance: $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Main Income", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made
	for i, income := range incomes{
		go func(i int, income Income){
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				fmt.Println("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
				balance.Unlock()
			}
		}(i, income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()
}
