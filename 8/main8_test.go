package main

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestCustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup()

	wg.Add(3)

	var counter int

	go func() {
		counter++
		wg.Done()
	}()
	go func() {
		counter++
		wg.Done()
	}()
	go func() {
		counter++
		wg.Done()
	}()
	wg.Wait()

	if counter != 3 {
		t.Errorf("Expected counter 3, resived: %d", counter)
	}
}

func TestCustomWaitGroupNegativeCounter(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to negative counter")
		}
	}()

	wg := NewCustomWaitGroup()
	wg.Add(-1)
}

func TestCustomWaitGroupConcurrent(t *testing.T) {
	wg := NewCustomWaitGroup()
	var counter int32
	const numberGoroutines = 100

	wg.Add(numberGoroutines)

	for i := 0; i < numberGoroutines; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
			fmt.Println(counter)
		}()
	}

	wg.Wait()

	if counter != numberGoroutines {
		t.Errorf("Expected counter %d, resived %d", numberGoroutines, counter)
	}
}
