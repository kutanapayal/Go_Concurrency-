package main

import (
	"fmt"
	"sync"
)

type Income struct {
	source string
	amount int
}

func main() {

	var bacance sync.Mutex
	var bankBalance int
	var wg sync.WaitGroup

	incomes := []Income{
		{source: "main job", amount: 500},
		{source: "gift", amount: 50},
		{source: "part time job", amount: 100},
		{source: "investment", amount: 10},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {

		go func(i int, income Income) {

			defer wg.Done()
			for w := 0; w < 52; w++ {

				bacance.Lock()
				bankBalance += income.amount
				fmt.Printf("on week %d, you earned $%d.00 from %s\n", i, bankBalance, income.source)
				bacance.Unlock()
			}

		}(i, income)
	}

	wg.Wait()
	fmt.Printf("Total Bank Balance is : %d", bankBalance)
}
