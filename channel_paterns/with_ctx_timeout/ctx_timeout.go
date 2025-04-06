package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	resultCh := make(chan int)

	go RPCCall(ctx, resultCh)

	val, ok := <-resultCh
	if !ok {
		fmt.Println("Call is too long")
	}
	fmt.Println("Call is handled", val)
}

func RPCCall(ctx context.Context, ch chan<- int) {
	work := time.After(time.Duration(rand.Intn(4)) * time.Second)

	select {

	case <-work:
		select {

		case ch <- rand.Int():
			fmt.Println("work is done")
		case <-ctx.Done():
			fmt.Println("work finished but context cancelled")
		}

	case <-ctx.Done():
		close(ch)
		fmt.Println("work is not done")
	}

}
