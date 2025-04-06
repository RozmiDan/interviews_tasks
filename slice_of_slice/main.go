package main

import "fmt"

func main() {
	test1 := []int{1, 2, 3, 4, 5}
	test1 = test1[:3] // 1 2 3 len = 3 cap = 5
	fmt.Println(test1, len(test1), cap(test1))

	test2 := test1[3:] // 4 5 len = 0 cap = 2
	fmt.Println(test2[:2], len(test2), cap(test2))
}
