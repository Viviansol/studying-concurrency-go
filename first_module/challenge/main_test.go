package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {

	wg.Add(1)
	go updateMessage("episolon")
	wg.Wait()
	if msg != "episolon" {
		t.Errorf("expected to find episolon, but its not there")
	}

}

func Test_printMessage(t *testing.T) {

	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "episolon"

	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if !strings.Contains(output, "episolon") {
		t.Errorf("expected episolon")
	}

}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "hello, universe") {
		t.Errorf("error")
	}
	if !strings.Contains(output, "hello, cosmos") {
		t.Errorf("error")
	}
	if !strings.Contains(output, "hello, world") {
		t.Errorf("error")
	}

}