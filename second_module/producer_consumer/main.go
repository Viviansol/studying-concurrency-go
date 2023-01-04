package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const numberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {

	pizzaNumber++
	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("Making pizza #%d! It will take %d seconds...", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** the cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("pizza order #%d is ready", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {

	//keep track of wich pizza we are making
	var i = 0

	//run forever or until we receive a quit notification
	//try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}

	}
}

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//print ou a message
	color.Cyan("The pizzeria is open for business")
	color.Cyan("---------------------------------")

	//create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in the background
	go pizzeria(pizzaJob)

	//create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("order #%d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The costumer is really mad ")
			}
		} else {
			color.Cyan("done making pizzas ...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** error closing channel", err)
			}
		}
	}

	//print out the ending message
	color.Cyan("--------------")
	color.Cyan("Done for today")
	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day ...")
	case pizzasFailed >= 6:
		color.Red("it was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow(" it was an okay day...")
	case pizzasFailed >= 2:
		color.Yellow("it was a pretty good day...")
	default:
		color.Green("it was a great day...")
	}
}
