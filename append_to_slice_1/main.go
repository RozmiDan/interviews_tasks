package main

import "fmt"

// Нужно понимать, что при присваивании в другой слайс результат append не увеличивает len(x), он
// остается также 3, однако len(y) будет уже ноль

func main() {
	var x []int
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y, z)
}

// func main() {

// 	var x []int
// 	x = append(x, 0)
// 	x = append(x, 1)
// 	x = append(x, 2)
// 	y := append(x, 3)
// 	z := append(y, 4)
// 	fmt.Println(y, z)

// [0 1 2 3] [0 1 2 3 4]
// }
