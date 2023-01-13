package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length")
		}
	}
}

func Test_dinningVaryingDelays(t *testing.T) {

	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{
			name:  "zero delay",
			delay: time.Second * 0,
		},
		{
			name:  "quarter second delay",
			delay: time.Millisecond * 250,
		},
		{
			name:  "half second delay",
			delay: time.Millisecond * 500,
		},
	}

	for _, a := range theTests {
		orderFinished = []string{}
		eatTime = a.delay
		sleepTime = a.delay
		thinkTime = a.delay
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length")
		}
	}
}
