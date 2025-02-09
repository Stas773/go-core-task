package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxRandNum = 100

func main() {
	randChan := randomGenerator()

	for i := 0; i < 10; i++ {
		num := <-randChan
		fmt.Println(num)
	}
}
func randomGenerator() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			ch <- rand.Intn(MaxRandNum)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return ch
}
