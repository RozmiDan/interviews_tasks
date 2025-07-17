package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

// частая задачка когда нужно обернуть долгий вызов

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println(callLongFunction(ctx, longFunction))
}

func callLongFunction(ctx context.Context, lfunc func() int) (int, error) {

	resChan := make(chan int)

	go func() {
		defer close(resChan)
		resChan <- lfunc()
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, errors.New("timeout")
		case val := <-resChan:
			return val, nil
		}
	}
}

func longFunction() int {
	time.Sleep(time.Duration(rand.N[int64](5)) * time.Second)
	return 52
}
