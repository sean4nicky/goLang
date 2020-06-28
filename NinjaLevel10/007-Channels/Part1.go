package main

import (
	"fmt"
	"sync"
	"time"
)

//write a program that
//○ launches 10 goroutines
//■ each goroutine adds 10 numbers to a channel
//○ pull the numbers off the channel and print them
var wg sync.WaitGroup

func routine(i int)chan int{
	c10 := make(chan int)
	go func(){
		for i := 1; i <= 10; i++ {
				c10 <- i
		defer wg.Done()
		}
	}()
	fmt.Printf("Goroutine: %v started\n Channel %d%d%d%d%d%d%d%d%d%d added \n ", i, <-c10, <-c10+1, <-c10+1, <-c10+1, <-c10+1, <-c10, <-c10+1, <-c10+1, <-c10+1, <-c10+1)
	//fmt.Printf("Channel %d added \n", c10)
	return c10
}

func main() {
	wg.Add(100)
	for i := 1; i <= 10; i++ {
		time.NewTimer(time.Second *30)
		go routine(i)
		}
	wg.Wait()
	fmt.Println("main finished")
}