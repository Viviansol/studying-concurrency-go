package main

import (
	"fmt"
	"strings"
)

func shout(ping, pong chan string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)
	go shout(ping, pong)
	fmt.Println("type sometihing and press enter(enter q to quit)")
	for {
		fmt.Println("->")
		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("all done. Closing channels")
	close(ping)
	close(pong)
}
