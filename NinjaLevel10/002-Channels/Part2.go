package main

import (
	"fmt"
)

func main() {
	//make the below code work by making chan a receive and send
	//	cs := make(<-chan int)
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
