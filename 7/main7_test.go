package main

import (
	"testing"
)

func TestMergeChans(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)
	}()
	go func() {
		ch2 <- 4
		ch2 <- 5
		ch2 <- 6
		close(ch2)
	}()
	go func() {
		ch3 <- 7
		ch3 <- 8
		ch3 <- 9
		close(ch3)
	}()

	merged := mergeChans(ch1, ch2, ch3)
	expected := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}
	received := make(map[int]bool)
	for v := range merged {
		received[v] = true
	}

	for key := range expected {
		if !received[key] {
			t.Errorf("Expected value %d don't received", key)
		}
	}
}
