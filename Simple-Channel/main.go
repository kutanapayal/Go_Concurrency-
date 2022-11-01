package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string) {

	for {

		s := <-ping

		pong <- fmt.Sprintf("%s !!!", strings.ToUpper(s))
	}

}

func main() {

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)
	fmt.Printf("Type something and Press Enter (Q for quit)")

	for {

		fmt.Println("->")
		var UserInput string

		_, _ = fmt.Scanln(&UserInput)

		if UserInput == strings.ToUpper("q") {
			break

		}

		ping <- UserInput

		//waiting for response
		response := <-pong

		fmt.Printf("response : %s\n", response)
	}

	fmt.Println("Closing The Channels.")
	close(ping)
	close(pong)

}
