package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("First to be printed")
	time.Sleep(1 * time.Second)

	printSomething("Second to be printed")
}
