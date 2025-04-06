package main

import (
	"fmt"
	"time"
)

// Задачка на внимательность - важно помнить что тут два вызова будут последовательными, соответственно 6 сек

func main() {
	timer := time.Now()

	_, _ = <-worker(), <-worker()

	fmt.Println(time.Since(timer).Seconds())
}

func worker() <-chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	return ch
}
