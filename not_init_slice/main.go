package main

import (
	"fmt"
	"sync"
)

// Здесь нет wg, канал небуф, поэтому после первого прочтения из канала произойдет дедлок, так
// чтение происходит один раз
// А также ОЧЕНЬ ВНИМАТЕЛЬНО на инициализацию канала - это nil канал

// func run() {
// 	var ch chan int
// 	for i := 0; i < 3; i++ {
// 		go func(idx int) {
// 			ch <- (idx + 1) * 2
// 		}(i)
// 	}
// 	fmt.Println("result:", <-ch)
// 	time.Sleep(2 * time.Second)
// }

// func main() {
// 	run()
// }

func run() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			ch <- (idx + 1) * 2
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	run()
}
