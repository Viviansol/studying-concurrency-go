package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {

	msg = "Hello, world!"
	wg.Add(1)
	go updateMessage("hello, universe")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("hello, cosmos")
	wg.Wait()
	printMessage()
	wg.Add(1)
	go updateMessage("hello, world")
	wg.Wait()
	printMessage()
}
