package main

import "fmt"

func splitChan(channel <-chan int, count int) []chan int {
	arr := make([]chan int, count)

	for i := 0; i < count; i++ {
		arr[i] = make(chan int)
	}

	go func() {
		defer func() {
			for _, i := range arr {
				close(i)
			}
		}()

		iter := 0
		for value := range channel {
			arr[iter] <- value
			iter = (iter + 1) % count
		}
	}()

	return arr
}

func main() {
	someChan := make(chan int)

	go func() {
		defer close(someChan)

		for i := 0; i < 100; i++ {
			someChan <- i
		}
	}()

	var fstChClosed, scndChClosed, thdChClosed bool

	channels := splitChan(someChan, 3)

	for !fstChClosed || !scndChClosed || !thdChClosed {
		select {
		case val, ok := <-channels[0]:
			if !ok {
				fstChClosed = true
				continue
			}
			fmt.Printf("value: %v from %v channel\n", val, 0)

		case val, ok := <-channels[1]:
			if !ok {
				scndChClosed = true
				continue
			}
			fmt.Printf("value: %v from %v channel\n", val, 1)

		case val, ok := <-channels[2]:
			if !ok {
				thdChClosed = true
				continue
			}
			fmt.Printf("value: %v from %v channel\n", val, 2)
		}
	}

}
