package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main() {
	c := make(chan int)
wg.Add(10)
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
		wg.Done()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(<-c, k)
	}
wg.Wait()
	fmt.Println("about to exit")
}
