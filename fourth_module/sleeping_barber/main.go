package main

import (
	"github.com/fatih/color"
	"math/rand"
	"time"
)

// variables
var seatingCapacity = 10
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

	// start the barber shop as a go routine

	//add clients

	// block until the barber shop is closed

}
