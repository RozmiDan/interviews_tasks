package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4

	close(ch1)
	close(ch2)

	ch3 := merge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

func merge[T any](chans ...chan T) <-chan T {
	resultCh := make(chan T)

	wg := &sync.WaitGroup{}

	for _, ch := range chans {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for val := range ch {
				resultCh <- val
			}

		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

// not good realisation
// func merge[T any](chans ...chan T) <-chan T {
// 	resultCh := make(chan T)

// 	go func() {
// 		ch1Open, ch2Open := true, true
// 		defer close(resultCh)

// 		for ch1Open || ch2Open {

// 			select {

// 			case val, ok := <-chans[0]:
// 				if !ok {
// 					ch1Open = false
// 				} else {
// 					resultCh <- val
// 				}

// 			case val, ok := <-chans[1]:
// 				if !ok {
// 					ch2Open = false
// 				} else {
// 					resultCh <- val
// 				}
// 			}

// 		}
// 	}()

// 	return resultCh
// }
