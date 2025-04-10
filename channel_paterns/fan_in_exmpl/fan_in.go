package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	outputCh := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(len(channels))

	for _, channel := range channels {
		go func() {
			defer wg.Done()

			for msg := range channel {
				outputCh <- msg
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	return outputCh
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		for i := 0; i < 100; i += 3 {
			ch1 <- i
			ch2 <- i + 1
			ch3 <- i + 2
		}
	}()

	for i := range mergeChannels(ch1, ch2, ch3) {
		fmt.Println(i)
	}

}
