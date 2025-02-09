package main

import (
	"sync"
	"testing"
)

func TestToCubeFloat(t *testing.T) {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(1)
	go toCubeFloat(ch1, ch2, &wg)

	go func() {
		ch1 <- 2
		ch1 <- 3
		ch1 <- 4
		close(ch1)
	}()

	go func() {
		wg.Wait()
		close(ch2)
	}()

	result := []float64{}
	for v := range ch2 {
		result = append(result, v)
	}

	expected := []float64{8, 27, 64}

	if len(result) != len(expected) {
		t.Errorf("Expected lenght: %d, resived lenght: %d", len(expected), len(result))
		return
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %v, resived %v", expected[i], v)
		}
	}
}
