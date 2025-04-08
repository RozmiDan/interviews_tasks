package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	channel chan struct{}
}

func NewSemaphore(cap int) *Semaphore {
	return &Semaphore{
		channel: make(chan struct{}, cap),
	}
}

func (s *Semaphore) Acquire() {
	s.channel <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.channel
}

func main() {
	sem := NewSemaphore(3)
	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		sem.Acquire()
		wg.Add(1)
		go func(i int) {
			defer func() {
				sem.Release()
				wg.Done()
			}()
			time.Sleep(time.Second)
			fmt.Println("Doing some work", i)
		}(i)
	}

	wg.Wait()
}
