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
	data chan PizaaOrder
	quit chan chan error
}

type PizaaOrder struct {
	pizaNumber int
	message    string
	success    bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizaaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received Order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		total++
		fmt.Printf("Making pizza #%d. It will take #%d seconds.....\n", pizzaNumber, delay)
		//delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order $%d is ready!", pizzaNumber)
		}

		p := PizaaOrder{
			pizaNumber: pizzaNumber,
			message:    msg,
			success:    success,
		}

		return &p

	}
	return &PizaaOrder{
		pizaNumber: pizzaNumber,
	}
}

func Pizzaria(pizzMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// run forever or until we receive a quit notification
	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		//try to make a pizza
		//decision
		if currentPizza != nil {
			i = currentPizza.pizaNumber

			select {
			// we try to make a pizza (we sent something to the data channel)
			case pizzMaker.data <- *currentPizza:

			case quitChan := <-pizzMaker.quit:
				//close channel
				close(pizzMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {

	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//print out a message

	color.Cyan("THe Pizzeria is for business!")
	color.Cyan("----------------------------")

	//create a procedure
	pizzaJob := &Producer{
		data: make(chan PizaaOrder),
		quit: make(chan chan error),
	}

	//run the procedure in the background
	go Pizzaria(pizzaJob)

	//create and run consumer
	for i := range pizzaJob.data {
		if i.pizaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is Out for delivery!", i.pizaNumber)
			} else {
				color.Red(i.message)
				color.Red("The Custmer is really mad!")
			}
		} else {
			color.Cyan("Done Making Pizzas....")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error Closing Channel!", err)
			}

		}
	}

	//print out the ending message
	color.Cyan("--------------")
	color.Cyan("Done for the day.")

	color.Cyan("we made %d pizzas, but failed to make %d, with %d attempts in total.", pizzasMade, pizzasFailed, total)

}
