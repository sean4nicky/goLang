package main

import (
	"fmt"
)

func main() {
	// using a buffer to get the below working
	//c := make(chan int)
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
