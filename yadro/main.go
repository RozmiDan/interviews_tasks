package main

import "fmt"

func main() {

	c := make(chan int, 7)
	a := cap(c)

	fmt.Println(a)
}
