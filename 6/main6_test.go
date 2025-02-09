package main

import (
	"testing"
	"time"
)

func TestRandomGenerator(t *testing.T) {
	randChan := randomGenerator()

	for i := 0; i < 10; i++ {
		select {
		case num := <-randChan:
			if num < 0 || num >= MaxRandNum {
				t.Errorf("Expected number in the range [0, %d), resived: %d", MaxRandNum, num)
			}
		case <-time.After(1 * time.Second):
			t.Fatal("Timeout reading from a channel")
		}
	}
}
