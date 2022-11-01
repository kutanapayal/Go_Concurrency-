package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

// func UpdateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()

// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {

	// var mutex sync.Mutex
	msg = "Hello, World"

	wg.Add(2)
	go updateMessage("Hello, Universe!")
	go updateMessage("Hello, Cosmos!")
	wg.Wait()

	fmt.Println(msg)
}
