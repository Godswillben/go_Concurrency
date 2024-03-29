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
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	PizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <- ch
}

func makePizza(pizzaNumber int) *PizzaOrder{
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		}else {
			pizzasMade++
		}
		total++
	
		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		}else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quite while making pizza #%d!", pizzaNumber)
		}else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}
		p := PizzaOrder{
			PizzaNumber: pizzaNumber,
			message: msg,
			success: success,
		}

		return &p
	}
	
	return &PizzaOrder{
		PizzaNumber: pizzaNumber,
	}
}

func pizzaria(pizzaMaker *Producer){
	// keep track of which pizza we are making
	var i = 0

	// run forever or until we receive a quit notifcation
	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		// try to make a pizza
		// 
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzaria(pizzaJob)
}
