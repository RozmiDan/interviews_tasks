package main

import (
	"fmt"
)

func main() {
	fmt.Println(arr())
}

func arr() []int {
	arr := []int{0, 1, 2, 3, 4}

	fmt.Println(arr, "\n", len(arr), cap(arr))

	arr[0] = 1

	fmt.Println(arr, "\n", len(arr), cap(arr))

	arr = append(arr, 5) // 1 1 2 3 4 5

	fmt.Println(arr, "\n", len(arr), cap(arr)) //l=6 c=10

	newArr := append(arr, 6) // 1 1 2 3 4 5 6

	fmt.Println(arr, "\n", len(arr), cap(arr))          //l=6 c=10
	fmt.Println(newArr, "\n", len(newArr), cap(newArr)) //l=7 c=10

	arr = append(arr, 7)

	fmt.Println(arr[6])
	fmt.Println(newArr[6])

	arr[0] = 10

	fmt.Println(arr, "\n", len(arr), cap(arr))
	fmt.Println(newArr, "\n", len(newArr), cap(newArr))

	return newArr
}
