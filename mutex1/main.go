package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updatedMessage(s string){
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello, world!"

	wg.Add(2)
	go updatedMessage("Hello World")
	go updatedMessage("Hello, cosmos!")
	wg.Wait()

	fmt.Println(msg)
}

// func updatedMessage(s string, m *sync.Mutex){
// 	defer wg.Done()
// 	msg = s
// }

// func main() {
// 	msg = "Hello, world!"

// 	var mutex sync.Mutex

// 	wg.Add(2)
// 	go updatedMessage("Hello World", &mutex)
// 	go updatedMessage("Hello, cosmos!", &mutex)
// 	wg.Wait()

// 	fmt.Println(msg)
// }
