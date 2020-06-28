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
var w2 sync.WaitGroup


func main() {
	var gr = 10
	wg.Add(10)
	w2.Add(100)
	go func() {
		for i := 1; i <= gr; i++{
		fmt.Printf("GoRoutines: %v\n", i)
		wg.Done()
		go func(){
			ch := make(chan int)
			for j := 1; j <= 10; j++{
				fmt.Printf("Channel: %v\n", j)
				ch <- j
			}
		}()
		w2.Done()
	}
	}()
w2.Wait()
wg.Wait()
}