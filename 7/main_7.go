package main

import (
	"fmt"
	"sync"
)

func main() {
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
	for v := range merged {
		fmt.Println(v)
	}
}

func mergeChans(chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	resultCh := make(chan int)

	output := func(ch <-chan int) {
		defer wg.Done()
		for v := range ch {
			resultCh <- v
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}
