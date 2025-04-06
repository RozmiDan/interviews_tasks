package main

import "fmt"

func main() {

	tasks := make(chan int, 1)

	tasks <- 1
	fmt.Println("dsfsd")
	tasks <- 2
	fmt.Println("dsfsd")

	fmt.Println()
}
