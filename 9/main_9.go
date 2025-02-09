package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)
	var wg sync.WaitGroup

	wg.Add(1)
	go toCubeFloat(ch1, ch2, &wg)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- uint8(i)
		}
		close(ch1)
	}()

	go func() {
		wg.Wait()
		close(ch2)
	}()

	for v := range ch2 {
		fmt.Println(v)
	}
}

func toCubeFloat(ch1 <-chan uint8, ch2 chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch1 {
		ch2 <- math.Pow(float64(v), 3)
	}
}
