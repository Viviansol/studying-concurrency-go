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

func main() {

	//variable for weekly balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Inicial account balance: $%d.00", bankBalance)
	fmt.Println()

	//define weekly revenue
	incomes := []Income{
		{
			Source: "main job",
			Amount: 7800,
		},
		{
			Source: "second job",
			Amount: 1500,
		},
		{
			Source: "part time",
			Amount: 800,
		},
		{
			Source: "investments",
			Amount: 500,
		},
	}
	wg.Add(len(incomes))

	//loop through 52 weeks and print out how much is mada; keep running total
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("on week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	//print out final balance
	fmt.Printf("final balance: $%d.00", bankBalance)
	fmt.Println()

}
