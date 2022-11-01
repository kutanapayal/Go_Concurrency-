package main

import (
	"fmt"
	"sync"
)

func PrintSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() //this will decrease the wg count by 1

	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup //waitGroup use to insure that the all goroutine must execute

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"lamda",
		"pi",
		"delta",
	}

	wg.Add(6) //this will six as counter

	for i, x := range words {
		go PrintSomething(fmt.Sprintf("%d %v", i, x), &wg)
	}
	wg.Wait() //it will wait till counter get to 0

	wg.Add(7) //add seven counter for next goroutine

	for i, x := range words {
		go PrintSomething(fmt.Sprintf("%d %v", i, x), &wg)
	}

	PrintSomething("THis Has to be printed Second!", &wg)

	wg.Wait() //it will wait till counter get to 0
}
