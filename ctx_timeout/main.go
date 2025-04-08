package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	NewCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for {
		select {
		case <-NewCtx.Done():
			fmt.Println("ctx done")
			return
		default:
			fmt.Println("working...")
			time.Sleep(1 * time.Second)
		}

	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)

	defer cancel()

	doWork(ctx)
}
