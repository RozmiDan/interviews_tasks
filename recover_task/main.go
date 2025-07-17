package main

import "fmt"

func someFunc() {
	panic("some panic")
}

func main() {
	fmt.Println("start")
	recoverWraper(someFunc)
}

func recoverWraper(funcWithPanic func()) {
	defer func() {
		if ok := recover(); ok != nil {
			fmt.Println("catch panic")
		}
	}()
	funcWithPanic()
}
