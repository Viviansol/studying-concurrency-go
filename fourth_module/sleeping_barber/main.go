package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

// variables
var seatingCapacity = 30
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {

	// seed our random number generator
	rand.Seed(time.Now().UnixNano())

	//print welcome message
	color.Yellow("The sleeping barber problem")
	color.Yellow("---------------------------")

	//create channels if we need any
	clientsChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	//create the barber shop

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientsChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day")

	// add barbers

	shop.addBarber("Frank")
	shop.addBarber("genaro")
	shop.addBarber("ana")
	shop.addBarber("filipe")
	shop.addBarber("anderson")
	shop.addBarber("nath")

	// start the barber shop as a go routine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	//add clients

	i := 1
	go func() {
		for {
			randmMillseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randmMillseconds)):
				shop.addClient(fmt.Sprintf("Client %d", i))
				i++
			}
		}
	}()

	// block until the barber shop is closed
	<-closed
}
