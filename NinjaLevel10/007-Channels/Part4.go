package main

import (
	"fmt"
	"sync"
)

//write a program that
//○ launches 10 goroutines
//■ each goroutine adds 10 numbers to a channel
//○ pull the numbers off the channel and print them
var wg sync.WaitGroup

func groutine(i int)chan int{
	wg.Add(100)
	c10 := make(chan int)
	go func(){
		for i := 0; i < 10; i++{
			wg.Done()
			c10 <- i
		}
		close(c10)
	}()
	wg.Wait()
	fmt.Printf("Channel %d added \n", c10)
	return c10
}

func main() {
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		wg.Done()
		go groutine(i)
		fmt.Println("Goroutine: ", i)
	}
	wg.Wait()
	fmt.Println("main finished")
}