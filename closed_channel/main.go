package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 23
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
