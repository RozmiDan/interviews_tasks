package main

import (
	"fmt"
	"sync"
)

func main() {
	// exmpl
	// {
	// 	b := new(bool)
	// 	fmt.Println(b)
	// 	modify(b)
	// 	fmt.Println(b)
	// }
	// return
	//exmpl
	// {
	// 	a := [2]int{0, 0}

	// 	if a == nil {
	// 		fmt.Println("true")
	// 	} else {
	// 		fmt.Println("false")
	// 	}
	// }
	// return
	// exmpl
	{
		var counter = 20
		wg := &sync.WaitGroup{}

		for i := range counter {
			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Println(i * i)
			}()
		}

		wg.Wait()
		fmt.Println("done")
	}
}

func modify(b *bool) {
	b = nil
}
