package main

import (
	"fmt"
	"sync"
	"time"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha", "beta", "delta", "gama", "pi", "zeta", "eta", "theta", "episolon",
	}
	wg.Add(len(words))
	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	wg.Wait()
	wg.Add(1)
	time.Sleep(1 * time.Second)
	printSomething("second", &wg)
}
