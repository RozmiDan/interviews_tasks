package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Исходная ф-я merge выглядит так:

// func merge(ch ...<-chan int) <-chan int {
// 	out := make(chan int)

// 	// Имеется 2 входных канала in1 и in2 и один выходной out.
// 	// Требуется реализовать функцию merge, которая будет сливать данные из входных каналов в один

// 	return out
// }

// По сути тут нужно написать Fan-in

func merge(ch ...<-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}

	for _, channel := range ch {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for val := range channel {
				out <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func source(sourcefunc func(int) int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- sourcefunc(i)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}()

	return ch
}

func main() {
	in1 := source(func(_ int) int { return rand.Int() })
	in2 := source(func(i int) int { return i })

	out := merge(in1, in2)

	for val := range out {
		fmt.Println(val)
	}
}
