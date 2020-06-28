package main

import (
	"fmt"
)

func main() {
//make the below code work by making chan a receive and send
	//	cs := make(chan<- int)
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
