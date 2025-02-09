package main

import (
	"sync/atomic"
)

func main() {

}

type CustomWaitGroup struct {
	counter int32
	done    chan struct{}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		done: make(chan struct{}),
	}
}

func (wg *CustomWaitGroup) Add(delta int) {
	atomic.AddInt32(&wg.counter, int32(delta))
	if atomic.LoadInt32(&wg.counter) < 0 {
		panic("Negative counter")
	}
}
func (wg *CustomWaitGroup) Done() {
	if atomic.AddInt32(&wg.counter, -1) == 0 {
		close(wg.done)
	}
}
func (wg *CustomWaitGroup) Wait() {
	<-wg.done
}
