package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	var wg sync.WaitGroup
	wg.Add(1)
	go printSomething("episolon", &wg)
	wg.Wait()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if !strings.Contains(output, "episolon") {
		t.Errorf("expected to find episolon, but its not there")
	}
}
