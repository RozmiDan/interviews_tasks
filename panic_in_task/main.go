package main

import (
	"fmt"
	"strconv"
	"sync"
)

// В данной задаче канал никогда не закроется и горутина main останется заблокированной в select

// Исходное решение

// func main() {
// 	var wg = &sync.WaitGroup{}
// 	ch := make(chan string, 3)

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(ch chan<- string, i int) {
// 			defer wg.Done()

// 			ch <- fmt.Sprintf("Goroutine %s", strconv.Itoa(i))
// 		}(ch, i)
// 	}

// 	for {
// 		select {
// 		case v := <-ch:
// 			fmt.Println(v)
// 		}
// 	}

// 	wg.Wait()
// 	close(ch)
// }

// Исправленное решение

func main() {
	var wg = &sync.WaitGroup{}
	ch := make(chan string, 3)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(ch chan<- string, i int) {
			defer wg.Done()

			ch <- fmt.Sprintf("Goroutine %s", strconv.Itoa(i))
		}(ch, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
