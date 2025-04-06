package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan int)
	cancelCh := make(chan struct{})
	timer := time.After(2 * time.Second)

	RPCCall(channel, cancelCh)

	select {
	case val, ok := <-channel:
		if !ok {
			fmt.Println("Kanal zakrit")
		}
		fmt.Println(val)
		break
	case <-timer:
		close(cancelCh)
		fmt.Println("Timeout")
	}

	fmt.Println("Programm finished")
}

func RPCCall(channel chan<- int, cancChan <-chan struct{}) {

	val := time.Second * (time.Duration(rand.Intn(5)))
	timer := time.After(val)
	fmt.Println("RPCCall timer is ", val)

	go func() {
		select {
		case <-timer:
			fmt.Println("Request is done")
			channel <- rand.Int()
		case <-cancChan:
			fmt.Println("Cancel operation handled")
		}
		close(channel)
	}()
}
