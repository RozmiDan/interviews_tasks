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

	// str := "привет"
	// str1 := "hello1"
	// str3 := str + str1
	// str4 := []rune(str3)
	// fmt.Println(len(str), len(str1), len(str3), len(str4))

	// 	arr := make([]int, 4, 10)
	// 	for i := range 4 {
	// 		arr[i] = i
	// 	}

	// 	arr2 := somefunc(&arr)
	// 	fmt.Println(arr, arr2)
	// 	wg := sync.WaitGroup{}
	// 	mu := sync.Mutex{}
	// 	mu.Lock()
	// }

	// func somefunc(arr *[]int) []int {
	// 	*arr = append(*arr, 10)
	// 	fmt.Println(arr)
	// 	return *arr

	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 45, 3, 2, 5, 63, 45, 64, 3, 4, 57, 5, 5, 4, 3, 2}
	//mp := make(map[int]int)

	// for val, i := range arr {
	// 	fmt.Println(i, val)
	// }

	// var arr2 []int
	// fmt.Println(arr2 == nil)
	// fmt.Println(len(arr2))
}
