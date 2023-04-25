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
	
	time.Sleep(10 * time.Second)
}
