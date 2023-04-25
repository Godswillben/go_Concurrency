package main

import (
	"fmt"
	"strings"
	"time"
)

func shout(ping chan string, pong chan string){
	for {
		s := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main(){
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)
	
	fmt.Println("Type something and press Enter (Enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q"){ // code to be changed
			break
		}

		ping <- userInput
		// wait for a response
		response := <- pong
		fmt.Println("Response: ", response)
	}

	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}
