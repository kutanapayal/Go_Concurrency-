package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

//constants
const hunger = 3

//variables
var philosophers = []string{"Locke", "Socrates", "Aristoles", "Pascal", "Plato"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second
var orderFinished []string
var orderFinishedMutex sync.Mutex

func dinningProblem(philosopher string, forkLeft, forkRight *sync.Mutex) {

	defer wg.Done()

	//print a message
	fmt.Println(philosopher, "is seated.")
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Println(philosopher, "is hungry.")
		time.Sleep(sleepTime)

		//Lock a froks
		forkLeft.Lock()
		fmt.Printf("\t%s picked up the fork to his left.\n", philosopher)
		forkRight.Lock()
		fmt.Printf("\t%s picked up the fork to his right.\n", philosopher)

		//print a message
		fmt.Println(philosopher, "has both forks, and is eating.")
		time.Sleep(eatTime)

		//give some time to philosopher for think
		fmt.Println(philosopher, "is thinking.")
		time.Sleep(thinkTime)

		//unlock the mutexes
		forkLeft.Unlock()
		fmt.Printf("\t%s put the left fork on table.\n", philosopher)
		forkRight.Unlock()
		fmt.Printf("\t%s put the right fork on table.\n", philosopher)

		time.Sleep(sleepTime)
	}

	fmt.Println(philosopher, "is satisfied.")
	time.Sleep(sleepTime)

	fmt.Println(philosopher, "has left the table. ")

	orderFinishedMutex.Lock()
	orderFinished = append(orderFinished, philosopher)
	orderFinishedMutex.Unlock()
}

func main() {

	//print intro
	fmt.Println("The Dinning Philosopher Problem")
	fmt.Println("------------------------------")

	//add 5 (number of philosopher) to wait for
	wg.Add(len(philosophers))

	//we need to create a mutex for the very first fork( the one to
	//the left of the first philosopher). we create it as a pointer,
	//since a sync.Mutex must not be copied after its initial use.
	Leftfork := &sync.Mutex{}

	//spawn one go routine for each philosopher
	for i := 0; i < len(philosophers); i++ {

		Rightfork := &sync.Mutex{}
		go dinningProblem(philosophers[i], Leftfork, Rightfork)
		Leftfork = Rightfork
	}

	wg.Wait()

	fmt.Println("The table is empty.")
	fmt.Println("--------------------")
	fmt.Printf("Order Finished : %s\n", strings.Join(orderFinished, ","))
}
