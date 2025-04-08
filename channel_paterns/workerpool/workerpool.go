package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func fetching(url string) int {
	_, err := http.Get(url)

	if err != nil {
		return 0
	}

	return 200
}

func WorkerPool(ctx context.Context, countWorkers int, workerFunc func(string) int, dataCh <-chan string) <-chan int {
	resultCh := make(chan int, 3)

	wg := &sync.WaitGroup{}

	for i := 0; i < countWorkers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-dataCh:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case resultCh <- workerFunc(val):
					}
				}

			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

func main() {
	urls := []string{"http://google.com", "http://stepik.org", "http://ya.ru", "https://translate.yandex.ru",
		"https://e.mail.ru", "https://github.com", "https://www.youtube.com"}

	dataCh := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for _, val := range urls {
			dataCh <- val
		}
		close(dataCh)
	}()

	resCh := WorkerPool(ctx, 3, fetching, dataCh)

	for val := range resCh {
		fmt.Println(val)
	}

	cancel()

}
